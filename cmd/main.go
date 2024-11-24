// main.go
package main

import (
	_ "github.com/ice-bergtech/dnh/src/internal/model_ent"
	"github.com/spf13/cobra"
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

	_ = rootCmd.Execute()
}
