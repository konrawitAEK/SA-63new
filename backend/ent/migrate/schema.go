// Code generated by entc, DO NOT EDIT.

package migrate

import (
	"github.com/facebookincubator/ent/dialect/sql/schema"
	"github.com/facebookincubator/ent/schema/field"
)

var (
	// DepartmentsColumns holds the columns for the "departments" table.
	DepartmentsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "departmentname", Type: field.TypeString},
	}
	// DepartmentsTable holds the schema information for the "departments" table.
	DepartmentsTable = &schema.Table{
		Name:        "departments",
		Columns:     DepartmentsColumns,
		PrimaryKey:  []*schema.Column{DepartmentsColumns[0]},
		ForeignKeys: []*schema.ForeignKey{},
	}
	// PhysiciansColumns holds the columns for the "physicians" table.
	PhysiciansColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "name", Type: field.TypeString},
		{Name: "email", Type: field.TypeString, Unique: true},
	}
	// PhysiciansTable holds the schema information for the "physicians" table.
	PhysiciansTable = &schema.Table{
		Name:        "physicians",
		Columns:     PhysiciansColumns,
		PrimaryKey:  []*schema.Column{PhysiciansColumns[0]},
		ForeignKeys: []*schema.ForeignKey{},
	}
	// PositionsColumns holds the columns for the "positions" table.
	PositionsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "nameposition", Type: field.TypeString},
	}
	// PositionsTable holds the schema information for the "positions" table.
	PositionsTable = &schema.Table{
		Name:        "positions",
		Columns:     PositionsColumns,
		PrimaryKey:  []*schema.Column{PositionsColumns[0]},
		ForeignKeys: []*schema.ForeignKey{},
	}
	// PositionassingmentsColumns holds the columns for the "positionassingments" table.
	PositionassingmentsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "day_start", Type: field.TypeTime},
		{Name: "DEPARTMENT_ID", Type: field.TypeInt, Nullable: true},
		{Name: "PHYSICIAN_ID", Type: field.TypeInt, Unique: true, Nullable: true},
		{Name: "POSITION_ID", Type: field.TypeInt, Nullable: true},
	}
	// PositionassingmentsTable holds the schema information for the "positionassingments" table.
	PositionassingmentsTable = &schema.Table{
		Name:       "positionassingments",
		Columns:    PositionassingmentsColumns,
		PrimaryKey: []*schema.Column{PositionassingmentsColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:  "positionassingments_departments_formdepartment",
				Columns: []*schema.Column{PositionassingmentsColumns[2]},

				RefColumns: []*schema.Column{DepartmentsColumns[0]},
				OnDelete:   schema.SetNull,
			},
			{
				Symbol:  "positionassingments_physicians_formuser",
				Columns: []*schema.Column{PositionassingmentsColumns[3]},

				RefColumns: []*schema.Column{PhysiciansColumns[0]},
				OnDelete:   schema.SetNull,
			},
			{
				Symbol:  "positionassingments_positions_formposition",
				Columns: []*schema.Column{PositionassingmentsColumns[4]},

				RefColumns: []*schema.Column{PositionsColumns[0]},
				OnDelete:   schema.SetNull,
			},
		},
	}
	// Tables holds all the tables in the schema.
	Tables = []*schema.Table{
		DepartmentsTable,
		PhysiciansTable,
		PositionsTable,
		PositionassingmentsTable,
	}
)

func init() {
	PositionassingmentsTable.ForeignKeys[0].RefTable = DepartmentsTable
	PositionassingmentsTable.ForeignKeys[1].RefTable = PhysiciansTable
	PositionassingmentsTable.ForeignKeys[2].RefTable = PositionsTable
}
