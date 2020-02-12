/*
Copyright © 2020 takumakume

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
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var (
	dryrun   bool
	interval int
)

// cleanupCmd represents the cleanup command
var cleanupCmd = &cobra.Command{
	Use:   "cleanup",
	Short: "clean up broken telepresence resources",

	Run: func(cmd *cobra.Command, args []string) {
		initCommand()
		if interval == 0 {
			if err := telepoliceObj.Cleanup(dryrun); err != nil {
				fmt.Println(err.Error())
				os.Exit(1)
			}
		} else {
			if err := telepoliceObj.WatchWithCleanup(dryrun, interval); err != nil {
				fmt.Println(err.Error())
				os.Exit(1)
			}
		}
	},
}

func init() {
	rootCmd.AddCommand(cleanupCmd)
	cleanupCmd.PersistentFlags().BoolVar(&dryrun, "dry-run", false, "dry run mode")
	cleanupCmd.PersistentFlags().IntVarP(&interval, "interval", "i", 0, "Start in daemon mode and process every specified number of seconds")
}
