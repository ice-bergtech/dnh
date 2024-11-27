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
	"github.com/ice-bergtech/dnh/src/internal/model_ent/dnsentry"
	"github.com/ice-bergtech/dnh/src/internal/model_ent/domain"
	"github.com/ice-bergtech/dnh/src/internal/model_ent/ipaddress"
	"github.com/ice-bergtech/dnh/src/internal/model_ent/nameserver"
	"github.com/ice-bergtech/dnh/src/internal/model_ent/predicate"
	"github.com/ice-bergtech/dnh/src/internal/model_ent/scanjob"
	"github.com/ice-bergtech/dnh/src/internal/model_ent/whois"
)

// NameserverUpdate is the builder for updating Nameserver entities.
type NameserverUpdate struct {
	config
	hooks    []Hook
	mutation *NameserverMutation
}

// Where appends a list predicates to the NameserverUpdate builder.
func (nu *NameserverUpdate) Where(ps ...predicate.Nameserver) *NameserverUpdate {
	nu.mutation.Where(ps...)
	return nu
}

// SetName sets the "name" field.
func (nu *NameserverUpdate) SetName(s string) *NameserverUpdate {
	nu.mutation.SetName(s)
	return nu
}

// SetNillableName sets the "name" field if the given value is not nil.
func (nu *NameserverUpdate) SetNillableName(s *string) *NameserverUpdate {
	if s != nil {
		nu.SetName(*s)
	}
	return nu
}

// SetTimeFirst sets the "time_first" field.
func (nu *NameserverUpdate) SetTimeFirst(t time.Time) *NameserverUpdate {
	nu.mutation.SetTimeFirst(t)
	return nu
}

// SetNillableTimeFirst sets the "time_first" field if the given value is not nil.
func (nu *NameserverUpdate) SetNillableTimeFirst(t *time.Time) *NameserverUpdate {
	if t != nil {
		nu.SetTimeFirst(*t)
	}
	return nu
}

// SetTimeLast sets the "time_last" field.
func (nu *NameserverUpdate) SetTimeLast(t time.Time) *NameserverUpdate {
	nu.mutation.SetTimeLast(t)
	return nu
}

// SetNillableTimeLast sets the "time_last" field if the given value is not nil.
func (nu *NameserverUpdate) SetNillableTimeLast(t *time.Time) *NameserverUpdate {
	if t != nil {
		nu.SetTimeLast(*t)
	}
	return nu
}

// AddIpaddresIDs adds the "ipaddress" edge to the IPAddress entity by IDs.
func (nu *NameserverUpdate) AddIpaddresIDs(ids ...int) *NameserverUpdate {
	nu.mutation.AddIpaddresIDs(ids...)
	return nu
}

// AddIpaddress adds the "ipaddress" edges to the IPAddress entity.
func (nu *NameserverUpdate) AddIpaddress(i ...*IPAddress) *NameserverUpdate {
	ids := make([]int, len(i))
	for j := range i {
		ids[j] = i[j].ID
	}
	return nu.AddIpaddresIDs(ids...)
}

// AddScanIDs adds the "scan" edge to the ScanJob entity by IDs.
func (nu *NameserverUpdate) AddScanIDs(ids ...int) *NameserverUpdate {
	nu.mutation.AddScanIDs(ids...)
	return nu
}

// AddScan adds the "scan" edges to the ScanJob entity.
func (nu *NameserverUpdate) AddScan(s ...*ScanJob) *NameserverUpdate {
	ids := make([]int, len(s))
	for i := range s {
		ids[i] = s[i].ID
	}
	return nu.AddScanIDs(ids...)
}

// AddDnsentryIDs adds the "dnsentry" edge to the DNSEntry entity by IDs.
func (nu *NameserverUpdate) AddDnsentryIDs(ids ...int) *NameserverUpdate {
	nu.mutation.AddDnsentryIDs(ids...)
	return nu
}

// AddDnsentry adds the "dnsentry" edges to the DNSEntry entity.
func (nu *NameserverUpdate) AddDnsentry(d ...*DNSEntry) *NameserverUpdate {
	ids := make([]int, len(d))
	for i := range d {
		ids[i] = d[i].ID
	}
	return nu.AddDnsentryIDs(ids...)
}

// AddDomainIDs adds the "domain" edge to the Domain entity by IDs.
func (nu *NameserverUpdate) AddDomainIDs(ids ...int) *NameserverUpdate {
	nu.mutation.AddDomainIDs(ids...)
	return nu
}

// AddDomain adds the "domain" edges to the Domain entity.
func (nu *NameserverUpdate) AddDomain(d ...*Domain) *NameserverUpdate {
	ids := make([]int, len(d))
	for i := range d {
		ids[i] = d[i].ID
	}
	return nu.AddDomainIDs(ids...)
}

// AddWhoiIDs adds the "whois" edge to the Whois entity by IDs.
func (nu *NameserverUpdate) AddWhoiIDs(ids ...int) *NameserverUpdate {
	nu.mutation.AddWhoiIDs(ids...)
	return nu
}

// AddWhois adds the "whois" edges to the Whois entity.
func (nu *NameserverUpdate) AddWhois(w ...*Whois) *NameserverUpdate {
	ids := make([]int, len(w))
	for i := range w {
		ids[i] = w[i].ID
	}
	return nu.AddWhoiIDs(ids...)
}

// Mutation returns the NameserverMutation object of the builder.
func (nu *NameserverUpdate) Mutation() *NameserverMutation {
	return nu.mutation
}

// ClearIpaddress clears all "ipaddress" edges to the IPAddress entity.
func (nu *NameserverUpdate) ClearIpaddress() *NameserverUpdate {
	nu.mutation.ClearIpaddress()
	return nu
}

// RemoveIpaddresIDs removes the "ipaddress" edge to IPAddress entities by IDs.
func (nu *NameserverUpdate) RemoveIpaddresIDs(ids ...int) *NameserverUpdate {
	nu.mutation.RemoveIpaddresIDs(ids...)
	return nu
}

// RemoveIpaddress removes "ipaddress" edges to IPAddress entities.
func (nu *NameserverUpdate) RemoveIpaddress(i ...*IPAddress) *NameserverUpdate {
	ids := make([]int, len(i))
	for j := range i {
		ids[j] = i[j].ID
	}
	return nu.RemoveIpaddresIDs(ids...)
}

// ClearScan clears all "scan" edges to the ScanJob entity.
func (nu *NameserverUpdate) ClearScan() *NameserverUpdate {
	nu.mutation.ClearScan()
	return nu
}

// RemoveScanIDs removes the "scan" edge to ScanJob entities by IDs.
func (nu *NameserverUpdate) RemoveScanIDs(ids ...int) *NameserverUpdate {
	nu.mutation.RemoveScanIDs(ids...)
	return nu
}

// RemoveScan removes "scan" edges to ScanJob entities.
func (nu *NameserverUpdate) RemoveScan(s ...*ScanJob) *NameserverUpdate {
	ids := make([]int, len(s))
	for i := range s {
		ids[i] = s[i].ID
	}
	return nu.RemoveScanIDs(ids...)
}

// ClearDnsentry clears all "dnsentry" edges to the DNSEntry entity.
func (nu *NameserverUpdate) ClearDnsentry() *NameserverUpdate {
	nu.mutation.ClearDnsentry()
	return nu
}

// RemoveDnsentryIDs removes the "dnsentry" edge to DNSEntry entities by IDs.
func (nu *NameserverUpdate) RemoveDnsentryIDs(ids ...int) *NameserverUpdate {
	nu.mutation.RemoveDnsentryIDs(ids...)
	return nu
}

// RemoveDnsentry removes "dnsentry" edges to DNSEntry entities.
func (nu *NameserverUpdate) RemoveDnsentry(d ...*DNSEntry) *NameserverUpdate {
	ids := make([]int, len(d))
	for i := range d {
		ids[i] = d[i].ID
	}
	return nu.RemoveDnsentryIDs(ids...)
}

// ClearDomain clears all "domain" edges to the Domain entity.
func (nu *NameserverUpdate) ClearDomain() *NameserverUpdate {
	nu.mutation.ClearDomain()
	return nu
}

// RemoveDomainIDs removes the "domain" edge to Domain entities by IDs.
func (nu *NameserverUpdate) RemoveDomainIDs(ids ...int) *NameserverUpdate {
	nu.mutation.RemoveDomainIDs(ids...)
	return nu
}

// RemoveDomain removes "domain" edges to Domain entities.
func (nu *NameserverUpdate) RemoveDomain(d ...*Domain) *NameserverUpdate {
	ids := make([]int, len(d))
	for i := range d {
		ids[i] = d[i].ID
	}
	return nu.RemoveDomainIDs(ids...)
}

// ClearWhois clears all "whois" edges to the Whois entity.
func (nu *NameserverUpdate) ClearWhois() *NameserverUpdate {
	nu.mutation.ClearWhois()
	return nu
}

// RemoveWhoiIDs removes the "whois" edge to Whois entities by IDs.
func (nu *NameserverUpdate) RemoveWhoiIDs(ids ...int) *NameserverUpdate {
	nu.mutation.RemoveWhoiIDs(ids...)
	return nu
}

// RemoveWhois removes "whois" edges to Whois entities.
func (nu *NameserverUpdate) RemoveWhois(w ...*Whois) *NameserverUpdate {
	ids := make([]int, len(w))
	for i := range w {
		ids[i] = w[i].ID
	}
	return nu.RemoveWhoiIDs(ids...)
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (nu *NameserverUpdate) Save(ctx context.Context) (int, error) {
	return withHooks(ctx, nu.sqlSave, nu.mutation, nu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (nu *NameserverUpdate) SaveX(ctx context.Context) int {
	affected, err := nu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (nu *NameserverUpdate) Exec(ctx context.Context) error {
	_, err := nu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (nu *NameserverUpdate) ExecX(ctx context.Context) {
	if err := nu.Exec(ctx); err != nil {
		panic(err)
	}
}

func (nu *NameserverUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := sqlgraph.NewUpdateSpec(nameserver.Table, nameserver.Columns, sqlgraph.NewFieldSpec(nameserver.FieldID, field.TypeInt))
	if ps := nu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := nu.mutation.Name(); ok {
		_spec.SetField(nameserver.FieldName, field.TypeString, value)
	}
	if value, ok := nu.mutation.TimeFirst(); ok {
		_spec.SetField(nameserver.FieldTimeFirst, field.TypeTime, value)
	}
	if value, ok := nu.mutation.TimeLast(); ok {
		_spec.SetField(nameserver.FieldTimeLast, field.TypeTime, value)
	}
	if nu.mutation.IpaddressCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   nameserver.IpaddressTable,
			Columns: nameserver.IpaddressPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(ipaddress.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := nu.mutation.RemovedIpaddressIDs(); len(nodes) > 0 && !nu.mutation.IpaddressCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   nameserver.IpaddressTable,
			Columns: nameserver.IpaddressPrimaryKey,
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
	if nodes := nu.mutation.IpaddressIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   nameserver.IpaddressTable,
			Columns: nameserver.IpaddressPrimaryKey,
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
	if nu.mutation.ScanCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   nameserver.ScanTable,
			Columns: nameserver.ScanPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(scanjob.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := nu.mutation.RemovedScanIDs(); len(nodes) > 0 && !nu.mutation.ScanCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   nameserver.ScanTable,
			Columns: nameserver.ScanPrimaryKey,
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
	if nodes := nu.mutation.ScanIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   nameserver.ScanTable,
			Columns: nameserver.ScanPrimaryKey,
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
	if nu.mutation.DnsentryCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   nameserver.DnsentryTable,
			Columns: nameserver.DnsentryPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(dnsentry.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := nu.mutation.RemovedDnsentryIDs(); len(nodes) > 0 && !nu.mutation.DnsentryCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   nameserver.DnsentryTable,
			Columns: nameserver.DnsentryPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(dnsentry.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := nu.mutation.DnsentryIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   nameserver.DnsentryTable,
			Columns: nameserver.DnsentryPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(dnsentry.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if nu.mutation.DomainCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   nameserver.DomainTable,
			Columns: nameserver.DomainPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(domain.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := nu.mutation.RemovedDomainIDs(); len(nodes) > 0 && !nu.mutation.DomainCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   nameserver.DomainTable,
			Columns: nameserver.DomainPrimaryKey,
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
	if nodes := nu.mutation.DomainIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   nameserver.DomainTable,
			Columns: nameserver.DomainPrimaryKey,
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
	if nu.mutation.WhoisCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   nameserver.WhoisTable,
			Columns: nameserver.WhoisPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(whois.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := nu.mutation.RemovedWhoisIDs(); len(nodes) > 0 && !nu.mutation.WhoisCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   nameserver.WhoisTable,
			Columns: nameserver.WhoisPrimaryKey,
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
	if nodes := nu.mutation.WhoisIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   nameserver.WhoisTable,
			Columns: nameserver.WhoisPrimaryKey,
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
	if n, err = sqlgraph.UpdateNodes(ctx, nu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{nameserver.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	nu.mutation.done = true
	return n, nil
}

// NameserverUpdateOne is the builder for updating a single Nameserver entity.
type NameserverUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *NameserverMutation
}

// SetName sets the "name" field.
func (nuo *NameserverUpdateOne) SetName(s string) *NameserverUpdateOne {
	nuo.mutation.SetName(s)
	return nuo
}

// SetNillableName sets the "name" field if the given value is not nil.
func (nuo *NameserverUpdateOne) SetNillableName(s *string) *NameserverUpdateOne {
	if s != nil {
		nuo.SetName(*s)
	}
	return nuo
}

// SetTimeFirst sets the "time_first" field.
func (nuo *NameserverUpdateOne) SetTimeFirst(t time.Time) *NameserverUpdateOne {
	nuo.mutation.SetTimeFirst(t)
	return nuo
}

// SetNillableTimeFirst sets the "time_first" field if the given value is not nil.
func (nuo *NameserverUpdateOne) SetNillableTimeFirst(t *time.Time) *NameserverUpdateOne {
	if t != nil {
		nuo.SetTimeFirst(*t)
	}
	return nuo
}

// SetTimeLast sets the "time_last" field.
func (nuo *NameserverUpdateOne) SetTimeLast(t time.Time) *NameserverUpdateOne {
	nuo.mutation.SetTimeLast(t)
	return nuo
}

// SetNillableTimeLast sets the "time_last" field if the given value is not nil.
func (nuo *NameserverUpdateOne) SetNillableTimeLast(t *time.Time) *NameserverUpdateOne {
	if t != nil {
		nuo.SetTimeLast(*t)
	}
	return nuo
}

// AddIpaddresIDs adds the "ipaddress" edge to the IPAddress entity by IDs.
func (nuo *NameserverUpdateOne) AddIpaddresIDs(ids ...int) *NameserverUpdateOne {
	nuo.mutation.AddIpaddresIDs(ids...)
	return nuo
}

// AddIpaddress adds the "ipaddress" edges to the IPAddress entity.
func (nuo *NameserverUpdateOne) AddIpaddress(i ...*IPAddress) *NameserverUpdateOne {
	ids := make([]int, len(i))
	for j := range i {
		ids[j] = i[j].ID
	}
	return nuo.AddIpaddresIDs(ids...)
}

// AddScanIDs adds the "scan" edge to the ScanJob entity by IDs.
func (nuo *NameserverUpdateOne) AddScanIDs(ids ...int) *NameserverUpdateOne {
	nuo.mutation.AddScanIDs(ids...)
	return nuo
}

// AddScan adds the "scan" edges to the ScanJob entity.
func (nuo *NameserverUpdateOne) AddScan(s ...*ScanJob) *NameserverUpdateOne {
	ids := make([]int, len(s))
	for i := range s {
		ids[i] = s[i].ID
	}
	return nuo.AddScanIDs(ids...)
}

// AddDnsentryIDs adds the "dnsentry" edge to the DNSEntry entity by IDs.
func (nuo *NameserverUpdateOne) AddDnsentryIDs(ids ...int) *NameserverUpdateOne {
	nuo.mutation.AddDnsentryIDs(ids...)
	return nuo
}

// AddDnsentry adds the "dnsentry" edges to the DNSEntry entity.
func (nuo *NameserverUpdateOne) AddDnsentry(d ...*DNSEntry) *NameserverUpdateOne {
	ids := make([]int, len(d))
	for i := range d {
		ids[i] = d[i].ID
	}
	return nuo.AddDnsentryIDs(ids...)
}

// AddDomainIDs adds the "domain" edge to the Domain entity by IDs.
func (nuo *NameserverUpdateOne) AddDomainIDs(ids ...int) *NameserverUpdateOne {
	nuo.mutation.AddDomainIDs(ids...)
	return nuo
}

// AddDomain adds the "domain" edges to the Domain entity.
func (nuo *NameserverUpdateOne) AddDomain(d ...*Domain) *NameserverUpdateOne {
	ids := make([]int, len(d))
	for i := range d {
		ids[i] = d[i].ID
	}
	return nuo.AddDomainIDs(ids...)
}

// AddWhoiIDs adds the "whois" edge to the Whois entity by IDs.
func (nuo *NameserverUpdateOne) AddWhoiIDs(ids ...int) *NameserverUpdateOne {
	nuo.mutation.AddWhoiIDs(ids...)
	return nuo
}

// AddWhois adds the "whois" edges to the Whois entity.
func (nuo *NameserverUpdateOne) AddWhois(w ...*Whois) *NameserverUpdateOne {
	ids := make([]int, len(w))
	for i := range w {
		ids[i] = w[i].ID
	}
	return nuo.AddWhoiIDs(ids...)
}

// Mutation returns the NameserverMutation object of the builder.
func (nuo *NameserverUpdateOne) Mutation() *NameserverMutation {
	return nuo.mutation
}

// ClearIpaddress clears all "ipaddress" edges to the IPAddress entity.
func (nuo *NameserverUpdateOne) ClearIpaddress() *NameserverUpdateOne {
	nuo.mutation.ClearIpaddress()
	return nuo
}

// RemoveIpaddresIDs removes the "ipaddress" edge to IPAddress entities by IDs.
func (nuo *NameserverUpdateOne) RemoveIpaddresIDs(ids ...int) *NameserverUpdateOne {
	nuo.mutation.RemoveIpaddresIDs(ids...)
	return nuo
}

// RemoveIpaddress removes "ipaddress" edges to IPAddress entities.
func (nuo *NameserverUpdateOne) RemoveIpaddress(i ...*IPAddress) *NameserverUpdateOne {
	ids := make([]int, len(i))
	for j := range i {
		ids[j] = i[j].ID
	}
	return nuo.RemoveIpaddresIDs(ids...)
}

// ClearScan clears all "scan" edges to the ScanJob entity.
func (nuo *NameserverUpdateOne) ClearScan() *NameserverUpdateOne {
	nuo.mutation.ClearScan()
	return nuo
}

// RemoveScanIDs removes the "scan" edge to ScanJob entities by IDs.
func (nuo *NameserverUpdateOne) RemoveScanIDs(ids ...int) *NameserverUpdateOne {
	nuo.mutation.RemoveScanIDs(ids...)
	return nuo
}

// RemoveScan removes "scan" edges to ScanJob entities.
func (nuo *NameserverUpdateOne) RemoveScan(s ...*ScanJob) *NameserverUpdateOne {
	ids := make([]int, len(s))
	for i := range s {
		ids[i] = s[i].ID
	}
	return nuo.RemoveScanIDs(ids...)
}

// ClearDnsentry clears all "dnsentry" edges to the DNSEntry entity.
func (nuo *NameserverUpdateOne) ClearDnsentry() *NameserverUpdateOne {
	nuo.mutation.ClearDnsentry()
	return nuo
}

// RemoveDnsentryIDs removes the "dnsentry" edge to DNSEntry entities by IDs.
func (nuo *NameserverUpdateOne) RemoveDnsentryIDs(ids ...int) *NameserverUpdateOne {
	nuo.mutation.RemoveDnsentryIDs(ids...)
	return nuo
}

// RemoveDnsentry removes "dnsentry" edges to DNSEntry entities.
func (nuo *NameserverUpdateOne) RemoveDnsentry(d ...*DNSEntry) *NameserverUpdateOne {
	ids := make([]int, len(d))
	for i := range d {
		ids[i] = d[i].ID
	}
	return nuo.RemoveDnsentryIDs(ids...)
}

// ClearDomain clears all "domain" edges to the Domain entity.
func (nuo *NameserverUpdateOne) ClearDomain() *NameserverUpdateOne {
	nuo.mutation.ClearDomain()
	return nuo
}

// RemoveDomainIDs removes the "domain" edge to Domain entities by IDs.
func (nuo *NameserverUpdateOne) RemoveDomainIDs(ids ...int) *NameserverUpdateOne {
	nuo.mutation.RemoveDomainIDs(ids...)
	return nuo
}

// RemoveDomain removes "domain" edges to Domain entities.
func (nuo *NameserverUpdateOne) RemoveDomain(d ...*Domain) *NameserverUpdateOne {
	ids := make([]int, len(d))
	for i := range d {
		ids[i] = d[i].ID
	}
	return nuo.RemoveDomainIDs(ids...)
}

// ClearWhois clears all "whois" edges to the Whois entity.
func (nuo *NameserverUpdateOne) ClearWhois() *NameserverUpdateOne {
	nuo.mutation.ClearWhois()
	return nuo
}

// RemoveWhoiIDs removes the "whois" edge to Whois entities by IDs.
func (nuo *NameserverUpdateOne) RemoveWhoiIDs(ids ...int) *NameserverUpdateOne {
	nuo.mutation.RemoveWhoiIDs(ids...)
	return nuo
}

// RemoveWhois removes "whois" edges to Whois entities.
func (nuo *NameserverUpdateOne) RemoveWhois(w ...*Whois) *NameserverUpdateOne {
	ids := make([]int, len(w))
	for i := range w {
		ids[i] = w[i].ID
	}
	return nuo.RemoveWhoiIDs(ids...)
}

// Where appends a list predicates to the NameserverUpdate builder.
func (nuo *NameserverUpdateOne) Where(ps ...predicate.Nameserver) *NameserverUpdateOne {
	nuo.mutation.Where(ps...)
	return nuo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (nuo *NameserverUpdateOne) Select(field string, fields ...string) *NameserverUpdateOne {
	nuo.fields = append([]string{field}, fields...)
	return nuo
}

// Save executes the query and returns the updated Nameserver entity.
func (nuo *NameserverUpdateOne) Save(ctx context.Context) (*Nameserver, error) {
	return withHooks(ctx, nuo.sqlSave, nuo.mutation, nuo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (nuo *NameserverUpdateOne) SaveX(ctx context.Context) *Nameserver {
	node, err := nuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (nuo *NameserverUpdateOne) Exec(ctx context.Context) error {
	_, err := nuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (nuo *NameserverUpdateOne) ExecX(ctx context.Context) {
	if err := nuo.Exec(ctx); err != nil {
		panic(err)
	}
}

func (nuo *NameserverUpdateOne) sqlSave(ctx context.Context) (_node *Nameserver, err error) {
	_spec := sqlgraph.NewUpdateSpec(nameserver.Table, nameserver.Columns, sqlgraph.NewFieldSpec(nameserver.FieldID, field.TypeInt))
	id, ok := nuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`model_ent: missing "Nameserver.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := nuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, nameserver.FieldID)
		for _, f := range fields {
			if !nameserver.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("model_ent: invalid field %q for query", f)}
			}
			if f != nameserver.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := nuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := nuo.mutation.Name(); ok {
		_spec.SetField(nameserver.FieldName, field.TypeString, value)
	}
	if value, ok := nuo.mutation.TimeFirst(); ok {
		_spec.SetField(nameserver.FieldTimeFirst, field.TypeTime, value)
	}
	if value, ok := nuo.mutation.TimeLast(); ok {
		_spec.SetField(nameserver.FieldTimeLast, field.TypeTime, value)
	}
	if nuo.mutation.IpaddressCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   nameserver.IpaddressTable,
			Columns: nameserver.IpaddressPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(ipaddress.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := nuo.mutation.RemovedIpaddressIDs(); len(nodes) > 0 && !nuo.mutation.IpaddressCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   nameserver.IpaddressTable,
			Columns: nameserver.IpaddressPrimaryKey,
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
	if nodes := nuo.mutation.IpaddressIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   nameserver.IpaddressTable,
			Columns: nameserver.IpaddressPrimaryKey,
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
	if nuo.mutation.ScanCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   nameserver.ScanTable,
			Columns: nameserver.ScanPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(scanjob.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := nuo.mutation.RemovedScanIDs(); len(nodes) > 0 && !nuo.mutation.ScanCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   nameserver.ScanTable,
			Columns: nameserver.ScanPrimaryKey,
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
	if nodes := nuo.mutation.ScanIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   nameserver.ScanTable,
			Columns: nameserver.ScanPrimaryKey,
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
	if nuo.mutation.DnsentryCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   nameserver.DnsentryTable,
			Columns: nameserver.DnsentryPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(dnsentry.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := nuo.mutation.RemovedDnsentryIDs(); len(nodes) > 0 && !nuo.mutation.DnsentryCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   nameserver.DnsentryTable,
			Columns: nameserver.DnsentryPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(dnsentry.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := nuo.mutation.DnsentryIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   nameserver.DnsentryTable,
			Columns: nameserver.DnsentryPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(dnsentry.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if nuo.mutation.DomainCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   nameserver.DomainTable,
			Columns: nameserver.DomainPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(domain.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := nuo.mutation.RemovedDomainIDs(); len(nodes) > 0 && !nuo.mutation.DomainCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   nameserver.DomainTable,
			Columns: nameserver.DomainPrimaryKey,
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
	if nodes := nuo.mutation.DomainIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   nameserver.DomainTable,
			Columns: nameserver.DomainPrimaryKey,
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
	if nuo.mutation.WhoisCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   nameserver.WhoisTable,
			Columns: nameserver.WhoisPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(whois.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := nuo.mutation.RemovedWhoisIDs(); len(nodes) > 0 && !nuo.mutation.WhoisCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   nameserver.WhoisTable,
			Columns: nameserver.WhoisPrimaryKey,
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
	if nodes := nuo.mutation.WhoisIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   nameserver.WhoisTable,
			Columns: nameserver.WhoisPrimaryKey,
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
	_node = &Nameserver{config: nuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, nuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{nameserver.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	nuo.mutation.done = true
	return _node, nil
}
