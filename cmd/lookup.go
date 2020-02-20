package cmd

import (
	"net/url"
	"regexp"
	"strings"

	"github.com/OPSWAT/mdcloud-go/pkg/lookup"
	"github.com/OPSWAT/mdcloud-go/pkg/utils"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

var download bool
var (
	ipRegex, _ = regexp.Compile(`^(([0-9]|[1-9][0-9]|1[0-9]{2}|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9]{2}|2[0-4][0-9]|25[0-5])$`)
)

func IsIpv4(ipAddress string) bool {
	ipAddress = strings.Trim(ipAddress, " ")
	return ipRegex.MatchString(ipAddress)
}

// lookupCmd represents the lookup command
var lookupCmd = &cobra.Command{
	Use:   "lookup [hash, ip, domain]",
	Short: "Lookup or download file or IPs/Domains",
	Long:  "Lookup or download file by md5, sha1, sha256 or IPs/Domains",
	Run: func(cmd *cobra.Command, args []string) {
		var ips []string
		var domains []string
		var urls []string
		var hashes []string
		for _, arg := range args {
			if strings.Contains(arg, ".") {
				// todo: add a more valid way to parse input
				if IsIpv4(arg) {
					ips = append(ips, arg)
				} else {
					u, _ := url.Parse(arg)
					if u.Path != "" || u.Host != "" {
						if u.Host != "" {
							url := u.Host + u.Path
							if u.RawQuery != "" {
								url = url + "/" + u.RawQuery
							}
							urls = append(urls, url)
						} else {
							domains = append(domains, arg)
						}
					} else {
						logrus.Fatalln("args not valid")
					}
				}
			} else {
				hashes = append(hashes, arg)
			}
		}
		utils.VerifyArgsOrRun(args, 0, func() {
			if len(ips) > 0 {
				lookup.ByIP(API, ips)
			}
			if len(domains) > 0 {
				lookup.ByDomain(API, domains)
			}
			if len(urls) > 0 {
				lookup.ByUrl(API, urls)
			}
			if len(hashes) > 0 {
				lookup.ByHash(API, hashes, download)
			}
		}, func() { cmd.Help() })
	},
}

func init() {
	RootCmd.AddCommand(lookupCmd)
	lookupCmd.PersistentFlags().BoolVarP(&download, "download", "d", false, "get download url")
}
