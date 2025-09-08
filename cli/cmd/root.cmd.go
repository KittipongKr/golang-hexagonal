package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

var (
	zone string // ตัวแปร global สำหรับเก็บค่า flag
)

var rootCommand = &cobra.Command{
	Use:   "app",
	Short: "root cli",
}

func Execute() {
	if err := rootCommand.Execute(); err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCommand.PersistentFlags().StringVarP(&zone, "zone", "z", "", "Environment profile (dev|uat|prd)")
}
