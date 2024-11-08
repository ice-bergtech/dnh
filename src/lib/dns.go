package dnh

type DNSEntry struct {
    Name string
    Type string
    TTL  int
}

func NewDNSEntry(name, dnsType string, ttl int) *DNSEntry {
    return &DNSEntry{
        Name: name,
        Type: dnsType,
        TTL:  ttl,
    }
}

func ProcessDnsEntry(){
	
}