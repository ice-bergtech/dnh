package api

//go:generate go run -modfile=../../tools/go.mod github.com/oapi-codegen/oapi-codegen/v2/cmd/oapi-codegen --config=config.yaml ../../api/api.yml

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/ice-bergtech/dnh/src/internal/config"
	models "github.com/ice-bergtech/dnh/src/internal/model_ent"
)

// optional code omitted

type Server struct{}

func NewServer() Server {
	return Server{}
}

// (GET /ping)
func (Server) GetPing(w http.ResponseWriter, r *http.Request) {
	resp := models.IPAddress{}

	w.WriteHeader(http.StatusOK)
	_ = json.NewEncoder(w).Encode(resp)
}

func StartServer(cfg config.AppConfig) {
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
