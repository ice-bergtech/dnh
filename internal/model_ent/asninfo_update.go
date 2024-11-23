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
	"github.com/ice-bergtech/dnh/src/internal/model_ent/predicate"
)

// ASNInfoUpdate is the builder for updating ASNInfo entities.
type ASNInfoUpdate struct {
	config
	hooks    []Hook
	mutation *ASNInfoMutation
}

// Where appends a list predicates to the ASNInfoUpdate builder.
func (aiu *ASNInfoUpdate) Where(ps ...predicate.ASNInfo) *ASNInfoUpdate {
	aiu.mutation.Where(ps...)
	return aiu
}

// SetAsn sets the "asn" field.
func (aiu *ASNInfoUpdate) SetAsn(i int) *ASNInfoUpdate {
	aiu.mutation.ResetAsn()
	aiu.mutation.SetAsn(i)
	return aiu
}

// SetNillableAsn sets the "asn" field if the given value is not nil.
func (aiu *ASNInfoUpdate) SetNillableAsn(i *int) *ASNInfoUpdate {
	if i != nil {
		aiu.SetAsn(*i)
	}
	return aiu
}

// AddAsn adds i to the "asn" field.
func (aiu *ASNInfoUpdate) AddAsn(i int) *ASNInfoUpdate {
	aiu.mutation.AddAsn(i)
	return aiu
}

// SetCountry sets the "country" field.
func (aiu *ASNInfoUpdate) SetCountry(s string) *ASNInfoUpdate {
	aiu.mutation.SetCountry(s)
	return aiu
}

// SetNillableCountry sets the "country" field if the given value is not nil.
func (aiu *ASNInfoUpdate) SetNillableCountry(s *string) *ASNInfoUpdate {
	if s != nil {
		aiu.SetCountry(*s)
	}
	return aiu
}

// SetRegistry sets the "registry" field.
func (aiu *ASNInfoUpdate) SetRegistry(s string) *ASNInfoUpdate {
	aiu.mutation.SetRegistry(s)
	return aiu
}

// SetNillableRegistry sets the "registry" field if the given value is not nil.
func (aiu *ASNInfoUpdate) SetNillableRegistry(s *string) *ASNInfoUpdate {
	if s != nil {
		aiu.SetRegistry(*s)
	}
	return aiu
}

// Mutation returns the ASNInfoMutation object of the builder.
func (aiu *ASNInfoUpdate) Mutation() *ASNInfoMutation {
	return aiu.mutation
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (aiu *ASNInfoUpdate) Save(ctx context.Context) (int, error) {
	return withHooks(ctx, aiu.sqlSave, aiu.mutation, aiu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (aiu *ASNInfoUpdate) SaveX(ctx context.Context) int {
	affected, err := aiu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (aiu *ASNInfoUpdate) Exec(ctx context.Context) error {
	_, err := aiu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (aiu *ASNInfoUpdate) ExecX(ctx context.Context) {
	if err := aiu.Exec(ctx); err != nil {
		panic(err)
	}
}

func (aiu *ASNInfoUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := sqlgraph.NewUpdateSpec(asninfo.Table, asninfo.Columns, sqlgraph.NewFieldSpec(asninfo.FieldID, field.TypeInt))
	if ps := aiu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := aiu.mutation.Asn(); ok {
		_spec.SetField(asninfo.FieldAsn, field.TypeInt, value)
	}
	if value, ok := aiu.mutation.AddedAsn(); ok {
		_spec.AddField(asninfo.FieldAsn, field.TypeInt, value)
	}
	if value, ok := aiu.mutation.Country(); ok {
		_spec.SetField(asninfo.FieldCountry, field.TypeString, value)
	}
	if value, ok := aiu.mutation.Registry(); ok {
		_spec.SetField(asninfo.FieldRegistry, field.TypeString, value)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, aiu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{asninfo.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	aiu.mutation.done = true
	return n, nil
}

// ASNInfoUpdateOne is the builder for updating a single ASNInfo entity.
type ASNInfoUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *ASNInfoMutation
}

// SetAsn sets the "asn" field.
func (aiuo *ASNInfoUpdateOne) SetAsn(i int) *ASNInfoUpdateOne {
	aiuo.mutation.ResetAsn()
	aiuo.mutation.SetAsn(i)
	return aiuo
}

// SetNillableAsn sets the "asn" field if the given value is not nil.
func (aiuo *ASNInfoUpdateOne) SetNillableAsn(i *int) *ASNInfoUpdateOne {
	if i != nil {
		aiuo.SetAsn(*i)
	}
	return aiuo
}

// AddAsn adds i to the "asn" field.
func (aiuo *ASNInfoUpdateOne) AddAsn(i int) *ASNInfoUpdateOne {
	aiuo.mutation.AddAsn(i)
	return aiuo
}

// SetCountry sets the "country" field.
func (aiuo *ASNInfoUpdateOne) SetCountry(s string) *ASNInfoUpdateOne {
	aiuo.mutation.SetCountry(s)
	return aiuo
}

// SetNillableCountry sets the "country" field if the given value is not nil.
func (aiuo *ASNInfoUpdateOne) SetNillableCountry(s *string) *ASNInfoUpdateOne {
	if s != nil {
		aiuo.SetCountry(*s)
	}
	return aiuo
}

// SetRegistry sets the "registry" field.
func (aiuo *ASNInfoUpdateOne) SetRegistry(s string) *ASNInfoUpdateOne {
	aiuo.mutation.SetRegistry(s)
	return aiuo
}

// SetNillableRegistry sets the "registry" field if the given value is not nil.
func (aiuo *ASNInfoUpdateOne) SetNillableRegistry(s *string) *ASNInfoUpdateOne {
	if s != nil {
		aiuo.SetRegistry(*s)
	}
	return aiuo
}

// Mutation returns the ASNInfoMutation object of the builder.
func (aiuo *ASNInfoUpdateOne) Mutation() *ASNInfoMutation {
	return aiuo.mutation
}

// Where appends a list predicates to the ASNInfoUpdate builder.
func (aiuo *ASNInfoUpdateOne) Where(ps ...predicate.ASNInfo) *ASNInfoUpdateOne {
	aiuo.mutation.Where(ps...)
	return aiuo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (aiuo *ASNInfoUpdateOne) Select(field string, fields ...string) *ASNInfoUpdateOne {
	aiuo.fields = append([]string{field}, fields...)
	return aiuo
}

// Save executes the query and returns the updated ASNInfo entity.
func (aiuo *ASNInfoUpdateOne) Save(ctx context.Context) (*ASNInfo, error) {
	return withHooks(ctx, aiuo.sqlSave, aiuo.mutation, aiuo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (aiuo *ASNInfoUpdateOne) SaveX(ctx context.Context) *ASNInfo {
	node, err := aiuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (aiuo *ASNInfoUpdateOne) Exec(ctx context.Context) error {
	_, err := aiuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (aiuo *ASNInfoUpdateOne) ExecX(ctx context.Context) {
	if err := aiuo.Exec(ctx); err != nil {
		panic(err)
	}
}

func (aiuo *ASNInfoUpdateOne) sqlSave(ctx context.Context) (_node *ASNInfo, err error) {
	_spec := sqlgraph.NewUpdateSpec(asninfo.Table, asninfo.Columns, sqlgraph.NewFieldSpec(asninfo.FieldID, field.TypeInt))
	id, ok := aiuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`model_ent: missing "ASNInfo.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := aiuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, asninfo.FieldID)
		for _, f := range fields {
			if !asninfo.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("model_ent: invalid field %q for query", f)}
			}
			if f != asninfo.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := aiuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := aiuo.mutation.Asn(); ok {
		_spec.SetField(asninfo.FieldAsn, field.TypeInt, value)
	}
	if value, ok := aiuo.mutation.AddedAsn(); ok {
		_spec.AddField(asninfo.FieldAsn, field.TypeInt, value)
	}
	if value, ok := aiuo.mutation.Country(); ok {
		_spec.SetField(asninfo.FieldCountry, field.TypeString, value)
	}
	if value, ok := aiuo.mutation.Registry(); ok {
		_spec.SetField(asninfo.FieldRegistry, field.TypeString, value)
	}
	_node = &ASNInfo{config: aiuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, aiuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{asninfo.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	aiuo.mutation.done = true
	return _node, nil
}
