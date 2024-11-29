package jobs

import (
	"context"
	"fmt"
	"net"
	"time"

	"github.com/ice-bergtech/dnh/src/internal/model_ent"
)

type Record int

const (
	A Record = iota
	CNAME
	Reverse
	TXT
	MX
	NS
	SRV
)

func (r Record) String() string {
	switch r {
	case A:
		return "A"
	case CNAME:
		return "CNAME"
	case Reverse:
		return "Reverse"
	case TXT:
		return "TXT"
	case MX:
		return "MX"
	case NS:
		return "NS"
	case SRV:
		return "SRV"
	default:
		return "Unknown"
	}
}

// DNSLookup performs a DNS lookup for the given address using the specified resolver (includes port).
func DNSLookup(address string, resolver string, types []Record) ([]model_ent.DNSEntry, map[string][]interface{}, error) {
	r := &net.Resolver{
		PreferGo: true,
		Dial: func(ctx context.Context, network, address string) (net.Conn, error) {
			d := net.Dialer{
				Timeout: time.Millisecond * time.Duration(10000),
			}
			return d.DialContext(ctx, network, resolver)
		},
	}
	result := []model_ent.DNSEntry{}
	edges := map[string][]interface{}{}
	for _, t := range types {
		records := []string{}
		switch t {
		case A:
			records, _ = r.LookupHost(context.Background(), address)
			for _, record := range records {
				edges[model_ent.TypeIPAddress] = append(edges[model_ent.TypeIPAddress],
					model_ent.IPAddress{
						IP:   record,
						Mask: "32",
					})
			}
		case CNAME:
			cname, err := r.LookupCNAME(context.Background(), address)
			if err != nil {
				records = []string{cname}
				edges[model_ent.TypeDomain] = append(edges[model_ent.TypeDomain], model_ent.Domain{
					Name:      cname,
					TimeFirst: time.Now(),
					TimeLast:  time.Now(),
				})
			}
		case Reverse:
			records, _ = r.LookupAddr(context.Background(), address)
			for _, record := range records {
				edges[model_ent.TypeDomain] = append(edges[model_ent.TypeDomain],
					model_ent.Domain{
						Name:      record,
						TimeFirst: time.Now(),
						TimeLast:  time.Now(),
					})
			}
		case TXT:
			records, _ = r.LookupTXT(context.Background(), address)
		case MX:
			mxRecs, _ := r.LookupMX(context.Background(), address)
			for _, mx := range mxRecs {
				records = append(records, fmt.Sprintf("%d %s", mx.Pref, mx.Host))
			}
		case NS:
			nsRecs, _ := r.LookupNS(context.Background(), address)
			for _, ns := range nsRecs {
				records = append(records, ns.Host)
				edges[model_ent.TypeNameserver] = append(edges[model_ent.TypeNameserver], model_ent.Nameserver{
					Name:      ns.Host,
					TimeFirst: time.Now(),
					TimeLast:  time.Now(),
				})
			}
		// case SRV:
		// 	r.LookupSRV(context.Background(), address)
		default:
			return result, edges, fmt.Errorf("unknown DNS record type: %d", t)
		}
		for _, rc := range records {
			result = append(result, model_ent.DNSEntry{
				Name:  address,
				Type:  t.String(),
				Value: rc,
				TTL:   0,
			})
		}
	}
	return result, edges, nil
}
