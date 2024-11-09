package dnh

import (
	"net"
	"time"
)

type IPaddress struct {
	IP          net.IP            `json:"ip"`
	Mask        net.IPMask        `json:"mask"`
	Advertisers []ASNinfo         `json:"advertisers"`
	Tags        map[string]string `json:"tags"`
}

type ASNinfo struct {
	ASN      int               `json:"asn"`
	Country  string            `json:"country"`
	Registry string            `json:"registry"`
	Tags     map[string]string `json:"tags"`
}

type DNSEntry struct {
	Name      string            `json:"name"`
	Type      string            `json:"type"`
	Value     string            `json:"value"`
	TTL       int               `json:"ttl"`
	Timestamp time.Time         `json:"timestamp"`
	Tags      map[string]string `json:"tags"`
}

type Domain struct {
	Name      string            `json:"name"`
	Records   []DNSEntry        `json:"records"`
	Parent    string            `json:"parent,omitempty"`
	Paths     []Path            `json:"paths"`
	Addresses []IPaddress       `json:"addresses"`
	Timestamp time.Time         `json:"timestamp"`
	Tags      map[string]string `json:"tags"`
}

type Path struct {
	Path string            `json:"path"`
	Tags map[string]string `json:"tags"`
}

type Nameserver struct {
	Name      string            `json:"name"`
	IP        IPaddress         `json:"ip"`
	Timestamp time.Time         `json:"timestamp"`
	Tags      map[string]string `json:"tags"`
}

type RegistrarData struct {
	Name        string            `json:"name"`
	URL         string            `json:"url"`
	CountryCode string            `json:"country_code"`
	Phone       string            `json:"phone,omitempty"`
	Fax         string            `json:"fax,omitempty"`
	Address     []IPaddress       `json:"address"`
	Timestamp   time.Time         `json:"timestamp"`
	Tags        map[string]string `json:"tags"`
}

// Struct to store WHOIS data for an IP range, domain, or AS.
type WhoisData struct {
	IPRange     IPaddress         `json:"ip_range,omitempty"`    // Optional: IP range in CIDR notation
	DomainName  string            `json:"domain_name,omitempty"` // Optional: Domain name
	ASN         uint32            `json:"asn,omitempty"`         // Optional: Autonomous System Number
	Country     string            `json:"country,omitempty"`     // Optional: Country code
	Created     time.Time         `json:"created,omitempty"`     // Optional: WHOIS data creation timestamp
	Updated     time.Time         `json:"updated,omitempty"`     // Optional: WHOIS data last update timestamp
	Registrar   string            `json:"registrar,omitempty"`   // Optional: Domain registrar
	Nameservers []Nameserver      `json:"nameservers,omitempty"` // Optional: List of nameservers associated with the domain or IP range
	Timestamp   time.Time         `json:"timestamp"`
	Tags        map[string]string `json:"tags"`
}
