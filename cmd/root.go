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
	Use:     "geolite2lookup <ipaddress>",
	Short:   "look up country using IP Address",
	Long:    `look up country using IP Address`,
	Version: version.Version,
	RunE: func(cmd *cobra.Command, args []string) error {
		if len(args) != 1 {
			return errors.New("ipaddress required")
		}

		ip := args[0]

		file, _ := cmd.Flags().GetString("file")
		if file == "" {
			file = os.Getenv("GEOLITE2_MMDB_FILE")
			if file == "" {
				return errors.New("GEOLITE2_MMDB_FILE not set")
			}
		}

		reader, err := geoip2.NewCountryReaderFromFile(file)
		if err != nil {
			return err
		}

		record, err := reader.Lookup(net.ParseIP(ip))
		if err != nil {
			println("GeoLite2 Country Edition: IP Address not found")
		} else {
			fmt.Printf("GeoLite2 Country Edition: %s, %s\n", record.Country.ISOCode, record.Country.Names["en"])
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
	rootCmd.Flags().StringP("file", "f", "", "mmdb_file")
}
