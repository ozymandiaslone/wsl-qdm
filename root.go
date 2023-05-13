package main

import (
	"fmt"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "qdm",
	Short: "Move files between Windows/WSL filesystems",
	Long:  `Move files between Windows/WSL filesystems`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("root called")
	},
}

func init() {
	rootCmd.AddCommand(sendCmd())
}
