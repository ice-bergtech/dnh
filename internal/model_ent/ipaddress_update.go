// Code generated by ent, DO NOT EDIT.

package model_ent

import (
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/ice-bergtech/dnh/src/internal/model_ent/asninfo"
	"github.com/ice-bergtech/dnh/src/internal/model_ent/ipaddress"
	"github.com/ice-bergtech/dnh/src/internal/model_ent/predicate"
)

// IPAddressUpdate is the builder for updating IPAddress entities.
type IPAddressUpdate struct {
	config
	hooks    []Hook
	mutation *IPAddressMutation
}

// Where appends a list predicates to the IPAddressUpdate builder.
func (iau *IPAddressUpdate) Where(ps ...predicate.IPAddress) *IPAddressUpdate {
	iau.mutation.Where(ps...)
	return iau
}

// SetIP sets the "ip" field.
func (iau *IPAddressUpdate) SetIP(s string) *IPAddressUpdate {
	iau.mutation.SetIP(s)
	return iau
}

// SetNillableIP sets the "ip" field if the given value is not nil.
func (iau *IPAddressUpdate) SetNillableIP(s *string) *IPAddressUpdate {
	if s != nil {
		iau.SetIP(*s)
	}
	return iau
}

// SetMask sets the "mask" field.
func (iau *IPAddressUpdate) SetMask(s string) *IPAddressUpdate {
	iau.mutation.SetMask(s)
	return iau
}

// SetNillableMask sets the "mask" field if the given value is not nil.
func (iau *IPAddressUpdate) SetNillableMask(s *string) *IPAddressUpdate {
	if s != nil {
		iau.SetMask(*s)
	}
	return iau
}

// AddAsninfoIDs adds the "asninfo" edge to the ASNInfo entity by IDs.
func (iau *IPAddressUpdate) AddAsninfoIDs(ids ...int) *IPAddressUpdate {
	iau.mutation.AddAsninfoIDs(ids...)
	return iau
}

// AddAsninfo adds the "asninfo" edges to the ASNInfo entity.
func (iau *IPAddressUpdate) AddAsninfo(a ...*ASNInfo) *IPAddressUpdate {
	ids := make([]int, len(a))
	for i := range a {
		ids[i] = a[i].ID
	}
	return iau.AddAsninfoIDs(ids...)
}

// Mutation returns the IPAddressMutation object of the builder.
func (iau *IPAddressUpdate) Mutation() *IPAddressMutation {
	return iau.mutation
}

// ClearAsninfo clears all "asninfo" edges to the ASNInfo entity.
func (iau *IPAddressUpdate) ClearAsninfo() *IPAddressUpdate {
	iau.mutation.ClearAsninfo()
	return iau
}

// RemoveAsninfoIDs removes the "asninfo" edge to ASNInfo entities by IDs.
func (iau *IPAddressUpdate) RemoveAsninfoIDs(ids ...int) *IPAddressUpdate {
	iau.mutation.RemoveAsninfoIDs(ids...)
	return iau
}

// RemoveAsninfo removes "asninfo" edges to ASNInfo entities.
func (iau *IPAddressUpdate) RemoveAsninfo(a ...*ASNInfo) *IPAddressUpdate {
	ids := make([]int, len(a))
	for i := range a {
		ids[i] = a[i].ID
	}
	return iau.RemoveAsninfoIDs(ids...)
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (iau *IPAddressUpdate) Save(ctx context.Context) (int, error) {
	return withHooks(ctx, iau.sqlSave, iau.mutation, iau.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (iau *IPAddressUpdate) SaveX(ctx context.Context) int {
	affected, err := iau.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (iau *IPAddressUpdate) Exec(ctx context.Context) error {
	_, err := iau.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (iau *IPAddressUpdate) ExecX(ctx context.Context) {
	if err := iau.Exec(ctx); err != nil {
		panic(err)
	}
}

func (iau *IPAddressUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := sqlgraph.NewUpdateSpec(ipaddress.Table, ipaddress.Columns, sqlgraph.NewFieldSpec(ipaddress.FieldID, field.TypeInt))
	if ps := iau.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := iau.mutation.IP(); ok {
		_spec.SetField(ipaddress.FieldIP, field.TypeString, value)
	}
	if value, ok := iau.mutation.Mask(); ok {
		_spec.SetField(ipaddress.FieldMask, field.TypeString, value)
	}
	if iau.mutation.AsninfoCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   ipaddress.AsninfoTable,
			Columns: []string{ipaddress.AsninfoColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(asninfo.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := iau.mutation.RemovedAsninfoIDs(); len(nodes) > 0 && !iau.mutation.AsninfoCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   ipaddress.AsninfoTable,
			Columns: []string{ipaddress.AsninfoColumn},
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
	if nodes := iau.mutation.AsninfoIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   ipaddress.AsninfoTable,
			Columns: []string{ipaddress.AsninfoColumn},
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
	if n, err = sqlgraph.UpdateNodes(ctx, iau.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{ipaddress.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	iau.mutation.done = true
	return n, nil
}

// IPAddressUpdateOne is the builder for updating a single IPAddress entity.
type IPAddressUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *IPAddressMutation
}

// SetIP sets the "ip" field.
func (iauo *IPAddressUpdateOne) SetIP(s string) *IPAddressUpdateOne {
	iauo.mutation.SetIP(s)
	return iauo
}

// SetNillableIP sets the "ip" field if the given value is not nil.
func (iauo *IPAddressUpdateOne) SetNillableIP(s *string) *IPAddressUpdateOne {
	if s != nil {
		iauo.SetIP(*s)
	}
	return iauo
}

// SetMask sets the "mask" field.
func (iauo *IPAddressUpdateOne) SetMask(s string) *IPAddressUpdateOne {
	iauo.mutation.SetMask(s)
	return iauo
}

// SetNillableMask sets the "mask" field if the given value is not nil.
func (iauo *IPAddressUpdateOne) SetNillableMask(s *string) *IPAddressUpdateOne {
	if s != nil {
		iauo.SetMask(*s)
	}
	return iauo
}

// AddAsninfoIDs adds the "asninfo" edge to the ASNInfo entity by IDs.
func (iauo *IPAddressUpdateOne) AddAsninfoIDs(ids ...int) *IPAddressUpdateOne {
	iauo.mutation.AddAsninfoIDs(ids...)
	return iauo
}

// AddAsninfo adds the "asninfo" edges to the ASNInfo entity.
func (iauo *IPAddressUpdateOne) AddAsninfo(a ...*ASNInfo) *IPAddressUpdateOne {
	ids := make([]int, len(a))
	for i := range a {
		ids[i] = a[i].ID
	}
	return iauo.AddAsninfoIDs(ids...)
}

// Mutation returns the IPAddressMutation object of the builder.
func (iauo *IPAddressUpdateOne) Mutation() *IPAddressMutation {
	return iauo.mutation
}

// ClearAsninfo clears all "asninfo" edges to the ASNInfo entity.
func (iauo *IPAddressUpdateOne) ClearAsninfo() *IPAddressUpdateOne {
	iauo.mutation.ClearAsninfo()
	return iauo
}

// RemoveAsninfoIDs removes the "asninfo" edge to ASNInfo entities by IDs.
func (iauo *IPAddressUpdateOne) RemoveAsninfoIDs(ids ...int) *IPAddressUpdateOne {
	iauo.mutation.RemoveAsninfoIDs(ids...)
	return iauo
}

// RemoveAsninfo removes "asninfo" edges to ASNInfo entities.
func (iauo *IPAddressUpdateOne) RemoveAsninfo(a ...*ASNInfo) *IPAddressUpdateOne {
	ids := make([]int, len(a))
	for i := range a {
		ids[i] = a[i].ID
	}
	return iauo.RemoveAsninfoIDs(ids...)
}

// Where appends a list predicates to the IPAddressUpdate builder.
func (iauo *IPAddressUpdateOne) Where(ps ...predicate.IPAddress) *IPAddressUpdateOne {
	iauo.mutation.Where(ps...)
	return iauo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (iauo *IPAddressUpdateOne) Select(field string, fields ...string) *IPAddressUpdateOne {
	iauo.fields = append([]string{field}, fields...)
	return iauo
}

// Save executes the query and returns the updated IPAddress entity.
func (iauo *IPAddressUpdateOne) Save(ctx context.Context) (*IPAddress, error) {
	return withHooks(ctx, iauo.sqlSave, iauo.mutation, iauo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (iauo *IPAddressUpdateOne) SaveX(ctx context.Context) *IPAddress {
	node, err := iauo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (iauo *IPAddressUpdateOne) Exec(ctx context.Context) error {
	_, err := iauo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (iauo *IPAddressUpdateOne) ExecX(ctx context.Context) {
	if err := iauo.Exec(ctx); err != nil {
		panic(err)
	}
}

func (iauo *IPAddressUpdateOne) sqlSave(ctx context.Context) (_node *IPAddress, err error) {
	_spec := sqlgraph.NewUpdateSpec(ipaddress.Table, ipaddress.Columns, sqlgraph.NewFieldSpec(ipaddress.FieldID, field.TypeInt))
	id, ok := iauo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`model_ent: missing "IPAddress.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := iauo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, ipaddress.FieldID)
		for _, f := range fields {
			if !ipaddress.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("model_ent: invalid field %q for query", f)}
			}
			if f != ipaddress.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := iauo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := iauo.mutation.IP(); ok {
		_spec.SetField(ipaddress.FieldIP, field.TypeString, value)
	}
	if value, ok := iauo.mutation.Mask(); ok {
		_spec.SetField(ipaddress.FieldMask, field.TypeString, value)
	}
	if iauo.mutation.AsninfoCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   ipaddress.AsninfoTable,
			Columns: []string{ipaddress.AsninfoColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(asninfo.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := iauo.mutation.RemovedAsninfoIDs(); len(nodes) > 0 && !iauo.mutation.AsninfoCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   ipaddress.AsninfoTable,
			Columns: []string{ipaddress.AsninfoColumn},
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
	if nodes := iauo.mutation.AsninfoIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   ipaddress.AsninfoTable,
			Columns: []string{ipaddress.AsninfoColumn},
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
	_node = &IPAddress{config: iauo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, iauo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{ipaddress.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	iauo.mutation.done = true
	return _node, nil
}