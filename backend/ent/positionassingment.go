// Code generated by entc, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"
	"time"

	"github.com/facebookincubator/ent/dialect/sql"
	"github.com/konrawitAEK/app/ent/department"
	"github.com/konrawitAEK/app/ent/physician"
	"github.com/konrawitAEK/app/ent/position"
	"github.com/konrawitAEK/app/ent/positionassingment"
)

// Positionassingment is the model entity for the Positionassingment schema.
type Positionassingment struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// DayStart holds the value of the "DayStart" field.
	DayStart time.Time `json:"DayStart,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the PositionassingmentQuery when eager-loading is set.
	Edges         PositionassingmentEdges `json:"edges"`
	DEPARTMENT_ID *int
	PHYSICIAN_ID  *int
	POSITION_ID   *int
}

// PositionassingmentEdges holds the relations/edges for other nodes in the graph.
type PositionassingmentEdges struct {
	// User holds the value of the user edge.
	User *Physician
	// Department holds the value of the department edge.
	Department *Department
	// Position holds the value of the position edge.
	Position *Position
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [3]bool
}

// UserOrErr returns the User value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e PositionassingmentEdges) UserOrErr() (*Physician, error) {
	if e.loadedTypes[0] {
		if e.User == nil {
			// The edge user was loaded in eager-loading,
			// but was not found.
			return nil, &NotFoundError{label: physician.Label}
		}
		return e.User, nil
	}
	return nil, &NotLoadedError{edge: "user"}
}

// DepartmentOrErr returns the Department value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e PositionassingmentEdges) DepartmentOrErr() (*Department, error) {
	if e.loadedTypes[1] {
		if e.Department == nil {
			// The edge department was loaded in eager-loading,
			// but was not found.
			return nil, &NotFoundError{label: department.Label}
		}
		return e.Department, nil
	}
	return nil, &NotLoadedError{edge: "department"}
}

// PositionOrErr returns the Position value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e PositionassingmentEdges) PositionOrErr() (*Position, error) {
	if e.loadedTypes[2] {
		if e.Position == nil {
			// The edge position was loaded in eager-loading,
			// but was not found.
			return nil, &NotFoundError{label: position.Label}
		}
		return e.Position, nil
	}
	return nil, &NotLoadedError{edge: "position"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Positionassingment) scanValues() []interface{} {
	return []interface{}{
		&sql.NullInt64{}, // id
		&sql.NullTime{},  // DayStart
	}
}

// fkValues returns the types for scanning foreign-keys values from sql.Rows.
func (*Positionassingment) fkValues() []interface{} {
	return []interface{}{
		&sql.NullInt64{}, // DEPARTMENT_ID
		&sql.NullInt64{}, // PHYSICIAN_ID
		&sql.NullInt64{}, // POSITION_ID
	}
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Positionassingment fields.
func (po *Positionassingment) assignValues(values ...interface{}) error {
	if m, n := len(values), len(positionassingment.Columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	value, ok := values[0].(*sql.NullInt64)
	if !ok {
		return fmt.Errorf("unexpected type %T for field id", value)
	}
	po.ID = int(value.Int64)
	values = values[1:]
	if value, ok := values[0].(*sql.NullTime); !ok {
		return fmt.Errorf("unexpected type %T for field DayStart", values[0])
	} else if value.Valid {
		po.DayStart = value.Time
	}
	values = values[1:]
	if len(values) == len(positionassingment.ForeignKeys) {
		if value, ok := values[0].(*sql.NullInt64); !ok {
			return fmt.Errorf("unexpected type %T for edge-field DEPARTMENT_ID", value)
		} else if value.Valid {
			po.DEPARTMENT_ID = new(int)
			*po.DEPARTMENT_ID = int(value.Int64)
		}
		if value, ok := values[1].(*sql.NullInt64); !ok {
			return fmt.Errorf("unexpected type %T for edge-field PHYSICIAN_ID", value)
		} else if value.Valid {
			po.PHYSICIAN_ID = new(int)
			*po.PHYSICIAN_ID = int(value.Int64)
		}
		if value, ok := values[2].(*sql.NullInt64); !ok {
			return fmt.Errorf("unexpected type %T for edge-field POSITION_ID", value)
		} else if value.Valid {
			po.POSITION_ID = new(int)
			*po.POSITION_ID = int(value.Int64)
		}
	}
	return nil
}

// QueryUser queries the user edge of the Positionassingment.
func (po *Positionassingment) QueryUser() *PhysicianQuery {
	return (&PositionassingmentClient{config: po.config}).QueryUser(po)
}

// QueryDepartment queries the department edge of the Positionassingment.
func (po *Positionassingment) QueryDepartment() *DepartmentQuery {
	return (&PositionassingmentClient{config: po.config}).QueryDepartment(po)
}

// QueryPosition queries the position edge of the Positionassingment.
func (po *Positionassingment) QueryPosition() *PositionQuery {
	return (&PositionassingmentClient{config: po.config}).QueryPosition(po)
}

// Update returns a builder for updating this Positionassingment.
// Note that, you need to call Positionassingment.Unwrap() before calling this method, if this Positionassingment
// was returned from a transaction, and the transaction was committed or rolled back.
func (po *Positionassingment) Update() *PositionassingmentUpdateOne {
	return (&PositionassingmentClient{config: po.config}).UpdateOne(po)
}

// Unwrap unwraps the entity that was returned from a transaction after it was closed,
// so that all next queries will be executed through the driver which created the transaction.
func (po *Positionassingment) Unwrap() *Positionassingment {
	tx, ok := po.config.driver.(*txDriver)
	if !ok {
		panic("ent: Positionassingment is not a transactional entity")
	}
	po.config.driver = tx.drv
	return po
}

// String implements the fmt.Stringer.
func (po *Positionassingment) String() string {
	var builder strings.Builder
	builder.WriteString("Positionassingment(")
	builder.WriteString(fmt.Sprintf("id=%v", po.ID))
	builder.WriteString(", DayStart=")
	builder.WriteString(po.DayStart.Format(time.ANSIC))
	builder.WriteByte(')')
	return builder.String()
}

// Positionassingments is a parsable slice of Positionassingment.
type Positionassingments []*Positionassingment

func (po Positionassingments) config(cfg config) {
	for _i := range po {
		po[_i].config = cfg
	}
}
