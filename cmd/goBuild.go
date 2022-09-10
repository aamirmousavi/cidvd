package cmd

import (
	"cidvd/action"
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var build = &cobra.Command{
	Use:   "build",
	Short: "build your go project",
	Long:  `build your go project`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("\tGoBuild\n")
		if out, err := action.GoBuild(BuildAddress, BuildTarget); err != nil {
			fmt.Printf("Error: %v", err)
			os.Exit(1)
		} else {
			fmt.Printf("%s", out)
		}
	},
}

var (
	BuildAddress string
	BuildTarget  string
)
