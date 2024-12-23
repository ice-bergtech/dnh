// Code generated by ent, DO NOT EDIT.

package model_ent

import (
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/ice-bergtech/dnh/src/internal/model_ent/asninfo"
	"github.com/ice-bergtech/dnh/src/internal/model_ent/domain"
	"github.com/ice-bergtech/dnh/src/internal/model_ent/ipaddress"
	"github.com/ice-bergtech/dnh/src/internal/model_ent/predicate"
	"github.com/ice-bergtech/dnh/src/internal/model_ent/registrar"
	"github.com/ice-bergtech/dnh/src/internal/model_ent/scanjob"
	"github.com/ice-bergtech/dnh/src/internal/model_ent/whois"
)

// RegistrarUpdate is the builder for updating Registrar entities.
type RegistrarUpdate struct {
	config
	hooks    []Hook
	mutation *RegistrarMutation
}

// Where appends a list predicates to the RegistrarUpdate builder.
func (ru *RegistrarUpdate) Where(ps ...predicate.Registrar) *RegistrarUpdate {
	ru.mutation.Where(ps...)
	return ru
}

// SetName sets the "name" field.
func (ru *RegistrarUpdate) SetName(s string) *RegistrarUpdate {
	ru.mutation.SetName(s)
	return ru
}

// SetNillableName sets the "name" field if the given value is not nil.
func (ru *RegistrarUpdate) SetNillableName(s *string) *RegistrarUpdate {
	if s != nil {
		ru.SetName(*s)
	}
	return ru
}

// SetURL sets the "url" field.
func (ru *RegistrarUpdate) SetURL(s string) *RegistrarUpdate {
	ru.mutation.SetURL(s)
	return ru
}

// SetNillableURL sets the "url" field if the given value is not nil.
func (ru *RegistrarUpdate) SetNillableURL(s *string) *RegistrarUpdate {
	if s != nil {
		ru.SetURL(*s)
	}
	return ru
}

// SetCountryCode sets the "country_code" field.
func (ru *RegistrarUpdate) SetCountryCode(s string) *RegistrarUpdate {
	ru.mutation.SetCountryCode(s)
	return ru
}

// SetNillableCountryCode sets the "country_code" field if the given value is not nil.
func (ru *RegistrarUpdate) SetNillableCountryCode(s *string) *RegistrarUpdate {
	if s != nil {
		ru.SetCountryCode(*s)
	}
	return ru
}

// SetPhone sets the "phone" field.
func (ru *RegistrarUpdate) SetPhone(s string) *RegistrarUpdate {
	ru.mutation.SetPhone(s)
	return ru
}

// SetNillablePhone sets the "phone" field if the given value is not nil.
func (ru *RegistrarUpdate) SetNillablePhone(s *string) *RegistrarUpdate {
	if s != nil {
		ru.SetPhone(*s)
	}
	return ru
}

// SetFax sets the "fax" field.
func (ru *RegistrarUpdate) SetFax(s string) *RegistrarUpdate {
	ru.mutation.SetFax(s)
	return ru
}

// SetNillableFax sets the "fax" field if the given value is not nil.
func (ru *RegistrarUpdate) SetNillableFax(s *string) *RegistrarUpdate {
	if s != nil {
		ru.SetFax(*s)
	}
	return ru
}

// SetAddress sets the "address" field.
func (ru *RegistrarUpdate) SetAddress(s string) *RegistrarUpdate {
	ru.mutation.SetAddress(s)
	return ru
}

// SetNillableAddress sets the "address" field if the given value is not nil.
func (ru *RegistrarUpdate) SetNillableAddress(s *string) *RegistrarUpdate {
	if s != nil {
		ru.SetAddress(*s)
	}
	return ru
}

// SetSource sets the "source" field.
func (ru *RegistrarUpdate) SetSource(s string) *RegistrarUpdate {
	ru.mutation.SetSource(s)
	return ru
}

// SetNillableSource sets the "source" field if the given value is not nil.
func (ru *RegistrarUpdate) SetNillableSource(s *string) *RegistrarUpdate {
	if s != nil {
		ru.SetSource(*s)
	}
	return ru
}

// SetTimeFirst sets the "time_first" field.
func (ru *RegistrarUpdate) SetTimeFirst(t time.Time) *RegistrarUpdate {
	ru.mutation.SetTimeFirst(t)
	return ru
}

// SetNillableTimeFirst sets the "time_first" field if the given value is not nil.
func (ru *RegistrarUpdate) SetNillableTimeFirst(t *time.Time) *RegistrarUpdate {
	if t != nil {
		ru.SetTimeFirst(*t)
	}
	return ru
}

// SetTimeLast sets the "time_last" field.
func (ru *RegistrarUpdate) SetTimeLast(t time.Time) *RegistrarUpdate {
	ru.mutation.SetTimeLast(t)
	return ru
}

// SetNillableTimeLast sets the "time_last" field if the given value is not nil.
func (ru *RegistrarUpdate) SetNillableTimeLast(t *time.Time) *RegistrarUpdate {
	if t != nil {
		ru.SetTimeLast(*t)
	}
	return ru
}

// AddIpaddresIDs adds the "ipaddress" edge to the IPAddress entity by IDs.
func (ru *RegistrarUpdate) AddIpaddresIDs(ids ...int) *RegistrarUpdate {
	ru.mutation.AddIpaddresIDs(ids...)
	return ru
}

// AddIpaddress adds the "ipaddress" edges to the IPAddress entity.
func (ru *RegistrarUpdate) AddIpaddress(i ...*IPAddress) *RegistrarUpdate {
	ids := make([]int, len(i))
	for j := range i {
		ids[j] = i[j].ID
	}
	return ru.AddIpaddresIDs(ids...)
}

// AddDomainIDs adds the "domain" edge to the Domain entity by IDs.
func (ru *RegistrarUpdate) AddDomainIDs(ids ...int) *RegistrarUpdate {
	ru.mutation.AddDomainIDs(ids...)
	return ru
}

// AddDomain adds the "domain" edges to the Domain entity.
func (ru *RegistrarUpdate) AddDomain(d ...*Domain) *RegistrarUpdate {
	ids := make([]int, len(d))
	for i := range d {
		ids[i] = d[i].ID
	}
	return ru.AddDomainIDs(ids...)
}

// AddAsninfoIDs adds the "asninfo" edge to the ASNInfo entity by IDs.
func (ru *RegistrarUpdate) AddAsninfoIDs(ids ...int) *RegistrarUpdate {
	ru.mutation.AddAsninfoIDs(ids...)
	return ru
}

// AddAsninfo adds the "asninfo" edges to the ASNInfo entity.
func (ru *RegistrarUpdate) AddAsninfo(a ...*ASNInfo) *RegistrarUpdate {
	ids := make([]int, len(a))
	for i := range a {
		ids[i] = a[i].ID
	}
	return ru.AddAsninfoIDs(ids...)
}

// AddScanIDs adds the "scan" edge to the ScanJob entity by IDs.
func (ru *RegistrarUpdate) AddScanIDs(ids ...int) *RegistrarUpdate {
	ru.mutation.AddScanIDs(ids...)
	return ru
}

// AddScan adds the "scan" edges to the ScanJob entity.
func (ru *RegistrarUpdate) AddScan(s ...*ScanJob) *RegistrarUpdate {
	ids := make([]int, len(s))
	for i := range s {
		ids[i] = s[i].ID
	}
	return ru.AddScanIDs(ids...)
}

// AddWhoiIDs adds the "whois" edge to the Whois entity by IDs.
func (ru *RegistrarUpdate) AddWhoiIDs(ids ...int) *RegistrarUpdate {
	ru.mutation.AddWhoiIDs(ids...)
	return ru
}

// AddWhois adds the "whois" edges to the Whois entity.
func (ru *RegistrarUpdate) AddWhois(w ...*Whois) *RegistrarUpdate {
	ids := make([]int, len(w))
	for i := range w {
		ids[i] = w[i].ID
	}
	return ru.AddWhoiIDs(ids...)
}

// Mutation returns the RegistrarMutation object of the builder.
func (ru *RegistrarUpdate) Mutation() *RegistrarMutation {
	return ru.mutation
}

// ClearIpaddress clears all "ipaddress" edges to the IPAddress entity.
func (ru *RegistrarUpdate) ClearIpaddress() *RegistrarUpdate {
	ru.mutation.ClearIpaddress()
	return ru
}

// RemoveIpaddresIDs removes the "ipaddress" edge to IPAddress entities by IDs.
func (ru *RegistrarUpdate) RemoveIpaddresIDs(ids ...int) *RegistrarUpdate {
	ru.mutation.RemoveIpaddresIDs(ids...)
	return ru
}

// RemoveIpaddress removes "ipaddress" edges to IPAddress entities.
func (ru *RegistrarUpdate) RemoveIpaddress(i ...*IPAddress) *RegistrarUpdate {
	ids := make([]int, len(i))
	for j := range i {
		ids[j] = i[j].ID
	}
	return ru.RemoveIpaddresIDs(ids...)
}

// ClearDomain clears all "domain" edges to the Domain entity.
func (ru *RegistrarUpdate) ClearDomain() *RegistrarUpdate {
	ru.mutation.ClearDomain()
	return ru
}

// RemoveDomainIDs removes the "domain" edge to Domain entities by IDs.
func (ru *RegistrarUpdate) RemoveDomainIDs(ids ...int) *RegistrarUpdate {
	ru.mutation.RemoveDomainIDs(ids...)
	return ru
}

// RemoveDomain removes "domain" edges to Domain entities.
func (ru *RegistrarUpdate) RemoveDomain(d ...*Domain) *RegistrarUpdate {
	ids := make([]int, len(d))
	for i := range d {
		ids[i] = d[i].ID
	}
	return ru.RemoveDomainIDs(ids...)
}

// ClearAsninfo clears all "asninfo" edges to the ASNInfo entity.
func (ru *RegistrarUpdate) ClearAsninfo() *RegistrarUpdate {
	ru.mutation.ClearAsninfo()
	return ru
}

// RemoveAsninfoIDs removes the "asninfo" edge to ASNInfo entities by IDs.
func (ru *RegistrarUpdate) RemoveAsninfoIDs(ids ...int) *RegistrarUpdate {
	ru.mutation.RemoveAsninfoIDs(ids...)
	return ru
}

// RemoveAsninfo removes "asninfo" edges to ASNInfo entities.
func (ru *RegistrarUpdate) RemoveAsninfo(a ...*ASNInfo) *RegistrarUpdate {
	ids := make([]int, len(a))
	for i := range a {
		ids[i] = a[i].ID
	}
	return ru.RemoveAsninfoIDs(ids...)
}

// ClearScan clears all "scan" edges to the ScanJob entity.
func (ru *RegistrarUpdate) ClearScan() *RegistrarUpdate {
	ru.mutation.ClearScan()
	return ru
}

// RemoveScanIDs removes the "scan" edge to ScanJob entities by IDs.
func (ru *RegistrarUpdate) RemoveScanIDs(ids ...int) *RegistrarUpdate {
	ru.mutation.RemoveScanIDs(ids...)
	return ru
}

// RemoveScan removes "scan" edges to ScanJob entities.
func (ru *RegistrarUpdate) RemoveScan(s ...*ScanJob) *RegistrarUpdate {
	ids := make([]int, len(s))
	for i := range s {
		ids[i] = s[i].ID
	}
	return ru.RemoveScanIDs(ids...)
}

// ClearWhois clears all "whois" edges to the Whois entity.
func (ru *RegistrarUpdate) ClearWhois() *RegistrarUpdate {
	ru.mutation.ClearWhois()
	return ru
}

// RemoveWhoiIDs removes the "whois" edge to Whois entities by IDs.
func (ru *RegistrarUpdate) RemoveWhoiIDs(ids ...int) *RegistrarUpdate {
	ru.mutation.RemoveWhoiIDs(ids...)
	return ru
}

// RemoveWhois removes "whois" edges to Whois entities.
func (ru *RegistrarUpdate) RemoveWhois(w ...*Whois) *RegistrarUpdate {
	ids := make([]int, len(w))
	for i := range w {
		ids[i] = w[i].ID
	}
	return ru.RemoveWhoiIDs(ids...)
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (ru *RegistrarUpdate) Save(ctx context.Context) (int, error) {
	return withHooks(ctx, ru.sqlSave, ru.mutation, ru.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (ru *RegistrarUpdate) SaveX(ctx context.Context) int {
	affected, err := ru.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (ru *RegistrarUpdate) Exec(ctx context.Context) error {
	_, err := ru.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ru *RegistrarUpdate) ExecX(ctx context.Context) {
	if err := ru.Exec(ctx); err != nil {
		panic(err)
	}
}

func (ru *RegistrarUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := sqlgraph.NewUpdateSpec(registrar.Table, registrar.Columns, sqlgraph.NewFieldSpec(registrar.FieldID, field.TypeInt))
	if ps := ru.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := ru.mutation.Name(); ok {
		_spec.SetField(registrar.FieldName, field.TypeString, value)
	}
	if value, ok := ru.mutation.URL(); ok {
		_spec.SetField(registrar.FieldURL, field.TypeString, value)
	}
	if value, ok := ru.mutation.CountryCode(); ok {
		_spec.SetField(registrar.FieldCountryCode, field.TypeString, value)
	}
	if value, ok := ru.mutation.Phone(); ok {
		_spec.SetField(registrar.FieldPhone, field.TypeString, value)
	}
	if value, ok := ru.mutation.Fax(); ok {
		_spec.SetField(registrar.FieldFax, field.TypeString, value)
	}
	if value, ok := ru.mutation.Address(); ok {
		_spec.SetField(registrar.FieldAddress, field.TypeString, value)
	}
	if value, ok := ru.mutation.Source(); ok {
		_spec.SetField(registrar.FieldSource, field.TypeString, value)
	}
	if value, ok := ru.mutation.TimeFirst(); ok {
		_spec.SetField(registrar.FieldTimeFirst, field.TypeTime, value)
	}
	if value, ok := ru.mutation.TimeLast(); ok {
		_spec.SetField(registrar.FieldTimeLast, field.TypeTime, value)
	}
	if ru.mutation.IpaddressCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   registrar.IpaddressTable,
			Columns: registrar.IpaddressPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(ipaddress.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := ru.mutation.RemovedIpaddressIDs(); len(nodes) > 0 && !ru.mutation.IpaddressCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   registrar.IpaddressTable,
			Columns: registrar.IpaddressPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(ipaddress.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := ru.mutation.IpaddressIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   registrar.IpaddressTable,
			Columns: registrar.IpaddressPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(ipaddress.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if ru.mutation.DomainCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   registrar.DomainTable,
			Columns: registrar.DomainPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(domain.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := ru.mutation.RemovedDomainIDs(); len(nodes) > 0 && !ru.mutation.DomainCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   registrar.DomainTable,
			Columns: registrar.DomainPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(domain.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := ru.mutation.DomainIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   registrar.DomainTable,
			Columns: registrar.DomainPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(domain.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if ru.mutation.AsninfoCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   registrar.AsninfoTable,
			Columns: registrar.AsninfoPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(asninfo.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := ru.mutation.RemovedAsninfoIDs(); len(nodes) > 0 && !ru.mutation.AsninfoCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   registrar.AsninfoTable,
			Columns: registrar.AsninfoPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(asninfo.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := ru.mutation.AsninfoIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   registrar.AsninfoTable,
			Columns: registrar.AsninfoPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(asninfo.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if ru.mutation.ScanCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   registrar.ScanTable,
			Columns: registrar.ScanPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(scanjob.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := ru.mutation.RemovedScanIDs(); len(nodes) > 0 && !ru.mutation.ScanCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   registrar.ScanTable,
			Columns: registrar.ScanPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(scanjob.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := ru.mutation.ScanIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   registrar.ScanTable,
			Columns: registrar.ScanPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(scanjob.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if ru.mutation.WhoisCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   registrar.WhoisTable,
			Columns: registrar.WhoisPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(whois.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := ru.mutation.RemovedWhoisIDs(); len(nodes) > 0 && !ru.mutation.WhoisCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   registrar.WhoisTable,
			Columns: registrar.WhoisPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(whois.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := ru.mutation.WhoisIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   registrar.WhoisTable,
			Columns: registrar.WhoisPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(whois.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, ru.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{registrar.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	ru.mutation.done = true
	return n, nil
}

// RegistrarUpdateOne is the builder for updating a single Registrar entity.
type RegistrarUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *RegistrarMutation
}

// SetName sets the "name" field.
func (ruo *RegistrarUpdateOne) SetName(s string) *RegistrarUpdateOne {
	ruo.mutation.SetName(s)
	return ruo
}

// SetNillableName sets the "name" field if the given value is not nil.
func (ruo *RegistrarUpdateOne) SetNillableName(s *string) *RegistrarUpdateOne {
	if s != nil {
		ruo.SetName(*s)
	}
	return ruo
}

// SetURL sets the "url" field.
func (ruo *RegistrarUpdateOne) SetURL(s string) *RegistrarUpdateOne {
	ruo.mutation.SetURL(s)
	return ruo
}

// SetNillableURL sets the "url" field if the given value is not nil.
func (ruo *RegistrarUpdateOne) SetNillableURL(s *string) *RegistrarUpdateOne {
	if s != nil {
		ruo.SetURL(*s)
	}
	return ruo
}

// SetCountryCode sets the "country_code" field.
func (ruo *RegistrarUpdateOne) SetCountryCode(s string) *RegistrarUpdateOne {
	ruo.mutation.SetCountryCode(s)
	return ruo
}

// SetNillableCountryCode sets the "country_code" field if the given value is not nil.
func (ruo *RegistrarUpdateOne) SetNillableCountryCode(s *string) *RegistrarUpdateOne {
	if s != nil {
		ruo.SetCountryCode(*s)
	}
	return ruo
}

// SetPhone sets the "phone" field.
func (ruo *RegistrarUpdateOne) SetPhone(s string) *RegistrarUpdateOne {
	ruo.mutation.SetPhone(s)
	return ruo
}

// SetNillablePhone sets the "phone" field if the given value is not nil.
func (ruo *RegistrarUpdateOne) SetNillablePhone(s *string) *RegistrarUpdateOne {
	if s != nil {
		ruo.SetPhone(*s)
	}
	return ruo
}

// SetFax sets the "fax" field.
func (ruo *RegistrarUpdateOne) SetFax(s string) *RegistrarUpdateOne {
	ruo.mutation.SetFax(s)
	return ruo
}

// SetNillableFax sets the "fax" field if the given value is not nil.
func (ruo *RegistrarUpdateOne) SetNillableFax(s *string) *RegistrarUpdateOne {
	if s != nil {
		ruo.SetFax(*s)
	}
	return ruo
}

// SetAddress sets the "address" field.
func (ruo *RegistrarUpdateOne) SetAddress(s string) *RegistrarUpdateOne {
	ruo.mutation.SetAddress(s)
	return ruo
}

// SetNillableAddress sets the "address" field if the given value is not nil.
func (ruo *RegistrarUpdateOne) SetNillableAddress(s *string) *RegistrarUpdateOne {
	if s != nil {
		ruo.SetAddress(*s)
	}
	return ruo
}

// SetSource sets the "source" field.
func (ruo *RegistrarUpdateOne) SetSource(s string) *RegistrarUpdateOne {
	ruo.mutation.SetSource(s)
	return ruo
}

// SetNillableSource sets the "source" field if the given value is not nil.
func (ruo *RegistrarUpdateOne) SetNillableSource(s *string) *RegistrarUpdateOne {
	if s != nil {
		ruo.SetSource(*s)
	}
	return ruo
}

// SetTimeFirst sets the "time_first" field.
func (ruo *RegistrarUpdateOne) SetTimeFirst(t time.Time) *RegistrarUpdateOne {
	ruo.mutation.SetTimeFirst(t)
	return ruo
}

// SetNillableTimeFirst sets the "time_first" field if the given value is not nil.
func (ruo *RegistrarUpdateOne) SetNillableTimeFirst(t *time.Time) *RegistrarUpdateOne {
	if t != nil {
		ruo.SetTimeFirst(*t)
	}
	return ruo
}

// SetTimeLast sets the "time_last" field.
func (ruo *RegistrarUpdateOne) SetTimeLast(t time.Time) *RegistrarUpdateOne {
	ruo.mutation.SetTimeLast(t)
	return ruo
}

// SetNillableTimeLast sets the "time_last" field if the given value is not nil.
func (ruo *RegistrarUpdateOne) SetNillableTimeLast(t *time.Time) *RegistrarUpdateOne {
	if t != nil {
		ruo.SetTimeLast(*t)
	}
	return ruo
}

// AddIpaddresIDs adds the "ipaddress" edge to the IPAddress entity by IDs.
func (ruo *RegistrarUpdateOne) AddIpaddresIDs(ids ...int) *RegistrarUpdateOne {
	ruo.mutation.AddIpaddresIDs(ids...)
	return ruo
}

// AddIpaddress adds the "ipaddress" edges to the IPAddress entity.
func (ruo *RegistrarUpdateOne) AddIpaddress(i ...*IPAddress) *RegistrarUpdateOne {
	ids := make([]int, len(i))
	for j := range i {
		ids[j] = i[j].ID
	}
	return ruo.AddIpaddresIDs(ids...)
}

// AddDomainIDs adds the "domain" edge to the Domain entity by IDs.
func (ruo *RegistrarUpdateOne) AddDomainIDs(ids ...int) *RegistrarUpdateOne {
	ruo.mutation.AddDomainIDs(ids...)
	return ruo
}

// AddDomain adds the "domain" edges to the Domain entity.
func (ruo *RegistrarUpdateOne) AddDomain(d ...*Domain) *RegistrarUpdateOne {
	ids := make([]int, len(d))
	for i := range d {
		ids[i] = d[i].ID
	}
	return ruo.AddDomainIDs(ids...)
}

// AddAsninfoIDs adds the "asninfo" edge to the ASNInfo entity by IDs.
func (ruo *RegistrarUpdateOne) AddAsninfoIDs(ids ...int) *RegistrarUpdateOne {
	ruo.mutation.AddAsninfoIDs(ids...)
	return ruo
}

// AddAsninfo adds the "asninfo" edges to the ASNInfo entity.
func (ruo *RegistrarUpdateOne) AddAsninfo(a ...*ASNInfo) *RegistrarUpdateOne {
	ids := make([]int, len(a))
	for i := range a {
		ids[i] = a[i].ID
	}
	return ruo.AddAsninfoIDs(ids...)
}

// AddScanIDs adds the "scan" edge to the ScanJob entity by IDs.
func (ruo *RegistrarUpdateOne) AddScanIDs(ids ...int) *RegistrarUpdateOne {
	ruo.mutation.AddScanIDs(ids...)
	return ruo
}

// AddScan adds the "scan" edges to the ScanJob entity.
func (ruo *RegistrarUpdateOne) AddScan(s ...*ScanJob) *RegistrarUpdateOne {
	ids := make([]int, len(s))
	for i := range s {
		ids[i] = s[i].ID
	}
	return ruo.AddScanIDs(ids...)
}

// AddWhoiIDs adds the "whois" edge to the Whois entity by IDs.
func (ruo *RegistrarUpdateOne) AddWhoiIDs(ids ...int) *RegistrarUpdateOne {
	ruo.mutation.AddWhoiIDs(ids...)
	return ruo
}

// AddWhois adds the "whois" edges to the Whois entity.
func (ruo *RegistrarUpdateOne) AddWhois(w ...*Whois) *RegistrarUpdateOne {
	ids := make([]int, len(w))
	for i := range w {
		ids[i] = w[i].ID
	}
	return ruo.AddWhoiIDs(ids...)
}

// Mutation returns the RegistrarMutation object of the builder.
func (ruo *RegistrarUpdateOne) Mutation() *RegistrarMutation {
	return ruo.mutation
}

// ClearIpaddress clears all "ipaddress" edges to the IPAddress entity.
func (ruo *RegistrarUpdateOne) ClearIpaddress() *RegistrarUpdateOne {
	ruo.mutation.ClearIpaddress()
	return ruo
}

// RemoveIpaddresIDs removes the "ipaddress" edge to IPAddress entities by IDs.
func (ruo *RegistrarUpdateOne) RemoveIpaddresIDs(ids ...int) *RegistrarUpdateOne {
	ruo.mutation.RemoveIpaddresIDs(ids...)
	return ruo
}

// RemoveIpaddress removes "ipaddress" edges to IPAddress entities.
func (ruo *RegistrarUpdateOne) RemoveIpaddress(i ...*IPAddress) *RegistrarUpdateOne {
	ids := make([]int, len(i))
	for j := range i {
		ids[j] = i[j].ID
	}
	return ruo.RemoveIpaddresIDs(ids...)
}

// ClearDomain clears all "domain" edges to the Domain entity.
func (ruo *RegistrarUpdateOne) ClearDomain() *RegistrarUpdateOne {
	ruo.mutation.ClearDomain()
	return ruo
}

// RemoveDomainIDs removes the "domain" edge to Domain entities by IDs.
func (ruo *RegistrarUpdateOne) RemoveDomainIDs(ids ...int) *RegistrarUpdateOne {
	ruo.mutation.RemoveDomainIDs(ids...)
	return ruo
}

// RemoveDomain removes "domain" edges to Domain entities.
func (ruo *RegistrarUpdateOne) RemoveDomain(d ...*Domain) *RegistrarUpdateOne {
	ids := make([]int, len(d))
	for i := range d {
		ids[i] = d[i].ID
	}
	return ruo.RemoveDomainIDs(ids...)
}

// ClearAsninfo clears all "asninfo" edges to the ASNInfo entity.
func (ruo *RegistrarUpdateOne) ClearAsninfo() *RegistrarUpdateOne {
	ruo.mutation.ClearAsninfo()
	return ruo
}

// RemoveAsninfoIDs removes the "asninfo" edge to ASNInfo entities by IDs.
func (ruo *RegistrarUpdateOne) RemoveAsninfoIDs(ids ...int) *RegistrarUpdateOne {
	ruo.mutation.RemoveAsninfoIDs(ids...)
	return ruo
}

// RemoveAsninfo removes "asninfo" edges to ASNInfo entities.
func (ruo *RegistrarUpdateOne) RemoveAsninfo(a ...*ASNInfo) *RegistrarUpdateOne {
	ids := make([]int, len(a))
	for i := range a {
		ids[i] = a[i].ID
	}
	return ruo.RemoveAsninfoIDs(ids...)
}

// ClearScan clears all "scan" edges to the ScanJob entity.
func (ruo *RegistrarUpdateOne) ClearScan() *RegistrarUpdateOne {
	ruo.mutation.ClearScan()
	return ruo
}

// RemoveScanIDs removes the "scan" edge to ScanJob entities by IDs.
func (ruo *RegistrarUpdateOne) RemoveScanIDs(ids ...int) *RegistrarUpdateOne {
	ruo.mutation.RemoveScanIDs(ids...)
	return ruo
}

// RemoveScan removes "scan" edges to ScanJob entities.
func (ruo *RegistrarUpdateOne) RemoveScan(s ...*ScanJob) *RegistrarUpdateOne {
	ids := make([]int, len(s))
	for i := range s {
		ids[i] = s[i].ID
	}
	return ruo.RemoveScanIDs(ids...)
}

// ClearWhois clears all "whois" edges to the Whois entity.
func (ruo *RegistrarUpdateOne) ClearWhois() *RegistrarUpdateOne {
	ruo.mutation.ClearWhois()
	return ruo
}

// RemoveWhoiIDs removes the "whois" edge to Whois entities by IDs.
func (ruo *RegistrarUpdateOne) RemoveWhoiIDs(ids ...int) *RegistrarUpdateOne {
	ruo.mutation.RemoveWhoiIDs(ids...)
	return ruo
}

// RemoveWhois removes "whois" edges to Whois entities.
func (ruo *RegistrarUpdateOne) RemoveWhois(w ...*Whois) *RegistrarUpdateOne {
	ids := make([]int, len(w))
	for i := range w {
		ids[i] = w[i].ID
	}
	return ruo.RemoveWhoiIDs(ids...)
}

// Where appends a list predicates to the RegistrarUpdate builder.
func (ruo *RegistrarUpdateOne) Where(ps ...predicate.Registrar) *RegistrarUpdateOne {
	ruo.mutation.Where(ps...)
	return ruo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (ruo *RegistrarUpdateOne) Select(field string, fields ...string) *RegistrarUpdateOne {
	ruo.fields = append([]string{field}, fields...)
	return ruo
}

// Save executes the query and returns the updated Registrar entity.
func (ruo *RegistrarUpdateOne) Save(ctx context.Context) (*Registrar, error) {
	return withHooks(ctx, ruo.sqlSave, ruo.mutation, ruo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (ruo *RegistrarUpdateOne) SaveX(ctx context.Context) *Registrar {
	node, err := ruo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (ruo *RegistrarUpdateOne) Exec(ctx context.Context) error {
	_, err := ruo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ruo *RegistrarUpdateOne) ExecX(ctx context.Context) {
	if err := ruo.Exec(ctx); err != nil {
		panic(err)
	}
}

func (ruo *RegistrarUpdateOne) sqlSave(ctx context.Context) (_node *Registrar, err error) {
	_spec := sqlgraph.NewUpdateSpec(registrar.Table, registrar.Columns, sqlgraph.NewFieldSpec(registrar.FieldID, field.TypeInt))
	id, ok := ruo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`model_ent: missing "Registrar.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := ruo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, registrar.FieldID)
		for _, f := range fields {
			if !registrar.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("model_ent: invalid field %q for query", f)}
			}
			if f != registrar.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := ruo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := ruo.mutation.Name(); ok {
		_spec.SetField(registrar.FieldName, field.TypeString, value)
	}
	if value, ok := ruo.mutation.URL(); ok {
		_spec.SetField(registrar.FieldURL, field.TypeString, value)
	}
	if value, ok := ruo.mutation.CountryCode(); ok {
		_spec.SetField(registrar.FieldCountryCode, field.TypeString, value)
	}
	if value, ok := ruo.mutation.Phone(); ok {
		_spec.SetField(registrar.FieldPhone, field.TypeString, value)
	}
	if value, ok := ruo.mutation.Fax(); ok {
		_spec.SetField(registrar.FieldFax, field.TypeString, value)
	}
	if value, ok := ruo.mutation.Address(); ok {
		_spec.SetField(registrar.FieldAddress, field.TypeString, value)
	}
	if value, ok := ruo.mutation.Source(); ok {
		_spec.SetField(registrar.FieldSource, field.TypeString, value)
	}
	if value, ok := ruo.mutation.TimeFirst(); ok {
		_spec.SetField(registrar.FieldTimeFirst, field.TypeTime, value)
	}
	if value, ok := ruo.mutation.TimeLast(); ok {
		_spec.SetField(registrar.FieldTimeLast, field.TypeTime, value)
	}
	if ruo.mutation.IpaddressCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   registrar.IpaddressTable,
			Columns: registrar.IpaddressPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(ipaddress.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := ruo.mutation.RemovedIpaddressIDs(); len(nodes) > 0 && !ruo.mutation.IpaddressCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   registrar.IpaddressTable,
			Columns: registrar.IpaddressPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(ipaddress.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := ruo.mutation.IpaddressIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   registrar.IpaddressTable,
			Columns: registrar.IpaddressPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(ipaddress.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if ruo.mutation.DomainCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   registrar.DomainTable,
			Columns: registrar.DomainPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(domain.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := ruo.mutation.RemovedDomainIDs(); len(nodes) > 0 && !ruo.mutation.DomainCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   registrar.DomainTable,
			Columns: registrar.DomainPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(domain.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := ruo.mutation.DomainIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   registrar.DomainTable,
			Columns: registrar.DomainPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(domain.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if ruo.mutation.AsninfoCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   registrar.AsninfoTable,
			Columns: registrar.AsninfoPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(asninfo.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := ruo.mutation.RemovedAsninfoIDs(); len(nodes) > 0 && !ruo.mutation.AsninfoCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   registrar.AsninfoTable,
			Columns: registrar.AsninfoPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(asninfo.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := ruo.mutation.AsninfoIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   registrar.AsninfoTable,
			Columns: registrar.AsninfoPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(asninfo.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if ruo.mutation.ScanCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   registrar.ScanTable,
			Columns: registrar.ScanPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(scanjob.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := ruo.mutation.RemovedScanIDs(); len(nodes) > 0 && !ruo.mutation.ScanCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   registrar.ScanTable,
			Columns: registrar.ScanPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(scanjob.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := ruo.mutation.ScanIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   registrar.ScanTable,
			Columns: registrar.ScanPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(scanjob.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if ruo.mutation.WhoisCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   registrar.WhoisTable,
			Columns: registrar.WhoisPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(whois.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := ruo.mutation.RemovedWhoisIDs(); len(nodes) > 0 && !ruo.mutation.WhoisCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   registrar.WhoisTable,
			Columns: registrar.WhoisPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(whois.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := ruo.mutation.WhoisIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   registrar.WhoisTable,
			Columns: registrar.WhoisPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(whois.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &Registrar{config: ruo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, ruo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{registrar.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	ruo.mutation.done = true
	return _node, nil
}
