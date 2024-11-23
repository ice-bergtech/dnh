// Code generated by ent, DO NOT EDIT.

package model_ent

import (
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/ice-bergtech/dnh/src/internal/model_ent/asninfo"
	"github.com/ice-bergtech/dnh/src/internal/model_ent/dnsentry"
	"github.com/ice-bergtech/dnh/src/internal/model_ent/domain"
	"github.com/ice-bergtech/dnh/src/internal/model_ent/ipaddress"
	"github.com/ice-bergtech/dnh/src/internal/model_ent/nameserver"
	"github.com/ice-bergtech/dnh/src/internal/model_ent/path"
	"github.com/ice-bergtech/dnh/src/internal/model_ent/registrar"
	"github.com/ice-bergtech/dnh/src/internal/model_ent/scan"
	"github.com/ice-bergtech/dnh/src/internal/model_ent/whois"
)

// ScanCreate is the builder for creating a Scan entity.
type ScanCreate struct {
	config
	mutation *ScanMutation
	hooks    []Hook
}

// SetScanid sets the "scanid" field.
func (sc *ScanCreate) SetScanid(s string) *ScanCreate {
	sc.mutation.SetScanid(s)
	return sc
}

// SetTimestamp sets the "timestamp" field.
func (sc *ScanCreate) SetTimestamp(t time.Time) *ScanCreate {
	sc.mutation.SetTimestamp(t)
	return sc
}

// AddIpaddresIDs adds the "ipaddress" edge to the IPAddress entity by IDs.
func (sc *ScanCreate) AddIpaddresIDs(ids ...int) *ScanCreate {
	sc.mutation.AddIpaddresIDs(ids...)
	return sc
}

// AddIpaddress adds the "ipaddress" edges to the IPAddress entity.
func (sc *ScanCreate) AddIpaddress(i ...*IPAddress) *ScanCreate {
	ids := make([]int, len(i))
	for j := range i {
		ids[j] = i[j].ID
	}
	return sc.AddIpaddresIDs(ids...)
}

// AddAsninfoIDs adds the "asninfo" edge to the ASNInfo entity by IDs.
func (sc *ScanCreate) AddAsninfoIDs(ids ...int) *ScanCreate {
	sc.mutation.AddAsninfoIDs(ids...)
	return sc
}

// AddAsninfo adds the "asninfo" edges to the ASNInfo entity.
func (sc *ScanCreate) AddAsninfo(a ...*ASNInfo) *ScanCreate {
	ids := make([]int, len(a))
	for i := range a {
		ids[i] = a[i].ID
	}
	return sc.AddAsninfoIDs(ids...)
}

// AddDnsentryIDs adds the "dnsentry" edge to the DNSEntry entity by IDs.
func (sc *ScanCreate) AddDnsentryIDs(ids ...int) *ScanCreate {
	sc.mutation.AddDnsentryIDs(ids...)
	return sc
}

// AddDnsentry adds the "dnsentry" edges to the DNSEntry entity.
func (sc *ScanCreate) AddDnsentry(d ...*DNSEntry) *ScanCreate {
	ids := make([]int, len(d))
	for i := range d {
		ids[i] = d[i].ID
	}
	return sc.AddDnsentryIDs(ids...)
}

// AddDomainIDs adds the "domain" edge to the Domain entity by IDs.
func (sc *ScanCreate) AddDomainIDs(ids ...int) *ScanCreate {
	sc.mutation.AddDomainIDs(ids...)
	return sc
}

// AddDomain adds the "domain" edges to the Domain entity.
func (sc *ScanCreate) AddDomain(d ...*Domain) *ScanCreate {
	ids := make([]int, len(d))
	for i := range d {
		ids[i] = d[i].ID
	}
	return sc.AddDomainIDs(ids...)
}

// AddPathIDs adds the "paths" edge to the Path entity by IDs.
func (sc *ScanCreate) AddPathIDs(ids ...int) *ScanCreate {
	sc.mutation.AddPathIDs(ids...)
	return sc
}

// AddPaths adds the "paths" edges to the Path entity.
func (sc *ScanCreate) AddPaths(p ...*Path) *ScanCreate {
	ids := make([]int, len(p))
	for i := range p {
		ids[i] = p[i].ID
	}
	return sc.AddPathIDs(ids...)
}

// AddNameserverIDs adds the "nameserver" edge to the Nameserver entity by IDs.
func (sc *ScanCreate) AddNameserverIDs(ids ...int) *ScanCreate {
	sc.mutation.AddNameserverIDs(ids...)
	return sc
}

// AddNameserver adds the "nameserver" edges to the Nameserver entity.
func (sc *ScanCreate) AddNameserver(n ...*Nameserver) *ScanCreate {
	ids := make([]int, len(n))
	for i := range n {
		ids[i] = n[i].ID
	}
	return sc.AddNameserverIDs(ids...)
}

// AddRegistrarIDs adds the "registrar" edge to the Registrar entity by IDs.
func (sc *ScanCreate) AddRegistrarIDs(ids ...int) *ScanCreate {
	sc.mutation.AddRegistrarIDs(ids...)
	return sc
}

// AddRegistrar adds the "registrar" edges to the Registrar entity.
func (sc *ScanCreate) AddRegistrar(r ...*Registrar) *ScanCreate {
	ids := make([]int, len(r))
	for i := range r {
		ids[i] = r[i].ID
	}
	return sc.AddRegistrarIDs(ids...)
}

// AddWhoiIDs adds the "whois" edge to the Whois entity by IDs.
func (sc *ScanCreate) AddWhoiIDs(ids ...int) *ScanCreate {
	sc.mutation.AddWhoiIDs(ids...)
	return sc
}

// AddWhois adds the "whois" edges to the Whois entity.
func (sc *ScanCreate) AddWhois(w ...*Whois) *ScanCreate {
	ids := make([]int, len(w))
	for i := range w {
		ids[i] = w[i].ID
	}
	return sc.AddWhoiIDs(ids...)
}

// Mutation returns the ScanMutation object of the builder.
func (sc *ScanCreate) Mutation() *ScanMutation {
	return sc.mutation
}

// Save creates the Scan in the database.
func (sc *ScanCreate) Save(ctx context.Context) (*Scan, error) {
	return withHooks(ctx, sc.sqlSave, sc.mutation, sc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (sc *ScanCreate) SaveX(ctx context.Context) *Scan {
	v, err := sc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (sc *ScanCreate) Exec(ctx context.Context) error {
	_, err := sc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (sc *ScanCreate) ExecX(ctx context.Context) {
	if err := sc.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (sc *ScanCreate) check() error {
	if _, ok := sc.mutation.Scanid(); !ok {
		return &ValidationError{Name: "scanid", err: errors.New(`model_ent: missing required field "Scan.scanid"`)}
	}
	if _, ok := sc.mutation.Timestamp(); !ok {
		return &ValidationError{Name: "timestamp", err: errors.New(`model_ent: missing required field "Scan.timestamp"`)}
	}
	return nil
}

func (sc *ScanCreate) sqlSave(ctx context.Context) (*Scan, error) {
	if err := sc.check(); err != nil {
		return nil, err
	}
	_node, _spec := sc.createSpec()
	if err := sqlgraph.CreateNode(ctx, sc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	_node.ID = int(id)
	sc.mutation.id = &_node.ID
	sc.mutation.done = true
	return _node, nil
}

func (sc *ScanCreate) createSpec() (*Scan, *sqlgraph.CreateSpec) {
	var (
		_node = &Scan{config: sc.config}
		_spec = sqlgraph.NewCreateSpec(scan.Table, sqlgraph.NewFieldSpec(scan.FieldID, field.TypeInt))
	)
	if value, ok := sc.mutation.Scanid(); ok {
		_spec.SetField(scan.FieldScanid, field.TypeString, value)
		_node.Scanid = value
	}
	if value, ok := sc.mutation.Timestamp(); ok {
		_spec.SetField(scan.FieldTimestamp, field.TypeTime, value)
		_node.Timestamp = value
	}
	if nodes := sc.mutation.IpaddressIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   scan.IpaddressTable,
			Columns: []string{scan.IpaddressColumn},
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
	if nodes := sc.mutation.AsninfoIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   scan.AsninfoTable,
			Columns: []string{scan.AsninfoColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(asninfo.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := sc.mutation.DnsentryIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   scan.DnsentryTable,
			Columns: []string{scan.DnsentryColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(dnsentry.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := sc.mutation.DomainIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   scan.DomainTable,
			Columns: []string{scan.DomainColumn},
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
	if nodes := sc.mutation.PathsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   scan.PathsTable,
			Columns: []string{scan.PathsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(path.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := sc.mutation.NameserverIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   scan.NameserverTable,
			Columns: []string{scan.NameserverColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(nameserver.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := sc.mutation.RegistrarIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   scan.RegistrarTable,
			Columns: []string{scan.RegistrarColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(registrar.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := sc.mutation.WhoisIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   scan.WhoisTable,
			Columns: []string{scan.WhoisColumn},
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

// ScanCreateBulk is the builder for creating many Scan entities in bulk.
type ScanCreateBulk struct {
	config
	err      error
	builders []*ScanCreate
}

// Save creates the Scan entities in the database.
func (scb *ScanCreateBulk) Save(ctx context.Context) ([]*Scan, error) {
	if scb.err != nil {
		return nil, scb.err
	}
	specs := make([]*sqlgraph.CreateSpec, len(scb.builders))
	nodes := make([]*Scan, len(scb.builders))
	mutators := make([]Mutator, len(scb.builders))
	for i := range scb.builders {
		func(i int, root context.Context) {
			builder := scb.builders[i]
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*ScanMutation)
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
					_, err = mutators[i+1].Mutate(root, scb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, scb.driver, spec); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, scb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (scb *ScanCreateBulk) SaveX(ctx context.Context) []*Scan {
	v, err := scb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (scb *ScanCreateBulk) Exec(ctx context.Context) error {
	_, err := scb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (scb *ScanCreateBulk) ExecX(ctx context.Context) {
	if err := scb.Exec(ctx); err != nil {
		panic(err)
	}
}
