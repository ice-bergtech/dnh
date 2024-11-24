package main

import (
	"github.com/ice-bergtech/dnh/src/internal/api"
	"github.com/ice-bergtech/dnh/src/internal/config"
	"github.com/spf13/cobra"
)

func ServerCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "server",
		Short: "Start a http server to serve/process data",
		Run: func(cmd *cobra.Command, args []string) {
			cfg, _ := config.LoadConfig()
			api.Start(*cfg)
		},
	}

	return cmd
}
