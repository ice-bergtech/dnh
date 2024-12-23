// Code generated by ent, DO NOT EDIT.

package asninfo

import (
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"github.com/ice-bergtech/dnh/src/internal/model_ent/predicate"
)

// ID filters vertices based on their ID field.
func ID(id int) predicate.ASNInfo {
	return predicate.ASNInfo(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int) predicate.ASNInfo {
	return predicate.ASNInfo(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int) predicate.ASNInfo {
	return predicate.ASNInfo(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int) predicate.ASNInfo {
	return predicate.ASNInfo(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...int) predicate.ASNInfo {
	return predicate.ASNInfo(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id int) predicate.ASNInfo {
	return predicate.ASNInfo(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int) predicate.ASNInfo {
	return predicate.ASNInfo(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int) predicate.ASNInfo {
	return predicate.ASNInfo(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int) predicate.ASNInfo {
	return predicate.ASNInfo(sql.FieldLTE(FieldID, id))
}

// Asn applies equality check predicate on the "asn" field. It's identical to AsnEQ.
func Asn(v int) predicate.ASNInfo {
	return predicate.ASNInfo(sql.FieldEQ(FieldAsn, v))
}

// Country applies equality check predicate on the "country" field. It's identical to CountryEQ.
func Country(v string) predicate.ASNInfo {
	return predicate.ASNInfo(sql.FieldEQ(FieldCountry, v))
}

// Registry applies equality check predicate on the "registry" field. It's identical to RegistryEQ.
func Registry(v string) predicate.ASNInfo {
	return predicate.ASNInfo(sql.FieldEQ(FieldRegistry, v))
}

// AsnEQ applies the EQ predicate on the "asn" field.
func AsnEQ(v int) predicate.ASNInfo {
	return predicate.ASNInfo(sql.FieldEQ(FieldAsn, v))
}

// AsnNEQ applies the NEQ predicate on the "asn" field.
func AsnNEQ(v int) predicate.ASNInfo {
	return predicate.ASNInfo(sql.FieldNEQ(FieldAsn, v))
}

// AsnIn applies the In predicate on the "asn" field.
func AsnIn(vs ...int) predicate.ASNInfo {
	return predicate.ASNInfo(sql.FieldIn(FieldAsn, vs...))
}

// AsnNotIn applies the NotIn predicate on the "asn" field.
func AsnNotIn(vs ...int) predicate.ASNInfo {
	return predicate.ASNInfo(sql.FieldNotIn(FieldAsn, vs...))
}

// AsnGT applies the GT predicate on the "asn" field.
func AsnGT(v int) predicate.ASNInfo {
	return predicate.ASNInfo(sql.FieldGT(FieldAsn, v))
}

// AsnGTE applies the GTE predicate on the "asn" field.
func AsnGTE(v int) predicate.ASNInfo {
	return predicate.ASNInfo(sql.FieldGTE(FieldAsn, v))
}

// AsnLT applies the LT predicate on the "asn" field.
func AsnLT(v int) predicate.ASNInfo {
	return predicate.ASNInfo(sql.FieldLT(FieldAsn, v))
}

// AsnLTE applies the LTE predicate on the "asn" field.
func AsnLTE(v int) predicate.ASNInfo {
	return predicate.ASNInfo(sql.FieldLTE(FieldAsn, v))
}

// CountryEQ applies the EQ predicate on the "country" field.
func CountryEQ(v string) predicate.ASNInfo {
	return predicate.ASNInfo(sql.FieldEQ(FieldCountry, v))
}

// CountryNEQ applies the NEQ predicate on the "country" field.
func CountryNEQ(v string) predicate.ASNInfo {
	return predicate.ASNInfo(sql.FieldNEQ(FieldCountry, v))
}

// CountryIn applies the In predicate on the "country" field.
func CountryIn(vs ...string) predicate.ASNInfo {
	return predicate.ASNInfo(sql.FieldIn(FieldCountry, vs...))
}

// CountryNotIn applies the NotIn predicate on the "country" field.
func CountryNotIn(vs ...string) predicate.ASNInfo {
	return predicate.ASNInfo(sql.FieldNotIn(FieldCountry, vs...))
}

// CountryGT applies the GT predicate on the "country" field.
func CountryGT(v string) predicate.ASNInfo {
	return predicate.ASNInfo(sql.FieldGT(FieldCountry, v))
}

// CountryGTE applies the GTE predicate on the "country" field.
func CountryGTE(v string) predicate.ASNInfo {
	return predicate.ASNInfo(sql.FieldGTE(FieldCountry, v))
}

// CountryLT applies the LT predicate on the "country" field.
func CountryLT(v string) predicate.ASNInfo {
	return predicate.ASNInfo(sql.FieldLT(FieldCountry, v))
}

// CountryLTE applies the LTE predicate on the "country" field.
func CountryLTE(v string) predicate.ASNInfo {
	return predicate.ASNInfo(sql.FieldLTE(FieldCountry, v))
}

// CountryContains applies the Contains predicate on the "country" field.
func CountryContains(v string) predicate.ASNInfo {
	return predicate.ASNInfo(sql.FieldContains(FieldCountry, v))
}

// CountryHasPrefix applies the HasPrefix predicate on the "country" field.
func CountryHasPrefix(v string) predicate.ASNInfo {
	return predicate.ASNInfo(sql.FieldHasPrefix(FieldCountry, v))
}

// CountryHasSuffix applies the HasSuffix predicate on the "country" field.
func CountryHasSuffix(v string) predicate.ASNInfo {
	return predicate.ASNInfo(sql.FieldHasSuffix(FieldCountry, v))
}

// CountryEqualFold applies the EqualFold predicate on the "country" field.
func CountryEqualFold(v string) predicate.ASNInfo {
	return predicate.ASNInfo(sql.FieldEqualFold(FieldCountry, v))
}

// CountryContainsFold applies the ContainsFold predicate on the "country" field.
func CountryContainsFold(v string) predicate.ASNInfo {
	return predicate.ASNInfo(sql.FieldContainsFold(FieldCountry, v))
}

// RegistryEQ applies the EQ predicate on the "registry" field.
func RegistryEQ(v string) predicate.ASNInfo {
	return predicate.ASNInfo(sql.FieldEQ(FieldRegistry, v))
}

// RegistryNEQ applies the NEQ predicate on the "registry" field.
func RegistryNEQ(v string) predicate.ASNInfo {
	return predicate.ASNInfo(sql.FieldNEQ(FieldRegistry, v))
}

// RegistryIn applies the In predicate on the "registry" field.
func RegistryIn(vs ...string) predicate.ASNInfo {
	return predicate.ASNInfo(sql.FieldIn(FieldRegistry, vs...))
}

// RegistryNotIn applies the NotIn predicate on the "registry" field.
func RegistryNotIn(vs ...string) predicate.ASNInfo {
	return predicate.ASNInfo(sql.FieldNotIn(FieldRegistry, vs...))
}

// RegistryGT applies the GT predicate on the "registry" field.
func RegistryGT(v string) predicate.ASNInfo {
	return predicate.ASNInfo(sql.FieldGT(FieldRegistry, v))
}

// RegistryGTE applies the GTE predicate on the "registry" field.
func RegistryGTE(v string) predicate.ASNInfo {
	return predicate.ASNInfo(sql.FieldGTE(FieldRegistry, v))
}

// RegistryLT applies the LT predicate on the "registry" field.
func RegistryLT(v string) predicate.ASNInfo {
	return predicate.ASNInfo(sql.FieldLT(FieldRegistry, v))
}

// RegistryLTE applies the LTE predicate on the "registry" field.
func RegistryLTE(v string) predicate.ASNInfo {
	return predicate.ASNInfo(sql.FieldLTE(FieldRegistry, v))
}

// RegistryContains applies the Contains predicate on the "registry" field.
func RegistryContains(v string) predicate.ASNInfo {
	return predicate.ASNInfo(sql.FieldContains(FieldRegistry, v))
}

// RegistryHasPrefix applies the HasPrefix predicate on the "registry" field.
func RegistryHasPrefix(v string) predicate.ASNInfo {
	return predicate.ASNInfo(sql.FieldHasPrefix(FieldRegistry, v))
}

// RegistryHasSuffix applies the HasSuffix predicate on the "registry" field.
func RegistryHasSuffix(v string) predicate.ASNInfo {
	return predicate.ASNInfo(sql.FieldHasSuffix(FieldRegistry, v))
}

// RegistryEqualFold applies the EqualFold predicate on the "registry" field.
func RegistryEqualFold(v string) predicate.ASNInfo {
	return predicate.ASNInfo(sql.FieldEqualFold(FieldRegistry, v))
}

// RegistryContainsFold applies the ContainsFold predicate on the "registry" field.
func RegistryContainsFold(v string) predicate.ASNInfo {
	return predicate.ASNInfo(sql.FieldContainsFold(FieldRegistry, v))
}

// HasScan applies the HasEdge predicate on the "scan" edge.
func HasScan() predicate.ASNInfo {
	return predicate.ASNInfo(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.M2M, true, ScanTable, ScanPrimaryKey...),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasScanWith applies the HasEdge predicate on the "scan" edge with a given conditions (other predicates).
func HasScanWith(preds ...predicate.ScanJob) predicate.ASNInfo {
	return predicate.ASNInfo(func(s *sql.Selector) {
		step := newScanStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasIpaddress applies the HasEdge predicate on the "ipaddress" edge.
func HasIpaddress() predicate.ASNInfo {
	return predicate.ASNInfo(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.M2M, true, IpaddressTable, IpaddressPrimaryKey...),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasIpaddressWith applies the HasEdge predicate on the "ipaddress" edge with a given conditions (other predicates).
func HasIpaddressWith(preds ...predicate.IPAddress) predicate.ASNInfo {
	return predicate.ASNInfo(func(s *sql.Selector) {
		step := newIpaddressStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasRegistrar applies the HasEdge predicate on the "registrar" edge.
func HasRegistrar() predicate.ASNInfo {
	return predicate.ASNInfo(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.M2M, true, RegistrarTable, RegistrarPrimaryKey...),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasRegistrarWith applies the HasEdge predicate on the "registrar" edge with a given conditions (other predicates).
func HasRegistrarWith(preds ...predicate.Registrar) predicate.ASNInfo {
	return predicate.ASNInfo(func(s *sql.Selector) {
		step := newRegistrarStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasWhois applies the HasEdge predicate on the "whois" edge.
func HasWhois() predicate.ASNInfo {
	return predicate.ASNInfo(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.M2M, true, WhoisTable, WhoisPrimaryKey...),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasWhoisWith applies the HasEdge predicate on the "whois" edge with a given conditions (other predicates).
func HasWhoisWith(preds ...predicate.Whois) predicate.ASNInfo {
	return predicate.ASNInfo(func(s *sql.Selector) {
		step := newWhoisStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.ASNInfo) predicate.ASNInfo {
	return predicate.ASNInfo(sql.AndPredicates(predicates...))
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.ASNInfo) predicate.ASNInfo {
	return predicate.ASNInfo(sql.OrPredicates(predicates...))
}

// Not applies the not operator on the given predicate.
func Not(p predicate.ASNInfo) predicate.ASNInfo {
	return predicate.ASNInfo(sql.NotPredicates(p))
}
