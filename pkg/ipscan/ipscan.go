package ipscan

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"math"
	"net/http"
	"os"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/OPSWAT/mdcloud-go/pkg/api"
	"github.com/OPSWAT/mdcloud-go/pkg/aws"
	"github.com/OPSWAT/mdcloud-go/pkg/utils"
	logger "github.com/sirupsen/logrus"

	"github.com/pkg/errors"

	"github.com/aws/aws-sdk-go/service/ec2"
)

// Request struct
type Request struct {
	Address []string `json:"address"`
}

// Response result of ipscan request
type Response struct {
	StartTime  time.Time    `json:"start_time"`
	DetectedBy int          `json:"detected_by"`
	Address    string       `json:"address"`
	Error      api.ApiError `json:"error"`
}

var (
	token  string
	ips    []string
	groups *ec2.DescribeSecurityGroupsOutput
	rp     Response
)

// ScanSGs starts scan for users AWS SGs
func ScanSGs(api api.API, sgs []string) {
	url := fmt.Sprintf("%s/ip", api.URL)
	if sgs != nil {
		getGroups(sgs)
	} else {
		getGroups(nil)
	}
	if len(groups.SecurityGroups) > 0 {
		for _, g := range groups.SecurityGroups {
			getSGIPs(g.IpPermissions)
			getSGIPs(g.IpPermissionsEgress)
		}
	}
	if len(ips) == 0 {
		logger.Warningln("security group empty")
		os.Exit(1)
	}

	var wg sync.WaitGroup
	if len(ips) < 5 {
		wg.Add(1)
	} else {
		wg.Add(int(len(ips) / 5))
	}
	errc := make(chan error, 1)
	done := make(chan bool, 1)

	for i := 0; i < len(ips); i += 5 {
		batch := ips[i:int(math.Min(float64(i+5), float64(len(ips))))]
		go func(batch []string, url string, wg *sync.WaitGroup) {
			mips, err := json.Marshal(&Request{Address: batch})
			if err != nil {
				errc <- err
			}
			req, err := http.NewRequest(http.MethodPost, url, bytes.NewBuffer(mips))
			if err != nil {
				errc <- err
			}
			req.Header.Set("Authorization", api.Token)
			req.Header.Set("Content-Type", "application/json")

			resp, err := api.Client.Do(req)
			if err != nil {
				errc <- errors.WithMessage(err, "request error")
			}
			defer resp.Body.Close()

			if resp.StatusCode >= 400 && resp.StatusCode <= 500 {
				if body, err := ioutil.ReadAll(resp.Body); err == nil {
					if e := json.Unmarshal(body, &rp); e != nil {
						errc <- errors.New("parsing response")
					}
					logger.WithFields(logger.Fields{
						"code":  rp.Error.Code,
						"msg":   strings.Join(rp.Error.Messages, " "),
						"batch": strings.Join(batch, ","),
					}).Fatalln(resp.Status)
				}
			} else {
				if remaining, err := strconv.Atoi(resp.Header.Get("x-ratelimit-remaining")); err == nil {
					if remaining > 0 {
						for _, item := range batch {
							logger.WithField("ip", item).Infoln("OK")
						}
					} else {
						errc <- errors.New("limit reached")
					}
				}
			}
			time.Sleep(5 * time.Second)
			defer wg.Done()
		}(batch, url, &wg)
	}

	go func() {
		wg.Wait()
		close(done)
	}()

	select {
	case <-done:
	case err := <-errc:
		if err != nil {
			logger.Errorln(err)
			return
		}
	}
}

// ListIPs starts scan for users AWS SGs
func ListIPs(sgs []string) {
	if sgs != nil && len(sgs) > 0 {
		getGroups(sgs)
	} else {
		getGroups(nil)
	}
	if len(groups.SecurityGroups) > 0 {
		for _, g := range groups.SecurityGroups {
			getSGIPs(g.IpPermissions)
			getSGIPs(g.IpPermissionsEgress)
		}
	}
	if len(ips) == 0 {
		logger.Fatalln(errors.New("Security group empty"))
	}

	// TODO: display security groups before
	for _, ip := range ips {
		logger.Infoln(ip)
	}
}

func getSGIPs(ipPermissions []*ec2.IpPermission) {
	for _, ipPermission := range ipPermissions {
		for _, ipRange := range ipPermission.IpRanges {
			item := strings.Split(*ipRange.CidrIp, "/")
			if item[1] == "32" && !utils.StringInSlice(item[0], ips) {
				ips = append(ips, item[0])
			}
		}
	}
}

func getGroups(sgs []string) {
	var err error
	svc := ec2.New(aws.Session)
	if sgs != nil {
		groups, err = svc.DescribeSecurityGroups(&ec2.DescribeSecurityGroupsInput{GroupIds: utils.StringSlice(sgs)})
	} else {
		groups, err = svc.DescribeSecurityGroups(&ec2.DescribeSecurityGroupsInput{})
	}
	if err != nil {
		logger.WithField("security_groups", sgs).Fatalln(errors.WithMessage(err, "Error requesting group description"))
	}
}
