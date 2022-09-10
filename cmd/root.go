package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var (
	// Used for flags.
	cfgFileAddress string

	rootCmd = &cobra.Command{
		Use:   "cidvd",
		Short: "CIDVD is a tool to update your local repository and build and restart manully your system services",
		Long:  `CIDVD is a tool to update your local repository and build and restart manully your system services`,
	}
)

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func init() {
	mydir, err := os.Getwd()
	if err != nil {
		panic(err)
	}
	auto.PersistentFlags().StringVar(&cfgFileAddress, "config", mydir+"/cidvd.json", "config file (default is ./cidvd.json)")
	rootCmd.AddCommand(auto)
	pull.Flags().StringVar(&PullAddress, "address", mydir, "address of your repository default is current directory")
	rootCmd.AddCommand(pull)
	build.Flags().StringVar(&BuildAddress, "address", mydir+"/cmd", "address of your repository default is current directory")
	build.Flags().StringVar(&BuildTarget, "target", "/data/go-services/api/api", "target of your build")
	rootCmd.AddCommand(build)
	restart.Flags().StringVar(&RestartServiceName, "service", "", "serviceName of your service")
	rootCmd.AddCommand(restart)
}
