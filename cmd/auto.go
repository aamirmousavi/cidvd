package cmd

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/spf13/cobra"
)

var auto = &cobra.Command{
	Use:   "auto",
	Short: "auto your go project",
	Long:  `auto your go project`,
	Run: func(cmd *cobra.Command, args []string) {
		if err := initAppConfig(cfgFileAddress); err != nil {
			fmt.Printf("Error In finding and parsing config: %v\n", err)
			os.Exit(1)
		}
		fmt.Printf("\tAuto\n")
		pull.Run(pull, args)
		build.Run(build, args)
		restart.Run(restart, args)
	},
}

func initAppConfig(addres string) error {
	type Config struct {
		Here           bool    `json:"here"`
		PullAddress    string  `json:"pullAddress"`
		BuildAddress   string  `json:"buildAddress"`
		BuildTarget    string  `json:"buildTarget"`
		RestartService *string `json:"restartService"`
	}
	var config Config
	fileJson, err := ioutil.ReadFile(addres)
	if err != nil {
		return err
	}
	if err := json.Unmarshal(fileJson, &config); err != nil {
		return err
	}

	PullAddress = config.PullAddress
	BuildAddress = config.BuildAddress
	BuildTarget = config.BuildTarget
	if config.RestartService != nil {
		RestartServiceName = *config.RestartService
	}
	if config.Here {
		mydir, err := os.Getwd()
		if err != nil {
			panic(err)
		}
		PullAddress = mydir
		BuildAddress = mydir + "/cmd"
	}
	return nil
}
