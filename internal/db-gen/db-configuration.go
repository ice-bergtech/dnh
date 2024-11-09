//go:generate go run db-configuration.go

// configuration.go
package main

import (
	dnh "github.com/ice-bergtech/dnh/src/internal/lib"
	"gorm.io/driver/sqlite"
	"gorm.io/gen"
	"gorm.io/gorm"
)

func main() {
	// Initialize the generator with configuration
	g := gen.NewGenerator(gen.Config{
		OutPath:       "../db", // output directory, default value is ./query
		Mode:          gen.WithDefaultQuery | gen.WithQueryInterface,
		FieldNullable: true,
	})

	// Initialize a *gorm.DB instance
	db, _ := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})

	// Use the above `*gorm.DB` instance to initialize the generator,
	// which is required to generate structs from db when using `GenerateModel/GenerateModelAs`
	g.UseDB(db)

	// Generate default DAO interface for those specified structs
	g.ApplyBasic(
		dnh.IPAddress{},
		dnh.ASNInfo{},
		dnh.DNSEntry{},
		dnh.Domain{},
		dnh.Path{},
		dnh.Nameserver{},
		dnh.Registrar{},
		dnh.Whois{},
	)

	// Generate default DAO interface for those generated structs from database
	// companyGenerator := g.GenerateModelAs("company", "MyCompany")
	// g.ApplyBasic(
	// 	g.GenerateModel("users"),
	// 	companyGenerator,
	// 	g.GenerateModelAs("people", "Person",
	// 		gen.FieldIgnore("deleted_at"),
	// 		gen.FieldNewTag("age", `json:"-"`),
	// 	),
	// )

	// Execute the generator
	g.Execute()
}
