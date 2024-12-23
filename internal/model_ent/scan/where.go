// Code generated by ent, DO NOT EDIT.

package scan

import (
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"github.com/ice-bergtech/dnh/src/internal/model_ent/predicate"
)

// ID filters vertices based on their ID field.
func ID(id int) predicate.Scan {
	return predicate.Scan(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int) predicate.Scan {
	return predicate.Scan(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int) predicate.Scan {
	return predicate.Scan(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int) predicate.Scan {
	return predicate.Scan(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...int) predicate.Scan {
	return predicate.Scan(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id int) predicate.Scan {
	return predicate.Scan(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int) predicate.Scan {
	return predicate.Scan(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int) predicate.Scan {
	return predicate.Scan(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int) predicate.Scan {
	return predicate.Scan(sql.FieldLTE(FieldID, id))
}

// Scanid applies equality check predicate on the "scanid" field. It's identical to ScanidEQ.
func Scanid(v string) predicate.Scan {
	return predicate.Scan(sql.FieldEQ(FieldScanid, v))
}

// Input applies equality check predicate on the "input" field. It's identical to InputEQ.
func Input(v string) predicate.Scan {
	return predicate.Scan(sql.FieldEQ(FieldInput, v))
}

// Type applies equality check predicate on the "type" field. It's identical to TypeEQ.
func Type(v string) predicate.Scan {
	return predicate.Scan(sql.FieldEQ(FieldType, v))
}

// Timestamp applies equality check predicate on the "timestamp" field. It's identical to TimestampEQ.
func Timestamp(v time.Time) predicate.Scan {
	return predicate.Scan(sql.FieldEQ(FieldTimestamp, v))
}

// ScanidEQ applies the EQ predicate on the "scanid" field.
func ScanidEQ(v string) predicate.Scan {
	return predicate.Scan(sql.FieldEQ(FieldScanid, v))
}

// ScanidNEQ applies the NEQ predicate on the "scanid" field.
func ScanidNEQ(v string) predicate.Scan {
	return predicate.Scan(sql.FieldNEQ(FieldScanid, v))
}

// ScanidIn applies the In predicate on the "scanid" field.
func ScanidIn(vs ...string) predicate.Scan {
	return predicate.Scan(sql.FieldIn(FieldScanid, vs...))
}

// ScanidNotIn applies the NotIn predicate on the "scanid" field.
func ScanidNotIn(vs ...string) predicate.Scan {
	return predicate.Scan(sql.FieldNotIn(FieldScanid, vs...))
}

// ScanidGT applies the GT predicate on the "scanid" field.
func ScanidGT(v string) predicate.Scan {
	return predicate.Scan(sql.FieldGT(FieldScanid, v))
}

// ScanidGTE applies the GTE predicate on the "scanid" field.
func ScanidGTE(v string) predicate.Scan {
	return predicate.Scan(sql.FieldGTE(FieldScanid, v))
}

// ScanidLT applies the LT predicate on the "scanid" field.
func ScanidLT(v string) predicate.Scan {
	return predicate.Scan(sql.FieldLT(FieldScanid, v))
}

// ScanidLTE applies the LTE predicate on the "scanid" field.
func ScanidLTE(v string) predicate.Scan {
	return predicate.Scan(sql.FieldLTE(FieldScanid, v))
}

// ScanidContains applies the Contains predicate on the "scanid" field.
func ScanidContains(v string) predicate.Scan {
	return predicate.Scan(sql.FieldContains(FieldScanid, v))
}

// ScanidHasPrefix applies the HasPrefix predicate on the "scanid" field.
func ScanidHasPrefix(v string) predicate.Scan {
	return predicate.Scan(sql.FieldHasPrefix(FieldScanid, v))
}

// ScanidHasSuffix applies the HasSuffix predicate on the "scanid" field.
func ScanidHasSuffix(v string) predicate.Scan {
	return predicate.Scan(sql.FieldHasSuffix(FieldScanid, v))
}

// ScanidEqualFold applies the EqualFold predicate on the "scanid" field.
func ScanidEqualFold(v string) predicate.Scan {
	return predicate.Scan(sql.FieldEqualFold(FieldScanid, v))
}

// ScanidContainsFold applies the ContainsFold predicate on the "scanid" field.
func ScanidContainsFold(v string) predicate.Scan {
	return predicate.Scan(sql.FieldContainsFold(FieldScanid, v))
}

// InputEQ applies the EQ predicate on the "input" field.
func InputEQ(v string) predicate.Scan {
	return predicate.Scan(sql.FieldEQ(FieldInput, v))
}

// InputNEQ applies the NEQ predicate on the "input" field.
func InputNEQ(v string) predicate.Scan {
	return predicate.Scan(sql.FieldNEQ(FieldInput, v))
}

// InputIn applies the In predicate on the "input" field.
func InputIn(vs ...string) predicate.Scan {
	return predicate.Scan(sql.FieldIn(FieldInput, vs...))
}

// InputNotIn applies the NotIn predicate on the "input" field.
func InputNotIn(vs ...string) predicate.Scan {
	return predicate.Scan(sql.FieldNotIn(FieldInput, vs...))
}

// InputGT applies the GT predicate on the "input" field.
func InputGT(v string) predicate.Scan {
	return predicate.Scan(sql.FieldGT(FieldInput, v))
}

// InputGTE applies the GTE predicate on the "input" field.
func InputGTE(v string) predicate.Scan {
	return predicate.Scan(sql.FieldGTE(FieldInput, v))
}

// InputLT applies the LT predicate on the "input" field.
func InputLT(v string) predicate.Scan {
	return predicate.Scan(sql.FieldLT(FieldInput, v))
}

// InputLTE applies the LTE predicate on the "input" field.
func InputLTE(v string) predicate.Scan {
	return predicate.Scan(sql.FieldLTE(FieldInput, v))
}

// InputContains applies the Contains predicate on the "input" field.
func InputContains(v string) predicate.Scan {
	return predicate.Scan(sql.FieldContains(FieldInput, v))
}

// InputHasPrefix applies the HasPrefix predicate on the "input" field.
func InputHasPrefix(v string) predicate.Scan {
	return predicate.Scan(sql.FieldHasPrefix(FieldInput, v))
}

// InputHasSuffix applies the HasSuffix predicate on the "input" field.
func InputHasSuffix(v string) predicate.Scan {
	return predicate.Scan(sql.FieldHasSuffix(FieldInput, v))
}

// InputEqualFold applies the EqualFold predicate on the "input" field.
func InputEqualFold(v string) predicate.Scan {
	return predicate.Scan(sql.FieldEqualFold(FieldInput, v))
}

// InputContainsFold applies the ContainsFold predicate on the "input" field.
func InputContainsFold(v string) predicate.Scan {
	return predicate.Scan(sql.FieldContainsFold(FieldInput, v))
}

// TypeEQ applies the EQ predicate on the "type" field.
func TypeEQ(v string) predicate.Scan {
	return predicate.Scan(sql.FieldEQ(FieldType, v))
}

// TypeNEQ applies the NEQ predicate on the "type" field.
func TypeNEQ(v string) predicate.Scan {
	return predicate.Scan(sql.FieldNEQ(FieldType, v))
}

// TypeIn applies the In predicate on the "type" field.
func TypeIn(vs ...string) predicate.Scan {
	return predicate.Scan(sql.FieldIn(FieldType, vs...))
}

// TypeNotIn applies the NotIn predicate on the "type" field.
func TypeNotIn(vs ...string) predicate.Scan {
	return predicate.Scan(sql.FieldNotIn(FieldType, vs...))
}

// TypeGT applies the GT predicate on the "type" field.
func TypeGT(v string) predicate.Scan {
	return predicate.Scan(sql.FieldGT(FieldType, v))
}

// TypeGTE applies the GTE predicate on the "type" field.
func TypeGTE(v string) predicate.Scan {
	return predicate.Scan(sql.FieldGTE(FieldType, v))
}

// TypeLT applies the LT predicate on the "type" field.
func TypeLT(v string) predicate.Scan {
	return predicate.Scan(sql.FieldLT(FieldType, v))
}

// TypeLTE applies the LTE predicate on the "type" field.
func TypeLTE(v string) predicate.Scan {
	return predicate.Scan(sql.FieldLTE(FieldType, v))
}

// TypeContains applies the Contains predicate on the "type" field.
func TypeContains(v string) predicate.Scan {
	return predicate.Scan(sql.FieldContains(FieldType, v))
}

// TypeHasPrefix applies the HasPrefix predicate on the "type" field.
func TypeHasPrefix(v string) predicate.Scan {
	return predicate.Scan(sql.FieldHasPrefix(FieldType, v))
}

// TypeHasSuffix applies the HasSuffix predicate on the "type" field.
func TypeHasSuffix(v string) predicate.Scan {
	return predicate.Scan(sql.FieldHasSuffix(FieldType, v))
}

// TypeEqualFold applies the EqualFold predicate on the "type" field.
func TypeEqualFold(v string) predicate.Scan {
	return predicate.Scan(sql.FieldEqualFold(FieldType, v))
}

// TypeContainsFold applies the ContainsFold predicate on the "type" field.
func TypeContainsFold(v string) predicate.Scan {
	return predicate.Scan(sql.FieldContainsFold(FieldType, v))
}

// TimestampEQ applies the EQ predicate on the "timestamp" field.
func TimestampEQ(v time.Time) predicate.Scan {
	return predicate.Scan(sql.FieldEQ(FieldTimestamp, v))
}

// TimestampNEQ applies the NEQ predicate on the "timestamp" field.
func TimestampNEQ(v time.Time) predicate.Scan {
	return predicate.Scan(sql.FieldNEQ(FieldTimestamp, v))
}

// TimestampIn applies the In predicate on the "timestamp" field.
func TimestampIn(vs ...time.Time) predicate.Scan {
	return predicate.Scan(sql.FieldIn(FieldTimestamp, vs...))
}

// TimestampNotIn applies the NotIn predicate on the "timestamp" field.
func TimestampNotIn(vs ...time.Time) predicate.Scan {
	return predicate.Scan(sql.FieldNotIn(FieldTimestamp, vs...))
}

// TimestampGT applies the GT predicate on the "timestamp" field.
func TimestampGT(v time.Time) predicate.Scan {
	return predicate.Scan(sql.FieldGT(FieldTimestamp, v))
}

// TimestampGTE applies the GTE predicate on the "timestamp" field.
func TimestampGTE(v time.Time) predicate.Scan {
	return predicate.Scan(sql.FieldGTE(FieldTimestamp, v))
}

// TimestampLT applies the LT predicate on the "timestamp" field.
func TimestampLT(v time.Time) predicate.Scan {
	return predicate.Scan(sql.FieldLT(FieldTimestamp, v))
}

// TimestampLTE applies the LTE predicate on the "timestamp" field.
func TimestampLTE(v time.Time) predicate.Scan {
	return predicate.Scan(sql.FieldLTE(FieldTimestamp, v))
}

// HasScanjob applies the HasEdge predicate on the "scanjob" edge.
func HasScanjob() predicate.Scan {
	return predicate.Scan(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.M2M, false, ScanjobTable, ScanjobPrimaryKey...),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasScanjobWith applies the HasEdge predicate on the "scanjob" edge with a given conditions (other predicates).
func HasScanjobWith(preds ...predicate.ScanJob) predicate.Scan {
	return predicate.Scan(func(s *sql.Selector) {
		step := newScanjobStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.Scan) predicate.Scan {
	return predicate.Scan(sql.AndPredicates(predicates...))
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.Scan) predicate.Scan {
	return predicate.Scan(sql.OrPredicates(predicates...))
}

// Not applies the not operator on the given predicate.
func Not(p predicate.Scan) predicate.Scan {
	return predicate.Scan(sql.NotPredicates(p))
}
