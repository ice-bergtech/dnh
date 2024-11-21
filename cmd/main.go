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
		Short: "A brief description of your application",
		Long: `A longer description that spans multiple lines and likely contains
examples and usage of using your application.`,
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
