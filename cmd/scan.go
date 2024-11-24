package main

import (
	"github.com/ice-bergtech/dnh/src/internal/config"
	"github.com/spf13/cobra"
)

func ScanCmd() *cobra.Command {

	cmd := &cobra.Command{
		Use:   "scan [flags] path",
		Short: "Run a scan against a provided input",
		Run: func(cmd *cobra.Command, args []string) {

			config.LoadConfig()
		},
	}
	config.CobraSetup(cmd)

	return cmd
}
