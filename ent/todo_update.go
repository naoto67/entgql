// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/naoto67/entgql/ent/predicate"
	"github.com/naoto67/entgql/ent/todo"
)

// TodoUpdate is the builder for updating Todo entities.
type TodoUpdate struct {
	config
	hooks    []Hook
	mutation *TodoMutation
}

// Where appends a list predicates to the TodoUpdate builder.
func (tu *TodoUpdate) Where(ps ...predicate.Todo) *TodoUpdate {
	tu.mutation.Where(ps...)
	return tu
}

// SetStatus sets the "status" field.
func (tu *TodoUpdate) SetStatus(t todo.Status) *TodoUpdate {
	tu.mutation.SetStatus(t)
	return tu
}

// SetNillableStatus sets the "status" field if the given value is not nil.
func (tu *TodoUpdate) SetNillableStatus(t *todo.Status) *TodoUpdate {
	if t != nil {
		tu.SetStatus(*t)
	}
	return tu
}

// SetPriority sets the "priority" field.
func (tu *TodoUpdate) SetPriority(i int) *TodoUpdate {
	tu.mutation.ResetPriority()
	tu.mutation.SetPriority(i)
	return tu
}

// SetNillablePriority sets the "priority" field if the given value is not nil.
func (tu *TodoUpdate) SetNillablePriority(i *int) *TodoUpdate {
	if i != nil {
		tu.SetPriority(*i)
	}
	return tu
}

// AddPriority adds i to the "priority" field.
func (tu *TodoUpdate) AddPriority(i int) *TodoUpdate {
	tu.mutation.AddPriority(i)
	return tu
}

// SetText sets the "text" field.
func (tu *TodoUpdate) SetText(s string) *TodoUpdate {
	tu.mutation.SetText(s)
	return tu
}

// SetNillableText sets the "text" field if the given value is not nil.
func (tu *TodoUpdate) SetNillableText(s *string) *TodoUpdate {
	if s != nil {
		tu.SetText(*s)
	}
	return tu
}

// SetBlob sets the "blob" field.
func (tu *TodoUpdate) SetBlob(b []byte) *TodoUpdate {
	tu.mutation.SetBlob(b)
	return tu
}

// ClearBlob clears the value of the "blob" field.
func (tu *TodoUpdate) ClearBlob() *TodoUpdate {
	tu.mutation.ClearBlob()
	return tu
}

// SetInit sets the "init" field.
func (tu *TodoUpdate) SetInit(m map[string]interface{}) *TodoUpdate {
	tu.mutation.SetInit(m)
	return tu
}

// ClearInit clears the value of the "init" field.
func (tu *TodoUpdate) ClearInit() *TodoUpdate {
	tu.mutation.ClearInit()
	return tu
}

// SetValue sets the "value" field.
func (tu *TodoUpdate) SetValue(i int) *TodoUpdate {
	tu.mutation.ResetValue()
	tu.mutation.SetValue(i)
	return tu
}

// SetNillableValue sets the "value" field if the given value is not nil.
func (tu *TodoUpdate) SetNillableValue(i *int) *TodoUpdate {
	if i != nil {
		tu.SetValue(*i)
	}
	return tu
}

// AddValue adds i to the "value" field.
func (tu *TodoUpdate) AddValue(i int) *TodoUpdate {
	tu.mutation.AddValue(i)
	return tu
}

// Mutation returns the TodoMutation object of the builder.
func (tu *TodoUpdate) Mutation() *TodoMutation {
	return tu.mutation
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (tu *TodoUpdate) Save(ctx context.Context) (int, error) {
	return withHooks(ctx, tu.sqlSave, tu.mutation, tu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (tu *TodoUpdate) SaveX(ctx context.Context) int {
	affected, err := tu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (tu *TodoUpdate) Exec(ctx context.Context) error {
	_, err := tu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (tu *TodoUpdate) ExecX(ctx context.Context) {
	if err := tu.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (tu *TodoUpdate) check() error {
	if v, ok := tu.mutation.Status(); ok {
		if err := todo.StatusValidator(v); err != nil {
			return &ValidationError{Name: "status", err: fmt.Errorf(`ent: validator failed for field "Todo.status": %w`, err)}
		}
	}
	if v, ok := tu.mutation.Text(); ok {
		if err := todo.TextValidator(v); err != nil {
			return &ValidationError{Name: "text", err: fmt.Errorf(`ent: validator failed for field "Todo.text": %w`, err)}
		}
	}
	return nil
}

func (tu *TodoUpdate) sqlSave(ctx context.Context) (n int, err error) {
	if err := tu.check(); err != nil {
		return n, err
	}
	_spec := sqlgraph.NewUpdateSpec(todo.Table, todo.Columns, sqlgraph.NewFieldSpec(todo.FieldID, field.TypeString))
	if ps := tu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := tu.mutation.Status(); ok {
		_spec.SetField(todo.FieldStatus, field.TypeEnum, value)
	}
	if value, ok := tu.mutation.Priority(); ok {
		_spec.SetField(todo.FieldPriority, field.TypeInt, value)
	}
	if value, ok := tu.mutation.AddedPriority(); ok {
		_spec.AddField(todo.FieldPriority, field.TypeInt, value)
	}
	if value, ok := tu.mutation.Text(); ok {
		_spec.SetField(todo.FieldText, field.TypeString, value)
	}
	if value, ok := tu.mutation.Blob(); ok {
		_spec.SetField(todo.FieldBlob, field.TypeBytes, value)
	}
	if tu.mutation.BlobCleared() {
		_spec.ClearField(todo.FieldBlob, field.TypeBytes)
	}
	if tu.mutation.CategoryIDCleared() {
		_spec.ClearField(todo.FieldCategoryID, field.TypeInt)
	}
	if value, ok := tu.mutation.Init(); ok {
		_spec.SetField(todo.FieldInit, field.TypeJSON, value)
	}
	if tu.mutation.InitCleared() {
		_spec.ClearField(todo.FieldInit, field.TypeJSON)
	}
	if value, ok := tu.mutation.Value(); ok {
		_spec.SetField(todo.FieldValue, field.TypeInt, value)
	}
	if value, ok := tu.mutation.AddedValue(); ok {
		_spec.AddField(todo.FieldValue, field.TypeInt, value)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, tu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{todo.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	tu.mutation.done = true
	return n, nil
}

// TodoUpdateOne is the builder for updating a single Todo entity.
type TodoUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *TodoMutation
}

// SetStatus sets the "status" field.
func (tuo *TodoUpdateOne) SetStatus(t todo.Status) *TodoUpdateOne {
	tuo.mutation.SetStatus(t)
	return tuo
}

// SetNillableStatus sets the "status" field if the given value is not nil.
func (tuo *TodoUpdateOne) SetNillableStatus(t *todo.Status) *TodoUpdateOne {
	if t != nil {
		tuo.SetStatus(*t)
	}
	return tuo
}

// SetPriority sets the "priority" field.
func (tuo *TodoUpdateOne) SetPriority(i int) *TodoUpdateOne {
	tuo.mutation.ResetPriority()
	tuo.mutation.SetPriority(i)
	return tuo
}

// SetNillablePriority sets the "priority" field if the given value is not nil.
func (tuo *TodoUpdateOne) SetNillablePriority(i *int) *TodoUpdateOne {
	if i != nil {
		tuo.SetPriority(*i)
	}
	return tuo
}

// AddPriority adds i to the "priority" field.
func (tuo *TodoUpdateOne) AddPriority(i int) *TodoUpdateOne {
	tuo.mutation.AddPriority(i)
	return tuo
}

// SetText sets the "text" field.
func (tuo *TodoUpdateOne) SetText(s string) *TodoUpdateOne {
	tuo.mutation.SetText(s)
	return tuo
}

// SetNillableText sets the "text" field if the given value is not nil.
func (tuo *TodoUpdateOne) SetNillableText(s *string) *TodoUpdateOne {
	if s != nil {
		tuo.SetText(*s)
	}
	return tuo
}

// SetBlob sets the "blob" field.
func (tuo *TodoUpdateOne) SetBlob(b []byte) *TodoUpdateOne {
	tuo.mutation.SetBlob(b)
	return tuo
}

// ClearBlob clears the value of the "blob" field.
func (tuo *TodoUpdateOne) ClearBlob() *TodoUpdateOne {
	tuo.mutation.ClearBlob()
	return tuo
}

// SetInit sets the "init" field.
func (tuo *TodoUpdateOne) SetInit(m map[string]interface{}) *TodoUpdateOne {
	tuo.mutation.SetInit(m)
	return tuo
}

// ClearInit clears the value of the "init" field.
func (tuo *TodoUpdateOne) ClearInit() *TodoUpdateOne {
	tuo.mutation.ClearInit()
	return tuo
}

// SetValue sets the "value" field.
func (tuo *TodoUpdateOne) SetValue(i int) *TodoUpdateOne {
	tuo.mutation.ResetValue()
	tuo.mutation.SetValue(i)
	return tuo
}

// SetNillableValue sets the "value" field if the given value is not nil.
func (tuo *TodoUpdateOne) SetNillableValue(i *int) *TodoUpdateOne {
	if i != nil {
		tuo.SetValue(*i)
	}
	return tuo
}

// AddValue adds i to the "value" field.
func (tuo *TodoUpdateOne) AddValue(i int) *TodoUpdateOne {
	tuo.mutation.AddValue(i)
	return tuo
}

// Mutation returns the TodoMutation object of the builder.
func (tuo *TodoUpdateOne) Mutation() *TodoMutation {
	return tuo.mutation
}

// Where appends a list predicates to the TodoUpdate builder.
func (tuo *TodoUpdateOne) Where(ps ...predicate.Todo) *TodoUpdateOne {
	tuo.mutation.Where(ps...)
	return tuo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (tuo *TodoUpdateOne) Select(field string, fields ...string) *TodoUpdateOne {
	tuo.fields = append([]string{field}, fields...)
	return tuo
}

// Save executes the query and returns the updated Todo entity.
func (tuo *TodoUpdateOne) Save(ctx context.Context) (*Todo, error) {
	return withHooks(ctx, tuo.sqlSave, tuo.mutation, tuo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (tuo *TodoUpdateOne) SaveX(ctx context.Context) *Todo {
	node, err := tuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (tuo *TodoUpdateOne) Exec(ctx context.Context) error {
	_, err := tuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (tuo *TodoUpdateOne) ExecX(ctx context.Context) {
	if err := tuo.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (tuo *TodoUpdateOne) check() error {
	if v, ok := tuo.mutation.Status(); ok {
		if err := todo.StatusValidator(v); err != nil {
			return &ValidationError{Name: "status", err: fmt.Errorf(`ent: validator failed for field "Todo.status": %w`, err)}
		}
	}
	if v, ok := tuo.mutation.Text(); ok {
		if err := todo.TextValidator(v); err != nil {
			return &ValidationError{Name: "text", err: fmt.Errorf(`ent: validator failed for field "Todo.text": %w`, err)}
		}
	}
	return nil
}

func (tuo *TodoUpdateOne) sqlSave(ctx context.Context) (_node *Todo, err error) {
	if err := tuo.check(); err != nil {
		return _node, err
	}
	_spec := sqlgraph.NewUpdateSpec(todo.Table, todo.Columns, sqlgraph.NewFieldSpec(todo.FieldID, field.TypeString))
	id, ok := tuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "Todo.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := tuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, todo.FieldID)
		for _, f := range fields {
			if !todo.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != todo.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := tuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := tuo.mutation.Status(); ok {
		_spec.SetField(todo.FieldStatus, field.TypeEnum, value)
	}
	if value, ok := tuo.mutation.Priority(); ok {
		_spec.SetField(todo.FieldPriority, field.TypeInt, value)
	}
	if value, ok := tuo.mutation.AddedPriority(); ok {
		_spec.AddField(todo.FieldPriority, field.TypeInt, value)
	}
	if value, ok := tuo.mutation.Text(); ok {
		_spec.SetField(todo.FieldText, field.TypeString, value)
	}
	if value, ok := tuo.mutation.Blob(); ok {
		_spec.SetField(todo.FieldBlob, field.TypeBytes, value)
	}
	if tuo.mutation.BlobCleared() {
		_spec.ClearField(todo.FieldBlob, field.TypeBytes)
	}
	if tuo.mutation.CategoryIDCleared() {
		_spec.ClearField(todo.FieldCategoryID, field.TypeInt)
	}
	if value, ok := tuo.mutation.Init(); ok {
		_spec.SetField(todo.FieldInit, field.TypeJSON, value)
	}
	if tuo.mutation.InitCleared() {
		_spec.ClearField(todo.FieldInit, field.TypeJSON)
	}
	if value, ok := tuo.mutation.Value(); ok {
		_spec.SetField(todo.FieldValue, field.TypeInt, value)
	}
	if value, ok := tuo.mutation.AddedValue(); ok {
		_spec.AddField(todo.FieldValue, field.TypeInt, value)
	}
	_node = &Todo{config: tuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, tuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{todo.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	tuo.mutation.done = true
	return _node, nil
}
