package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var randomCmd = &cobra.Command{
	Use:   "random",
	Short: "A brief description of your command",
	Long:  `A longer description that spans multiple lines and likely contains examples`,
	Run: func(cmd *cobra.Command, args []string) {
		getRandom()
	},
}

func init() {
	rootCmd.AddCommand(randomCmd)
}

func getRandom() {
	fmt.Println("hi bro")
}
