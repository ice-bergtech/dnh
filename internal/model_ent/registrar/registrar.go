// Code generated by ent, DO NOT EDIT.

package registrar

import (
	"entgo.io/ent/dialect/sql"
)

const (
	// Label holds the string label denoting the registrar type in the database.
	Label = "registrar"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldName holds the string denoting the name field in the database.
	FieldName = "name"
	// FieldURL holds the string denoting the url field in the database.
	FieldURL = "url"
	// FieldCountryCode holds the string denoting the country_code field in the database.
	FieldCountryCode = "country_code"
	// FieldPhone holds the string denoting the phone field in the database.
	FieldPhone = "phone"
	// FieldFax holds the string denoting the fax field in the database.
	FieldFax = "fax"
	// FieldAddress holds the string denoting the address field in the database.
	FieldAddress = "address"
	// FieldTimeFirst holds the string denoting the time_first field in the database.
	FieldTimeFirst = "time_first"
	// FieldTimeLast holds the string denoting the time_last field in the database.
	FieldTimeLast = "time_last"
	// Table holds the table name of the registrar in the database.
	Table = "registrars"
)

// Columns holds all SQL columns for registrar fields.
var Columns = []string{
	FieldID,
	FieldName,
	FieldURL,
	FieldCountryCode,
	FieldPhone,
	FieldFax,
	FieldAddress,
	FieldTimeFirst,
	FieldTimeLast,
}

// ForeignKeys holds the SQL foreign-keys that are owned by the "registrars"
// table and are not defined as standalone fields in the schema.
var ForeignKeys = []string{
	"scan_registrar",
	"whois_registrar",
}

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
			return true
		}
	}
	for i := range ForeignKeys {
		if column == ForeignKeys[i] {
			return true
		}
	}
	return false
}

// OrderOption defines the ordering options for the Registrar queries.
type OrderOption func(*sql.Selector)

// ByID orders the results by the id field.
func ByID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldID, opts...).ToFunc()
}

// ByName orders the results by the name field.
func ByName(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldName, opts...).ToFunc()
}

// ByURL orders the results by the url field.
func ByURL(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldURL, opts...).ToFunc()
}

// ByCountryCode orders the results by the country_code field.
func ByCountryCode(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldCountryCode, opts...).ToFunc()
}

// ByPhone orders the results by the phone field.
func ByPhone(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldPhone, opts...).ToFunc()
}

// ByFax orders the results by the fax field.
func ByFax(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldFax, opts...).ToFunc()
}

// ByAddress orders the results by the address field.
func ByAddress(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldAddress, opts...).ToFunc()
}

// ByTimeFirst orders the results by the time_first field.
func ByTimeFirst(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldTimeFirst, opts...).ToFunc()
}

// ByTimeLast orders the results by the time_last field.
func ByTimeLast(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldTimeLast, opts...).ToFunc()
}
