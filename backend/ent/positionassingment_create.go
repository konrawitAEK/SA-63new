// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/facebookincubator/ent/dialect/sql/sqlgraph"
	"github.com/facebookincubator/ent/schema/field"
	"github.com/konrawitAEK/app/ent/department"
	"github.com/konrawitAEK/app/ent/physician"
	"github.com/konrawitAEK/app/ent/position"
	"github.com/konrawitAEK/app/ent/positionassingment"
)

// PositionassingmentCreate is the builder for creating a Positionassingment entity.
type PositionassingmentCreate struct {
	config
	mutation *PositionassingmentMutation
	hooks    []Hook
}

// SetDayStart sets the DayStart field.
func (pc *PositionassingmentCreate) SetDayStart(t time.Time) *PositionassingmentCreate {
	pc.mutation.SetDayStart(t)
	return pc
}

// SetUserID sets the user edge to Physician by id.
func (pc *PositionassingmentCreate) SetUserID(id int) *PositionassingmentCreate {
	pc.mutation.SetUserID(id)
	return pc
}

// SetNillableUserID sets the user edge to Physician by id if the given value is not nil.
func (pc *PositionassingmentCreate) SetNillableUserID(id *int) *PositionassingmentCreate {
	if id != nil {
		pc = pc.SetUserID(*id)
	}
	return pc
}

// SetUser sets the user edge to Physician.
func (pc *PositionassingmentCreate) SetUser(p *Physician) *PositionassingmentCreate {
	return pc.SetUserID(p.ID)
}

// SetDepartmentID sets the department edge to Department by id.
func (pc *PositionassingmentCreate) SetDepartmentID(id int) *PositionassingmentCreate {
	pc.mutation.SetDepartmentID(id)
	return pc
}

// SetNillableDepartmentID sets the department edge to Department by id if the given value is not nil.
func (pc *PositionassingmentCreate) SetNillableDepartmentID(id *int) *PositionassingmentCreate {
	if id != nil {
		pc = pc.SetDepartmentID(*id)
	}
	return pc
}

// SetDepartment sets the department edge to Department.
func (pc *PositionassingmentCreate) SetDepartment(d *Department) *PositionassingmentCreate {
	return pc.SetDepartmentID(d.ID)
}

// SetPositionID sets the position edge to Position by id.
func (pc *PositionassingmentCreate) SetPositionID(id int) *PositionassingmentCreate {
	pc.mutation.SetPositionID(id)
	return pc
}

// SetNillablePositionID sets the position edge to Position by id if the given value is not nil.
func (pc *PositionassingmentCreate) SetNillablePositionID(id *int) *PositionassingmentCreate {
	if id != nil {
		pc = pc.SetPositionID(*id)
	}
	return pc
}

// SetPosition sets the position edge to Position.
func (pc *PositionassingmentCreate) SetPosition(p *Position) *PositionassingmentCreate {
	return pc.SetPositionID(p.ID)
}

// Mutation returns the PositionassingmentMutation object of the builder.
func (pc *PositionassingmentCreate) Mutation() *PositionassingmentMutation {
	return pc.mutation
}

// Save creates the Positionassingment in the database.
func (pc *PositionassingmentCreate) Save(ctx context.Context) (*Positionassingment, error) {
	if _, ok := pc.mutation.DayStart(); !ok {
		return nil, &ValidationError{Name: "DayStart", err: errors.New("ent: missing required field \"DayStart\"")}
	}
	var (
		err  error
		node *Positionassingment
	)
	if len(pc.hooks) == 0 {
		node, err = pc.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*PositionassingmentMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			pc.mutation = mutation
			node, err = pc.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(pc.hooks) - 1; i >= 0; i-- {
			mut = pc.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, pc.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX calls Save and panics if Save returns an error.
func (pc *PositionassingmentCreate) SaveX(ctx context.Context) *Positionassingment {
	v, err := pc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (pc *PositionassingmentCreate) sqlSave(ctx context.Context) (*Positionassingment, error) {
	po, _spec := pc.createSpec()
	if err := sqlgraph.CreateNode(ctx, pc.driver, _spec); err != nil {
		if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	po.ID = int(id)
	return po, nil
}

func (pc *PositionassingmentCreate) createSpec() (*Positionassingment, *sqlgraph.CreateSpec) {
	var (
		po    = &Positionassingment{config: pc.config}
		_spec = &sqlgraph.CreateSpec{
			Table: positionassingment.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: positionassingment.FieldID,
			},
		}
	)
	if value, ok := pc.mutation.DayStart(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: positionassingment.FieldDayStart,
		})
		po.DayStart = value
	}
	if nodes := pc.mutation.UserIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: true,
			Table:   positionassingment.UserTable,
			Columns: []string{positionassingment.UserColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: physician.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := pc.mutation.DepartmentIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   positionassingment.DepartmentTable,
			Columns: []string{positionassingment.DepartmentColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: department.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := pc.mutation.PositionIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   positionassingment.PositionTable,
			Columns: []string{positionassingment.PositionColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: position.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	return po, _spec
}
