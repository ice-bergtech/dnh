package main

import (
	"github.com/ice-bergtech/dnh/src/internal/config"
	"github.com/spf13/cobra"
)

func QueryCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "query [flags] value",
		Short: "Run a query against scanned data",
		Run: func(cmd *cobra.Command, args []string) {
			config.NewConfig()
		},
	}

	return cmd
}
