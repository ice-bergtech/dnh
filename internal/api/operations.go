package api

//go:generate go run -modfile=../../tools/go.mod github.com/oapi-codegen/oapi-codegen/v2/cmd/oapi-codegen --config=config.yaml ../../api/api.yml

import (
	"encoding/json"
	"net/http"

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

func (Server) GetDataMeta(w http.ResponseWriter, r *http.Request) {}

// Get data details
// (GET /data/{type})
func (Server) GetDataDetails(w http.ResponseWriter, r *http.Request, pType string, params GetDataDetailsParams) {
}

// Start a Scan
// (POST /scan)
func (Server) StartScan(w http.ResponseWriter, r *http.Request) {}

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
