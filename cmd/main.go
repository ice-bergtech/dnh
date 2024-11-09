// main.go
package main

import (
	"github.com/ice-bergtech/dnh/src/internal/config"
	"github.com/ice-bergtech/dnh/src/internal/db"
)

func main() {
	println("Ba dum, tss!")
	cfg, err := config.NewConfig()
	if err != nil {
		return
	}

	//setup db
	db, err := db.OpenSqlite(cfg)
	if err != nil {
		return
	}
}
