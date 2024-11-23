// main.go
package main

import (
	"github.com/ice-bergtech/dnh/src/internal/config"
	"github.com/ice-bergtech/dnh/src/internal/db"
	"github.com/ice-bergtech/dnh/src/internal/models"
	"github.com/spf13/cobra"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
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
	database, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})

	db.SetDefault(database)

	// query the first user
	database.AutoMigrate(
		models.Scan{},
		models.ASNInfo{},
		models.DNSEntry{},
		models.Domain{},
		models.IPAddress{},
		models.Nameserver{},
		models.Path{},
		models.Registrar{},
		models.Whois{},
	)
}
