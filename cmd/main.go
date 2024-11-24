// main.go
package main

import (
	"fmt"

	"github.com/ice-bergtech/dnh/src/internal/config"
	_ "github.com/ice-bergtech/dnh/src/internal/model_ent"
	"github.com/spf13/cobra"
	"golang.org/x/exp/slog"
)

func main() {
	rootCmd := &cobra.Command{
		Use:   "dnh",
		Short: "A tool to collect data about a domain",
		Long:  `Provided a property of a domain, collects various information around it.`,
	}

	rootCmd.AddCommand(
		//InitCmd(),
		ServerCmd(),
		ScanCmd(),
		QueryCmd(),
		//GenerateCmd(),
	)

	_, err := config.LoadConfig()
	if err != nil {
		slog.Error(fmt.Sprintf("Failed to read config file: %s", err.Error()))
	}

	_ = rootCmd.Execute()
}
