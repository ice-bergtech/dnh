package api

//go:generate go run -modfile=../../tools/go.mod github.com/oapi-codegen/oapi-codegen/v2/cmd/oapi-codegen --config=config.yaml ../../api/api.yml

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/ice-bergtech/dnh/src/internal/config"
	"github.com/ice-bergtech/dnh/src/internal/db_client"
	model "github.com/ice-bergtech/dnh/src/internal/model_ent"
	"github.com/ice-bergtech/dnh/src/internal/model_ent/domain"
	"github.com/ice-bergtech/dnh/src/internal/model_ent/ipaddress"
	"github.com/ice-bergtech/dnh/src/internal/model_ent/whois"
)

// optional code omitted

type Server struct{}

func NewServer() Server {
	return Server{}
}

// (GET /ping)
func (Server) GetPing(w http.ResponseWriter, r *http.Request) {
	resp := model.IPAddress{}

	w.WriteHeader(http.StatusOK)
	_ = json.NewEncoder(w).Encode(resp)
}

func (Server) GetDomainsDomainName(w http.ResponseWriter, r *http.Request, domainName string) {
	resp, _ := DBClient.Domain.Query().Where(domain.NameEQ(domainName)).First(r.Context())

	w.WriteHeader(http.StatusOK)
	_ = json.NewEncoder(w).Encode(resp)
}

// Get IP address details
// (GET /ipaddresses/{ip})
func (Server) GetIpaddressesIp(w http.ResponseWriter, r *http.Request, ip string) {
	resp, _ := DBClient.IPAddress.Query().Where(ipaddress.IPEQ(ip)).First(r.Context())

	w.WriteHeader(http.StatusOK)
	_ = json.NewEncoder(w).Encode(resp)
}

// Get WHOIS data for an IP range, domain, or AS
// (GET /whois/{query})
func (Server) GetWhoisQuery(w http.ResponseWriter, r *http.Request, query string) {
	resp, _ := DBClient.Whois.Query().Where(whois.QueryEQ(query)).First(r.Context())

	w.WriteHeader(http.StatusOK)
	_ = json.NewEncoder(w).Encode(resp)
}

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
