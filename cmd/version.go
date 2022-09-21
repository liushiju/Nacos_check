package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
)

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "查看版本",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Version: 0.5.1")
	},
}

func init() {
	rootCmd.AddCommand(versionCmd)
}
