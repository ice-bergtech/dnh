package dnh

func NewDNSEntry(name, dnsType string, ttl int) *DNSEntry {
	return &DNSEntry{
		Name: name,
		Type: dnsType,
		TTL:  ttl,
	}
}

func ProcessDnsEntry() {

}
