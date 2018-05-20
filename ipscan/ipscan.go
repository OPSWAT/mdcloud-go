package ipscan

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"math"
	"net/http"
	"os"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/OPSWAT/mdcloud-go/api"

	"github.com/pkg/errors"

	"github.com/OPSWAT/mdcloud-go/aws"
	"github.com/OPSWAT/mdcloud-go/utils"
	"github.com/aws/aws-sdk-go/service/ec2"
)

// Request struct
type Request struct {
	Address []string `json:"address"`
}

// Response result of ipscan request
type Response struct {
	Success bool `json:"success"`
	Error   struct {
		Code     int      `json:"code"`
		Messages []string `json:"messages"`
	} `json:"error"`
	Data []struct {
		StartTime  time.Time `json:"start_time"`
		DetectedBy int       `json:"detected_by"`
		Address    string    `json:"address"`
	} `json:"data"`
}

// Apikey used
var Apikey string

var (
	token  string
	ips    []string
	groups *ec2.DescribeSecurityGroupsOutput
	rp     Response
)

// ScanSGs starts scan for users AWS SGs
func ScanSGs(sgs []string) {
	url := api.URL + "ip"
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
		log.Printf("security group empty")
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
		go func(batch []string, url string, token string, wg *sync.WaitGroup) {
			mips, err := json.Marshal(&Request{Address: batch})
			if err != nil {
				errc <- err
			}
			req, err := http.NewRequest("POST", url, bytes.NewBuffer(mips))
			if err != nil {
				errc <- err
			}
			req.Header.Set("authorization", "apikey "+token)
			req.Header.Set("Content-Type", "application/json")

			client := &http.Client{Timeout: 5 * time.Minute}
			resp, err := client.Do(req)
			if err != nil {
				errc <- errors.WithMessage(err, "request error")
			}
			defer resp.Body.Close()

			if resp.StatusCode >= 400 && resp.StatusCode <= 500 {
				if body, err := ioutil.ReadAll(resp.Body); err == nil {
					if e := json.Unmarshal(body, &rp); e != nil {
						errc <- errors.New("parsing response")
					}
					log.Fatalln(fmt.Sprintf("error: %v - %s while scanning batch %s", rp.Error.Code, strings.Join(rp.Error.Messages, " "), strings.Join(batch, ",")))
				}
			} else {
				if remaining, err := strconv.Atoi(resp.Header.Get("x-ratelimit-remaining")); err == nil {
					if remaining > 0 {
						for _, item := range batch {
							fmt.Println(item + " : OK")
						}
					} else {
						errc <- errors.New("limit reached")
					}
				}
			}
			time.Sleep(5 * time.Second)
			defer wg.Done()
		}(batch, url, Apikey, &wg)
	}

	go func() {
		wg.Wait()
		close(done)
	}()

	select {
	case <-done:
	case err := <-errc:
		if err != nil {
			fmt.Println("error: ", err)
			return
		}
	}
}

// ListIPs starts scan for users AWS SGs
func ListIPs(sgs []string) {
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
		fmt.Println("security group empty")
		os.Exit(1)
	}

	// TODO: display security groups before
	for _, ip := range ips {
		fmt.Println(ip)
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
		log.Fatalf("Error requesting group description: %v", err)
	}
}
