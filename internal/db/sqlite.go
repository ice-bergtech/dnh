package db

import (
	"github.com/ice-bergtech/dnh/src/internal/config"
	"gorm.io/driver/sqlite" // Sqlite driver based on CGO

	// "github.com/glebarez/sqlite" // Pure go SQLite driver, checkout https://github.com/glebarez/sqlite for details
	"gorm.io/gorm"
)

// github.com/mattn/go-sqlite3
// db, err := gorm.Open(sqlite.Open("dnh.db"), &gorm.Config{})

// Create a function to generate the GORM config from the internal config
func GenerateGormConfig(internalConfig config.AppConfig) *gorm.Config {
	// Implement logic to convert the internal config into a GORM config
	gormConfig := &gorm.Config{}
	// https://gorm.io/docs/gorm_config.html
	// if internalConfig.DBType == "sqlite" {

	// } else if internalConfig.DBType == "postgresql" {

	// } else {
	// 	panic("unsupported database type")
	// }

	return gormConfig
}

func OpenSqlite(internalConfig *config.AppConfig) (*gorm.DB, error) {
	dbCfg := GenerateGormConfig(*internalConfig)
	return gorm.Open(sqlite.Open(internalConfig.DBHost), dbCfg)
}
