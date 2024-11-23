// main.go
package main

import (
	"github.com/ice-bergtech/dnh/src/internal/config"
	_ "github.com/ice-bergtech/dnh/src/internal/model_ent"
	"github.com/spf13/cobra"
)

func main() {
	rootCmd := &cobra.Command{
		Use:   "dnh",
		Short: "A tool to collect data about a domain",
		Long:  `Provided a property of a domain, will collect `,
		Run: func(cmd *cobra.Command, args []string) {
			println("Log line")

		},
	}

	rootCmd.Execute()

	// println("Ba dum, tss!")
	_, err := config.NewConfig()
	if err != nil {
		return
	}

	//setup db
	// Initialize a *gorm.DB instance
}
