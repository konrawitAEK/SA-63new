// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"

	"github.com/facebookincubator/ent/dialect/sql"
	"github.com/facebookincubator/ent/dialect/sql/sqlgraph"
	"github.com/facebookincubator/ent/schema/field"
	"github.com/konrawitAEK/app/ent/department"
	"github.com/konrawitAEK/app/ent/positionassingment"
	"github.com/konrawitAEK/app/ent/predicate"
)

// DepartmentUpdate is the builder for updating Department entities.
type DepartmentUpdate struct {
	config
	hooks      []Hook
	mutation   *DepartmentMutation
	predicates []predicate.Department
}

// Where adds a new predicate for the builder.
func (du *DepartmentUpdate) Where(ps ...predicate.Department) *DepartmentUpdate {
	du.predicates = append(du.predicates, ps...)
	return du
}

// SetDepartmentname sets the Departmentname field.
func (du *DepartmentUpdate) SetDepartmentname(s string) *DepartmentUpdate {
	du.mutation.SetDepartmentname(s)
	return du
}

// AddFormdepartmentIDs adds the formdepartment edge to Positionassingment by ids.
func (du *DepartmentUpdate) AddFormdepartmentIDs(ids ...int) *DepartmentUpdate {
	du.mutation.AddFormdepartmentIDs(ids...)
	return du
}

// AddFormdepartment adds the formdepartment edges to Positionassingment.
func (du *DepartmentUpdate) AddFormdepartment(p ...*Positionassingment) *DepartmentUpdate {
	ids := make([]int, len(p))
	for i := range p {
		ids[i] = p[i].ID
	}
	return du.AddFormdepartmentIDs(ids...)
}

// Mutation returns the DepartmentMutation object of the builder.
func (du *DepartmentUpdate) Mutation() *DepartmentMutation {
	return du.mutation
}

// RemoveFormdepartmentIDs removes the formdepartment edge to Positionassingment by ids.
func (du *DepartmentUpdate) RemoveFormdepartmentIDs(ids ...int) *DepartmentUpdate {
	du.mutation.RemoveFormdepartmentIDs(ids...)
	return du
}

// RemoveFormdepartment removes formdepartment edges to Positionassingment.
func (du *DepartmentUpdate) RemoveFormdepartment(p ...*Positionassingment) *DepartmentUpdate {
	ids := make([]int, len(p))
	for i := range p {
		ids[i] = p[i].ID
	}
	return du.RemoveFormdepartmentIDs(ids...)
}

// Save executes the query and returns the number of rows/vertices matched by this operation.
func (du *DepartmentUpdate) Save(ctx context.Context) (int, error) {

	var (
		err      error
		affected int
	)
	if len(du.hooks) == 0 {
		affected, err = du.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*DepartmentMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			du.mutation = mutation
			affected, err = du.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(du.hooks) - 1; i >= 0; i-- {
			mut = du.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, du.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (du *DepartmentUpdate) SaveX(ctx context.Context) int {
	affected, err := du.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (du *DepartmentUpdate) Exec(ctx context.Context) error {
	_, err := du.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (du *DepartmentUpdate) ExecX(ctx context.Context) {
	if err := du.Exec(ctx); err != nil {
		panic(err)
	}
}

func (du *DepartmentUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   department.Table,
			Columns: department.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: department.FieldID,
			},
		},
	}
	if ps := du.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := du.mutation.Departmentname(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: department.FieldDepartmentname,
		})
	}
	if nodes := du.mutation.RemovedFormdepartmentIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   department.FormdepartmentTable,
			Columns: []string{department.FormdepartmentColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: positionassingment.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := du.mutation.FormdepartmentIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   department.FormdepartmentTable,
			Columns: []string{department.FormdepartmentColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: positionassingment.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, du.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{department.Label}
		} else if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return 0, err
	}
	return n, nil
}

// DepartmentUpdateOne is the builder for updating a single Department entity.
type DepartmentUpdateOne struct {
	config
	hooks    []Hook
	mutation *DepartmentMutation
}

// SetDepartmentname sets the Departmentname field.
func (duo *DepartmentUpdateOne) SetDepartmentname(s string) *DepartmentUpdateOne {
	duo.mutation.SetDepartmentname(s)
	return duo
}

// AddFormdepartmentIDs adds the formdepartment edge to Positionassingment by ids.
func (duo *DepartmentUpdateOne) AddFormdepartmentIDs(ids ...int) *DepartmentUpdateOne {
	duo.mutation.AddFormdepartmentIDs(ids...)
	return duo
}

// AddFormdepartment adds the formdepartment edges to Positionassingment.
func (duo *DepartmentUpdateOne) AddFormdepartment(p ...*Positionassingment) *DepartmentUpdateOne {
	ids := make([]int, len(p))
	for i := range p {
		ids[i] = p[i].ID
	}
	return duo.AddFormdepartmentIDs(ids...)
}

// Mutation returns the DepartmentMutation object of the builder.
func (duo *DepartmentUpdateOne) Mutation() *DepartmentMutation {
	return duo.mutation
}

// RemoveFormdepartmentIDs removes the formdepartment edge to Positionassingment by ids.
func (duo *DepartmentUpdateOne) RemoveFormdepartmentIDs(ids ...int) *DepartmentUpdateOne {
	duo.mutation.RemoveFormdepartmentIDs(ids...)
	return duo
}

// RemoveFormdepartment removes formdepartment edges to Positionassingment.
func (duo *DepartmentUpdateOne) RemoveFormdepartment(p ...*Positionassingment) *DepartmentUpdateOne {
	ids := make([]int, len(p))
	for i := range p {
		ids[i] = p[i].ID
	}
	return duo.RemoveFormdepartmentIDs(ids...)
}

// Save executes the query and returns the updated entity.
func (duo *DepartmentUpdateOne) Save(ctx context.Context) (*Department, error) {

	var (
		err  error
		node *Department
	)
	if len(duo.hooks) == 0 {
		node, err = duo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*DepartmentMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			duo.mutation = mutation
			node, err = duo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(duo.hooks) - 1; i >= 0; i-- {
			mut = duo.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, duo.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (duo *DepartmentUpdateOne) SaveX(ctx context.Context) *Department {
	d, err := duo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return d
}

// Exec executes the query on the entity.
func (duo *DepartmentUpdateOne) Exec(ctx context.Context) error {
	_, err := duo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (duo *DepartmentUpdateOne) ExecX(ctx context.Context) {
	if err := duo.Exec(ctx); err != nil {
		panic(err)
	}
}

func (duo *DepartmentUpdateOne) sqlSave(ctx context.Context) (d *Department, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   department.Table,
			Columns: department.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: department.FieldID,
			},
		},
	}
	id, ok := duo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "ID", err: fmt.Errorf("missing Department.ID for update")}
	}
	_spec.Node.ID.Value = id
	if value, ok := duo.mutation.Departmentname(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: department.FieldDepartmentname,
		})
	}
	if nodes := duo.mutation.RemovedFormdepartmentIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   department.FormdepartmentTable,
			Columns: []string{department.FormdepartmentColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: positionassingment.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := duo.mutation.FormdepartmentIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   department.FormdepartmentTable,
			Columns: []string{department.FormdepartmentColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: positionassingment.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	d = &Department{config: duo.config}
	_spec.Assign = d.assignValues
	_spec.ScanValues = d.scanValues()
	if err = sqlgraph.UpdateNode(ctx, duo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{department.Label}
		} else if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return nil, err
	}
	return d, nil
}
