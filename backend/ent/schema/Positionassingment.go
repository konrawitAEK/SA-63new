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

func (Positionassingment ) Edges() []ent.Edge {
    return []ent.Edge{
        edge.From("formdata", Physician.Type).
            Ref("user").
            Unique(),

        edge.From("formdepartment", Department.Type).
            Ref("department").
            Unique(),
        edge.From("formposition", Position.Type).
            Ref("position").
            Unique(),
    }
}