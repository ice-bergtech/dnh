package db_client

import (
	"context"
	"fmt"

	"github.com/ice-bergtech/dnh/src/internal/config"
	"github.com/ice-bergtech/dnh/src/internal/model_ent"
)

func NewClientSqlite(cfg *config.AppConfig) (*model_ent.Client, error) {
	return model_ent.Open("sqlite3",
		//dnh_data.db
		fmt.Sprintf("file:%s?cache=shared&_fk=1", cfg.DB.Name))
}

func NewClientPostgres(cfg *config.AppConfig) (*model_ent.Client, error) {
	return model_ent.Open("postgres",
		fmt.Sprintf("host=%s port=%d user=%s dbname=%s password=%s",
			cfg.DB.Host,
			cfg.DB.Port,
			cfg.DB.User,
			cfg.DB.Name,
			cfg.DB.Pass))
}

func NewClient(cfg *config.AppConfig) (*model_ent.Client, error) {
	var (
		client *model_ent.Client
		err    error
	)
	switch cfg.DB.Type {
	case "sqlite":
		fallthrough
	case "sqlite3":
		client, err = NewClientSqlite(cfg)
	case "postgresql":
		client, err = NewClientPostgres(cfg)
	}
	if err != nil {
		return nil, fmt.Errorf("failed to set db connections: %v", err)
	}

	// Run the auto migration tool.
	if err = client.Schema.Create(context.Background()); err != nil {
		return nil, fmt.Errorf("failed creating schema resources: %v", err)
	}

	return client, err
}
