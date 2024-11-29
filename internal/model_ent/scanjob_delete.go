// Code generated by ent, DO NOT EDIT.

package model_ent

import (
	"context"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/ice-bergtech/dnh/src/internal/model_ent/predicate"
	"github.com/ice-bergtech/dnh/src/internal/model_ent/scanjob"
)

// ScanJobDelete is the builder for deleting a ScanJob entity.
type ScanJobDelete struct {
	config
	hooks    []Hook
	mutation *ScanJobMutation
}

// Where appends a list predicates to the ScanJobDelete builder.
func (sjd *ScanJobDelete) Where(ps ...predicate.ScanJob) *ScanJobDelete {
	sjd.mutation.Where(ps...)
	return sjd
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (sjd *ScanJobDelete) Exec(ctx context.Context) (int, error) {
	return withHooks(ctx, sjd.sqlExec, sjd.mutation, sjd.hooks)
}

// ExecX is like Exec, but panics if an error occurs.
func (sjd *ScanJobDelete) ExecX(ctx context.Context) int {
	n, err := sjd.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (sjd *ScanJobDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := sqlgraph.NewDeleteSpec(scanjob.Table, sqlgraph.NewFieldSpec(scanjob.FieldID, field.TypeInt))
	if ps := sjd.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	affected, err := sqlgraph.DeleteNodes(ctx, sjd.driver, _spec)
	if err != nil && sqlgraph.IsConstraintError(err) {
		err = &ConstraintError{msg: err.Error(), wrap: err}
	}
	sjd.mutation.done = true
	return affected, err
}

// ScanJobDeleteOne is the builder for deleting a single ScanJob entity.
type ScanJobDeleteOne struct {
	sjd *ScanJobDelete
}

// Where appends a list predicates to the ScanJobDelete builder.
func (sjdo *ScanJobDeleteOne) Where(ps ...predicate.ScanJob) *ScanJobDeleteOne {
	sjdo.sjd.mutation.Where(ps...)
	return sjdo
}

// Exec executes the deletion query.
func (sjdo *ScanJobDeleteOne) Exec(ctx context.Context) error {
	n, err := sjdo.sjd.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{scanjob.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (sjdo *ScanJobDeleteOne) ExecX(ctx context.Context) {
	if err := sjdo.Exec(ctx); err != nil {
		panic(err)
	}
}