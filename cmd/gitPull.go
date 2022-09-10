package cmd

import (
	"cidvd/action"
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var pull = &cobra.Command{
	Use:   "pull",
	Short: "pull all branches from git",
	Long:  `pull all branches from git`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("\tGitPull\n")
		if out, err := action.GitPull(PullAddress); err != nil {
			fmt.Printf("Error: %v\n", err)
			os.Exit(1)
		} else {
			fmt.Printf("%s\n", out)
		}
	},
}
var (
	PullAddress string
)
