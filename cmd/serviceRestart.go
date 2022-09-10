package cmd

import (
	"cidvd/action"
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var restart = &cobra.Command{
	Use:   "restart",
	Short: "Restart a service",
	Long:  `Restart a service`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("\tServiceRestart\n")
		if RestartServiceName == "" {
			initAppConfig(cfgFileAddress)
		}
		if out, err := action.ServiceRestart(RestartServiceName); err != nil {
			fmt.Printf("Error: %v\n", err)
			os.Exit(1)
		} else {
			fmt.Printf("%s", out)
		}
	},
}
var (
	RestartServiceName string
)
