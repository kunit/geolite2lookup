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
	Short:   "look up country/city using IP Address or hostname",
	Long:    `look up country/city using IP Address or hostname`,
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

		locale, _ := cmd.Flags().GetString("locale")

		displayType := ""

		code, _ := cmd.Flags().GetBool("code")
		if code {
			displayType = "code"
		} else {
			name, _ := cmd.Flags().GetBool("name")
			if name {
				displayType = "name"
			} else {
				info, _ := cmd.Flags().GetBool("info")
				if info {
					displayType = "info"
				}
			}
		}

		mmdb := fmt.Sprintf("%s/%s", dir, file)

		countryCode := ""
		countryName := ""

		if mmdbType == "Country" {
			reader, err := geoip2.NewCountryReaderFromFile(mmdb)
			if err != nil {
				return err
			}

			record, err := reader.Lookup(net.ParseIP(ip))
			if err == nil {
				countryCode = record.Country.ISOCode
				countryName = record.Country.Names[locale]
			}

			if displayType == "code" {
				fmt.Printf("%s\n", countryCode)
			} else if displayType == "name" {
				fmt.Printf("%s\n", countryName)
			} else if countryCode != "" {
				fmt.Printf("GeoLite2 Country Edition: %s, %s\n", countryCode, countryName)
			} else {
				fmt.Printf("GeoLite2 Country Edition: IP Address not found\n")
			}
		} else {
			reader, err := geoip2.NewCityReaderFromFile(mmdb)
			if err != nil {
				return err
			}

			record, err := reader.Lookup(net.ParseIP(ip))
			if err == nil {
				countryCode = record.Country.ISOCode
				countryName = record.Country.Names[locale]
			}

			if displayType == "code" {
				fmt.Printf("%s\n", countryCode)
			} else if displayType == "name" {
				fmt.Printf("%s\n", countryName)
			} else if countryCode != "" {
				fmt.Printf("GeoLite2 City Edition: %s, %s\n", countryCode, countryName)
				if displayType == "info" {
					fmt.Printf("  Country Code: %s\n", record.Country.ISOCode)
					fmt.Print("  Location:")
					if record.City.Names[locale] != "" {
						fmt.Printf(" %s,", record.City.Names[locale])
					}
					if len(record.Subdivisions) != 0 {
						for _, s := range record.Subdivisions {
							fmt.Printf(" %s,", s.Names[locale])
						}
					}
					fmt.Printf(" %s,", record.Country.Names[locale])
					fmt.Printf(" %s\n", record.Continent.Names[locale])

					if record.Postal.Code != "" {
						fmt.Printf("  Postal Code: %s\n", record.Postal.Code)
					}
					fmt.Printf("  TimeZone: %s\n", record.Location.TimeZone)
					fmt.Printf("  Approximate Coordinates: %.4f, %.4f\n", record.Location.Latitude, record.Location.Longitude)
					fmt.Printf("  Accuracy Radius: %d\n", record.Location.AccuracyRadius)
					if record.Country.ISOCode == "US" {
						fmt.Printf("  MetroCode: %d\n", record.Location.MetroCode)
					}
				}
			} else {
				fmt.Printf("GeoLite2 City Edition: IP Address not found\n")
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
	defaultLocale := os.Getenv("GEOLITE2LOOKUP_LOCALE")
	if defaultLocale == "" {
		defaultLocale = "en"
	}

	defaultMMDBType := os.Getenv("GEOLITE2LOOKUP_MMDB_TYPE")
	if defaultMMDBType == "" {
		defaultMMDBType = "Country"
	}

	defaultMMDBDir := os.Getenv("GEOLITE2LOOKUP_MMDB_DIR")
	if defaultMMDBDir == "" {
		defaultMMDBDir = "/usr/share/GeoIP2"
	}

	rootCmd.Flags().StringP("locale", "l", defaultLocale, "locale to use when display names")
	rootCmd.Flags().StringP("type", "t", defaultMMDBType, "MMDB Edition")
	rootCmd.Flags().StringP("dir", "d", defaultMMDBDir, "MMDB directory")
	rootCmd.Flags().StringP("file", "f", "", "MMDB filename (default \"GeoLite2-[type].mmdb\")")
	rootCmd.Flags().BoolP("info", "i", false, "show additional information (only type \"City\")")
	rootCmd.Flags().BoolP("code", "c", false, "show country code only")
	rootCmd.Flags().BoolP("name", "n", false, "show country name only")
}
