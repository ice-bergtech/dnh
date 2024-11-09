// main.go
package main

import (
	"github.com/ice-bergtech/dnh/src/internal/config"
	"github.com/ice-bergtech/dnh/src/internal/db"
	"github.com/spf13/cobra"
)

func main() {
	rootCmd := &cobra.Command{
		Use:   "dnh",
		Short: "A brief description of your application",
		Long: `A longer description that spans multiple lines and likely contains
examples and usage of using your application.`,
		Run: func(cmd *cobra.Command, args []string) {
			println("Log line")

		},
	}

	rootCmd.Execute()

	// println("Ba dum, tss!")
	cfg, err := config.NewConfig()
	if err != nil {
		return
	}

	//setup db
	_, err = db.OpenSqlite(cfg)
	if err != nil {
		return
	}
}
