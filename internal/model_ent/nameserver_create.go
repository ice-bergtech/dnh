// Code generated by ent, DO NOT EDIT.

package model_ent

import (
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/ice-bergtech/dnh/src/internal/model_ent/domain"
	"github.com/ice-bergtech/dnh/src/internal/model_ent/ipaddress"
	"github.com/ice-bergtech/dnh/src/internal/model_ent/nameserver"
	"github.com/ice-bergtech/dnh/src/internal/model_ent/scan"
	"github.com/ice-bergtech/dnh/src/internal/model_ent/whois"
)

// NameserverCreate is the builder for creating a Nameserver entity.
type NameserverCreate struct {
	config
	mutation *NameserverMutation
	hooks    []Hook
}

// SetName sets the "name" field.
func (nc *NameserverCreate) SetName(s string) *NameserverCreate {
	nc.mutation.SetName(s)
	return nc
}

// SetTimeFirst sets the "time_first" field.
func (nc *NameserverCreate) SetTimeFirst(t time.Time) *NameserverCreate {
	nc.mutation.SetTimeFirst(t)
	return nc
}

// SetTimeLast sets the "time_last" field.
func (nc *NameserverCreate) SetTimeLast(t time.Time) *NameserverCreate {
	nc.mutation.SetTimeLast(t)
	return nc
}

// AddIpaddresIDs adds the "ipaddress" edge to the IPAddress entity by IDs.
func (nc *NameserverCreate) AddIpaddresIDs(ids ...int) *NameserverCreate {
	nc.mutation.AddIpaddresIDs(ids...)
	return nc
}

// AddIpaddress adds the "ipaddress" edges to the IPAddress entity.
func (nc *NameserverCreate) AddIpaddress(i ...*IPAddress) *NameserverCreate {
	ids := make([]int, len(i))
	for j := range i {
		ids[j] = i[j].ID
	}
	return nc.AddIpaddresIDs(ids...)
}

// AddScanIDs adds the "scan" edge to the Scan entity by IDs.
func (nc *NameserverCreate) AddScanIDs(ids ...int) *NameserverCreate {
	nc.mutation.AddScanIDs(ids...)
	return nc
}

// AddScan adds the "scan" edges to the Scan entity.
func (nc *NameserverCreate) AddScan(s ...*Scan) *NameserverCreate {
	ids := make([]int, len(s))
	for i := range s {
		ids[i] = s[i].ID
	}
	return nc.AddScanIDs(ids...)
}

// AddDnsentryIDs adds the "dnsentry" edge to the Scan entity by IDs.
func (nc *NameserverCreate) AddDnsentryIDs(ids ...int) *NameserverCreate {
	nc.mutation.AddDnsentryIDs(ids...)
	return nc
}

// AddDnsentry adds the "dnsentry" edges to the Scan entity.
func (nc *NameserverCreate) AddDnsentry(s ...*Scan) *NameserverCreate {
	ids := make([]int, len(s))
	for i := range s {
		ids[i] = s[i].ID
	}
	return nc.AddDnsentryIDs(ids...)
}

// AddDomainIDs adds the "domain" edge to the Domain entity by IDs.
func (nc *NameserverCreate) AddDomainIDs(ids ...int) *NameserverCreate {
	nc.mutation.AddDomainIDs(ids...)
	return nc
}

// AddDomain adds the "domain" edges to the Domain entity.
func (nc *NameserverCreate) AddDomain(d ...*Domain) *NameserverCreate {
	ids := make([]int, len(d))
	for i := range d {
		ids[i] = d[i].ID
	}
	return nc.AddDomainIDs(ids...)
}

// AddWhoiIDs adds the "whois" edge to the Whois entity by IDs.
func (nc *NameserverCreate) AddWhoiIDs(ids ...int) *NameserverCreate {
	nc.mutation.AddWhoiIDs(ids...)
	return nc
}

// AddWhois adds the "whois" edges to the Whois entity.
func (nc *NameserverCreate) AddWhois(w ...*Whois) *NameserverCreate {
	ids := make([]int, len(w))
	for i := range w {
		ids[i] = w[i].ID
	}
	return nc.AddWhoiIDs(ids...)
}

// Mutation returns the NameserverMutation object of the builder.
func (nc *NameserverCreate) Mutation() *NameserverMutation {
	return nc.mutation
}

// Save creates the Nameserver in the database.
func (nc *NameserverCreate) Save(ctx context.Context) (*Nameserver, error) {
	return withHooks(ctx, nc.sqlSave, nc.mutation, nc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (nc *NameserverCreate) SaveX(ctx context.Context) *Nameserver {
	v, err := nc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (nc *NameserverCreate) Exec(ctx context.Context) error {
	_, err := nc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (nc *NameserverCreate) ExecX(ctx context.Context) {
	if err := nc.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (nc *NameserverCreate) check() error {
	if _, ok := nc.mutation.Name(); !ok {
		return &ValidationError{Name: "name", err: errors.New(`model_ent: missing required field "Nameserver.name"`)}
	}
	if _, ok := nc.mutation.TimeFirst(); !ok {
		return &ValidationError{Name: "time_first", err: errors.New(`model_ent: missing required field "Nameserver.time_first"`)}
	}
	if _, ok := nc.mutation.TimeLast(); !ok {
		return &ValidationError{Name: "time_last", err: errors.New(`model_ent: missing required field "Nameserver.time_last"`)}
	}
	return nil
}

func (nc *NameserverCreate) sqlSave(ctx context.Context) (*Nameserver, error) {
	if err := nc.check(); err != nil {
		return nil, err
	}
	_node, _spec := nc.createSpec()
	if err := sqlgraph.CreateNode(ctx, nc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	_node.ID = int(id)
	nc.mutation.id = &_node.ID
	nc.mutation.done = true
	return _node, nil
}

func (nc *NameserverCreate) createSpec() (*Nameserver, *sqlgraph.CreateSpec) {
	var (
		_node = &Nameserver{config: nc.config}
		_spec = sqlgraph.NewCreateSpec(nameserver.Table, sqlgraph.NewFieldSpec(nameserver.FieldID, field.TypeInt))
	)
	if value, ok := nc.mutation.Name(); ok {
		_spec.SetField(nameserver.FieldName, field.TypeString, value)
		_node.Name = value
	}
	if value, ok := nc.mutation.TimeFirst(); ok {
		_spec.SetField(nameserver.FieldTimeFirst, field.TypeTime, value)
		_node.TimeFirst = value
	}
	if value, ok := nc.mutation.TimeLast(); ok {
		_spec.SetField(nameserver.FieldTimeLast, field.TypeTime, value)
		_node.TimeLast = value
	}
	if nodes := nc.mutation.IpaddressIDs(); len(nodes) > 0 {
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
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := nc.mutation.ScanIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   nameserver.ScanTable,
			Columns: nameserver.ScanPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(scan.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := nc.mutation.DnsentryIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   nameserver.DnsentryTable,
			Columns: nameserver.DnsentryPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(scan.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := nc.mutation.DomainIDs(); len(nodes) > 0 {
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
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := nc.mutation.WhoisIDs(); len(nodes) > 0 {
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
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// NameserverCreateBulk is the builder for creating many Nameserver entities in bulk.
type NameserverCreateBulk struct {
	config
	err      error
	builders []*NameserverCreate
}

// Save creates the Nameserver entities in the database.
func (ncb *NameserverCreateBulk) Save(ctx context.Context) ([]*Nameserver, error) {
	if ncb.err != nil {
		return nil, ncb.err
	}
	specs := make([]*sqlgraph.CreateSpec, len(ncb.builders))
	nodes := make([]*Nameserver, len(ncb.builders))
	mutators := make([]Mutator, len(ncb.builders))
	for i := range ncb.builders {
		func(i int, root context.Context) {
			builder := ncb.builders[i]
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*NameserverMutation)
				if !ok {
					return nil, fmt.Errorf("unexpected mutation type %T", m)
				}
				if err := builder.check(); err != nil {
					return nil, err
				}
				builder.mutation = mutation
				var err error
				nodes[i], specs[i] = builder.createSpec()
				if i < len(mutators)-1 {
					_, err = mutators[i+1].Mutate(root, ncb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, ncb.driver, spec); err != nil {
						if sqlgraph.IsConstraintError(err) {
							err = &ConstraintError{msg: err.Error(), wrap: err}
						}
					}
				}
				if err != nil {
					return nil, err
				}
				mutation.id = &nodes[i].ID
				if specs[i].ID.Value != nil {
					id := specs[i].ID.Value.(int64)
					nodes[i].ID = int(id)
				}
				mutation.done = true
				return nodes[i], nil
			})
			for i := len(builder.hooks) - 1; i >= 0; i-- {
				mut = builder.hooks[i](mut)
			}
			mutators[i] = mut
		}(i, ctx)
	}
	if len(mutators) > 0 {
		if _, err := mutators[0].Mutate(ctx, ncb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (ncb *NameserverCreateBulk) SaveX(ctx context.Context) []*Nameserver {
	v, err := ncb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (ncb *NameserverCreateBulk) Exec(ctx context.Context) error {
	_, err := ncb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ncb *NameserverCreateBulk) ExecX(ctx context.Context) {
	if err := ncb.Exec(ctx); err != nil {
		panic(err)
	}
}
