// Code generated by ent, DO NOT EDIT.

package ipaddress

import (
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"github.com/ice-bergtech/dnh/src/internal/model_ent/predicate"
)

// ID filters vertices based on their ID field.
func ID(id int) predicate.IPAddress {
	return predicate.IPAddress(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int) predicate.IPAddress {
	return predicate.IPAddress(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int) predicate.IPAddress {
	return predicate.IPAddress(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int) predicate.IPAddress {
	return predicate.IPAddress(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...int) predicate.IPAddress {
	return predicate.IPAddress(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id int) predicate.IPAddress {
	return predicate.IPAddress(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int) predicate.IPAddress {
	return predicate.IPAddress(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int) predicate.IPAddress {
	return predicate.IPAddress(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int) predicate.IPAddress {
	return predicate.IPAddress(sql.FieldLTE(FieldID, id))
}

// IP applies equality check predicate on the "ip" field. It's identical to IPEQ.
func IP(v string) predicate.IPAddress {
	return predicate.IPAddress(sql.FieldEQ(FieldIP, v))
}

// Mask applies equality check predicate on the "mask" field. It's identical to MaskEQ.
func Mask(v string) predicate.IPAddress {
	return predicate.IPAddress(sql.FieldEQ(FieldMask, v))
}

// IPEQ applies the EQ predicate on the "ip" field.
func IPEQ(v string) predicate.IPAddress {
	return predicate.IPAddress(sql.FieldEQ(FieldIP, v))
}

// IPNEQ applies the NEQ predicate on the "ip" field.
func IPNEQ(v string) predicate.IPAddress {
	return predicate.IPAddress(sql.FieldNEQ(FieldIP, v))
}

// IPIn applies the In predicate on the "ip" field.
func IPIn(vs ...string) predicate.IPAddress {
	return predicate.IPAddress(sql.FieldIn(FieldIP, vs...))
}

// IPNotIn applies the NotIn predicate on the "ip" field.
func IPNotIn(vs ...string) predicate.IPAddress {
	return predicate.IPAddress(sql.FieldNotIn(FieldIP, vs...))
}

// IPGT applies the GT predicate on the "ip" field.
func IPGT(v string) predicate.IPAddress {
	return predicate.IPAddress(sql.FieldGT(FieldIP, v))
}

// IPGTE applies the GTE predicate on the "ip" field.
func IPGTE(v string) predicate.IPAddress {
	return predicate.IPAddress(sql.FieldGTE(FieldIP, v))
}

// IPLT applies the LT predicate on the "ip" field.
func IPLT(v string) predicate.IPAddress {
	return predicate.IPAddress(sql.FieldLT(FieldIP, v))
}

// IPLTE applies the LTE predicate on the "ip" field.
func IPLTE(v string) predicate.IPAddress {
	return predicate.IPAddress(sql.FieldLTE(FieldIP, v))
}

// IPContains applies the Contains predicate on the "ip" field.
func IPContains(v string) predicate.IPAddress {
	return predicate.IPAddress(sql.FieldContains(FieldIP, v))
}

// IPHasPrefix applies the HasPrefix predicate on the "ip" field.
func IPHasPrefix(v string) predicate.IPAddress {
	return predicate.IPAddress(sql.FieldHasPrefix(FieldIP, v))
}

// IPHasSuffix applies the HasSuffix predicate on the "ip" field.
func IPHasSuffix(v string) predicate.IPAddress {
	return predicate.IPAddress(sql.FieldHasSuffix(FieldIP, v))
}

// IPEqualFold applies the EqualFold predicate on the "ip" field.
func IPEqualFold(v string) predicate.IPAddress {
	return predicate.IPAddress(sql.FieldEqualFold(FieldIP, v))
}

// IPContainsFold applies the ContainsFold predicate on the "ip" field.
func IPContainsFold(v string) predicate.IPAddress {
	return predicate.IPAddress(sql.FieldContainsFold(FieldIP, v))
}

// MaskEQ applies the EQ predicate on the "mask" field.
func MaskEQ(v string) predicate.IPAddress {
	return predicate.IPAddress(sql.FieldEQ(FieldMask, v))
}

// MaskNEQ applies the NEQ predicate on the "mask" field.
func MaskNEQ(v string) predicate.IPAddress {
	return predicate.IPAddress(sql.FieldNEQ(FieldMask, v))
}

// MaskIn applies the In predicate on the "mask" field.
func MaskIn(vs ...string) predicate.IPAddress {
	return predicate.IPAddress(sql.FieldIn(FieldMask, vs...))
}

// MaskNotIn applies the NotIn predicate on the "mask" field.
func MaskNotIn(vs ...string) predicate.IPAddress {
	return predicate.IPAddress(sql.FieldNotIn(FieldMask, vs...))
}

// MaskGT applies the GT predicate on the "mask" field.
func MaskGT(v string) predicate.IPAddress {
	return predicate.IPAddress(sql.FieldGT(FieldMask, v))
}

// MaskGTE applies the GTE predicate on the "mask" field.
func MaskGTE(v string) predicate.IPAddress {
	return predicate.IPAddress(sql.FieldGTE(FieldMask, v))
}

// MaskLT applies the LT predicate on the "mask" field.
func MaskLT(v string) predicate.IPAddress {
	return predicate.IPAddress(sql.FieldLT(FieldMask, v))
}

// MaskLTE applies the LTE predicate on the "mask" field.
func MaskLTE(v string) predicate.IPAddress {
	return predicate.IPAddress(sql.FieldLTE(FieldMask, v))
}

// MaskContains applies the Contains predicate on the "mask" field.
func MaskContains(v string) predicate.IPAddress {
	return predicate.IPAddress(sql.FieldContains(FieldMask, v))
}

// MaskHasPrefix applies the HasPrefix predicate on the "mask" field.
func MaskHasPrefix(v string) predicate.IPAddress {
	return predicate.IPAddress(sql.FieldHasPrefix(FieldMask, v))
}

// MaskHasSuffix applies the HasSuffix predicate on the "mask" field.
func MaskHasSuffix(v string) predicate.IPAddress {
	return predicate.IPAddress(sql.FieldHasSuffix(FieldMask, v))
}

// MaskEqualFold applies the EqualFold predicate on the "mask" field.
func MaskEqualFold(v string) predicate.IPAddress {
	return predicate.IPAddress(sql.FieldEqualFold(FieldMask, v))
}

// MaskContainsFold applies the ContainsFold predicate on the "mask" field.
func MaskContainsFold(v string) predicate.IPAddress {
	return predicate.IPAddress(sql.FieldContainsFold(FieldMask, v))
}

// HasAsninfo applies the HasEdge predicate on the "asninfo" edge.
func HasAsninfo() predicate.IPAddress {
	return predicate.IPAddress(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.M2M, false, AsninfoTable, AsninfoPrimaryKey...),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasAsninfoWith applies the HasEdge predicate on the "asninfo" edge with a given conditions (other predicates).
func HasAsninfoWith(preds ...predicate.ASNInfo) predicate.IPAddress {
	return predicate.IPAddress(func(s *sql.Selector) {
		step := newAsninfoStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasScan applies the HasEdge predicate on the "scan" edge.
func HasScan() predicate.IPAddress {
	return predicate.IPAddress(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.M2M, true, ScanTable, ScanPrimaryKey...),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasScanWith applies the HasEdge predicate on the "scan" edge with a given conditions (other predicates).
func HasScanWith(preds ...predicate.ScanJob) predicate.IPAddress {
	return predicate.IPAddress(func(s *sql.Selector) {
		step := newScanStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasDnsentry applies the HasEdge predicate on the "dnsentry" edge.
func HasDnsentry() predicate.IPAddress {
	return predicate.IPAddress(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.M2M, true, DnsentryTable, DnsentryPrimaryKey...),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasDnsentryWith applies the HasEdge predicate on the "dnsentry" edge with a given conditions (other predicates).
func HasDnsentryWith(preds ...predicate.DNSEntry) predicate.IPAddress {
	return predicate.IPAddress(func(s *sql.Selector) {
		step := newDnsentryStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasDomain applies the HasEdge predicate on the "domain" edge.
func HasDomain() predicate.IPAddress {
	return predicate.IPAddress(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.M2M, true, DomainTable, DomainPrimaryKey...),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasDomainWith applies the HasEdge predicate on the "domain" edge with a given conditions (other predicates).
func HasDomainWith(preds ...predicate.Domain) predicate.IPAddress {
	return predicate.IPAddress(func(s *sql.Selector) {
		step := newDomainStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasNameserver applies the HasEdge predicate on the "nameserver" edge.
func HasNameserver() predicate.IPAddress {
	return predicate.IPAddress(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.M2M, true, NameserverTable, NameserverPrimaryKey...),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasNameserverWith applies the HasEdge predicate on the "nameserver" edge with a given conditions (other predicates).
func HasNameserverWith(preds ...predicate.Nameserver) predicate.IPAddress {
	return predicate.IPAddress(func(s *sql.Selector) {
		step := newNameserverStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasRegistrar applies the HasEdge predicate on the "registrar" edge.
func HasRegistrar() predicate.IPAddress {
	return predicate.IPAddress(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.M2M, true, RegistrarTable, RegistrarPrimaryKey...),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasRegistrarWith applies the HasEdge predicate on the "registrar" edge with a given conditions (other predicates).
func HasRegistrarWith(preds ...predicate.Registrar) predicate.IPAddress {
	return predicate.IPAddress(func(s *sql.Selector) {
		step := newRegistrarStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasWhois applies the HasEdge predicate on the "whois" edge.
func HasWhois() predicate.IPAddress {
	return predicate.IPAddress(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.M2M, true, WhoisTable, WhoisPrimaryKey...),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasWhoisWith applies the HasEdge predicate on the "whois" edge with a given conditions (other predicates).
func HasWhoisWith(preds ...predicate.Whois) predicate.IPAddress {
	return predicate.IPAddress(func(s *sql.Selector) {
		step := newWhoisStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.IPAddress) predicate.IPAddress {
	return predicate.IPAddress(sql.AndPredicates(predicates...))
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.IPAddress) predicate.IPAddress {
	return predicate.IPAddress(sql.OrPredicates(predicates...))
}

// Not applies the not operator on the given predicate.
func Not(p predicate.IPAddress) predicate.IPAddress {
	return predicate.IPAddress(sql.NotPredicates(p))
}
