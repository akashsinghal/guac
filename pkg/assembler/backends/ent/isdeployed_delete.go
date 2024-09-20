// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/guacsec/guac/pkg/assembler/backends/ent/isdeployed"
	"github.com/guacsec/guac/pkg/assembler/backends/ent/predicate"
)

// IsDeployedDelete is the builder for deleting a IsDeployed entity.
type IsDeployedDelete struct {
	config
	hooks    []Hook
	mutation *IsDeployedMutation
}

// Where appends a list predicates to the IsDeployedDelete builder.
func (idd *IsDeployedDelete) Where(ps ...predicate.IsDeployed) *IsDeployedDelete {
	idd.mutation.Where(ps...)
	return idd
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (idd *IsDeployedDelete) Exec(ctx context.Context) (int, error) {
	return withHooks(ctx, idd.sqlExec, idd.mutation, idd.hooks)
}

// ExecX is like Exec, but panics if an error occurs.
func (idd *IsDeployedDelete) ExecX(ctx context.Context) int {
	n, err := idd.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (idd *IsDeployedDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := sqlgraph.NewDeleteSpec(isdeployed.Table, sqlgraph.NewFieldSpec(isdeployed.FieldID, field.TypeUUID))
	if ps := idd.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	affected, err := sqlgraph.DeleteNodes(ctx, idd.driver, _spec)
	if err != nil && sqlgraph.IsConstraintError(err) {
		err = &ConstraintError{msg: err.Error(), wrap: err}
	}
	idd.mutation.done = true
	return affected, err
}

// IsDeployedDeleteOne is the builder for deleting a single IsDeployed entity.
type IsDeployedDeleteOne struct {
	idd *IsDeployedDelete
}

// Where appends a list predicates to the IsDeployedDelete builder.
func (iddo *IsDeployedDeleteOne) Where(ps ...predicate.IsDeployed) *IsDeployedDeleteOne {
	iddo.idd.mutation.Where(ps...)
	return iddo
}

// Exec executes the deletion query.
func (iddo *IsDeployedDeleteOne) Exec(ctx context.Context) error {
	n, err := iddo.idd.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{isdeployed.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (iddo *IsDeployedDeleteOne) ExecX(ctx context.Context) {
	if err := iddo.Exec(ctx); err != nil {
		panic(err)
	}
}
