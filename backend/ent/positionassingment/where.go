// Code generated by entc, DO NOT EDIT.

package positionassingment

import (
	"time"

	"github.com/facebookincubator/ent/dialect/sql"
	"github.com/facebookincubator/ent/dialect/sql/sqlgraph"
	"github.com/konrawitAEK/app/ent/predicate"
)

// ID filters vertices based on their identifier.
func ID(id int) predicate.Positionassingment {
	return predicate.Positionassingment(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int) predicate.Positionassingment {
	return predicate.Positionassingment(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int) predicate.Positionassingment {
	return predicate.Positionassingment(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldID), id))
	})
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int) predicate.Positionassingment {
	return predicate.Positionassingment(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(ids) == 0 {
			s.Where(sql.False())
			return
		}
		v := make([]interface{}, len(ids))
		for i := range v {
			v[i] = ids[i]
		}
		s.Where(sql.In(s.C(FieldID), v...))
	})
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...int) predicate.Positionassingment {
	return predicate.Positionassingment(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(ids) == 0 {
			s.Where(sql.False())
			return
		}
		v := make([]interface{}, len(ids))
		for i := range v {
			v[i] = ids[i]
		}
		s.Where(sql.NotIn(s.C(FieldID), v...))
	})
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id int) predicate.Positionassingment {
	return predicate.Positionassingment(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldID), id))
	})
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int) predicate.Positionassingment {
	return predicate.Positionassingment(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldID), id))
	})
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int) predicate.Positionassingment {
	return predicate.Positionassingment(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldID), id))
	})
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int) predicate.Positionassingment {
	return predicate.Positionassingment(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldID), id))
	})
}

// DayStart applies equality check predicate on the "DayStart" field. It's identical to DayStartEQ.
func DayStart(v time.Time) predicate.Positionassingment {
	return predicate.Positionassingment(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldDayStart), v))
	})
}

// DayStartEQ applies the EQ predicate on the "DayStart" field.
func DayStartEQ(v time.Time) predicate.Positionassingment {
	return predicate.Positionassingment(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldDayStart), v))
	})
}

// DayStartNEQ applies the NEQ predicate on the "DayStart" field.
func DayStartNEQ(v time.Time) predicate.Positionassingment {
	return predicate.Positionassingment(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldDayStart), v))
	})
}

// DayStartIn applies the In predicate on the "DayStart" field.
func DayStartIn(vs ...time.Time) predicate.Positionassingment {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Positionassingment(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldDayStart), v...))
	})
}

// DayStartNotIn applies the NotIn predicate on the "DayStart" field.
func DayStartNotIn(vs ...time.Time) predicate.Positionassingment {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Positionassingment(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldDayStart), v...))
	})
}

// DayStartGT applies the GT predicate on the "DayStart" field.
func DayStartGT(v time.Time) predicate.Positionassingment {
	return predicate.Positionassingment(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldDayStart), v))
	})
}

// DayStartGTE applies the GTE predicate on the "DayStart" field.
func DayStartGTE(v time.Time) predicate.Positionassingment {
	return predicate.Positionassingment(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldDayStart), v))
	})
}

// DayStartLT applies the LT predicate on the "DayStart" field.
func DayStartLT(v time.Time) predicate.Positionassingment {
	return predicate.Positionassingment(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldDayStart), v))
	})
}

// DayStartLTE applies the LTE predicate on the "DayStart" field.
func DayStartLTE(v time.Time) predicate.Positionassingment {
	return predicate.Positionassingment(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldDayStart), v))
	})
}

// HasUser applies the HasEdge predicate on the "user" edge.
func HasUser() predicate.Positionassingment {
	return predicate.Positionassingment(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(UserTable, FieldID),
			sqlgraph.Edge(sqlgraph.O2O, true, UserTable, UserColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasUserWith applies the HasEdge predicate on the "user" edge with a given conditions (other predicates).
func HasUserWith(preds ...predicate.Physician) predicate.Positionassingment {
	return predicate.Positionassingment(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(UserInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.O2O, true, UserTable, UserColumn),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasDepartment applies the HasEdge predicate on the "department" edge.
func HasDepartment() predicate.Positionassingment {
	return predicate.Positionassingment(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(DepartmentTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, DepartmentTable, DepartmentColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasDepartmentWith applies the HasEdge predicate on the "department" edge with a given conditions (other predicates).
func HasDepartmentWith(preds ...predicate.Department) predicate.Positionassingment {
	return predicate.Positionassingment(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(DepartmentInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, DepartmentTable, DepartmentColumn),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasPosition applies the HasEdge predicate on the "position" edge.
func HasPosition() predicate.Positionassingment {
	return predicate.Positionassingment(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(PositionTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, PositionTable, PositionColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasPositionWith applies the HasEdge predicate on the "position" edge with a given conditions (other predicates).
func HasPositionWith(preds ...predicate.Position) predicate.Positionassingment {
	return predicate.Positionassingment(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(PositionInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, PositionTable, PositionColumn),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups list of predicates with the AND operator between them.
func And(predicates ...predicate.Positionassingment) predicate.Positionassingment {
	return predicate.Positionassingment(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups list of predicates with the OR operator between them.
func Or(predicates ...predicate.Positionassingment) predicate.Positionassingment {
	return predicate.Positionassingment(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for i, p := range predicates {
			if i > 0 {
				s1.Or()
			}
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Not applies the not operator on the given predicate.
func Not(p predicate.Positionassingment) predicate.Positionassingment {
	return predicate.Positionassingment(func(s *sql.Selector) {
		p(s.Not())
	})
}
