// Code generated by ent, DO NOT EDIT.

package model_ent

import (
	"context"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/ice-bergtech/dnh/src/internal/model_ent/nameserver"
	"github.com/ice-bergtech/dnh/src/internal/model_ent/predicate"
)

// NameserverDelete is the builder for deleting a Nameserver entity.
type NameserverDelete struct {
	config
	hooks    []Hook
	mutation *NameserverMutation
}

// Where appends a list predicates to the NameserverDelete builder.
func (nd *NameserverDelete) Where(ps ...predicate.Nameserver) *NameserverDelete {
	nd.mutation.Where(ps...)
	return nd
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (nd *NameserverDelete) Exec(ctx context.Context) (int, error) {
	return withHooks(ctx, nd.sqlExec, nd.mutation, nd.hooks)
}

// ExecX is like Exec, but panics if an error occurs.
func (nd *NameserverDelete) ExecX(ctx context.Context) int {
	n, err := nd.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (nd *NameserverDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := sqlgraph.NewDeleteSpec(nameserver.Table, sqlgraph.NewFieldSpec(nameserver.FieldID, field.TypeInt))
	if ps := nd.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	affected, err := sqlgraph.DeleteNodes(ctx, nd.driver, _spec)
	if err != nil && sqlgraph.IsConstraintError(err) {
		err = &ConstraintError{msg: err.Error(), wrap: err}
	}
	nd.mutation.done = true
	return affected, err
}

// NameserverDeleteOne is the builder for deleting a single Nameserver entity.
type NameserverDeleteOne struct {
	nd *NameserverDelete
}

// Where appends a list predicates to the NameserverDelete builder.
func (ndo *NameserverDeleteOne) Where(ps ...predicate.Nameserver) *NameserverDeleteOne {
	ndo.nd.mutation.Where(ps...)
	return ndo
}

// Exec executes the deletion query.
func (ndo *NameserverDeleteOne) Exec(ctx context.Context) error {
	n, err := ndo.nd.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{nameserver.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (ndo *NameserverDeleteOne) ExecX(ctx context.Context) {
	if err := ndo.Exec(ctx); err != nil {
		panic(err)
	}
}