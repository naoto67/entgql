// Code generated by ent, DO NOT EDIT.

package ent

import (
	"github.com/naoto67/entgql/ent/todo"
)

// CreateTodoInput represents a mutation input for creating todos.
type CreateTodoInput struct {
	Status     todo.Status
	Priority   *int
	Text       string
	CategoryID *int
	Init       map[string]interface{}
	Value      *int
}

// Mutate applies the CreateTodoInput on the TodoMutation builder.
func (i *CreateTodoInput) Mutate(m *TodoMutation) {
	m.SetStatus(i.Status)
	if v := i.Priority; v != nil {
		m.SetPriority(*v)
	}
	m.SetText(i.Text)
	if v := i.CategoryID; v != nil {
		m.SetCategoryID(*v)
	}
	if v := i.Init; v != nil {
		m.SetInit(v)
	}
	if v := i.Value; v != nil {
		m.SetValue(*v)
	}
}

// SetInput applies the change-set in the CreateTodoInput on the TodoCreate builder.
func (c *TodoCreate) SetInput(i CreateTodoInput) *TodoCreate {
	i.Mutate(c.Mutation())
	return c
}

// UpdateTodoInput represents a mutation input for updating todos.
type UpdateTodoInput struct {
	Status    *todo.Status
	Priority  *int
	Text      *string
	ClearInit bool
	Init      map[string]interface{}
	Value     *int
}

// Mutate applies the UpdateTodoInput on the TodoMutation builder.
func (i *UpdateTodoInput) Mutate(m *TodoMutation) {
	if v := i.Status; v != nil {
		m.SetStatus(*v)
	}
	if v := i.Priority; v != nil {
		m.SetPriority(*v)
	}
	if v := i.Text; v != nil {
		m.SetText(*v)
	}
	if i.ClearInit {
		m.ClearInit()
	}
	if v := i.Init; v != nil {
		m.SetInit(v)
	}
	if v := i.Value; v != nil {
		m.SetValue(*v)
	}
}

// SetInput applies the change-set in the UpdateTodoInput on the TodoUpdate builder.
func (c *TodoUpdate) SetInput(i UpdateTodoInput) *TodoUpdate {
	i.Mutate(c.Mutation())
	return c
}

// SetInput applies the change-set in the UpdateTodoInput on the TodoUpdateOne builder.
func (c *TodoUpdateOne) SetInput(i UpdateTodoInput) *TodoUpdateOne {
	i.Mutate(c.Mutation())
	return c
}
