package schema
import (
   "github.com/facebookincubator/ent"
   "github.com/facebookincubator/ent/schema/edge"
   "github.com/facebookincubator/ent/schema/field"
)
 
type Physician struct {
   ent.Schema
}
 
func (Physician) Fields() []ent.Field {
   return []ent.Field{
       
       field.String("NAME"),
       field.String("EMAIL").
       Unique(),
   }
}
 
func (Physician) Edges() []ent.Edge {
   return []ent.Edge{
      edge.To("formuser", Positionassingment .Type).Unique().
      StorageKey(edge.Column("PHYSICIAN_ID")),
  }
}
