package models

import (
	"net"
	"time"

	"gorm.io/gorm"
)

type Scan struct {
	gorm.Model
	ScanID    string      `json:"scanid"`
	Timestamp time.Time   `json:"timestamp"`
	Results   ScanResults `json:"results" gorm:"serializer:json"`
}

func (Scan) TableName() string {
	return "scan"
}

type ScanResults struct {
	IPAddress     IPAddress  `json:"ip_address"`
	ASNInfo       ASNInfo    `json:"asn_info"`
	DNSEntries    []DNSEntry `json:"dns_entries"`
	Domain        Domain     `json:"domain"`
	Paths         []Path     `json:"paths"`
	Nameserver    Nameserver `json:"nameserver"`
	RegistrarData Registrar  `json:"registrar"`
	Whois         Whois      `json:"whois"`
}

type IPAddress struct {
	gorm.Model
	IP          net.IP     `json:"ip"`
	Mask        net.IPMask `json:"mask"`
	Advertisers []ASNInfo  `json:"advertisers" gorm:"many2many:asninfo_asn;"`
	Tags        []Tags     `json:"tags" gorm:"serializer:json"`
}

type ASNInfo struct {
	gorm.Model
	ASN      int    `json:"asn"`
	Country  string `json:"country"`
	Registry string `json:"registry"`
	Tags     []Tags `json:"tags" gorm:"serializer:json"`
}

type DNSEntry struct {
	gorm.Model
	Name     string    `json:"name"`
	Type     string    `json:"type"`
	Value    string    `json:"value"`
	TTL      int       `json:"ttl"`
	TimeFist time.Time `json:"time_first"`
	TimeLast time.Time `json:"time_last"`
	Tags     []Tags    `json:"tags"  gorm:"serializer:json"`
}

type Domain struct {
	gorm.Model
	Name        string       `json:"name"`
	Records     []DNSEntry   `json:"records" gorm:"many2many:domain_dns"`
	Parent      string       `json:"parent,omitempty"`
	Paths       []Path       `json:"paths" gorm:"many2many:domain_path"`
	Addresses   []IPAddress  `json:"addresses" gorm:"many2many:domain_address"`
	Nameservers []Nameserver `json:"nameservers" gorm:"many2many:domain_ns"`
	TimeFist    time.Time    `json:"time_first"`
	TimeLast    time.Time    `json:"time_last"`
	Tags        []Tags       `json:"tags"  gorm:"serializer:json"`
}

type Path struct {
	gorm.Model
	Path string `json:"path"`
	Tags []Tags `json:"tags" gorm:"serializer:json"`
}

type Nameserver struct {
	gorm.Model
	Name     string    `json:"name"`
	IP       IPAddress `json:"ip" gorm:"many2many:ns_ip"`
	TimeFist time.Time `json:"time_first"`
	TimeLast time.Time `json:"time_last"`
	Tags     []Tags    `json:"tags"  gorm:"serializer:json"`
}

type Registrar struct {
	gorm.Model
	Name        string      `json:"name"`
	URL         string      `json:"url"`
	CountryCode string      `json:"country_code"`
	Phone       string      `json:"phone,omitempty"`
	Fax         string      `json:"fax,omitempty"`
	Address     []IPAddress `json:"address" gorm:"many2many:registrar_ip"`
	TimeFist    time.Time   `json:"time_first"`
	TimeLast    time.Time   `json:"time_last"`
	Tags        []Tags      `json:"tags" gorm:"serializer:json"`
}

// Struct to store WHOIS data for an IP range, domain, or AS.
type Whois struct {
	gorm.Model
	IPRange     IPAddress    `json:"ip_range,omitempty" gorm:"foreignKey:IP"`         // Optional: IP range in CIDR notation
	DomainName  string       `json:"domain_name,omitempty"`                           // Optional: Domain name
	ASN         uint32       `json:"asn,omitempty"`                                   // Optional: Autonomous System Number
	Country     string       `json:"country,omitempty"`                               // Optional: Country code
	Created     time.Time    `json:"created,omitempty"`                               // Optional: WHOIS data creation timestamp
	Updated     time.Time    `json:"updated,omitempty"`                               // Optional: WHOIS data last update timestamp
	Registrar   string       `json:"registrar,omitempty"`                             // Optional: Domain registrar
	Nameservers []Nameserver `json:"nameservers,omitempty" gorm:"many2many:whois_ns"` // Optional: List of nameservers associated with the domain or IP range
	TimeFist    time.Time    `json:"time_first"`
	TimeLast    time.Time    `json:"time_last"`
	Tags        []Tags       `json:"tags"  gorm:"serializer:json"`
}

type Tags struct {
	gorm.Model
	Tag       string `json:"tag"`
	Reference string `json:"reference"`
	Notes     string `json:"notes"`
}
