// Code generated by ent, DO NOT EDIT.

package ent

import (
	"encoding/json"
	"fmt"
	"strings"
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"github.com/naoto67/entgql/ent/schema/puuid"
	"github.com/naoto67/entgql/ent/todo"
)

// Todo is the model entity for the Todo schema.
type Todo struct {
	config `json:"-"`
	// ID of the ent.
	ID puuid.ID `json:"id,omitempty"`
	// CreatedAt holds the value of the "created_at" field.
	CreatedAt time.Time `json:"created_at,omitempty"`
	// Status holds the value of the "status" field.
	Status todo.Status `json:"status,omitempty"`
	// Priority holds the value of the "priority" field.
	Priority int `json:"priority,omitempty"`
	// Text holds the value of the "text" field.
	Text string `json:"text,omitempty"`
	// Blob holds the value of the "blob" field.
	Blob []byte `json:"blob,omitempty"`
	// Init holds the value of the "init" field.
	Init map[string]interface{} `json:"init,omitempty"`
	// Value holds the value of the "value" field.
	Value        int `json:"value,omitempty"`
	selectValues sql.SelectValues
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Todo) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case todo.FieldBlob, todo.FieldInit:
			values[i] = new([]byte)
		case todo.FieldID:
			values[i] = new(puuid.ID)
		case todo.FieldPriority, todo.FieldValue:
			values[i] = new(sql.NullInt64)
		case todo.FieldStatus, todo.FieldText:
			values[i] = new(sql.NullString)
		case todo.FieldCreatedAt:
			values[i] = new(sql.NullTime)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Todo fields.
func (t *Todo) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case todo.FieldID:
			if value, ok := values[i].(*puuid.ID); !ok {
				return fmt.Errorf("unexpected type %T for field id", values[i])
			} else if value != nil {
				t.ID = *value
			}
		case todo.FieldCreatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field created_at", values[i])
			} else if value.Valid {
				t.CreatedAt = value.Time
			}
		case todo.FieldStatus:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field status", values[i])
			} else if value.Valid {
				t.Status = todo.Status(value.String)
			}
		case todo.FieldPriority:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field priority", values[i])
			} else if value.Valid {
				t.Priority = int(value.Int64)
			}
		case todo.FieldText:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field text", values[i])
			} else if value.Valid {
				t.Text = value.String
			}
		case todo.FieldBlob:
			if value, ok := values[i].(*[]byte); !ok {
				return fmt.Errorf("unexpected type %T for field blob", values[i])
			} else if value != nil {
				t.Blob = *value
			}
		case todo.FieldInit:
			if value, ok := values[i].(*[]byte); !ok {
				return fmt.Errorf("unexpected type %T for field init", values[i])
			} else if value != nil && len(*value) > 0 {
				if err := json.Unmarshal(*value, &t.Init); err != nil {
					return fmt.Errorf("unmarshal field init: %w", err)
				}
			}
		case todo.FieldValue:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field value", values[i])
			} else if value.Valid {
				t.Value = int(value.Int64)
			}
		default:
			t.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// GetValue returns the ent.Value that was dynamically selected and assigned to the Todo.
// This includes values selected through modifiers, order, etc.
func (t *Todo) GetValue(name string) (ent.Value, error) {
	return t.selectValues.Get(name)
}

// Update returns a builder for updating this Todo.
// Note that you need to call Todo.Unwrap() before calling this method if this Todo
// was returned from a transaction, and the transaction was committed or rolled back.
func (t *Todo) Update() *TodoUpdateOne {
	return NewTodoClient(t.config).UpdateOne(t)
}

// Unwrap unwraps the Todo entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (t *Todo) Unwrap() *Todo {
	_tx, ok := t.config.driver.(*txDriver)
	if !ok {
		panic("ent: Todo is not a transactional entity")
	}
	t.config.driver = _tx.drv
	return t
}

// String implements the fmt.Stringer.
func (t *Todo) String() string {
	var builder strings.Builder
	builder.WriteString("Todo(")
	builder.WriteString(fmt.Sprintf("id=%v, ", t.ID))
	builder.WriteString("created_at=")
	builder.WriteString(t.CreatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("status=")
	builder.WriteString(fmt.Sprintf("%v", t.Status))
	builder.WriteString(", ")
	builder.WriteString("priority=")
	builder.WriteString(fmt.Sprintf("%v", t.Priority))
	builder.WriteString(", ")
	builder.WriteString("text=")
	builder.WriteString(t.Text)
	builder.WriteString(", ")
	builder.WriteString("blob=")
	builder.WriteString(fmt.Sprintf("%v", t.Blob))
	builder.WriteString(", ")
	builder.WriteString("init=")
	builder.WriteString(fmt.Sprintf("%v", t.Init))
	builder.WriteString(", ")
	builder.WriteString("value=")
	builder.WriteString(fmt.Sprintf("%v", t.Value))
	builder.WriteByte(')')
	return builder.String()
}

// Todos is a parsable slice of Todo.
type Todos []*Todo
