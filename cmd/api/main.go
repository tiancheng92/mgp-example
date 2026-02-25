package main

import (
	"mgp_example/server"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:     "mgp",
	Version: "v1.0",
	Short:   "mgp",
	Long:    "mgp",
	Run: func(_ *cobra.Command, _ []string) {
		server.Run()
	},
}

func main() {
	if err := rootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}
