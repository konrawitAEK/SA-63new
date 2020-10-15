package schema
import (
    "github.com/facebookincubator/ent"
    "github.com/facebookincubator/ent/schema/edge"
    "github.com/facebookincubator/ent/schema/field"
)

type Position struct {
    ent.Schema
}

func (Position) Fields() []ent.Field {
    return []ent.Field{
        field.String("Nameposition"),        
    }
}

func (Position) Edges() []ent.Edge {
    return []ent.Edge{
        edge.To("formposition", Positionassingment .Type).StorageKey(edge.Column("POSITION_ID")),
    }
}
