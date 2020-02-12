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
	"runtime"
	"strings"

	"github.com/spf13/cobra"
	"github.com/takumakume/telepolice/telepolice"
)

var (
	useInClusterConfig           bool
	concurrency                  int
	ignorerablePodStartTimeOfSec int
	telepoliceObj                *telepolice.Telepolice
	namespaces                   string
	allNamespaces                bool
	verbose                      bool
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "telepolice",
	Short: "A tool to clean up broken telepresence resources",
	Long:  `A tool to clean up broken telepresence resources`,
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	numCPU := runtime.NumCPU()
	rootCmd.PersistentFlags().IntVarP(&concurrency, "concurrency", "c", numCPU, fmt.Sprintf("Multiple processing default: %d (runtime.NumCPU())", numCPU))
	rootCmd.PersistentFlags().IntVar(&ignorerablePodStartTimeOfSec, "ignorerable-pod-start-time-of-sec", 180, fmt.Sprintf("Pod immediately after startup is in preparation and passes health check for the specified number of seconds default: %d s", 60))
	rootCmd.PersistentFlags().BoolVar(&useInClusterConfig, "use-in-cluster-config", false, "Use the service account kubernetes gives to pods")
	rootCmd.PersistentFlags().StringVarP(&namespaces, "namespaces", "n", "default", "Target namespaces (e.g. ns1,ns2)")
	rootCmd.PersistentFlags().BoolVarP(&allNamespaces, "all-namespaces", "A", false, "Target all namespaces")
	rootCmd.PersistentFlags().BoolVar(&verbose, "verbose", false, "Output verbose logs")
}

func strToSlice(s string) []string {
	return strings.Split(strings.TrimRight(s, ","), ",")
}

func initCommand() {
	c := telepolice.NewConfig(concurrency, ignorerablePodStartTimeOfSec)
	var err error
	if useInClusterConfig {
		telepoliceObj, err = telepolice.NewByInClusterConfig(c)
	} else {
		telepoliceObj, err = telepolice.NewByKubeConfig(c)
	}

	if verbose {
		telepoliceObj.EnableVerbose()
	}

	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	if allNamespaces {
		telepoliceObj.SetAllNamespaces()
	} else {
		telepoliceObj.SetNamespaces(strToSlice(namespaces))
	}
}
