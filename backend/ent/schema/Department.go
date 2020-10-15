package schema
import (
    "github.com/facebookincubator/ent"
    "github.com/facebookincubator/ent/schema/edge"
    "github.com/facebookincubator/ent/schema/field"
)

type Department struct {
    ent.Schema
}

func (Department) Fields() []ent.Field {
    return []ent.Field{
        field.String("Departmentname"),    
	}
}

func (Department) Edges() []ent.Edge {
    return []ent.Edge{
        edge.To("formdepartment", Positionassingment .Type).StorageKey(edge.Column("DEPARTMENT_ID")),
    }
}
