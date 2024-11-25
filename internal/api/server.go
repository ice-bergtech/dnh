package api

//go:generate go run -modfile=../../tools/go.mod github.com/oapi-codegen/oapi-codegen/v2/cmd/oapi-codegen --config=config.yaml ../../api/api.yml

import (
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/ice-bergtech/dnh/src/internal/config"
	"github.com/ice-bergtech/dnh/src/internal/db_client"
	model "github.com/ice-bergtech/dnh/src/internal/model_ent"
)

var DBClient *model.Client

func Start(cfg config.AppConfig) {
	// Setup db connection
	var err error
	DBClient, err = db_client.NewClient(&cfg)
	if err != nil {
		log.Fatal("error creating db client: %w", err)
	}

	// create a type that satisfies the `api.ServerInterface`, which contains an implementation of every operation from the generated code
	server := NewServer()

	r := mux.NewRouter()
	// get an `http.Handler` that we can use
	h := HandlerFromMux(server, r)

	s := &http.Server{
		Handler: h,
		Addr:    "0.0.0.0:" + strconv.FormatInt(int64(cfg.Port), 10),
	}

	// And we serve HTTP until the world ends.
	log.Fatal(s.ListenAndServe())
}
