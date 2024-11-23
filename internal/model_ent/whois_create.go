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
	"github.com/ice-bergtech/dnh/src/internal/model_ent/domain"
	"github.com/ice-bergtech/dnh/src/internal/model_ent/ipaddress"
	"github.com/ice-bergtech/dnh/src/internal/model_ent/nameserver"
	"github.com/ice-bergtech/dnh/src/internal/model_ent/registrar"
	"github.com/ice-bergtech/dnh/src/internal/model_ent/whois"
)

// WhoisCreate is the builder for creating a Whois entity.
type WhoisCreate struct {
	config
	mutation *WhoisMutation
	hooks    []Hook
}

// SetQuery sets the "query" field.
func (wc *WhoisCreate) SetQuery(s string) *WhoisCreate {
	wc.mutation.SetQuery(s)
	return wc
}

// SetServer sets the "server" field.
func (wc *WhoisCreate) SetServer(s string) *WhoisCreate {
	wc.mutation.SetServer(s)
	return wc
}

// SetRaw sets the "raw" field.
func (wc *WhoisCreate) SetRaw(s string) *WhoisCreate {
	wc.mutation.SetRaw(s)
	return wc
}

// SetCountry sets the "country" field.
func (wc *WhoisCreate) SetCountry(s string) *WhoisCreate {
	wc.mutation.SetCountry(s)
	return wc
}

// SetCreated sets the "created" field.
func (wc *WhoisCreate) SetCreated(t time.Time) *WhoisCreate {
	wc.mutation.SetCreated(t)
	return wc
}

// SetUpdated sets the "updated" field.
func (wc *WhoisCreate) SetUpdated(t time.Time) *WhoisCreate {
	wc.mutation.SetUpdated(t)
	return wc
}

// SetTimeFirst sets the "time_first" field.
func (wc *WhoisCreate) SetTimeFirst(t time.Time) *WhoisCreate {
	wc.mutation.SetTimeFirst(t)
	return wc
}

// SetTimeLast sets the "time_last" field.
func (wc *WhoisCreate) SetTimeLast(t time.Time) *WhoisCreate {
	wc.mutation.SetTimeLast(t)
	return wc
}

// AddIprangeIDs adds the "iprange" edge to the IPAddress entity by IDs.
func (wc *WhoisCreate) AddIprangeIDs(ids ...int) *WhoisCreate {
	wc.mutation.AddIprangeIDs(ids...)
	return wc
}

// AddIprange adds the "iprange" edges to the IPAddress entity.
func (wc *WhoisCreate) AddIprange(i ...*IPAddress) *WhoisCreate {
	ids := make([]int, len(i))
	for j := range i {
		ids[j] = i[j].ID
	}
	return wc.AddIprangeIDs(ids...)
}

// AddDomainIDs adds the "domain" edge to the Domain entity by IDs.
func (wc *WhoisCreate) AddDomainIDs(ids ...int) *WhoisCreate {
	wc.mutation.AddDomainIDs(ids...)
	return wc
}

// AddDomain adds the "domain" edges to the Domain entity.
func (wc *WhoisCreate) AddDomain(d ...*Domain) *WhoisCreate {
	ids := make([]int, len(d))
	for i := range d {
		ids[i] = d[i].ID
	}
	return wc.AddDomainIDs(ids...)
}

// AddAsnIDs adds the "asn" edge to the ASNInfo entity by IDs.
func (wc *WhoisCreate) AddAsnIDs(ids ...int) *WhoisCreate {
	wc.mutation.AddAsnIDs(ids...)
	return wc
}

// AddAsn adds the "asn" edges to the ASNInfo entity.
func (wc *WhoisCreate) AddAsn(a ...*ASNInfo) *WhoisCreate {
	ids := make([]int, len(a))
	for i := range a {
		ids[i] = a[i].ID
	}
	return wc.AddAsnIDs(ids...)
}

// AddRegistrarIDs adds the "registrar" edge to the Registrar entity by IDs.
func (wc *WhoisCreate) AddRegistrarIDs(ids ...int) *WhoisCreate {
	wc.mutation.AddRegistrarIDs(ids...)
	return wc
}

// AddRegistrar adds the "registrar" edges to the Registrar entity.
func (wc *WhoisCreate) AddRegistrar(r ...*Registrar) *WhoisCreate {
	ids := make([]int, len(r))
	for i := range r {
		ids[i] = r[i].ID
	}
	return wc.AddRegistrarIDs(ids...)
}

// AddNameserverIDs adds the "nameservers" edge to the Nameserver entity by IDs.
func (wc *WhoisCreate) AddNameserverIDs(ids ...int) *WhoisCreate {
	wc.mutation.AddNameserverIDs(ids...)
	return wc
}

// AddNameservers adds the "nameservers" edges to the Nameserver entity.
func (wc *WhoisCreate) AddNameservers(n ...*Nameserver) *WhoisCreate {
	ids := make([]int, len(n))
	for i := range n {
		ids[i] = n[i].ID
	}
	return wc.AddNameserverIDs(ids...)
}

// Mutation returns the WhoisMutation object of the builder.
func (wc *WhoisCreate) Mutation() *WhoisMutation {
	return wc.mutation
}

// Save creates the Whois in the database.
func (wc *WhoisCreate) Save(ctx context.Context) (*Whois, error) {
	return withHooks(ctx, wc.sqlSave, wc.mutation, wc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (wc *WhoisCreate) SaveX(ctx context.Context) *Whois {
	v, err := wc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (wc *WhoisCreate) Exec(ctx context.Context) error {
	_, err := wc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (wc *WhoisCreate) ExecX(ctx context.Context) {
	if err := wc.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (wc *WhoisCreate) check() error {
	if _, ok := wc.mutation.Query(); !ok {
		return &ValidationError{Name: "query", err: errors.New(`model_ent: missing required field "Whois.query"`)}
	}
	if _, ok := wc.mutation.Server(); !ok {
		return &ValidationError{Name: "server", err: errors.New(`model_ent: missing required field "Whois.server"`)}
	}
	if _, ok := wc.mutation.Raw(); !ok {
		return &ValidationError{Name: "raw", err: errors.New(`model_ent: missing required field "Whois.raw"`)}
	}
	if _, ok := wc.mutation.Country(); !ok {
		return &ValidationError{Name: "country", err: errors.New(`model_ent: missing required field "Whois.country"`)}
	}
	if _, ok := wc.mutation.Created(); !ok {
		return &ValidationError{Name: "created", err: errors.New(`model_ent: missing required field "Whois.created"`)}
	}
	if _, ok := wc.mutation.Updated(); !ok {
		return &ValidationError{Name: "updated", err: errors.New(`model_ent: missing required field "Whois.updated"`)}
	}
	if _, ok := wc.mutation.TimeFirst(); !ok {
		return &ValidationError{Name: "time_first", err: errors.New(`model_ent: missing required field "Whois.time_first"`)}
	}
	if _, ok := wc.mutation.TimeLast(); !ok {
		return &ValidationError{Name: "time_last", err: errors.New(`model_ent: missing required field "Whois.time_last"`)}
	}
	return nil
}

func (wc *WhoisCreate) sqlSave(ctx context.Context) (*Whois, error) {
	if err := wc.check(); err != nil {
		return nil, err
	}
	_node, _spec := wc.createSpec()
	if err := sqlgraph.CreateNode(ctx, wc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	_node.ID = int(id)
	wc.mutation.id = &_node.ID
	wc.mutation.done = true
	return _node, nil
}

func (wc *WhoisCreate) createSpec() (*Whois, *sqlgraph.CreateSpec) {
	var (
		_node = &Whois{config: wc.config}
		_spec = sqlgraph.NewCreateSpec(whois.Table, sqlgraph.NewFieldSpec(whois.FieldID, field.TypeInt))
	)
	if value, ok := wc.mutation.Query(); ok {
		_spec.SetField(whois.FieldQuery, field.TypeString, value)
		_node.Query = value
	}
	if value, ok := wc.mutation.Server(); ok {
		_spec.SetField(whois.FieldServer, field.TypeString, value)
		_node.Server = value
	}
	if value, ok := wc.mutation.Raw(); ok {
		_spec.SetField(whois.FieldRaw, field.TypeString, value)
		_node.Raw = value
	}
	if value, ok := wc.mutation.Country(); ok {
		_spec.SetField(whois.FieldCountry, field.TypeString, value)
		_node.Country = value
	}
	if value, ok := wc.mutation.Created(); ok {
		_spec.SetField(whois.FieldCreated, field.TypeTime, value)
		_node.Created = value
	}
	if value, ok := wc.mutation.Updated(); ok {
		_spec.SetField(whois.FieldUpdated, field.TypeTime, value)
		_node.Updated = value
	}
	if value, ok := wc.mutation.TimeFirst(); ok {
		_spec.SetField(whois.FieldTimeFirst, field.TypeTime, value)
		_node.TimeFirst = value
	}
	if value, ok := wc.mutation.TimeLast(); ok {
		_spec.SetField(whois.FieldTimeLast, field.TypeTime, value)
		_node.TimeLast = value
	}
	if nodes := wc.mutation.IprangeIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   whois.IprangeTable,
			Columns: []string{whois.IprangeColumn},
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
	if nodes := wc.mutation.DomainIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   whois.DomainTable,
			Columns: []string{whois.DomainColumn},
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
	if nodes := wc.mutation.AsnIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   whois.AsnTable,
			Columns: []string{whois.AsnColumn},
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
	if nodes := wc.mutation.RegistrarIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   whois.RegistrarTable,
			Columns: []string{whois.RegistrarColumn},
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
	if nodes := wc.mutation.NameserversIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   whois.NameserversTable,
			Columns: []string{whois.NameserversColumn},
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
	return _node, _spec
}

// WhoisCreateBulk is the builder for creating many Whois entities in bulk.
type WhoisCreateBulk struct {
	config
	err      error
	builders []*WhoisCreate
}

// Save creates the Whois entities in the database.
func (wcb *WhoisCreateBulk) Save(ctx context.Context) ([]*Whois, error) {
	if wcb.err != nil {
		return nil, wcb.err
	}
	specs := make([]*sqlgraph.CreateSpec, len(wcb.builders))
	nodes := make([]*Whois, len(wcb.builders))
	mutators := make([]Mutator, len(wcb.builders))
	for i := range wcb.builders {
		func(i int, root context.Context) {
			builder := wcb.builders[i]
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*WhoisMutation)
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
					_, err = mutators[i+1].Mutate(root, wcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, wcb.driver, spec); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, wcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (wcb *WhoisCreateBulk) SaveX(ctx context.Context) []*Whois {
	v, err := wcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (wcb *WhoisCreateBulk) Exec(ctx context.Context) error {
	_, err := wcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (wcb *WhoisCreateBulk) ExecX(ctx context.Context) {
	if err := wcb.Exec(ctx); err != nil {
		panic(err)
	}
}