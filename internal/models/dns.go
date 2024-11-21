package models

import (
	"fmt"
	"net"
)

// dnsServers := map[string]string{
// 	"google-1":      "8.8.8.8",
// 	"google-2":    "8.8.4.4",
// 	"cloudflare-1":  "1.1.1.1",
// 	"quad9-1":       "9.9.9.9",
//     "quad9-2":       "64.64.64.64",
// 	"adguard-home":  "93.184.216.34",
// }

// dnsRecordTypes := map[string]string{
//     "A":      "IPv4 address",
//     "AAAA":   "IPv6 address",
//     "CNAME":  "Canonical name record",
//     "MX":     "Mail exchange record",
//     "NS":     "Name server record",
//     "SOA":    "Start of authority record",
//     "TXT":    "Text record",
//     "SRV":    "Service record",
//     "PTR":    "Pointer record",
//     "RP":     "Responsible person record",
//     "HINFO":  "Host information record",
//     "MINFO":  "Mailbox information record",
//     "MX":     "Mail exchange record",
//     "TXT":    "Text record",
//     "AAAA":   "IPv6 address",
// }

// type DNSEntry struct {
// 	Name string
// 	Type string
// 	TTL  int
// }

func NewDNSEntry(name, dnsType string, ttl int) *DNSEntry {
	return &DNSEntry{
		Name: name,
		Type: dnsType,
		TTL:  ttl,
	}
}

func ProcessDnsEntry(domain string) {
	var ips []net.IP
	// for _, resolver := range dnsServers {
	// 	addrs, err := net.LookupHost(resolver + ":" + domain)
	// 	if err == nil {
	// 		ips = append(ips, net.ParseIP(addrs[0]))
	// 	}
	// }
	fmt.Println("DNS records for", domain, ":", ips)
}
