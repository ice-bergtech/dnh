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
	"github.com/ice-bergtech/dnh/src/internal/model_ent/scanjob"
)

// DNSEntryCreate is the builder for creating a DNSEntry entity.
type DNSEntryCreate struct {
	config
	mutation *DNSEntryMutation
	hooks    []Hook
	conflict []sql.ConflictOption
}

// SetName sets the "name" field.
func (dec *DNSEntryCreate) SetName(s string) *DNSEntryCreate {
	dec.mutation.SetName(s)
	return dec
}

// SetType sets the "type" field.
func (dec *DNSEntryCreate) SetType(s string) *DNSEntryCreate {
	dec.mutation.SetType(s)
	return dec
}

// SetValue sets the "value" field.
func (dec *DNSEntryCreate) SetValue(s string) *DNSEntryCreate {
	dec.mutation.SetValue(s)
	return dec
}

// SetTTL sets the "ttl" field.
func (dec *DNSEntryCreate) SetTTL(i int) *DNSEntryCreate {
	dec.mutation.SetTTL(i)
	return dec
}

// SetTimeFirst sets the "time_first" field.
func (dec *DNSEntryCreate) SetTimeFirst(t time.Time) *DNSEntryCreate {
	dec.mutation.SetTimeFirst(t)
	return dec
}

// SetTimeLast sets the "time_last" field.
func (dec *DNSEntryCreate) SetTimeLast(t time.Time) *DNSEntryCreate {
	dec.mutation.SetTimeLast(t)
	return dec
}

// AddDomainIDs adds the "domain" edge to the Domain entity by IDs.
func (dec *DNSEntryCreate) AddDomainIDs(ids ...int) *DNSEntryCreate {
	dec.mutation.AddDomainIDs(ids...)
	return dec
}

// AddDomain adds the "domain" edges to the Domain entity.
func (dec *DNSEntryCreate) AddDomain(d ...*Domain) *DNSEntryCreate {
	ids := make([]int, len(d))
	for i := range d {
		ids[i] = d[i].ID
	}
	return dec.AddDomainIDs(ids...)
}

// AddIpaddresIDs adds the "ipaddress" edge to the IPAddress entity by IDs.
func (dec *DNSEntryCreate) AddIpaddresIDs(ids ...int) *DNSEntryCreate {
	dec.mutation.AddIpaddresIDs(ids...)
	return dec
}

// AddIpaddress adds the "ipaddress" edges to the IPAddress entity.
func (dec *DNSEntryCreate) AddIpaddress(i ...*IPAddress) *DNSEntryCreate {
	ids := make([]int, len(i))
	for j := range i {
		ids[j] = i[j].ID
	}
	return dec.AddIpaddresIDs(ids...)
}

// AddNameserverIDs adds the "nameserver" edge to the Nameserver entity by IDs.
func (dec *DNSEntryCreate) AddNameserverIDs(ids ...int) *DNSEntryCreate {
	dec.mutation.AddNameserverIDs(ids...)
	return dec
}

// AddNameserver adds the "nameserver" edges to the Nameserver entity.
func (dec *DNSEntryCreate) AddNameserver(n ...*Nameserver) *DNSEntryCreate {
	ids := make([]int, len(n))
	for i := range n {
		ids[i] = n[i].ID
	}
	return dec.AddNameserverIDs(ids...)
}

// AddScanIDs adds the "scan" edge to the ScanJob entity by IDs.
func (dec *DNSEntryCreate) AddScanIDs(ids ...int) *DNSEntryCreate {
	dec.mutation.AddScanIDs(ids...)
	return dec
}

// AddScan adds the "scan" edges to the ScanJob entity.
func (dec *DNSEntryCreate) AddScan(s ...*ScanJob) *DNSEntryCreate {
	ids := make([]int, len(s))
	for i := range s {
		ids[i] = s[i].ID
	}
	return dec.AddScanIDs(ids...)
}

// Mutation returns the DNSEntryMutation object of the builder.
func (dec *DNSEntryCreate) Mutation() *DNSEntryMutation {
	return dec.mutation
}

// Save creates the DNSEntry in the database.
func (dec *DNSEntryCreate) Save(ctx context.Context) (*DNSEntry, error) {
	return withHooks(ctx, dec.sqlSave, dec.mutation, dec.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (dec *DNSEntryCreate) SaveX(ctx context.Context) *DNSEntry {
	v, err := dec.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (dec *DNSEntryCreate) Exec(ctx context.Context) error {
	_, err := dec.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (dec *DNSEntryCreate) ExecX(ctx context.Context) {
	if err := dec.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (dec *DNSEntryCreate) check() error {
	if _, ok := dec.mutation.Name(); !ok {
		return &ValidationError{Name: "name", err: errors.New(`model_ent: missing required field "DNSEntry.name"`)}
	}
	if _, ok := dec.mutation.GetType(); !ok {
		return &ValidationError{Name: "type", err: errors.New(`model_ent: missing required field "DNSEntry.type"`)}
	}
	if _, ok := dec.mutation.Value(); !ok {
		return &ValidationError{Name: "value", err: errors.New(`model_ent: missing required field "DNSEntry.value"`)}
	}
	if _, ok := dec.mutation.TTL(); !ok {
		return &ValidationError{Name: "ttl", err: errors.New(`model_ent: missing required field "DNSEntry.ttl"`)}
	}
	if _, ok := dec.mutation.TimeFirst(); !ok {
		return &ValidationError{Name: "time_first", err: errors.New(`model_ent: missing required field "DNSEntry.time_first"`)}
	}
	if _, ok := dec.mutation.TimeLast(); !ok {
		return &ValidationError{Name: "time_last", err: errors.New(`model_ent: missing required field "DNSEntry.time_last"`)}
	}
	return nil
}

func (dec *DNSEntryCreate) sqlSave(ctx context.Context) (*DNSEntry, error) {
	if err := dec.check(); err != nil {
		return nil, err
	}
	_node, _spec := dec.createSpec()
	if err := sqlgraph.CreateNode(ctx, dec.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	_node.ID = int(id)
	dec.mutation.id = &_node.ID
	dec.mutation.done = true
	return _node, nil
}

func (dec *DNSEntryCreate) createSpec() (*DNSEntry, *sqlgraph.CreateSpec) {
	var (
		_node = &DNSEntry{config: dec.config}
		_spec = sqlgraph.NewCreateSpec(dnsentry.Table, sqlgraph.NewFieldSpec(dnsentry.FieldID, field.TypeInt))
	)
	_spec.OnConflict = dec.conflict
	if value, ok := dec.mutation.Name(); ok {
		_spec.SetField(dnsentry.FieldName, field.TypeString, value)
		_node.Name = value
	}
	if value, ok := dec.mutation.GetType(); ok {
		_spec.SetField(dnsentry.FieldType, field.TypeString, value)
		_node.Type = value
	}
	if value, ok := dec.mutation.Value(); ok {
		_spec.SetField(dnsentry.FieldValue, field.TypeString, value)
		_node.Value = value
	}
	if value, ok := dec.mutation.TTL(); ok {
		_spec.SetField(dnsentry.FieldTTL, field.TypeInt, value)
		_node.TTL = value
	}
	if value, ok := dec.mutation.TimeFirst(); ok {
		_spec.SetField(dnsentry.FieldTimeFirst, field.TypeTime, value)
		_node.TimeFirst = value
	}
	if value, ok := dec.mutation.TimeLast(); ok {
		_spec.SetField(dnsentry.FieldTimeLast, field.TypeTime, value)
		_node.TimeLast = value
	}
	if nodes := dec.mutation.DomainIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   dnsentry.DomainTable,
			Columns: dnsentry.DomainPrimaryKey,
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
	if nodes := dec.mutation.IpaddressIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   dnsentry.IpaddressTable,
			Columns: dnsentry.IpaddressPrimaryKey,
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
	if nodes := dec.mutation.NameserverIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   dnsentry.NameserverTable,
			Columns: dnsentry.NameserverPrimaryKey,
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
	if nodes := dec.mutation.ScanIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   dnsentry.ScanTable,
			Columns: dnsentry.ScanPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(scanjob.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// OnConflict allows configuring the `ON CONFLICT` / `ON DUPLICATE KEY` clause
// of the `INSERT` statement. For example:
//
//	client.DNSEntry.Create().
//		SetName(v).
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		// Override some of the fields with custom
//		// update values.
//		Update(func(u *ent.DNSEntryUpsert) {
//			SetName(v+v).
//		}).
//		Exec(ctx)
func (dec *DNSEntryCreate) OnConflict(opts ...sql.ConflictOption) *DNSEntryUpsertOne {
	dec.conflict = opts
	return &DNSEntryUpsertOne{
		create: dec,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.DNSEntry.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
func (dec *DNSEntryCreate) OnConflictColumns(columns ...string) *DNSEntryUpsertOne {
	dec.conflict = append(dec.conflict, sql.ConflictColumns(columns...))
	return &DNSEntryUpsertOne{
		create: dec,
	}
}

type (
	// DNSEntryUpsertOne is the builder for "upsert"-ing
	//  one DNSEntry node.
	DNSEntryUpsertOne struct {
		create *DNSEntryCreate
	}

	// DNSEntryUpsert is the "OnConflict" setter.
	DNSEntryUpsert struct {
		*sql.UpdateSet
	}
)

// SetName sets the "name" field.
func (u *DNSEntryUpsert) SetName(v string) *DNSEntryUpsert {
	u.Set(dnsentry.FieldName, v)
	return u
}

// UpdateName sets the "name" field to the value that was provided on create.
func (u *DNSEntryUpsert) UpdateName() *DNSEntryUpsert {
	u.SetExcluded(dnsentry.FieldName)
	return u
}

// SetType sets the "type" field.
func (u *DNSEntryUpsert) SetType(v string) *DNSEntryUpsert {
	u.Set(dnsentry.FieldType, v)
	return u
}

// UpdateType sets the "type" field to the value that was provided on create.
func (u *DNSEntryUpsert) UpdateType() *DNSEntryUpsert {
	u.SetExcluded(dnsentry.FieldType)
	return u
}

// SetValue sets the "value" field.
func (u *DNSEntryUpsert) SetValue(v string) *DNSEntryUpsert {
	u.Set(dnsentry.FieldValue, v)
	return u
}

// UpdateValue sets the "value" field to the value that was provided on create.
func (u *DNSEntryUpsert) UpdateValue() *DNSEntryUpsert {
	u.SetExcluded(dnsentry.FieldValue)
	return u
}

// SetTTL sets the "ttl" field.
func (u *DNSEntryUpsert) SetTTL(v int) *DNSEntryUpsert {
	u.Set(dnsentry.FieldTTL, v)
	return u
}

// UpdateTTL sets the "ttl" field to the value that was provided on create.
func (u *DNSEntryUpsert) UpdateTTL() *DNSEntryUpsert {
	u.SetExcluded(dnsentry.FieldTTL)
	return u
}

// AddTTL adds v to the "ttl" field.
func (u *DNSEntryUpsert) AddTTL(v int) *DNSEntryUpsert {
	u.Add(dnsentry.FieldTTL, v)
	return u
}

// SetTimeFirst sets the "time_first" field.
func (u *DNSEntryUpsert) SetTimeFirst(v time.Time) *DNSEntryUpsert {
	u.Set(dnsentry.FieldTimeFirst, v)
	return u
}

// UpdateTimeFirst sets the "time_first" field to the value that was provided on create.
func (u *DNSEntryUpsert) UpdateTimeFirst() *DNSEntryUpsert {
	u.SetExcluded(dnsentry.FieldTimeFirst)
	return u
}

// SetTimeLast sets the "time_last" field.
func (u *DNSEntryUpsert) SetTimeLast(v time.Time) *DNSEntryUpsert {
	u.Set(dnsentry.FieldTimeLast, v)
	return u
}

// UpdateTimeLast sets the "time_last" field to the value that was provided on create.
func (u *DNSEntryUpsert) UpdateTimeLast() *DNSEntryUpsert {
	u.SetExcluded(dnsentry.FieldTimeLast)
	return u
}

// UpdateNewValues updates the mutable fields using the new values that were set on create.
// Using this option is equivalent to using:
//
//	client.DNSEntry.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//		).
//		Exec(ctx)
func (u *DNSEntryUpsertOne) UpdateNewValues() *DNSEntryUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//	client.DNSEntry.Create().
//	    OnConflict(sql.ResolveWithIgnore()).
//	    Exec(ctx)
func (u *DNSEntryUpsertOne) Ignore() *DNSEntryUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *DNSEntryUpsertOne) DoNothing() *DNSEntryUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the DNSEntryCreate.OnConflict
// documentation for more info.
func (u *DNSEntryUpsertOne) Update(set func(*DNSEntryUpsert)) *DNSEntryUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&DNSEntryUpsert{UpdateSet: update})
	}))
	return u
}

// SetName sets the "name" field.
func (u *DNSEntryUpsertOne) SetName(v string) *DNSEntryUpsertOne {
	return u.Update(func(s *DNSEntryUpsert) {
		s.SetName(v)
	})
}

// UpdateName sets the "name" field to the value that was provided on create.
func (u *DNSEntryUpsertOne) UpdateName() *DNSEntryUpsertOne {
	return u.Update(func(s *DNSEntryUpsert) {
		s.UpdateName()
	})
}

// SetType sets the "type" field.
func (u *DNSEntryUpsertOne) SetType(v string) *DNSEntryUpsertOne {
	return u.Update(func(s *DNSEntryUpsert) {
		s.SetType(v)
	})
}

// UpdateType sets the "type" field to the value that was provided on create.
func (u *DNSEntryUpsertOne) UpdateType() *DNSEntryUpsertOne {
	return u.Update(func(s *DNSEntryUpsert) {
		s.UpdateType()
	})
}

// SetValue sets the "value" field.
func (u *DNSEntryUpsertOne) SetValue(v string) *DNSEntryUpsertOne {
	return u.Update(func(s *DNSEntryUpsert) {
		s.SetValue(v)
	})
}

// UpdateValue sets the "value" field to the value that was provided on create.
func (u *DNSEntryUpsertOne) UpdateValue() *DNSEntryUpsertOne {
	return u.Update(func(s *DNSEntryUpsert) {
		s.UpdateValue()
	})
}

// SetTTL sets the "ttl" field.
func (u *DNSEntryUpsertOne) SetTTL(v int) *DNSEntryUpsertOne {
	return u.Update(func(s *DNSEntryUpsert) {
		s.SetTTL(v)
	})
}

// AddTTL adds v to the "ttl" field.
func (u *DNSEntryUpsertOne) AddTTL(v int) *DNSEntryUpsertOne {
	return u.Update(func(s *DNSEntryUpsert) {
		s.AddTTL(v)
	})
}

// UpdateTTL sets the "ttl" field to the value that was provided on create.
func (u *DNSEntryUpsertOne) UpdateTTL() *DNSEntryUpsertOne {
	return u.Update(func(s *DNSEntryUpsert) {
		s.UpdateTTL()
	})
}

// SetTimeFirst sets the "time_first" field.
func (u *DNSEntryUpsertOne) SetTimeFirst(v time.Time) *DNSEntryUpsertOne {
	return u.Update(func(s *DNSEntryUpsert) {
		s.SetTimeFirst(v)
	})
}

// UpdateTimeFirst sets the "time_first" field to the value that was provided on create.
func (u *DNSEntryUpsertOne) UpdateTimeFirst() *DNSEntryUpsertOne {
	return u.Update(func(s *DNSEntryUpsert) {
		s.UpdateTimeFirst()
	})
}

// SetTimeLast sets the "time_last" field.
func (u *DNSEntryUpsertOne) SetTimeLast(v time.Time) *DNSEntryUpsertOne {
	return u.Update(func(s *DNSEntryUpsert) {
		s.SetTimeLast(v)
	})
}

// UpdateTimeLast sets the "time_last" field to the value that was provided on create.
func (u *DNSEntryUpsertOne) UpdateTimeLast() *DNSEntryUpsertOne {
	return u.Update(func(s *DNSEntryUpsert) {
		s.UpdateTimeLast()
	})
}

// Exec executes the query.
func (u *DNSEntryUpsertOne) Exec(ctx context.Context) error {
	if len(u.create.conflict) == 0 {
		return errors.New("model_ent: missing options for DNSEntryCreate.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *DNSEntryUpsertOne) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}

// Exec executes the UPSERT query and returns the inserted/updated ID.
func (u *DNSEntryUpsertOne) ID(ctx context.Context) (id int, err error) {
	node, err := u.create.Save(ctx)
	if err != nil {
		return id, err
	}
	return node.ID, nil
}

// IDX is like ID, but panics if an error occurs.
func (u *DNSEntryUpsertOne) IDX(ctx context.Context) int {
	id, err := u.ID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// DNSEntryCreateBulk is the builder for creating many DNSEntry entities in bulk.
type DNSEntryCreateBulk struct {
	config
	err      error
	builders []*DNSEntryCreate
	conflict []sql.ConflictOption
}

// Save creates the DNSEntry entities in the database.
func (decb *DNSEntryCreateBulk) Save(ctx context.Context) ([]*DNSEntry, error) {
	if decb.err != nil {
		return nil, decb.err
	}
	specs := make([]*sqlgraph.CreateSpec, len(decb.builders))
	nodes := make([]*DNSEntry, len(decb.builders))
	mutators := make([]Mutator, len(decb.builders))
	for i := range decb.builders {
		func(i int, root context.Context) {
			builder := decb.builders[i]
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*DNSEntryMutation)
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
					_, err = mutators[i+1].Mutate(root, decb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					spec.OnConflict = decb.conflict
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, decb.driver, spec); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, decb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (decb *DNSEntryCreateBulk) SaveX(ctx context.Context) []*DNSEntry {
	v, err := decb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (decb *DNSEntryCreateBulk) Exec(ctx context.Context) error {
	_, err := decb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (decb *DNSEntryCreateBulk) ExecX(ctx context.Context) {
	if err := decb.Exec(ctx); err != nil {
		panic(err)
	}
}

// OnConflict allows configuring the `ON CONFLICT` / `ON DUPLICATE KEY` clause
// of the `INSERT` statement. For example:
//
//	client.DNSEntry.CreateBulk(builders...).
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		// Override some of the fields with custom
//		// update values.
//		Update(func(u *ent.DNSEntryUpsert) {
//			SetName(v+v).
//		}).
//		Exec(ctx)
func (decb *DNSEntryCreateBulk) OnConflict(opts ...sql.ConflictOption) *DNSEntryUpsertBulk {
	decb.conflict = opts
	return &DNSEntryUpsertBulk{
		create: decb,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.DNSEntry.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
func (decb *DNSEntryCreateBulk) OnConflictColumns(columns ...string) *DNSEntryUpsertBulk {
	decb.conflict = append(decb.conflict, sql.ConflictColumns(columns...))
	return &DNSEntryUpsertBulk{
		create: decb,
	}
}

// DNSEntryUpsertBulk is the builder for "upsert"-ing
// a bulk of DNSEntry nodes.
type DNSEntryUpsertBulk struct {
	create *DNSEntryCreateBulk
}

// UpdateNewValues updates the mutable fields using the new values that
// were set on create. Using this option is equivalent to using:
//
//	client.DNSEntry.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//		).
//		Exec(ctx)
func (u *DNSEntryUpsertBulk) UpdateNewValues() *DNSEntryUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//	client.DNSEntry.Create().
//		OnConflict(sql.ResolveWithIgnore()).
//		Exec(ctx)
func (u *DNSEntryUpsertBulk) Ignore() *DNSEntryUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *DNSEntryUpsertBulk) DoNothing() *DNSEntryUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the DNSEntryCreateBulk.OnConflict
// documentation for more info.
func (u *DNSEntryUpsertBulk) Update(set func(*DNSEntryUpsert)) *DNSEntryUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&DNSEntryUpsert{UpdateSet: update})
	}))
	return u
}

// SetName sets the "name" field.
func (u *DNSEntryUpsertBulk) SetName(v string) *DNSEntryUpsertBulk {
	return u.Update(func(s *DNSEntryUpsert) {
		s.SetName(v)
	})
}

// UpdateName sets the "name" field to the value that was provided on create.
func (u *DNSEntryUpsertBulk) UpdateName() *DNSEntryUpsertBulk {
	return u.Update(func(s *DNSEntryUpsert) {
		s.UpdateName()
	})
}

// SetType sets the "type" field.
func (u *DNSEntryUpsertBulk) SetType(v string) *DNSEntryUpsertBulk {
	return u.Update(func(s *DNSEntryUpsert) {
		s.SetType(v)
	})
}

// UpdateType sets the "type" field to the value that was provided on create.
func (u *DNSEntryUpsertBulk) UpdateType() *DNSEntryUpsertBulk {
	return u.Update(func(s *DNSEntryUpsert) {
		s.UpdateType()
	})
}

// SetValue sets the "value" field.
func (u *DNSEntryUpsertBulk) SetValue(v string) *DNSEntryUpsertBulk {
	return u.Update(func(s *DNSEntryUpsert) {
		s.SetValue(v)
	})
}

// UpdateValue sets the "value" field to the value that was provided on create.
func (u *DNSEntryUpsertBulk) UpdateValue() *DNSEntryUpsertBulk {
	return u.Update(func(s *DNSEntryUpsert) {
		s.UpdateValue()
	})
}

// SetTTL sets the "ttl" field.
func (u *DNSEntryUpsertBulk) SetTTL(v int) *DNSEntryUpsertBulk {
	return u.Update(func(s *DNSEntryUpsert) {
		s.SetTTL(v)
	})
}

// AddTTL adds v to the "ttl" field.
func (u *DNSEntryUpsertBulk) AddTTL(v int) *DNSEntryUpsertBulk {
	return u.Update(func(s *DNSEntryUpsert) {
		s.AddTTL(v)
	})
}

// UpdateTTL sets the "ttl" field to the value that was provided on create.
func (u *DNSEntryUpsertBulk) UpdateTTL() *DNSEntryUpsertBulk {
	return u.Update(func(s *DNSEntryUpsert) {
		s.UpdateTTL()
	})
}

// SetTimeFirst sets the "time_first" field.
func (u *DNSEntryUpsertBulk) SetTimeFirst(v time.Time) *DNSEntryUpsertBulk {
	return u.Update(func(s *DNSEntryUpsert) {
		s.SetTimeFirst(v)
	})
}

// UpdateTimeFirst sets the "time_first" field to the value that was provided on create.
func (u *DNSEntryUpsertBulk) UpdateTimeFirst() *DNSEntryUpsertBulk {
	return u.Update(func(s *DNSEntryUpsert) {
		s.UpdateTimeFirst()
	})
}

// SetTimeLast sets the "time_last" field.
func (u *DNSEntryUpsertBulk) SetTimeLast(v time.Time) *DNSEntryUpsertBulk {
	return u.Update(func(s *DNSEntryUpsert) {
		s.SetTimeLast(v)
	})
}

// UpdateTimeLast sets the "time_last" field to the value that was provided on create.
func (u *DNSEntryUpsertBulk) UpdateTimeLast() *DNSEntryUpsertBulk {
	return u.Update(func(s *DNSEntryUpsert) {
		s.UpdateTimeLast()
	})
}

// Exec executes the query.
func (u *DNSEntryUpsertBulk) Exec(ctx context.Context) error {
	if u.create.err != nil {
		return u.create.err
	}
	for i, b := range u.create.builders {
		if len(b.conflict) != 0 {
			return fmt.Errorf("model_ent: OnConflict was set for builder %d. Set it on the DNSEntryCreateBulk instead", i)
		}
	}
	if len(u.create.conflict) == 0 {
		return errors.New("model_ent: missing options for DNSEntryCreateBulk.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *DNSEntryUpsertBulk) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}
