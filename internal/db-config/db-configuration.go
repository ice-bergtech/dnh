// configuration.go
package main

import (
	models "github.com/ice-bergtech/dnh/src/internal/model_ent"
	"gorm.io/driver/sqlite"
	"gorm.io/gen"
	"gorm.io/gorm"
)

// Dynamic SQL
type Querier interface {
	// SELECT * FROM @@table WHERE name = @name{{if role !=""}} AND role = @role{{end}}
	FilterWithNameAndRole(name, role string) ([]gen.T, error)
}

// This is ran to generate the gorm DB
func main() {
	// Initialize the generator with configuration
	g := gen.NewGenerator(gen.Config{
		OutPath:       "../db", // output directory, default value is ./query
		Mode:          gen.WithDefaultQuery | gen.WithQueryInterface | gen.WithoutContext,
		FieldNullable: true,
	})

	// Initialize a *gorm.DB instance
	db, _ := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})

	// Use the above `*gorm.DB` instance to initialize the generator,
	// which is required to generate structs from db when using `GenerateModel/GenerateModelAs`
	g.UseDB(db)

	// Generate default DAO interface for those specified structs
	g.ApplyBasic(
		models.Scan{},
		models.IPAddress{},
		models.ASNInfo{},
		models.DNSEntry{},
		models.Domain{},
		models.Path{},
		models.Nameserver{},
		models.Registrar{},
		models.Whois{},
	)

	// Generate Type Safe API with Dynamic SQL defined on Querier interface for `model.User` and `model.Company`
	g.ApplyInterface(
		func(Querier) {},
		models.Scan{},
		models.IPAddress{},
		models.ASNInfo{},
		models.DNSEntry{},
		models.Domain{},
		models.Path{},
		models.Nameserver{},
		models.Registrar{},
		models.Whois{})

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
