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
	"github.com/ice-bergtech/dnh/src/internal/model_ent/predicate"
	"github.com/ice-bergtech/dnh/src/internal/model_ent/registrar"
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

// Mutation returns the RegistrarMutation object of the builder.
func (ru *RegistrarUpdate) Mutation() *RegistrarMutation {
	return ru.mutation
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
	if value, ok := ru.mutation.TimeFirst(); ok {
		_spec.SetField(registrar.FieldTimeFirst, field.TypeTime, value)
	}
	if value, ok := ru.mutation.TimeLast(); ok {
		_spec.SetField(registrar.FieldTimeLast, field.TypeTime, value)
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

// Mutation returns the RegistrarMutation object of the builder.
func (ruo *RegistrarUpdateOne) Mutation() *RegistrarMutation {
	return ruo.mutation
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
	if value, ok := ruo.mutation.TimeFirst(); ok {
		_spec.SetField(registrar.FieldTimeFirst, field.TypeTime, value)
	}
	if value, ok := ruo.mutation.TimeLast(); ok {
		_spec.SetField(registrar.FieldTimeLast, field.TypeTime, value)
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