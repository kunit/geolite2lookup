/*
Copyright © 2021 TAKAHASHI Kunihiko <kunihiko.takahashi@gmail.com>

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:
The above copyright notice and this permission notice shall be included in
all copies or substantial portions of the Software.
THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
THE SOFTWARE.
*/
package cmd

import (
	"errors"
	"fmt"
	"net"
	"os"

	"github.com/IncSW/geoip2"
	"github.com/kunit/geolite2lookup/version"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:     "geolite2lookup <ipaddress|hostname>",
	Short:   "look up country using IP Address or hostname",
	Long:    `look up country using IP Address or hostname`,
	Version: version.Version,
	RunE: func(cmd *cobra.Command, args []string) error {
		if len(args) != 1 {
			return errors.New("ipaddress required")
		}

		ip_or_hostname := args[0]

		var ip string

		addr, err := net.ResolveIPAddr("ip", ip_or_hostname)

		if err != nil {
			ip = ip_or_hostname
		} else {
			ip = addr.String()
		}

		mmdbType, _ := cmd.Flags().GetString("type")
		dir, _ := cmd.Flags().GetString("dir")
		file, _ := cmd.Flags().GetString("file")
		if file == "" {
			if mmdbType == "Country" {
				file = "GeoLite2-Country.mmdb"
			} else {
				file = "GeoLite2-City.mmdb"
			}
		}

		mmdb := fmt.Sprintf("%s/%s", dir, file)

		if mmdbType == "Country" {
			reader, err := geoip2.NewCountryReaderFromFile(mmdb)
			if err != nil {
				return err
			}

			record, err := reader.Lookup(net.ParseIP(ip))
			if err != nil {
				fmt.Printf("GeoLite2 Country Edition: IP Address not found\n")
			} else {
				fmt.Printf("GeoLite2 Country Edition: %s, %s\n", record.Country.ISOCode, record.Country.Names["en"])
			}
		} else {
			reader, err := geoip2.NewCityReaderFromFile(mmdb)
			if err != nil {
				return err
			}

			record, err := reader.Lookup(net.ParseIP(ip))
			if err != nil {
				fmt.Printf("GeoLite2 City Edition: IP Address not found\n")
			} else {
				fmt.Printf("GeoLite2 City Edition: %s, %s\n", record.Country.ISOCode, record.Country.Names["en"])
			}
		}

		return nil
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	rootCmd.Flags().StringP("type", "t", "Country", "MMDB Edition")
	rootCmd.Flags().StringP("dir", "d", "/usr/share/GeoIP2", "MMDB directory")
	rootCmd.Flags().StringP("file", "f", "", "MMDB filename (default \"GeoLite2-[type].mmdb\")")
}
