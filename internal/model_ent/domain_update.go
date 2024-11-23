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
	"github.com/ice-bergtech/dnh/src/internal/model_ent/path"
	"github.com/ice-bergtech/dnh/src/internal/model_ent/predicate"
)

// DomainUpdate is the builder for updating Domain entities.
type DomainUpdate struct {
	config
	hooks    []Hook
	mutation *DomainMutation
}

// Where appends a list predicates to the DomainUpdate builder.
func (du *DomainUpdate) Where(ps ...predicate.Domain) *DomainUpdate {
	du.mutation.Where(ps...)
	return du
}

// SetName sets the "name" field.
func (du *DomainUpdate) SetName(s string) *DomainUpdate {
	du.mutation.SetName(s)
	return du
}

// SetNillableName sets the "name" field if the given value is not nil.
func (du *DomainUpdate) SetNillableName(s *string) *DomainUpdate {
	if s != nil {
		du.SetName(*s)
	}
	return du
}

// SetTimeFirst sets the "time_first" field.
func (du *DomainUpdate) SetTimeFirst(t time.Time) *DomainUpdate {
	du.mutation.SetTimeFirst(t)
	return du
}

// SetNillableTimeFirst sets the "time_first" field if the given value is not nil.
func (du *DomainUpdate) SetNillableTimeFirst(t *time.Time) *DomainUpdate {
	if t != nil {
		du.SetTimeFirst(*t)
	}
	return du
}

// SetTimeLast sets the "time_last" field.
func (du *DomainUpdate) SetTimeLast(t time.Time) *DomainUpdate {
	du.mutation.SetTimeLast(t)
	return du
}

// SetNillableTimeLast sets the "time_last" field if the given value is not nil.
func (du *DomainUpdate) SetNillableTimeLast(t *time.Time) *DomainUpdate {
	if t != nil {
		du.SetTimeLast(*t)
	}
	return du
}

// AddNameserverIDs adds the "nameserver" edge to the Nameserver entity by IDs.
func (du *DomainUpdate) AddNameserverIDs(ids ...int) *DomainUpdate {
	du.mutation.AddNameserverIDs(ids...)
	return du
}

// AddNameserver adds the "nameserver" edges to the Nameserver entity.
func (du *DomainUpdate) AddNameserver(n ...*Nameserver) *DomainUpdate {
	ids := make([]int, len(n))
	for i := range n {
		ids[i] = n[i].ID
	}
	return du.AddNameserverIDs(ids...)
}

// AddDomainIDs adds the "domain" edge to the Domain entity by IDs.
func (du *DomainUpdate) AddDomainIDs(ids ...int) *DomainUpdate {
	du.mutation.AddDomainIDs(ids...)
	return du
}

// AddDomain adds the "domain" edges to the Domain entity.
func (du *DomainUpdate) AddDomain(d ...*Domain) *DomainUpdate {
	ids := make([]int, len(d))
	for i := range d {
		ids[i] = d[i].ID
	}
	return du.AddDomainIDs(ids...)
}

// AddDnsentryIDs adds the "dnsentry" edge to the DNSEntry entity by IDs.
func (du *DomainUpdate) AddDnsentryIDs(ids ...int) *DomainUpdate {
	du.mutation.AddDnsentryIDs(ids...)
	return du
}

// AddDnsentry adds the "dnsentry" edges to the DNSEntry entity.
func (du *DomainUpdate) AddDnsentry(d ...*DNSEntry) *DomainUpdate {
	ids := make([]int, len(d))
	for i := range d {
		ids[i] = d[i].ID
	}
	return du.AddDnsentryIDs(ids...)
}

// AddIpaddresIDs adds the "ipaddress" edge to the IPAddress entity by IDs.
func (du *DomainUpdate) AddIpaddresIDs(ids ...int) *DomainUpdate {
	du.mutation.AddIpaddresIDs(ids...)
	return du
}

// AddIpaddress adds the "ipaddress" edges to the IPAddress entity.
func (du *DomainUpdate) AddIpaddress(i ...*IPAddress) *DomainUpdate {
	ids := make([]int, len(i))
	for j := range i {
		ids[j] = i[j].ID
	}
	return du.AddIpaddresIDs(ids...)
}

// AddPathIDs adds the "path" edge to the Path entity by IDs.
func (du *DomainUpdate) AddPathIDs(ids ...int) *DomainUpdate {
	du.mutation.AddPathIDs(ids...)
	return du
}

// AddPath adds the "path" edges to the Path entity.
func (du *DomainUpdate) AddPath(p ...*Path) *DomainUpdate {
	ids := make([]int, len(p))
	for i := range p {
		ids[i] = p[i].ID
	}
	return du.AddPathIDs(ids...)
}

// Mutation returns the DomainMutation object of the builder.
func (du *DomainUpdate) Mutation() *DomainMutation {
	return du.mutation
}

// ClearNameserver clears all "nameserver" edges to the Nameserver entity.
func (du *DomainUpdate) ClearNameserver() *DomainUpdate {
	du.mutation.ClearNameserver()
	return du
}

// RemoveNameserverIDs removes the "nameserver" edge to Nameserver entities by IDs.
func (du *DomainUpdate) RemoveNameserverIDs(ids ...int) *DomainUpdate {
	du.mutation.RemoveNameserverIDs(ids...)
	return du
}

// RemoveNameserver removes "nameserver" edges to Nameserver entities.
func (du *DomainUpdate) RemoveNameserver(n ...*Nameserver) *DomainUpdate {
	ids := make([]int, len(n))
	for i := range n {
		ids[i] = n[i].ID
	}
	return du.RemoveNameserverIDs(ids...)
}

// ClearDomain clears all "domain" edges to the Domain entity.
func (du *DomainUpdate) ClearDomain() *DomainUpdate {
	du.mutation.ClearDomain()
	return du
}

// RemoveDomainIDs removes the "domain" edge to Domain entities by IDs.
func (du *DomainUpdate) RemoveDomainIDs(ids ...int) *DomainUpdate {
	du.mutation.RemoveDomainIDs(ids...)
	return du
}

// RemoveDomain removes "domain" edges to Domain entities.
func (du *DomainUpdate) RemoveDomain(d ...*Domain) *DomainUpdate {
	ids := make([]int, len(d))
	for i := range d {
		ids[i] = d[i].ID
	}
	return du.RemoveDomainIDs(ids...)
}

// ClearDnsentry clears all "dnsentry" edges to the DNSEntry entity.
func (du *DomainUpdate) ClearDnsentry() *DomainUpdate {
	du.mutation.ClearDnsentry()
	return du
}

// RemoveDnsentryIDs removes the "dnsentry" edge to DNSEntry entities by IDs.
func (du *DomainUpdate) RemoveDnsentryIDs(ids ...int) *DomainUpdate {
	du.mutation.RemoveDnsentryIDs(ids...)
	return du
}

// RemoveDnsentry removes "dnsentry" edges to DNSEntry entities.
func (du *DomainUpdate) RemoveDnsentry(d ...*DNSEntry) *DomainUpdate {
	ids := make([]int, len(d))
	for i := range d {
		ids[i] = d[i].ID
	}
	return du.RemoveDnsentryIDs(ids...)
}

// ClearIpaddress clears all "ipaddress" edges to the IPAddress entity.
func (du *DomainUpdate) ClearIpaddress() *DomainUpdate {
	du.mutation.ClearIpaddress()
	return du
}

// RemoveIpaddresIDs removes the "ipaddress" edge to IPAddress entities by IDs.
func (du *DomainUpdate) RemoveIpaddresIDs(ids ...int) *DomainUpdate {
	du.mutation.RemoveIpaddresIDs(ids...)
	return du
}

// RemoveIpaddress removes "ipaddress" edges to IPAddress entities.
func (du *DomainUpdate) RemoveIpaddress(i ...*IPAddress) *DomainUpdate {
	ids := make([]int, len(i))
	for j := range i {
		ids[j] = i[j].ID
	}
	return du.RemoveIpaddresIDs(ids...)
}

// ClearPath clears all "path" edges to the Path entity.
func (du *DomainUpdate) ClearPath() *DomainUpdate {
	du.mutation.ClearPath()
	return du
}

// RemovePathIDs removes the "path" edge to Path entities by IDs.
func (du *DomainUpdate) RemovePathIDs(ids ...int) *DomainUpdate {
	du.mutation.RemovePathIDs(ids...)
	return du
}

// RemovePath removes "path" edges to Path entities.
func (du *DomainUpdate) RemovePath(p ...*Path) *DomainUpdate {
	ids := make([]int, len(p))
	for i := range p {
		ids[i] = p[i].ID
	}
	return du.RemovePathIDs(ids...)
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (du *DomainUpdate) Save(ctx context.Context) (int, error) {
	return withHooks(ctx, du.sqlSave, du.mutation, du.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (du *DomainUpdate) SaveX(ctx context.Context) int {
	affected, err := du.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (du *DomainUpdate) Exec(ctx context.Context) error {
	_, err := du.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (du *DomainUpdate) ExecX(ctx context.Context) {
	if err := du.Exec(ctx); err != nil {
		panic(err)
	}
}

func (du *DomainUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := sqlgraph.NewUpdateSpec(domain.Table, domain.Columns, sqlgraph.NewFieldSpec(domain.FieldID, field.TypeInt))
	if ps := du.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := du.mutation.Name(); ok {
		_spec.SetField(domain.FieldName, field.TypeString, value)
	}
	if value, ok := du.mutation.TimeFirst(); ok {
		_spec.SetField(domain.FieldTimeFirst, field.TypeTime, value)
	}
	if value, ok := du.mutation.TimeLast(); ok {
		_spec.SetField(domain.FieldTimeLast, field.TypeTime, value)
	}
	if du.mutation.NameserverCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   domain.NameserverTable,
			Columns: []string{domain.NameserverColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(nameserver.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := du.mutation.RemovedNameserverIDs(); len(nodes) > 0 && !du.mutation.NameserverCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   domain.NameserverTable,
			Columns: []string{domain.NameserverColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(nameserver.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := du.mutation.NameserverIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   domain.NameserverTable,
			Columns: []string{domain.NameserverColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(nameserver.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if du.mutation.DomainCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   domain.DomainTable,
			Columns: domain.DomainPrimaryKey,
			Bidi:    true,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(domain.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := du.mutation.RemovedDomainIDs(); len(nodes) > 0 && !du.mutation.DomainCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   domain.DomainTable,
			Columns: domain.DomainPrimaryKey,
			Bidi:    true,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(domain.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := du.mutation.DomainIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   domain.DomainTable,
			Columns: domain.DomainPrimaryKey,
			Bidi:    true,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(domain.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if du.mutation.DnsentryCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   domain.DnsentryTable,
			Columns: []string{domain.DnsentryColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(dnsentry.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := du.mutation.RemovedDnsentryIDs(); len(nodes) > 0 && !du.mutation.DnsentryCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   domain.DnsentryTable,
			Columns: []string{domain.DnsentryColumn},
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
	if nodes := du.mutation.DnsentryIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   domain.DnsentryTable,
			Columns: []string{domain.DnsentryColumn},
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
	if du.mutation.IpaddressCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   domain.IpaddressTable,
			Columns: []string{domain.IpaddressColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(ipaddress.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := du.mutation.RemovedIpaddressIDs(); len(nodes) > 0 && !du.mutation.IpaddressCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   domain.IpaddressTable,
			Columns: []string{domain.IpaddressColumn},
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
	if nodes := du.mutation.IpaddressIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   domain.IpaddressTable,
			Columns: []string{domain.IpaddressColumn},
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
	if du.mutation.PathCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   domain.PathTable,
			Columns: []string{domain.PathColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(path.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := du.mutation.RemovedPathIDs(); len(nodes) > 0 && !du.mutation.PathCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   domain.PathTable,
			Columns: []string{domain.PathColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(path.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := du.mutation.PathIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   domain.PathTable,
			Columns: []string{domain.PathColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(path.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, du.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{domain.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	du.mutation.done = true
	return n, nil
}

// DomainUpdateOne is the builder for updating a single Domain entity.
type DomainUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *DomainMutation
}

// SetName sets the "name" field.
func (duo *DomainUpdateOne) SetName(s string) *DomainUpdateOne {
	duo.mutation.SetName(s)
	return duo
}

// SetNillableName sets the "name" field if the given value is not nil.
func (duo *DomainUpdateOne) SetNillableName(s *string) *DomainUpdateOne {
	if s != nil {
		duo.SetName(*s)
	}
	return duo
}

// SetTimeFirst sets the "time_first" field.
func (duo *DomainUpdateOne) SetTimeFirst(t time.Time) *DomainUpdateOne {
	duo.mutation.SetTimeFirst(t)
	return duo
}

// SetNillableTimeFirst sets the "time_first" field if the given value is not nil.
func (duo *DomainUpdateOne) SetNillableTimeFirst(t *time.Time) *DomainUpdateOne {
	if t != nil {
		duo.SetTimeFirst(*t)
	}
	return duo
}

// SetTimeLast sets the "time_last" field.
func (duo *DomainUpdateOne) SetTimeLast(t time.Time) *DomainUpdateOne {
	duo.mutation.SetTimeLast(t)
	return duo
}

// SetNillableTimeLast sets the "time_last" field if the given value is not nil.
func (duo *DomainUpdateOne) SetNillableTimeLast(t *time.Time) *DomainUpdateOne {
	if t != nil {
		duo.SetTimeLast(*t)
	}
	return duo
}

// AddNameserverIDs adds the "nameserver" edge to the Nameserver entity by IDs.
func (duo *DomainUpdateOne) AddNameserverIDs(ids ...int) *DomainUpdateOne {
	duo.mutation.AddNameserverIDs(ids...)
	return duo
}

// AddNameserver adds the "nameserver" edges to the Nameserver entity.
func (duo *DomainUpdateOne) AddNameserver(n ...*Nameserver) *DomainUpdateOne {
	ids := make([]int, len(n))
	for i := range n {
		ids[i] = n[i].ID
	}
	return duo.AddNameserverIDs(ids...)
}

// AddDomainIDs adds the "domain" edge to the Domain entity by IDs.
func (duo *DomainUpdateOne) AddDomainIDs(ids ...int) *DomainUpdateOne {
	duo.mutation.AddDomainIDs(ids...)
	return duo
}

// AddDomain adds the "domain" edges to the Domain entity.
func (duo *DomainUpdateOne) AddDomain(d ...*Domain) *DomainUpdateOne {
	ids := make([]int, len(d))
	for i := range d {
		ids[i] = d[i].ID
	}
	return duo.AddDomainIDs(ids...)
}

// AddDnsentryIDs adds the "dnsentry" edge to the DNSEntry entity by IDs.
func (duo *DomainUpdateOne) AddDnsentryIDs(ids ...int) *DomainUpdateOne {
	duo.mutation.AddDnsentryIDs(ids...)
	return duo
}

// AddDnsentry adds the "dnsentry" edges to the DNSEntry entity.
func (duo *DomainUpdateOne) AddDnsentry(d ...*DNSEntry) *DomainUpdateOne {
	ids := make([]int, len(d))
	for i := range d {
		ids[i] = d[i].ID
	}
	return duo.AddDnsentryIDs(ids...)
}

// AddIpaddresIDs adds the "ipaddress" edge to the IPAddress entity by IDs.
func (duo *DomainUpdateOne) AddIpaddresIDs(ids ...int) *DomainUpdateOne {
	duo.mutation.AddIpaddresIDs(ids...)
	return duo
}

// AddIpaddress adds the "ipaddress" edges to the IPAddress entity.
func (duo *DomainUpdateOne) AddIpaddress(i ...*IPAddress) *DomainUpdateOne {
	ids := make([]int, len(i))
	for j := range i {
		ids[j] = i[j].ID
	}
	return duo.AddIpaddresIDs(ids...)
}

// AddPathIDs adds the "path" edge to the Path entity by IDs.
func (duo *DomainUpdateOne) AddPathIDs(ids ...int) *DomainUpdateOne {
	duo.mutation.AddPathIDs(ids...)
	return duo
}

// AddPath adds the "path" edges to the Path entity.
func (duo *DomainUpdateOne) AddPath(p ...*Path) *DomainUpdateOne {
	ids := make([]int, len(p))
	for i := range p {
		ids[i] = p[i].ID
	}
	return duo.AddPathIDs(ids...)
}

// Mutation returns the DomainMutation object of the builder.
func (duo *DomainUpdateOne) Mutation() *DomainMutation {
	return duo.mutation
}

// ClearNameserver clears all "nameserver" edges to the Nameserver entity.
func (duo *DomainUpdateOne) ClearNameserver() *DomainUpdateOne {
	duo.mutation.ClearNameserver()
	return duo
}

// RemoveNameserverIDs removes the "nameserver" edge to Nameserver entities by IDs.
func (duo *DomainUpdateOne) RemoveNameserverIDs(ids ...int) *DomainUpdateOne {
	duo.mutation.RemoveNameserverIDs(ids...)
	return duo
}

// RemoveNameserver removes "nameserver" edges to Nameserver entities.
func (duo *DomainUpdateOne) RemoveNameserver(n ...*Nameserver) *DomainUpdateOne {
	ids := make([]int, len(n))
	for i := range n {
		ids[i] = n[i].ID
	}
	return duo.RemoveNameserverIDs(ids...)
}

// ClearDomain clears all "domain" edges to the Domain entity.
func (duo *DomainUpdateOne) ClearDomain() *DomainUpdateOne {
	duo.mutation.ClearDomain()
	return duo
}

// RemoveDomainIDs removes the "domain" edge to Domain entities by IDs.
func (duo *DomainUpdateOne) RemoveDomainIDs(ids ...int) *DomainUpdateOne {
	duo.mutation.RemoveDomainIDs(ids...)
	return duo
}

// RemoveDomain removes "domain" edges to Domain entities.
func (duo *DomainUpdateOne) RemoveDomain(d ...*Domain) *DomainUpdateOne {
	ids := make([]int, len(d))
	for i := range d {
		ids[i] = d[i].ID
	}
	return duo.RemoveDomainIDs(ids...)
}

// ClearDnsentry clears all "dnsentry" edges to the DNSEntry entity.
func (duo *DomainUpdateOne) ClearDnsentry() *DomainUpdateOne {
	duo.mutation.ClearDnsentry()
	return duo
}

// RemoveDnsentryIDs removes the "dnsentry" edge to DNSEntry entities by IDs.
func (duo *DomainUpdateOne) RemoveDnsentryIDs(ids ...int) *DomainUpdateOne {
	duo.mutation.RemoveDnsentryIDs(ids...)
	return duo
}

// RemoveDnsentry removes "dnsentry" edges to DNSEntry entities.
func (duo *DomainUpdateOne) RemoveDnsentry(d ...*DNSEntry) *DomainUpdateOne {
	ids := make([]int, len(d))
	for i := range d {
		ids[i] = d[i].ID
	}
	return duo.RemoveDnsentryIDs(ids...)
}

// ClearIpaddress clears all "ipaddress" edges to the IPAddress entity.
func (duo *DomainUpdateOne) ClearIpaddress() *DomainUpdateOne {
	duo.mutation.ClearIpaddress()
	return duo
}

// RemoveIpaddresIDs removes the "ipaddress" edge to IPAddress entities by IDs.
func (duo *DomainUpdateOne) RemoveIpaddresIDs(ids ...int) *DomainUpdateOne {
	duo.mutation.RemoveIpaddresIDs(ids...)
	return duo
}

// RemoveIpaddress removes "ipaddress" edges to IPAddress entities.
func (duo *DomainUpdateOne) RemoveIpaddress(i ...*IPAddress) *DomainUpdateOne {
	ids := make([]int, len(i))
	for j := range i {
		ids[j] = i[j].ID
	}
	return duo.RemoveIpaddresIDs(ids...)
}

// ClearPath clears all "path" edges to the Path entity.
func (duo *DomainUpdateOne) ClearPath() *DomainUpdateOne {
	duo.mutation.ClearPath()
	return duo
}

// RemovePathIDs removes the "path" edge to Path entities by IDs.
func (duo *DomainUpdateOne) RemovePathIDs(ids ...int) *DomainUpdateOne {
	duo.mutation.RemovePathIDs(ids...)
	return duo
}

// RemovePath removes "path" edges to Path entities.
func (duo *DomainUpdateOne) RemovePath(p ...*Path) *DomainUpdateOne {
	ids := make([]int, len(p))
	for i := range p {
		ids[i] = p[i].ID
	}
	return duo.RemovePathIDs(ids...)
}

// Where appends a list predicates to the DomainUpdate builder.
func (duo *DomainUpdateOne) Where(ps ...predicate.Domain) *DomainUpdateOne {
	duo.mutation.Where(ps...)
	return duo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (duo *DomainUpdateOne) Select(field string, fields ...string) *DomainUpdateOne {
	duo.fields = append([]string{field}, fields...)
	return duo
}

// Save executes the query and returns the updated Domain entity.
func (duo *DomainUpdateOne) Save(ctx context.Context) (*Domain, error) {
	return withHooks(ctx, duo.sqlSave, duo.mutation, duo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (duo *DomainUpdateOne) SaveX(ctx context.Context) *Domain {
	node, err := duo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (duo *DomainUpdateOne) Exec(ctx context.Context) error {
	_, err := duo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (duo *DomainUpdateOne) ExecX(ctx context.Context) {
	if err := duo.Exec(ctx); err != nil {
		panic(err)
	}
}

func (duo *DomainUpdateOne) sqlSave(ctx context.Context) (_node *Domain, err error) {
	_spec := sqlgraph.NewUpdateSpec(domain.Table, domain.Columns, sqlgraph.NewFieldSpec(domain.FieldID, field.TypeInt))
	id, ok := duo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`model_ent: missing "Domain.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := duo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, domain.FieldID)
		for _, f := range fields {
			if !domain.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("model_ent: invalid field %q for query", f)}
			}
			if f != domain.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := duo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := duo.mutation.Name(); ok {
		_spec.SetField(domain.FieldName, field.TypeString, value)
	}
	if value, ok := duo.mutation.TimeFirst(); ok {
		_spec.SetField(domain.FieldTimeFirst, field.TypeTime, value)
	}
	if value, ok := duo.mutation.TimeLast(); ok {
		_spec.SetField(domain.FieldTimeLast, field.TypeTime, value)
	}
	if duo.mutation.NameserverCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   domain.NameserverTable,
			Columns: []string{domain.NameserverColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(nameserver.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := duo.mutation.RemovedNameserverIDs(); len(nodes) > 0 && !duo.mutation.NameserverCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   domain.NameserverTable,
			Columns: []string{domain.NameserverColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(nameserver.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := duo.mutation.NameserverIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   domain.NameserverTable,
			Columns: []string{domain.NameserverColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(nameserver.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if duo.mutation.DomainCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   domain.DomainTable,
			Columns: domain.DomainPrimaryKey,
			Bidi:    true,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(domain.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := duo.mutation.RemovedDomainIDs(); len(nodes) > 0 && !duo.mutation.DomainCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   domain.DomainTable,
			Columns: domain.DomainPrimaryKey,
			Bidi:    true,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(domain.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := duo.mutation.DomainIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   domain.DomainTable,
			Columns: domain.DomainPrimaryKey,
			Bidi:    true,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(domain.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if duo.mutation.DnsentryCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   domain.DnsentryTable,
			Columns: []string{domain.DnsentryColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(dnsentry.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := duo.mutation.RemovedDnsentryIDs(); len(nodes) > 0 && !duo.mutation.DnsentryCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   domain.DnsentryTable,
			Columns: []string{domain.DnsentryColumn},
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
	if nodes := duo.mutation.DnsentryIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   domain.DnsentryTable,
			Columns: []string{domain.DnsentryColumn},
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
	if duo.mutation.IpaddressCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   domain.IpaddressTable,
			Columns: []string{domain.IpaddressColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(ipaddress.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := duo.mutation.RemovedIpaddressIDs(); len(nodes) > 0 && !duo.mutation.IpaddressCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   domain.IpaddressTable,
			Columns: []string{domain.IpaddressColumn},
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
	if nodes := duo.mutation.IpaddressIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   domain.IpaddressTable,
			Columns: []string{domain.IpaddressColumn},
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
	if duo.mutation.PathCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   domain.PathTable,
			Columns: []string{domain.PathColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(path.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := duo.mutation.RemovedPathIDs(); len(nodes) > 0 && !duo.mutation.PathCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   domain.PathTable,
			Columns: []string{domain.PathColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(path.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := duo.mutation.PathIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   domain.PathTable,
			Columns: []string{domain.PathColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(path.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &Domain{config: duo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, duo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{domain.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	duo.mutation.done = true
	return _node, nil
}