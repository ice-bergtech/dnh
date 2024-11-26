// Code generated by ent, DO NOT EDIT.

package whois

import (
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"github.com/ice-bergtech/dnh/src/internal/model_ent/predicate"
)

// ID filters vertices based on their ID field.
func ID(id int) predicate.Whois {
	return predicate.Whois(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int) predicate.Whois {
	return predicate.Whois(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int) predicate.Whois {
	return predicate.Whois(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int) predicate.Whois {
	return predicate.Whois(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...int) predicate.Whois {
	return predicate.Whois(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id int) predicate.Whois {
	return predicate.Whois(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int) predicate.Whois {
	return predicate.Whois(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int) predicate.Whois {
	return predicate.Whois(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int) predicate.Whois {
	return predicate.Whois(sql.FieldLTE(FieldID, id))
}

// Query applies equality check predicate on the "query" field. It's identical to QueryEQ.
func Query(v string) predicate.Whois {
	return predicate.Whois(sql.FieldEQ(FieldQuery, v))
}

// Server applies equality check predicate on the "server" field. It's identical to ServerEQ.
func Server(v string) predicate.Whois {
	return predicate.Whois(sql.FieldEQ(FieldServer, v))
}

// Raw applies equality check predicate on the "raw" field. It's identical to RawEQ.
func Raw(v string) predicate.Whois {
	return predicate.Whois(sql.FieldEQ(FieldRaw, v))
}

// Country applies equality check predicate on the "country" field. It's identical to CountryEQ.
func Country(v string) predicate.Whois {
	return predicate.Whois(sql.FieldEQ(FieldCountry, v))
}

// Created applies equality check predicate on the "created" field. It's identical to CreatedEQ.
func Created(v time.Time) predicate.Whois {
	return predicate.Whois(sql.FieldEQ(FieldCreated, v))
}

// Updated applies equality check predicate on the "updated" field. It's identical to UpdatedEQ.
func Updated(v time.Time) predicate.Whois {
	return predicate.Whois(sql.FieldEQ(FieldUpdated, v))
}

// TimeFirst applies equality check predicate on the "time_first" field. It's identical to TimeFirstEQ.
func TimeFirst(v time.Time) predicate.Whois {
	return predicate.Whois(sql.FieldEQ(FieldTimeFirst, v))
}

// TimeLast applies equality check predicate on the "time_last" field. It's identical to TimeLastEQ.
func TimeLast(v time.Time) predicate.Whois {
	return predicate.Whois(sql.FieldEQ(FieldTimeLast, v))
}

// QueryEQ applies the EQ predicate on the "query" field.
func QueryEQ(v string) predicate.Whois {
	return predicate.Whois(sql.FieldEQ(FieldQuery, v))
}

// QueryNEQ applies the NEQ predicate on the "query" field.
func QueryNEQ(v string) predicate.Whois {
	return predicate.Whois(sql.FieldNEQ(FieldQuery, v))
}

// QueryIn applies the In predicate on the "query" field.
func QueryIn(vs ...string) predicate.Whois {
	return predicate.Whois(sql.FieldIn(FieldQuery, vs...))
}

// QueryNotIn applies the NotIn predicate on the "query" field.
func QueryNotIn(vs ...string) predicate.Whois {
	return predicate.Whois(sql.FieldNotIn(FieldQuery, vs...))
}

// QueryGT applies the GT predicate on the "query" field.
func QueryGT(v string) predicate.Whois {
	return predicate.Whois(sql.FieldGT(FieldQuery, v))
}

// QueryGTE applies the GTE predicate on the "query" field.
func QueryGTE(v string) predicate.Whois {
	return predicate.Whois(sql.FieldGTE(FieldQuery, v))
}

// QueryLT applies the LT predicate on the "query" field.
func QueryLT(v string) predicate.Whois {
	return predicate.Whois(sql.FieldLT(FieldQuery, v))
}

// QueryLTE applies the LTE predicate on the "query" field.
func QueryLTE(v string) predicate.Whois {
	return predicate.Whois(sql.FieldLTE(FieldQuery, v))
}

// QueryContains applies the Contains predicate on the "query" field.
func QueryContains(v string) predicate.Whois {
	return predicate.Whois(sql.FieldContains(FieldQuery, v))
}

// QueryHasPrefix applies the HasPrefix predicate on the "query" field.
func QueryHasPrefix(v string) predicate.Whois {
	return predicate.Whois(sql.FieldHasPrefix(FieldQuery, v))
}

// QueryHasSuffix applies the HasSuffix predicate on the "query" field.
func QueryHasSuffix(v string) predicate.Whois {
	return predicate.Whois(sql.FieldHasSuffix(FieldQuery, v))
}

// QueryEqualFold applies the EqualFold predicate on the "query" field.
func QueryEqualFold(v string) predicate.Whois {
	return predicate.Whois(sql.FieldEqualFold(FieldQuery, v))
}

// QueryContainsFold applies the ContainsFold predicate on the "query" field.
func QueryContainsFold(v string) predicate.Whois {
	return predicate.Whois(sql.FieldContainsFold(FieldQuery, v))
}

// ServerEQ applies the EQ predicate on the "server" field.
func ServerEQ(v string) predicate.Whois {
	return predicate.Whois(sql.FieldEQ(FieldServer, v))
}

// ServerNEQ applies the NEQ predicate on the "server" field.
func ServerNEQ(v string) predicate.Whois {
	return predicate.Whois(sql.FieldNEQ(FieldServer, v))
}

// ServerIn applies the In predicate on the "server" field.
func ServerIn(vs ...string) predicate.Whois {
	return predicate.Whois(sql.FieldIn(FieldServer, vs...))
}

// ServerNotIn applies the NotIn predicate on the "server" field.
func ServerNotIn(vs ...string) predicate.Whois {
	return predicate.Whois(sql.FieldNotIn(FieldServer, vs...))
}

// ServerGT applies the GT predicate on the "server" field.
func ServerGT(v string) predicate.Whois {
	return predicate.Whois(sql.FieldGT(FieldServer, v))
}

// ServerGTE applies the GTE predicate on the "server" field.
func ServerGTE(v string) predicate.Whois {
	return predicate.Whois(sql.FieldGTE(FieldServer, v))
}

// ServerLT applies the LT predicate on the "server" field.
func ServerLT(v string) predicate.Whois {
	return predicate.Whois(sql.FieldLT(FieldServer, v))
}

// ServerLTE applies the LTE predicate on the "server" field.
func ServerLTE(v string) predicate.Whois {
	return predicate.Whois(sql.FieldLTE(FieldServer, v))
}

// ServerContains applies the Contains predicate on the "server" field.
func ServerContains(v string) predicate.Whois {
	return predicate.Whois(sql.FieldContains(FieldServer, v))
}

// ServerHasPrefix applies the HasPrefix predicate on the "server" field.
func ServerHasPrefix(v string) predicate.Whois {
	return predicate.Whois(sql.FieldHasPrefix(FieldServer, v))
}

// ServerHasSuffix applies the HasSuffix predicate on the "server" field.
func ServerHasSuffix(v string) predicate.Whois {
	return predicate.Whois(sql.FieldHasSuffix(FieldServer, v))
}

// ServerEqualFold applies the EqualFold predicate on the "server" field.
func ServerEqualFold(v string) predicate.Whois {
	return predicate.Whois(sql.FieldEqualFold(FieldServer, v))
}

// ServerContainsFold applies the ContainsFold predicate on the "server" field.
func ServerContainsFold(v string) predicate.Whois {
	return predicate.Whois(sql.FieldContainsFold(FieldServer, v))
}

// RawEQ applies the EQ predicate on the "raw" field.
func RawEQ(v string) predicate.Whois {
	return predicate.Whois(sql.FieldEQ(FieldRaw, v))
}

// RawNEQ applies the NEQ predicate on the "raw" field.
func RawNEQ(v string) predicate.Whois {
	return predicate.Whois(sql.FieldNEQ(FieldRaw, v))
}

// RawIn applies the In predicate on the "raw" field.
func RawIn(vs ...string) predicate.Whois {
	return predicate.Whois(sql.FieldIn(FieldRaw, vs...))
}

// RawNotIn applies the NotIn predicate on the "raw" field.
func RawNotIn(vs ...string) predicate.Whois {
	return predicate.Whois(sql.FieldNotIn(FieldRaw, vs...))
}

// RawGT applies the GT predicate on the "raw" field.
func RawGT(v string) predicate.Whois {
	return predicate.Whois(sql.FieldGT(FieldRaw, v))
}

// RawGTE applies the GTE predicate on the "raw" field.
func RawGTE(v string) predicate.Whois {
	return predicate.Whois(sql.FieldGTE(FieldRaw, v))
}

// RawLT applies the LT predicate on the "raw" field.
func RawLT(v string) predicate.Whois {
	return predicate.Whois(sql.FieldLT(FieldRaw, v))
}

// RawLTE applies the LTE predicate on the "raw" field.
func RawLTE(v string) predicate.Whois {
	return predicate.Whois(sql.FieldLTE(FieldRaw, v))
}

// RawContains applies the Contains predicate on the "raw" field.
func RawContains(v string) predicate.Whois {
	return predicate.Whois(sql.FieldContains(FieldRaw, v))
}

// RawHasPrefix applies the HasPrefix predicate on the "raw" field.
func RawHasPrefix(v string) predicate.Whois {
	return predicate.Whois(sql.FieldHasPrefix(FieldRaw, v))
}

// RawHasSuffix applies the HasSuffix predicate on the "raw" field.
func RawHasSuffix(v string) predicate.Whois {
	return predicate.Whois(sql.FieldHasSuffix(FieldRaw, v))
}

// RawEqualFold applies the EqualFold predicate on the "raw" field.
func RawEqualFold(v string) predicate.Whois {
	return predicate.Whois(sql.FieldEqualFold(FieldRaw, v))
}

// RawContainsFold applies the ContainsFold predicate on the "raw" field.
func RawContainsFold(v string) predicate.Whois {
	return predicate.Whois(sql.FieldContainsFold(FieldRaw, v))
}

// CountryEQ applies the EQ predicate on the "country" field.
func CountryEQ(v string) predicate.Whois {
	return predicate.Whois(sql.FieldEQ(FieldCountry, v))
}

// CountryNEQ applies the NEQ predicate on the "country" field.
func CountryNEQ(v string) predicate.Whois {
	return predicate.Whois(sql.FieldNEQ(FieldCountry, v))
}

// CountryIn applies the In predicate on the "country" field.
func CountryIn(vs ...string) predicate.Whois {
	return predicate.Whois(sql.FieldIn(FieldCountry, vs...))
}

// CountryNotIn applies the NotIn predicate on the "country" field.
func CountryNotIn(vs ...string) predicate.Whois {
	return predicate.Whois(sql.FieldNotIn(FieldCountry, vs...))
}

// CountryGT applies the GT predicate on the "country" field.
func CountryGT(v string) predicate.Whois {
	return predicate.Whois(sql.FieldGT(FieldCountry, v))
}

// CountryGTE applies the GTE predicate on the "country" field.
func CountryGTE(v string) predicate.Whois {
	return predicate.Whois(sql.FieldGTE(FieldCountry, v))
}

// CountryLT applies the LT predicate on the "country" field.
func CountryLT(v string) predicate.Whois {
	return predicate.Whois(sql.FieldLT(FieldCountry, v))
}

// CountryLTE applies the LTE predicate on the "country" field.
func CountryLTE(v string) predicate.Whois {
	return predicate.Whois(sql.FieldLTE(FieldCountry, v))
}

// CountryContains applies the Contains predicate on the "country" field.
func CountryContains(v string) predicate.Whois {
	return predicate.Whois(sql.FieldContains(FieldCountry, v))
}

// CountryHasPrefix applies the HasPrefix predicate on the "country" field.
func CountryHasPrefix(v string) predicate.Whois {
	return predicate.Whois(sql.FieldHasPrefix(FieldCountry, v))
}

// CountryHasSuffix applies the HasSuffix predicate on the "country" field.
func CountryHasSuffix(v string) predicate.Whois {
	return predicate.Whois(sql.FieldHasSuffix(FieldCountry, v))
}

// CountryEqualFold applies the EqualFold predicate on the "country" field.
func CountryEqualFold(v string) predicate.Whois {
	return predicate.Whois(sql.FieldEqualFold(FieldCountry, v))
}

// CountryContainsFold applies the ContainsFold predicate on the "country" field.
func CountryContainsFold(v string) predicate.Whois {
	return predicate.Whois(sql.FieldContainsFold(FieldCountry, v))
}

// CreatedEQ applies the EQ predicate on the "created" field.
func CreatedEQ(v time.Time) predicate.Whois {
	return predicate.Whois(sql.FieldEQ(FieldCreated, v))
}

// CreatedNEQ applies the NEQ predicate on the "created" field.
func CreatedNEQ(v time.Time) predicate.Whois {
	return predicate.Whois(sql.FieldNEQ(FieldCreated, v))
}

// CreatedIn applies the In predicate on the "created" field.
func CreatedIn(vs ...time.Time) predicate.Whois {
	return predicate.Whois(sql.FieldIn(FieldCreated, vs...))
}

// CreatedNotIn applies the NotIn predicate on the "created" field.
func CreatedNotIn(vs ...time.Time) predicate.Whois {
	return predicate.Whois(sql.FieldNotIn(FieldCreated, vs...))
}

// CreatedGT applies the GT predicate on the "created" field.
func CreatedGT(v time.Time) predicate.Whois {
	return predicate.Whois(sql.FieldGT(FieldCreated, v))
}

// CreatedGTE applies the GTE predicate on the "created" field.
func CreatedGTE(v time.Time) predicate.Whois {
	return predicate.Whois(sql.FieldGTE(FieldCreated, v))
}

// CreatedLT applies the LT predicate on the "created" field.
func CreatedLT(v time.Time) predicate.Whois {
	return predicate.Whois(sql.FieldLT(FieldCreated, v))
}

// CreatedLTE applies the LTE predicate on the "created" field.
func CreatedLTE(v time.Time) predicate.Whois {
	return predicate.Whois(sql.FieldLTE(FieldCreated, v))
}

// UpdatedEQ applies the EQ predicate on the "updated" field.
func UpdatedEQ(v time.Time) predicate.Whois {
	return predicate.Whois(sql.FieldEQ(FieldUpdated, v))
}

// UpdatedNEQ applies the NEQ predicate on the "updated" field.
func UpdatedNEQ(v time.Time) predicate.Whois {
	return predicate.Whois(sql.FieldNEQ(FieldUpdated, v))
}

// UpdatedIn applies the In predicate on the "updated" field.
func UpdatedIn(vs ...time.Time) predicate.Whois {
	return predicate.Whois(sql.FieldIn(FieldUpdated, vs...))
}

// UpdatedNotIn applies the NotIn predicate on the "updated" field.
func UpdatedNotIn(vs ...time.Time) predicate.Whois {
	return predicate.Whois(sql.FieldNotIn(FieldUpdated, vs...))
}

// UpdatedGT applies the GT predicate on the "updated" field.
func UpdatedGT(v time.Time) predicate.Whois {
	return predicate.Whois(sql.FieldGT(FieldUpdated, v))
}

// UpdatedGTE applies the GTE predicate on the "updated" field.
func UpdatedGTE(v time.Time) predicate.Whois {
	return predicate.Whois(sql.FieldGTE(FieldUpdated, v))
}

// UpdatedLT applies the LT predicate on the "updated" field.
func UpdatedLT(v time.Time) predicate.Whois {
	return predicate.Whois(sql.FieldLT(FieldUpdated, v))
}

// UpdatedLTE applies the LTE predicate on the "updated" field.
func UpdatedLTE(v time.Time) predicate.Whois {
	return predicate.Whois(sql.FieldLTE(FieldUpdated, v))
}

// TimeFirstEQ applies the EQ predicate on the "time_first" field.
func TimeFirstEQ(v time.Time) predicate.Whois {
	return predicate.Whois(sql.FieldEQ(FieldTimeFirst, v))
}

// TimeFirstNEQ applies the NEQ predicate on the "time_first" field.
func TimeFirstNEQ(v time.Time) predicate.Whois {
	return predicate.Whois(sql.FieldNEQ(FieldTimeFirst, v))
}

// TimeFirstIn applies the In predicate on the "time_first" field.
func TimeFirstIn(vs ...time.Time) predicate.Whois {
	return predicate.Whois(sql.FieldIn(FieldTimeFirst, vs...))
}

// TimeFirstNotIn applies the NotIn predicate on the "time_first" field.
func TimeFirstNotIn(vs ...time.Time) predicate.Whois {
	return predicate.Whois(sql.FieldNotIn(FieldTimeFirst, vs...))
}

// TimeFirstGT applies the GT predicate on the "time_first" field.
func TimeFirstGT(v time.Time) predicate.Whois {
	return predicate.Whois(sql.FieldGT(FieldTimeFirst, v))
}

// TimeFirstGTE applies the GTE predicate on the "time_first" field.
func TimeFirstGTE(v time.Time) predicate.Whois {
	return predicate.Whois(sql.FieldGTE(FieldTimeFirst, v))
}

// TimeFirstLT applies the LT predicate on the "time_first" field.
func TimeFirstLT(v time.Time) predicate.Whois {
	return predicate.Whois(sql.FieldLT(FieldTimeFirst, v))
}

// TimeFirstLTE applies the LTE predicate on the "time_first" field.
func TimeFirstLTE(v time.Time) predicate.Whois {
	return predicate.Whois(sql.FieldLTE(FieldTimeFirst, v))
}

// TimeLastEQ applies the EQ predicate on the "time_last" field.
func TimeLastEQ(v time.Time) predicate.Whois {
	return predicate.Whois(sql.FieldEQ(FieldTimeLast, v))
}

// TimeLastNEQ applies the NEQ predicate on the "time_last" field.
func TimeLastNEQ(v time.Time) predicate.Whois {
	return predicate.Whois(sql.FieldNEQ(FieldTimeLast, v))
}

// TimeLastIn applies the In predicate on the "time_last" field.
func TimeLastIn(vs ...time.Time) predicate.Whois {
	return predicate.Whois(sql.FieldIn(FieldTimeLast, vs...))
}

// TimeLastNotIn applies the NotIn predicate on the "time_last" field.
func TimeLastNotIn(vs ...time.Time) predicate.Whois {
	return predicate.Whois(sql.FieldNotIn(FieldTimeLast, vs...))
}

// TimeLastGT applies the GT predicate on the "time_last" field.
func TimeLastGT(v time.Time) predicate.Whois {
	return predicate.Whois(sql.FieldGT(FieldTimeLast, v))
}

// TimeLastGTE applies the GTE predicate on the "time_last" field.
func TimeLastGTE(v time.Time) predicate.Whois {
	return predicate.Whois(sql.FieldGTE(FieldTimeLast, v))
}

// TimeLastLT applies the LT predicate on the "time_last" field.
func TimeLastLT(v time.Time) predicate.Whois {
	return predicate.Whois(sql.FieldLT(FieldTimeLast, v))
}

// TimeLastLTE applies the LTE predicate on the "time_last" field.
func TimeLastLTE(v time.Time) predicate.Whois {
	return predicate.Whois(sql.FieldLTE(FieldTimeLast, v))
}

// HasIprange applies the HasEdge predicate on the "iprange" edge.
func HasIprange() predicate.Whois {
	return predicate.Whois(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.M2M, false, IprangeTable, IprangePrimaryKey...),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasIprangeWith applies the HasEdge predicate on the "iprange" edge with a given conditions (other predicates).
func HasIprangeWith(preds ...predicate.IPAddress) predicate.Whois {
	return predicate.Whois(func(s *sql.Selector) {
		step := newIprangeStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasDomain applies the HasEdge predicate on the "domain" edge.
func HasDomain() predicate.Whois {
	return predicate.Whois(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.M2M, false, DomainTable, DomainPrimaryKey...),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasDomainWith applies the HasEdge predicate on the "domain" edge with a given conditions (other predicates).
func HasDomainWith(preds ...predicate.Domain) predicate.Whois {
	return predicate.Whois(func(s *sql.Selector) {
		step := newDomainStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasAsn applies the HasEdge predicate on the "asn" edge.
func HasAsn() predicate.Whois {
	return predicate.Whois(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.M2M, false, AsnTable, AsnPrimaryKey...),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasAsnWith applies the HasEdge predicate on the "asn" edge with a given conditions (other predicates).
func HasAsnWith(preds ...predicate.ASNInfo) predicate.Whois {
	return predicate.Whois(func(s *sql.Selector) {
		step := newAsnStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasRegistrar applies the HasEdge predicate on the "registrar" edge.
func HasRegistrar() predicate.Whois {
	return predicate.Whois(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.M2M, false, RegistrarTable, RegistrarPrimaryKey...),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasRegistrarWith applies the HasEdge predicate on the "registrar" edge with a given conditions (other predicates).
func HasRegistrarWith(preds ...predicate.Registrar) predicate.Whois {
	return predicate.Whois(func(s *sql.Selector) {
		step := newRegistrarStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasNameserver applies the HasEdge predicate on the "nameserver" edge.
func HasNameserver() predicate.Whois {
	return predicate.Whois(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.M2M, false, NameserverTable, NameserverPrimaryKey...),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasNameserverWith applies the HasEdge predicate on the "nameserver" edge with a given conditions (other predicates).
func HasNameserverWith(preds ...predicate.Nameserver) predicate.Whois {
	return predicate.Whois(func(s *sql.Selector) {
		step := newNameserverStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasScan applies the HasEdge predicate on the "scan" edge.
func HasScan() predicate.Whois {
	return predicate.Whois(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.M2M, true, ScanTable, ScanPrimaryKey...),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasScanWith applies the HasEdge predicate on the "scan" edge with a given conditions (other predicates).
func HasScanWith(preds ...predicate.Scan) predicate.Whois {
	return predicate.Whois(func(s *sql.Selector) {
		step := newScanStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.Whois) predicate.Whois {
	return predicate.Whois(sql.AndPredicates(predicates...))
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.Whois) predicate.Whois {
	return predicate.Whois(sql.OrPredicates(predicates...))
}

// Not applies the not operator on the given predicate.
func Not(p predicate.Whois) predicate.Whois {
	return predicate.Whois(sql.NotPredicates(p))
}
