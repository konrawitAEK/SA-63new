package schema
import (
    "github.com/facebookincubator/ent"
    "github.com/facebookincubator/ent/schema/edge"
    "github.com/facebookincubator/ent/schema/field"
)

type Positionassingment  struct {
    ent.Schema
}

func (Positionassingment ) Fields() []ent.Field {
    return []ent.Field{
        field.Time("DayStart"),
    }
}

func (Positionassingment) Edges() []ent.Edge {
    return []ent.Edge{
        edge.From("user", Physician.Type).
            Ref("formuser").
			Unique(),
        edge.From("department", Department.Type).
            Ref("formdepartment").
			Unique(),
        edge.From("position", Position.Type).
            Ref("formposition").
			Unique(),
    }
}