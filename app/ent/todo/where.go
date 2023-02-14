// Code generated by ent, DO NOT EDIT.

package todo

import (
	"app/ent/predicate"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
)

// ID filters vertices based on their ID field.
func ID(id int) predicate.Todo {
	return predicate.Todo(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int) predicate.Todo {
	return predicate.Todo(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int) predicate.Todo {
	return predicate.Todo(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldID), id))
	})
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int) predicate.Todo {
	return predicate.Todo(func(s *sql.Selector) {
		v := make([]any, len(ids))
		for i := range v {
			v[i] = ids[i]
		}
		s.Where(sql.In(s.C(FieldID), v...))
	})
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...int) predicate.Todo {
	return predicate.Todo(func(s *sql.Selector) {
		v := make([]any, len(ids))
		for i := range v {
			v[i] = ids[i]
		}
		s.Where(sql.NotIn(s.C(FieldID), v...))
	})
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id int) predicate.Todo {
	return predicate.Todo(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldID), id))
	})
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int) predicate.Todo {
	return predicate.Todo(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldID), id))
	})
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int) predicate.Todo {
	return predicate.Todo(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldID), id))
	})
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int) predicate.Todo {
	return predicate.Todo(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldID), id))
	})
}

// Title applies equality check predicate on the "title" field. It's identical to TitleEQ.
func Title(v string) predicate.Todo {
	return predicate.Todo(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldTitle), v))
	})
}

// Description applies equality check predicate on the "description" field. It's identical to DescriptionEQ.
func Description(v string) predicate.Todo {
	return predicate.Todo(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldDescription), v))
	})
}

// Status applies equality check predicate on the "status" field. It's identical to StatusEQ.
func Status(v string) predicate.Todo {
	return predicate.Todo(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldStatus), v))
	})
}

// DueDate applies equality check predicate on the "due_date" field. It's identical to DueDateEQ.
func DueDate(v string) predicate.Todo {
	return predicate.Todo(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldDueDate), v))
	})
}

// CompletedAt applies equality check predicate on the "completed_at" field. It's identical to CompletedAtEQ.
func CompletedAt(v string) predicate.Todo {
	return predicate.Todo(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldCompletedAt), v))
	})
}

// CreatedAt applies equality check predicate on the "created_at" field. It's identical to CreatedAtEQ.
func CreatedAt(v string) predicate.Todo {
	return predicate.Todo(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldCreatedAt), v))
	})
}

// TitleEQ applies the EQ predicate on the "title" field.
func TitleEQ(v string) predicate.Todo {
	return predicate.Todo(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldTitle), v))
	})
}

// TitleNEQ applies the NEQ predicate on the "title" field.
func TitleNEQ(v string) predicate.Todo {
	return predicate.Todo(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldTitle), v))
	})
}

// TitleIn applies the In predicate on the "title" field.
func TitleIn(vs ...string) predicate.Todo {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Todo(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldTitle), v...))
	})
}

// TitleNotIn applies the NotIn predicate on the "title" field.
func TitleNotIn(vs ...string) predicate.Todo {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Todo(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldTitle), v...))
	})
}

// TitleGT applies the GT predicate on the "title" field.
func TitleGT(v string) predicate.Todo {
	return predicate.Todo(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldTitle), v))
	})
}

// TitleGTE applies the GTE predicate on the "title" field.
func TitleGTE(v string) predicate.Todo {
	return predicate.Todo(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldTitle), v))
	})
}

// TitleLT applies the LT predicate on the "title" field.
func TitleLT(v string) predicate.Todo {
	return predicate.Todo(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldTitle), v))
	})
}

// TitleLTE applies the LTE predicate on the "title" field.
func TitleLTE(v string) predicate.Todo {
	return predicate.Todo(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldTitle), v))
	})
}

// TitleContains applies the Contains predicate on the "title" field.
func TitleContains(v string) predicate.Todo {
	return predicate.Todo(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldTitle), v))
	})
}

// TitleHasPrefix applies the HasPrefix predicate on the "title" field.
func TitleHasPrefix(v string) predicate.Todo {
	return predicate.Todo(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldTitle), v))
	})
}

// TitleHasSuffix applies the HasSuffix predicate on the "title" field.
func TitleHasSuffix(v string) predicate.Todo {
	return predicate.Todo(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldTitle), v))
	})
}

// TitleEqualFold applies the EqualFold predicate on the "title" field.
func TitleEqualFold(v string) predicate.Todo {
	return predicate.Todo(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldTitle), v))
	})
}

// TitleContainsFold applies the ContainsFold predicate on the "title" field.
func TitleContainsFold(v string) predicate.Todo {
	return predicate.Todo(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldTitle), v))
	})
}

// DescriptionEQ applies the EQ predicate on the "description" field.
func DescriptionEQ(v string) predicate.Todo {
	return predicate.Todo(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldDescription), v))
	})
}

// DescriptionNEQ applies the NEQ predicate on the "description" field.
func DescriptionNEQ(v string) predicate.Todo {
	return predicate.Todo(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldDescription), v))
	})
}

// DescriptionIn applies the In predicate on the "description" field.
func DescriptionIn(vs ...string) predicate.Todo {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Todo(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldDescription), v...))
	})
}

// DescriptionNotIn applies the NotIn predicate on the "description" field.
func DescriptionNotIn(vs ...string) predicate.Todo {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Todo(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldDescription), v...))
	})
}

// DescriptionGT applies the GT predicate on the "description" field.
func DescriptionGT(v string) predicate.Todo {
	return predicate.Todo(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldDescription), v))
	})
}

// DescriptionGTE applies the GTE predicate on the "description" field.
func DescriptionGTE(v string) predicate.Todo {
	return predicate.Todo(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldDescription), v))
	})
}

// DescriptionLT applies the LT predicate on the "description" field.
func DescriptionLT(v string) predicate.Todo {
	return predicate.Todo(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldDescription), v))
	})
}

// DescriptionLTE applies the LTE predicate on the "description" field.
func DescriptionLTE(v string) predicate.Todo {
	return predicate.Todo(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldDescription), v))
	})
}

// DescriptionContains applies the Contains predicate on the "description" field.
func DescriptionContains(v string) predicate.Todo {
	return predicate.Todo(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldDescription), v))
	})
}

// DescriptionHasPrefix applies the HasPrefix predicate on the "description" field.
func DescriptionHasPrefix(v string) predicate.Todo {
	return predicate.Todo(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldDescription), v))
	})
}

// DescriptionHasSuffix applies the HasSuffix predicate on the "description" field.
func DescriptionHasSuffix(v string) predicate.Todo {
	return predicate.Todo(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldDescription), v))
	})
}

// DescriptionEqualFold applies the EqualFold predicate on the "description" field.
func DescriptionEqualFold(v string) predicate.Todo {
	return predicate.Todo(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldDescription), v))
	})
}

// DescriptionContainsFold applies the ContainsFold predicate on the "description" field.
func DescriptionContainsFold(v string) predicate.Todo {
	return predicate.Todo(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldDescription), v))
	})
}

// StatusEQ applies the EQ predicate on the "status" field.
func StatusEQ(v string) predicate.Todo {
	return predicate.Todo(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldStatus), v))
	})
}

// StatusNEQ applies the NEQ predicate on the "status" field.
func StatusNEQ(v string) predicate.Todo {
	return predicate.Todo(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldStatus), v))
	})
}

// StatusIn applies the In predicate on the "status" field.
func StatusIn(vs ...string) predicate.Todo {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Todo(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldStatus), v...))
	})
}

// StatusNotIn applies the NotIn predicate on the "status" field.
func StatusNotIn(vs ...string) predicate.Todo {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Todo(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldStatus), v...))
	})
}

// StatusGT applies the GT predicate on the "status" field.
func StatusGT(v string) predicate.Todo {
	return predicate.Todo(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldStatus), v))
	})
}

// StatusGTE applies the GTE predicate on the "status" field.
func StatusGTE(v string) predicate.Todo {
	return predicate.Todo(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldStatus), v))
	})
}

// StatusLT applies the LT predicate on the "status" field.
func StatusLT(v string) predicate.Todo {
	return predicate.Todo(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldStatus), v))
	})
}

// StatusLTE applies the LTE predicate on the "status" field.
func StatusLTE(v string) predicate.Todo {
	return predicate.Todo(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldStatus), v))
	})
}

// StatusContains applies the Contains predicate on the "status" field.
func StatusContains(v string) predicate.Todo {
	return predicate.Todo(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldStatus), v))
	})
}

// StatusHasPrefix applies the HasPrefix predicate on the "status" field.
func StatusHasPrefix(v string) predicate.Todo {
	return predicate.Todo(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldStatus), v))
	})
}

// StatusHasSuffix applies the HasSuffix predicate on the "status" field.
func StatusHasSuffix(v string) predicate.Todo {
	return predicate.Todo(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldStatus), v))
	})
}

// StatusEqualFold applies the EqualFold predicate on the "status" field.
func StatusEqualFold(v string) predicate.Todo {
	return predicate.Todo(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldStatus), v))
	})
}

// StatusContainsFold applies the ContainsFold predicate on the "status" field.
func StatusContainsFold(v string) predicate.Todo {
	return predicate.Todo(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldStatus), v))
	})
}

// DueDateEQ applies the EQ predicate on the "due_date" field.
func DueDateEQ(v string) predicate.Todo {
	return predicate.Todo(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldDueDate), v))
	})
}

// DueDateNEQ applies the NEQ predicate on the "due_date" field.
func DueDateNEQ(v string) predicate.Todo {
	return predicate.Todo(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldDueDate), v))
	})
}

// DueDateIn applies the In predicate on the "due_date" field.
func DueDateIn(vs ...string) predicate.Todo {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Todo(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldDueDate), v...))
	})
}

// DueDateNotIn applies the NotIn predicate on the "due_date" field.
func DueDateNotIn(vs ...string) predicate.Todo {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Todo(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldDueDate), v...))
	})
}

// DueDateGT applies the GT predicate on the "due_date" field.
func DueDateGT(v string) predicate.Todo {
	return predicate.Todo(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldDueDate), v))
	})
}

// DueDateGTE applies the GTE predicate on the "due_date" field.
func DueDateGTE(v string) predicate.Todo {
	return predicate.Todo(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldDueDate), v))
	})
}

// DueDateLT applies the LT predicate on the "due_date" field.
func DueDateLT(v string) predicate.Todo {
	return predicate.Todo(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldDueDate), v))
	})
}

// DueDateLTE applies the LTE predicate on the "due_date" field.
func DueDateLTE(v string) predicate.Todo {
	return predicate.Todo(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldDueDate), v))
	})
}

// DueDateContains applies the Contains predicate on the "due_date" field.
func DueDateContains(v string) predicate.Todo {
	return predicate.Todo(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldDueDate), v))
	})
}

// DueDateHasPrefix applies the HasPrefix predicate on the "due_date" field.
func DueDateHasPrefix(v string) predicate.Todo {
	return predicate.Todo(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldDueDate), v))
	})
}

// DueDateHasSuffix applies the HasSuffix predicate on the "due_date" field.
func DueDateHasSuffix(v string) predicate.Todo {
	return predicate.Todo(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldDueDate), v))
	})
}

// DueDateEqualFold applies the EqualFold predicate on the "due_date" field.
func DueDateEqualFold(v string) predicate.Todo {
	return predicate.Todo(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldDueDate), v))
	})
}

// DueDateContainsFold applies the ContainsFold predicate on the "due_date" field.
func DueDateContainsFold(v string) predicate.Todo {
	return predicate.Todo(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldDueDate), v))
	})
}

// CompletedAtEQ applies the EQ predicate on the "completed_at" field.
func CompletedAtEQ(v string) predicate.Todo {
	return predicate.Todo(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldCompletedAt), v))
	})
}

// CompletedAtNEQ applies the NEQ predicate on the "completed_at" field.
func CompletedAtNEQ(v string) predicate.Todo {
	return predicate.Todo(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldCompletedAt), v))
	})
}

// CompletedAtIn applies the In predicate on the "completed_at" field.
func CompletedAtIn(vs ...string) predicate.Todo {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Todo(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldCompletedAt), v...))
	})
}

// CompletedAtNotIn applies the NotIn predicate on the "completed_at" field.
func CompletedAtNotIn(vs ...string) predicate.Todo {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Todo(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldCompletedAt), v...))
	})
}

// CompletedAtGT applies the GT predicate on the "completed_at" field.
func CompletedAtGT(v string) predicate.Todo {
	return predicate.Todo(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldCompletedAt), v))
	})
}

// CompletedAtGTE applies the GTE predicate on the "completed_at" field.
func CompletedAtGTE(v string) predicate.Todo {
	return predicate.Todo(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldCompletedAt), v))
	})
}

// CompletedAtLT applies the LT predicate on the "completed_at" field.
func CompletedAtLT(v string) predicate.Todo {
	return predicate.Todo(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldCompletedAt), v))
	})
}

// CompletedAtLTE applies the LTE predicate on the "completed_at" field.
func CompletedAtLTE(v string) predicate.Todo {
	return predicate.Todo(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldCompletedAt), v))
	})
}

// CompletedAtContains applies the Contains predicate on the "completed_at" field.
func CompletedAtContains(v string) predicate.Todo {
	return predicate.Todo(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldCompletedAt), v))
	})
}

// CompletedAtHasPrefix applies the HasPrefix predicate on the "completed_at" field.
func CompletedAtHasPrefix(v string) predicate.Todo {
	return predicate.Todo(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldCompletedAt), v))
	})
}

// CompletedAtHasSuffix applies the HasSuffix predicate on the "completed_at" field.
func CompletedAtHasSuffix(v string) predicate.Todo {
	return predicate.Todo(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldCompletedAt), v))
	})
}

// CompletedAtEqualFold applies the EqualFold predicate on the "completed_at" field.
func CompletedAtEqualFold(v string) predicate.Todo {
	return predicate.Todo(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldCompletedAt), v))
	})
}

// CompletedAtContainsFold applies the ContainsFold predicate on the "completed_at" field.
func CompletedAtContainsFold(v string) predicate.Todo {
	return predicate.Todo(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldCompletedAt), v))
	})
}

// CreatedAtEQ applies the EQ predicate on the "created_at" field.
func CreatedAtEQ(v string) predicate.Todo {
	return predicate.Todo(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldCreatedAt), v))
	})
}

// CreatedAtNEQ applies the NEQ predicate on the "created_at" field.
func CreatedAtNEQ(v string) predicate.Todo {
	return predicate.Todo(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldCreatedAt), v))
	})
}

// CreatedAtIn applies the In predicate on the "created_at" field.
func CreatedAtIn(vs ...string) predicate.Todo {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Todo(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldCreatedAt), v...))
	})
}

// CreatedAtNotIn applies the NotIn predicate on the "created_at" field.
func CreatedAtNotIn(vs ...string) predicate.Todo {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Todo(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldCreatedAt), v...))
	})
}

// CreatedAtGT applies the GT predicate on the "created_at" field.
func CreatedAtGT(v string) predicate.Todo {
	return predicate.Todo(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldCreatedAt), v))
	})
}

// CreatedAtGTE applies the GTE predicate on the "created_at" field.
func CreatedAtGTE(v string) predicate.Todo {
	return predicate.Todo(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldCreatedAt), v))
	})
}

// CreatedAtLT applies the LT predicate on the "created_at" field.
func CreatedAtLT(v string) predicate.Todo {
	return predicate.Todo(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldCreatedAt), v))
	})
}

// CreatedAtLTE applies the LTE predicate on the "created_at" field.
func CreatedAtLTE(v string) predicate.Todo {
	return predicate.Todo(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldCreatedAt), v))
	})
}

// CreatedAtContains applies the Contains predicate on the "created_at" field.
func CreatedAtContains(v string) predicate.Todo {
	return predicate.Todo(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldCreatedAt), v))
	})
}

// CreatedAtHasPrefix applies the HasPrefix predicate on the "created_at" field.
func CreatedAtHasPrefix(v string) predicate.Todo {
	return predicate.Todo(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldCreatedAt), v))
	})
}

// CreatedAtHasSuffix applies the HasSuffix predicate on the "created_at" field.
func CreatedAtHasSuffix(v string) predicate.Todo {
	return predicate.Todo(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldCreatedAt), v))
	})
}

// CreatedAtEqualFold applies the EqualFold predicate on the "created_at" field.
func CreatedAtEqualFold(v string) predicate.Todo {
	return predicate.Todo(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldCreatedAt), v))
	})
}

// CreatedAtContainsFold applies the ContainsFold predicate on the "created_at" field.
func CreatedAtContainsFold(v string) predicate.Todo {
	return predicate.Todo(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldCreatedAt), v))
	})
}

// HasUser applies the HasEdge predicate on the "user" edge.
func HasUser() predicate.Todo {
	return predicate.Todo(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(UserTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, UserTable, UserColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasUserWith applies the HasEdge predicate on the "user" edge with a given conditions (other predicates).
func HasUserWith(preds ...predicate.User) predicate.Todo {
	return predicate.Todo(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(UserInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, UserTable, UserColumn),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.Todo) predicate.Todo {
	return predicate.Todo(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.Todo) predicate.Todo {
	return predicate.Todo(func(s *sql.Selector) {
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
func Not(p predicate.Todo) predicate.Todo {
	return predicate.Todo(func(s *sql.Selector) {
		p(s.Not())
	})
}
