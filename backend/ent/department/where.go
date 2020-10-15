// Code generated by entc, DO NOT EDIT.

package department

import (
	"github.com/facebookincubator/ent/dialect/sql"
	"github.com/facebookincubator/ent/dialect/sql/sqlgraph"
	"github.com/konrawitAEK/app/ent/predicate"
)

// ID filters vertices based on their identifier.
func ID(id int) predicate.Department {
	return predicate.Department(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int) predicate.Department {
	return predicate.Department(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int) predicate.Department {
	return predicate.Department(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldID), id))
	})
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int) predicate.Department {
	return predicate.Department(func(s *sql.Selector) {
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
func IDNotIn(ids ...int) predicate.Department {
	return predicate.Department(func(s *sql.Selector) {
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
func IDGT(id int) predicate.Department {
	return predicate.Department(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldID), id))
	})
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int) predicate.Department {
	return predicate.Department(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldID), id))
	})
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int) predicate.Department {
	return predicate.Department(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldID), id))
	})
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int) predicate.Department {
	return predicate.Department(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldID), id))
	})
}

// Departmentname applies equality check predicate on the "Departmentname" field. It's identical to DepartmentnameEQ.
func Departmentname(v string) predicate.Department {
	return predicate.Department(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldDepartmentname), v))
	})
}

// DepartmentnameEQ applies the EQ predicate on the "Departmentname" field.
func DepartmentnameEQ(v string) predicate.Department {
	return predicate.Department(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldDepartmentname), v))
	})
}

// DepartmentnameNEQ applies the NEQ predicate on the "Departmentname" field.
func DepartmentnameNEQ(v string) predicate.Department {
	return predicate.Department(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldDepartmentname), v))
	})
}

// DepartmentnameIn applies the In predicate on the "Departmentname" field.
func DepartmentnameIn(vs ...string) predicate.Department {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Department(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldDepartmentname), v...))
	})
}

// DepartmentnameNotIn applies the NotIn predicate on the "Departmentname" field.
func DepartmentnameNotIn(vs ...string) predicate.Department {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Department(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldDepartmentname), v...))
	})
}

// DepartmentnameGT applies the GT predicate on the "Departmentname" field.
func DepartmentnameGT(v string) predicate.Department {
	return predicate.Department(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldDepartmentname), v))
	})
}

// DepartmentnameGTE applies the GTE predicate on the "Departmentname" field.
func DepartmentnameGTE(v string) predicate.Department {
	return predicate.Department(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldDepartmentname), v))
	})
}

// DepartmentnameLT applies the LT predicate on the "Departmentname" field.
func DepartmentnameLT(v string) predicate.Department {
	return predicate.Department(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldDepartmentname), v))
	})
}

// DepartmentnameLTE applies the LTE predicate on the "Departmentname" field.
func DepartmentnameLTE(v string) predicate.Department {
	return predicate.Department(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldDepartmentname), v))
	})
}

// DepartmentnameContains applies the Contains predicate on the "Departmentname" field.
func DepartmentnameContains(v string) predicate.Department {
	return predicate.Department(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldDepartmentname), v))
	})
}

// DepartmentnameHasPrefix applies the HasPrefix predicate on the "Departmentname" field.
func DepartmentnameHasPrefix(v string) predicate.Department {
	return predicate.Department(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldDepartmentname), v))
	})
}

// DepartmentnameHasSuffix applies the HasSuffix predicate on the "Departmentname" field.
func DepartmentnameHasSuffix(v string) predicate.Department {
	return predicate.Department(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldDepartmentname), v))
	})
}

// DepartmentnameEqualFold applies the EqualFold predicate on the "Departmentname" field.
func DepartmentnameEqualFold(v string) predicate.Department {
	return predicate.Department(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldDepartmentname), v))
	})
}

// DepartmentnameContainsFold applies the ContainsFold predicate on the "Departmentname" field.
func DepartmentnameContainsFold(v string) predicate.Department {
	return predicate.Department(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldDepartmentname), v))
	})
}

// HasFormdepartment applies the HasEdge predicate on the "formdepartment" edge.
func HasFormdepartment() predicate.Department {
	return predicate.Department(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(FormdepartmentTable, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, FormdepartmentTable, FormdepartmentColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasFormdepartmentWith applies the HasEdge predicate on the "formdepartment" edge with a given conditions (other predicates).
func HasFormdepartmentWith(preds ...predicate.Positionassingment) predicate.Department {
	return predicate.Department(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(FormdepartmentInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, FormdepartmentTable, FormdepartmentColumn),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups list of predicates with the AND operator between them.
func And(predicates ...predicate.Department) predicate.Department {
	return predicate.Department(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups list of predicates with the OR operator between them.
func Or(predicates ...predicate.Department) predicate.Department {
	return predicate.Department(func(s *sql.Selector) {
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
func Not(p predicate.Department) predicate.Department {
	return predicate.Department(func(s *sql.Selector) {
		p(s.Not())
	})
}
