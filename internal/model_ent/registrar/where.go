// Code generated by ent, DO NOT EDIT.

package registrar

import (
	"time"

	"entgo.io/ent/dialect/sql"
	"github.com/ice-bergtech/dnh/src/internal/model_ent/predicate"
)

// ID filters vertices based on their ID field.
func ID(id int) predicate.Registrar {
	return predicate.Registrar(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int) predicate.Registrar {
	return predicate.Registrar(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int) predicate.Registrar {
	return predicate.Registrar(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int) predicate.Registrar {
	return predicate.Registrar(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...int) predicate.Registrar {
	return predicate.Registrar(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id int) predicate.Registrar {
	return predicate.Registrar(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int) predicate.Registrar {
	return predicate.Registrar(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int) predicate.Registrar {
	return predicate.Registrar(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int) predicate.Registrar {
	return predicate.Registrar(sql.FieldLTE(FieldID, id))
}

// Name applies equality check predicate on the "name" field. It's identical to NameEQ.
func Name(v string) predicate.Registrar {
	return predicate.Registrar(sql.FieldEQ(FieldName, v))
}

// URL applies equality check predicate on the "url" field. It's identical to URLEQ.
func URL(v string) predicate.Registrar {
	return predicate.Registrar(sql.FieldEQ(FieldURL, v))
}

// CountryCode applies equality check predicate on the "country_code" field. It's identical to CountryCodeEQ.
func CountryCode(v string) predicate.Registrar {
	return predicate.Registrar(sql.FieldEQ(FieldCountryCode, v))
}

// Phone applies equality check predicate on the "phone" field. It's identical to PhoneEQ.
func Phone(v string) predicate.Registrar {
	return predicate.Registrar(sql.FieldEQ(FieldPhone, v))
}

// Fax applies equality check predicate on the "fax" field. It's identical to FaxEQ.
func Fax(v string) predicate.Registrar {
	return predicate.Registrar(sql.FieldEQ(FieldFax, v))
}

// Address applies equality check predicate on the "address" field. It's identical to AddressEQ.
func Address(v string) predicate.Registrar {
	return predicate.Registrar(sql.FieldEQ(FieldAddress, v))
}

// TimeFirst applies equality check predicate on the "time_first" field. It's identical to TimeFirstEQ.
func TimeFirst(v time.Time) predicate.Registrar {
	return predicate.Registrar(sql.FieldEQ(FieldTimeFirst, v))
}

// TimeLast applies equality check predicate on the "time_last" field. It's identical to TimeLastEQ.
func TimeLast(v time.Time) predicate.Registrar {
	return predicate.Registrar(sql.FieldEQ(FieldTimeLast, v))
}

// NameEQ applies the EQ predicate on the "name" field.
func NameEQ(v string) predicate.Registrar {
	return predicate.Registrar(sql.FieldEQ(FieldName, v))
}

// NameNEQ applies the NEQ predicate on the "name" field.
func NameNEQ(v string) predicate.Registrar {
	return predicate.Registrar(sql.FieldNEQ(FieldName, v))
}

// NameIn applies the In predicate on the "name" field.
func NameIn(vs ...string) predicate.Registrar {
	return predicate.Registrar(sql.FieldIn(FieldName, vs...))
}

// NameNotIn applies the NotIn predicate on the "name" field.
func NameNotIn(vs ...string) predicate.Registrar {
	return predicate.Registrar(sql.FieldNotIn(FieldName, vs...))
}

// NameGT applies the GT predicate on the "name" field.
func NameGT(v string) predicate.Registrar {
	return predicate.Registrar(sql.FieldGT(FieldName, v))
}

// NameGTE applies the GTE predicate on the "name" field.
func NameGTE(v string) predicate.Registrar {
	return predicate.Registrar(sql.FieldGTE(FieldName, v))
}

// NameLT applies the LT predicate on the "name" field.
func NameLT(v string) predicate.Registrar {
	return predicate.Registrar(sql.FieldLT(FieldName, v))
}

// NameLTE applies the LTE predicate on the "name" field.
func NameLTE(v string) predicate.Registrar {
	return predicate.Registrar(sql.FieldLTE(FieldName, v))
}

// NameContains applies the Contains predicate on the "name" field.
func NameContains(v string) predicate.Registrar {
	return predicate.Registrar(sql.FieldContains(FieldName, v))
}

// NameHasPrefix applies the HasPrefix predicate on the "name" field.
func NameHasPrefix(v string) predicate.Registrar {
	return predicate.Registrar(sql.FieldHasPrefix(FieldName, v))
}

// NameHasSuffix applies the HasSuffix predicate on the "name" field.
func NameHasSuffix(v string) predicate.Registrar {
	return predicate.Registrar(sql.FieldHasSuffix(FieldName, v))
}

// NameEqualFold applies the EqualFold predicate on the "name" field.
func NameEqualFold(v string) predicate.Registrar {
	return predicate.Registrar(sql.FieldEqualFold(FieldName, v))
}

// NameContainsFold applies the ContainsFold predicate on the "name" field.
func NameContainsFold(v string) predicate.Registrar {
	return predicate.Registrar(sql.FieldContainsFold(FieldName, v))
}

// URLEQ applies the EQ predicate on the "url" field.
func URLEQ(v string) predicate.Registrar {
	return predicate.Registrar(sql.FieldEQ(FieldURL, v))
}

// URLNEQ applies the NEQ predicate on the "url" field.
func URLNEQ(v string) predicate.Registrar {
	return predicate.Registrar(sql.FieldNEQ(FieldURL, v))
}

// URLIn applies the In predicate on the "url" field.
func URLIn(vs ...string) predicate.Registrar {
	return predicate.Registrar(sql.FieldIn(FieldURL, vs...))
}

// URLNotIn applies the NotIn predicate on the "url" field.
func URLNotIn(vs ...string) predicate.Registrar {
	return predicate.Registrar(sql.FieldNotIn(FieldURL, vs...))
}

// URLGT applies the GT predicate on the "url" field.
func URLGT(v string) predicate.Registrar {
	return predicate.Registrar(sql.FieldGT(FieldURL, v))
}

// URLGTE applies the GTE predicate on the "url" field.
func URLGTE(v string) predicate.Registrar {
	return predicate.Registrar(sql.FieldGTE(FieldURL, v))
}

// URLLT applies the LT predicate on the "url" field.
func URLLT(v string) predicate.Registrar {
	return predicate.Registrar(sql.FieldLT(FieldURL, v))
}

// URLLTE applies the LTE predicate on the "url" field.
func URLLTE(v string) predicate.Registrar {
	return predicate.Registrar(sql.FieldLTE(FieldURL, v))
}

// URLContains applies the Contains predicate on the "url" field.
func URLContains(v string) predicate.Registrar {
	return predicate.Registrar(sql.FieldContains(FieldURL, v))
}

// URLHasPrefix applies the HasPrefix predicate on the "url" field.
func URLHasPrefix(v string) predicate.Registrar {
	return predicate.Registrar(sql.FieldHasPrefix(FieldURL, v))
}

// URLHasSuffix applies the HasSuffix predicate on the "url" field.
func URLHasSuffix(v string) predicate.Registrar {
	return predicate.Registrar(sql.FieldHasSuffix(FieldURL, v))
}

// URLEqualFold applies the EqualFold predicate on the "url" field.
func URLEqualFold(v string) predicate.Registrar {
	return predicate.Registrar(sql.FieldEqualFold(FieldURL, v))
}

// URLContainsFold applies the ContainsFold predicate on the "url" field.
func URLContainsFold(v string) predicate.Registrar {
	return predicate.Registrar(sql.FieldContainsFold(FieldURL, v))
}

// CountryCodeEQ applies the EQ predicate on the "country_code" field.
func CountryCodeEQ(v string) predicate.Registrar {
	return predicate.Registrar(sql.FieldEQ(FieldCountryCode, v))
}

// CountryCodeNEQ applies the NEQ predicate on the "country_code" field.
func CountryCodeNEQ(v string) predicate.Registrar {
	return predicate.Registrar(sql.FieldNEQ(FieldCountryCode, v))
}

// CountryCodeIn applies the In predicate on the "country_code" field.
func CountryCodeIn(vs ...string) predicate.Registrar {
	return predicate.Registrar(sql.FieldIn(FieldCountryCode, vs...))
}

// CountryCodeNotIn applies the NotIn predicate on the "country_code" field.
func CountryCodeNotIn(vs ...string) predicate.Registrar {
	return predicate.Registrar(sql.FieldNotIn(FieldCountryCode, vs...))
}

// CountryCodeGT applies the GT predicate on the "country_code" field.
func CountryCodeGT(v string) predicate.Registrar {
	return predicate.Registrar(sql.FieldGT(FieldCountryCode, v))
}

// CountryCodeGTE applies the GTE predicate on the "country_code" field.
func CountryCodeGTE(v string) predicate.Registrar {
	return predicate.Registrar(sql.FieldGTE(FieldCountryCode, v))
}

// CountryCodeLT applies the LT predicate on the "country_code" field.
func CountryCodeLT(v string) predicate.Registrar {
	return predicate.Registrar(sql.FieldLT(FieldCountryCode, v))
}

// CountryCodeLTE applies the LTE predicate on the "country_code" field.
func CountryCodeLTE(v string) predicate.Registrar {
	return predicate.Registrar(sql.FieldLTE(FieldCountryCode, v))
}

// CountryCodeContains applies the Contains predicate on the "country_code" field.
func CountryCodeContains(v string) predicate.Registrar {
	return predicate.Registrar(sql.FieldContains(FieldCountryCode, v))
}

// CountryCodeHasPrefix applies the HasPrefix predicate on the "country_code" field.
func CountryCodeHasPrefix(v string) predicate.Registrar {
	return predicate.Registrar(sql.FieldHasPrefix(FieldCountryCode, v))
}

// CountryCodeHasSuffix applies the HasSuffix predicate on the "country_code" field.
func CountryCodeHasSuffix(v string) predicate.Registrar {
	return predicate.Registrar(sql.FieldHasSuffix(FieldCountryCode, v))
}

// CountryCodeEqualFold applies the EqualFold predicate on the "country_code" field.
func CountryCodeEqualFold(v string) predicate.Registrar {
	return predicate.Registrar(sql.FieldEqualFold(FieldCountryCode, v))
}

// CountryCodeContainsFold applies the ContainsFold predicate on the "country_code" field.
func CountryCodeContainsFold(v string) predicate.Registrar {
	return predicate.Registrar(sql.FieldContainsFold(FieldCountryCode, v))
}

// PhoneEQ applies the EQ predicate on the "phone" field.
func PhoneEQ(v string) predicate.Registrar {
	return predicate.Registrar(sql.FieldEQ(FieldPhone, v))
}

// PhoneNEQ applies the NEQ predicate on the "phone" field.
func PhoneNEQ(v string) predicate.Registrar {
	return predicate.Registrar(sql.FieldNEQ(FieldPhone, v))
}

// PhoneIn applies the In predicate on the "phone" field.
func PhoneIn(vs ...string) predicate.Registrar {
	return predicate.Registrar(sql.FieldIn(FieldPhone, vs...))
}

// PhoneNotIn applies the NotIn predicate on the "phone" field.
func PhoneNotIn(vs ...string) predicate.Registrar {
	return predicate.Registrar(sql.FieldNotIn(FieldPhone, vs...))
}

// PhoneGT applies the GT predicate on the "phone" field.
func PhoneGT(v string) predicate.Registrar {
	return predicate.Registrar(sql.FieldGT(FieldPhone, v))
}

// PhoneGTE applies the GTE predicate on the "phone" field.
func PhoneGTE(v string) predicate.Registrar {
	return predicate.Registrar(sql.FieldGTE(FieldPhone, v))
}

// PhoneLT applies the LT predicate on the "phone" field.
func PhoneLT(v string) predicate.Registrar {
	return predicate.Registrar(sql.FieldLT(FieldPhone, v))
}

// PhoneLTE applies the LTE predicate on the "phone" field.
func PhoneLTE(v string) predicate.Registrar {
	return predicate.Registrar(sql.FieldLTE(FieldPhone, v))
}

// PhoneContains applies the Contains predicate on the "phone" field.
func PhoneContains(v string) predicate.Registrar {
	return predicate.Registrar(sql.FieldContains(FieldPhone, v))
}

// PhoneHasPrefix applies the HasPrefix predicate on the "phone" field.
func PhoneHasPrefix(v string) predicate.Registrar {
	return predicate.Registrar(sql.FieldHasPrefix(FieldPhone, v))
}

// PhoneHasSuffix applies the HasSuffix predicate on the "phone" field.
func PhoneHasSuffix(v string) predicate.Registrar {
	return predicate.Registrar(sql.FieldHasSuffix(FieldPhone, v))
}

// PhoneEqualFold applies the EqualFold predicate on the "phone" field.
func PhoneEqualFold(v string) predicate.Registrar {
	return predicate.Registrar(sql.FieldEqualFold(FieldPhone, v))
}

// PhoneContainsFold applies the ContainsFold predicate on the "phone" field.
func PhoneContainsFold(v string) predicate.Registrar {
	return predicate.Registrar(sql.FieldContainsFold(FieldPhone, v))
}

// FaxEQ applies the EQ predicate on the "fax" field.
func FaxEQ(v string) predicate.Registrar {
	return predicate.Registrar(sql.FieldEQ(FieldFax, v))
}

// FaxNEQ applies the NEQ predicate on the "fax" field.
func FaxNEQ(v string) predicate.Registrar {
	return predicate.Registrar(sql.FieldNEQ(FieldFax, v))
}

// FaxIn applies the In predicate on the "fax" field.
func FaxIn(vs ...string) predicate.Registrar {
	return predicate.Registrar(sql.FieldIn(FieldFax, vs...))
}

// FaxNotIn applies the NotIn predicate on the "fax" field.
func FaxNotIn(vs ...string) predicate.Registrar {
	return predicate.Registrar(sql.FieldNotIn(FieldFax, vs...))
}

// FaxGT applies the GT predicate on the "fax" field.
func FaxGT(v string) predicate.Registrar {
	return predicate.Registrar(sql.FieldGT(FieldFax, v))
}

// FaxGTE applies the GTE predicate on the "fax" field.
func FaxGTE(v string) predicate.Registrar {
	return predicate.Registrar(sql.FieldGTE(FieldFax, v))
}

// FaxLT applies the LT predicate on the "fax" field.
func FaxLT(v string) predicate.Registrar {
	return predicate.Registrar(sql.FieldLT(FieldFax, v))
}

// FaxLTE applies the LTE predicate on the "fax" field.
func FaxLTE(v string) predicate.Registrar {
	return predicate.Registrar(sql.FieldLTE(FieldFax, v))
}

// FaxContains applies the Contains predicate on the "fax" field.
func FaxContains(v string) predicate.Registrar {
	return predicate.Registrar(sql.FieldContains(FieldFax, v))
}

// FaxHasPrefix applies the HasPrefix predicate on the "fax" field.
func FaxHasPrefix(v string) predicate.Registrar {
	return predicate.Registrar(sql.FieldHasPrefix(FieldFax, v))
}

// FaxHasSuffix applies the HasSuffix predicate on the "fax" field.
func FaxHasSuffix(v string) predicate.Registrar {
	return predicate.Registrar(sql.FieldHasSuffix(FieldFax, v))
}

// FaxEqualFold applies the EqualFold predicate on the "fax" field.
func FaxEqualFold(v string) predicate.Registrar {
	return predicate.Registrar(sql.FieldEqualFold(FieldFax, v))
}

// FaxContainsFold applies the ContainsFold predicate on the "fax" field.
func FaxContainsFold(v string) predicate.Registrar {
	return predicate.Registrar(sql.FieldContainsFold(FieldFax, v))
}

// AddressEQ applies the EQ predicate on the "address" field.
func AddressEQ(v string) predicate.Registrar {
	return predicate.Registrar(sql.FieldEQ(FieldAddress, v))
}

// AddressNEQ applies the NEQ predicate on the "address" field.
func AddressNEQ(v string) predicate.Registrar {
	return predicate.Registrar(sql.FieldNEQ(FieldAddress, v))
}

// AddressIn applies the In predicate on the "address" field.
func AddressIn(vs ...string) predicate.Registrar {
	return predicate.Registrar(sql.FieldIn(FieldAddress, vs...))
}

// AddressNotIn applies the NotIn predicate on the "address" field.
func AddressNotIn(vs ...string) predicate.Registrar {
	return predicate.Registrar(sql.FieldNotIn(FieldAddress, vs...))
}

// AddressGT applies the GT predicate on the "address" field.
func AddressGT(v string) predicate.Registrar {
	return predicate.Registrar(sql.FieldGT(FieldAddress, v))
}

// AddressGTE applies the GTE predicate on the "address" field.
func AddressGTE(v string) predicate.Registrar {
	return predicate.Registrar(sql.FieldGTE(FieldAddress, v))
}

// AddressLT applies the LT predicate on the "address" field.
func AddressLT(v string) predicate.Registrar {
	return predicate.Registrar(sql.FieldLT(FieldAddress, v))
}

// AddressLTE applies the LTE predicate on the "address" field.
func AddressLTE(v string) predicate.Registrar {
	return predicate.Registrar(sql.FieldLTE(FieldAddress, v))
}

// AddressContains applies the Contains predicate on the "address" field.
func AddressContains(v string) predicate.Registrar {
	return predicate.Registrar(sql.FieldContains(FieldAddress, v))
}

// AddressHasPrefix applies the HasPrefix predicate on the "address" field.
func AddressHasPrefix(v string) predicate.Registrar {
	return predicate.Registrar(sql.FieldHasPrefix(FieldAddress, v))
}

// AddressHasSuffix applies the HasSuffix predicate on the "address" field.
func AddressHasSuffix(v string) predicate.Registrar {
	return predicate.Registrar(sql.FieldHasSuffix(FieldAddress, v))
}

// AddressEqualFold applies the EqualFold predicate on the "address" field.
func AddressEqualFold(v string) predicate.Registrar {
	return predicate.Registrar(sql.FieldEqualFold(FieldAddress, v))
}

// AddressContainsFold applies the ContainsFold predicate on the "address" field.
func AddressContainsFold(v string) predicate.Registrar {
	return predicate.Registrar(sql.FieldContainsFold(FieldAddress, v))
}

// TimeFirstEQ applies the EQ predicate on the "time_first" field.
func TimeFirstEQ(v time.Time) predicate.Registrar {
	return predicate.Registrar(sql.FieldEQ(FieldTimeFirst, v))
}

// TimeFirstNEQ applies the NEQ predicate on the "time_first" field.
func TimeFirstNEQ(v time.Time) predicate.Registrar {
	return predicate.Registrar(sql.FieldNEQ(FieldTimeFirst, v))
}

// TimeFirstIn applies the In predicate on the "time_first" field.
func TimeFirstIn(vs ...time.Time) predicate.Registrar {
	return predicate.Registrar(sql.FieldIn(FieldTimeFirst, vs...))
}

// TimeFirstNotIn applies the NotIn predicate on the "time_first" field.
func TimeFirstNotIn(vs ...time.Time) predicate.Registrar {
	return predicate.Registrar(sql.FieldNotIn(FieldTimeFirst, vs...))
}

// TimeFirstGT applies the GT predicate on the "time_first" field.
func TimeFirstGT(v time.Time) predicate.Registrar {
	return predicate.Registrar(sql.FieldGT(FieldTimeFirst, v))
}

// TimeFirstGTE applies the GTE predicate on the "time_first" field.
func TimeFirstGTE(v time.Time) predicate.Registrar {
	return predicate.Registrar(sql.FieldGTE(FieldTimeFirst, v))
}

// TimeFirstLT applies the LT predicate on the "time_first" field.
func TimeFirstLT(v time.Time) predicate.Registrar {
	return predicate.Registrar(sql.FieldLT(FieldTimeFirst, v))
}

// TimeFirstLTE applies the LTE predicate on the "time_first" field.
func TimeFirstLTE(v time.Time) predicate.Registrar {
	return predicate.Registrar(sql.FieldLTE(FieldTimeFirst, v))
}

// TimeLastEQ applies the EQ predicate on the "time_last" field.
func TimeLastEQ(v time.Time) predicate.Registrar {
	return predicate.Registrar(sql.FieldEQ(FieldTimeLast, v))
}

// TimeLastNEQ applies the NEQ predicate on the "time_last" field.
func TimeLastNEQ(v time.Time) predicate.Registrar {
	return predicate.Registrar(sql.FieldNEQ(FieldTimeLast, v))
}

// TimeLastIn applies the In predicate on the "time_last" field.
func TimeLastIn(vs ...time.Time) predicate.Registrar {
	return predicate.Registrar(sql.FieldIn(FieldTimeLast, vs...))
}

// TimeLastNotIn applies the NotIn predicate on the "time_last" field.
func TimeLastNotIn(vs ...time.Time) predicate.Registrar {
	return predicate.Registrar(sql.FieldNotIn(FieldTimeLast, vs...))
}

// TimeLastGT applies the GT predicate on the "time_last" field.
func TimeLastGT(v time.Time) predicate.Registrar {
	return predicate.Registrar(sql.FieldGT(FieldTimeLast, v))
}

// TimeLastGTE applies the GTE predicate on the "time_last" field.
func TimeLastGTE(v time.Time) predicate.Registrar {
	return predicate.Registrar(sql.FieldGTE(FieldTimeLast, v))
}

// TimeLastLT applies the LT predicate on the "time_last" field.
func TimeLastLT(v time.Time) predicate.Registrar {
	return predicate.Registrar(sql.FieldLT(FieldTimeLast, v))
}

// TimeLastLTE applies the LTE predicate on the "time_last" field.
func TimeLastLTE(v time.Time) predicate.Registrar {
	return predicate.Registrar(sql.FieldLTE(FieldTimeLast, v))
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.Registrar) predicate.Registrar {
	return predicate.Registrar(sql.AndPredicates(predicates...))
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.Registrar) predicate.Registrar {
	return predicate.Registrar(sql.OrPredicates(predicates...))
}

// Not applies the not operator on the given predicate.
func Not(p predicate.Registrar) predicate.Registrar {
	return predicate.Registrar(sql.NotPredicates(p))
}