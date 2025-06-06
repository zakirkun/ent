// Copyright 2019-present Facebook Inc. All rights reserved.
// This source code is licensed under the Apache 2.0 license found
// in the LICENSE file in the root directory of this source tree.

// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"net/url"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/entc/integration/json/ent/schema"
	"entgo.io/ent/entc/integration/json/ent/user"
	"entgo.io/ent/schema/field"
)

// UserCreate is the builder for creating a User entity.
type UserCreate struct {
	config
	mutation *UserMutation
	hooks    []Hook
}

// SetT sets the "t" field.
func (_c *UserCreate) SetT(v *schema.T) *UserCreate {
	_c.mutation.SetT(v)
	return _c
}

// SetURL sets the "url" field.
func (_c *UserCreate) SetURL(v *url.URL) *UserCreate {
	_c.mutation.SetURL(v)
	return _c
}

// SetURLs sets the "URLs" field.
func (_c *UserCreate) SetURLs(v []*url.URL) *UserCreate {
	_c.mutation.SetURLs(v)
	return _c
}

// SetRaw sets the "raw" field.
func (_c *UserCreate) SetRaw(v json.RawMessage) *UserCreate {
	_c.mutation.SetRaw(v)
	return _c
}

// SetDirs sets the "dirs" field.
func (_c *UserCreate) SetDirs(v []http.Dir) *UserCreate {
	_c.mutation.SetDirs(v)
	return _c
}

// SetInts sets the "ints" field.
func (_c *UserCreate) SetInts(v []int) *UserCreate {
	_c.mutation.SetInts(v)
	return _c
}

// SetFloats sets the "floats" field.
func (_c *UserCreate) SetFloats(v []float64) *UserCreate {
	_c.mutation.SetFloats(v)
	return _c
}

// SetStrings sets the "strings" field.
func (_c *UserCreate) SetStrings(v []string) *UserCreate {
	_c.mutation.SetStrings(v)
	return _c
}

// SetIntsValidate sets the "ints_validate" field.
func (_c *UserCreate) SetIntsValidate(v []int) *UserCreate {
	_c.mutation.SetIntsValidate(v)
	return _c
}

// SetFloatsValidate sets the "floats_validate" field.
func (_c *UserCreate) SetFloatsValidate(v []float64) *UserCreate {
	_c.mutation.SetFloatsValidate(v)
	return _c
}

// SetStringsValidate sets the "strings_validate" field.
func (_c *UserCreate) SetStringsValidate(v []string) *UserCreate {
	_c.mutation.SetStringsValidate(v)
	return _c
}

// SetAddr sets the "addr" field.
func (_c *UserCreate) SetAddr(v schema.Addr) *UserCreate {
	_c.mutation.SetAddr(v)
	return _c
}

// SetNillableAddr sets the "addr" field if the given value is not nil.
func (_c *UserCreate) SetNillableAddr(v *schema.Addr) *UserCreate {
	if v != nil {
		_c.SetAddr(*v)
	}
	return _c
}

// SetUnknown sets the "unknown" field.
func (_c *UserCreate) SetUnknown(v any) *UserCreate {
	_c.mutation.SetUnknown(v)
	return _c
}

// Mutation returns the UserMutation object of the builder.
func (_c *UserCreate) Mutation() *UserMutation {
	return _c.mutation
}

// Save creates the User in the database.
func (_c *UserCreate) Save(ctx context.Context) (*User, error) {
	_c.defaults()
	return withHooks(ctx, _c.sqlSave, _c.mutation, _c.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (_c *UserCreate) SaveX(ctx context.Context) *User {
	v, err := _c.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (_c *UserCreate) Exec(ctx context.Context) error {
	_, err := _c.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (_c *UserCreate) ExecX(ctx context.Context) {
	if err := _c.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (_c *UserCreate) defaults() {
	if _, ok := _c.mutation.Dirs(); !ok {
		v := user.DefaultDirs()
		_c.mutation.SetDirs(v)
	}
	if _, ok := _c.mutation.Ints(); !ok {
		v := user.DefaultInts
		_c.mutation.SetInts(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (_c *UserCreate) check() error {
	if _, ok := _c.mutation.Dirs(); !ok {
		return &ValidationError{Name: "dirs", err: errors.New(`ent: missing required field "User.dirs"`)}
	}
	if v, ok := _c.mutation.IntsValidate(); ok {
		if err := user.IntsValidateValidator(v); err != nil {
			return &ValidationError{Name: "ints_validate", err: fmt.Errorf(`ent: validator failed for field "User.ints_validate": %w`, err)}
		}
	}
	if v, ok := _c.mutation.FloatsValidate(); ok {
		if err := user.FloatsValidateValidator(v); err != nil {
			return &ValidationError{Name: "floats_validate", err: fmt.Errorf(`ent: validator failed for field "User.floats_validate": %w`, err)}
		}
	}
	if v, ok := _c.mutation.StringsValidate(); ok {
		if err := user.StringsValidateValidator(v); err != nil {
			return &ValidationError{Name: "strings_validate", err: fmt.Errorf(`ent: validator failed for field "User.strings_validate": %w`, err)}
		}
	}
	return nil
}

func (_c *UserCreate) sqlSave(ctx context.Context) (*User, error) {
	if err := _c.check(); err != nil {
		return nil, err
	}
	_node, _spec := _c.createSpec()
	if err := sqlgraph.CreateNode(ctx, _c.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	_node.ID = int(id)
	_c.mutation.id = &_node.ID
	_c.mutation.done = true
	return _node, nil
}

func (_c *UserCreate) createSpec() (*User, *sqlgraph.CreateSpec) {
	var (
		_node = &User{config: _c.config}
		_spec = sqlgraph.NewCreateSpec(user.Table, sqlgraph.NewFieldSpec(user.FieldID, field.TypeInt))
	)
	if value, ok := _c.mutation.T(); ok {
		_spec.SetField(user.FieldT, field.TypeJSON, value)
		_node.T = value
	}
	if value, ok := _c.mutation.URL(); ok {
		_spec.SetField(user.FieldURL, field.TypeJSON, value)
		_node.URL = value
	}
	if value, ok := _c.mutation.URLs(); ok {
		_spec.SetField(user.FieldURLs, field.TypeJSON, value)
		_node.URLs = value
	}
	if value, ok := _c.mutation.Raw(); ok {
		_spec.SetField(user.FieldRaw, field.TypeJSON, value)
		_node.Raw = value
	}
	if value, ok := _c.mutation.Dirs(); ok {
		_spec.SetField(user.FieldDirs, field.TypeJSON, value)
		_node.Dirs = value
	}
	if value, ok := _c.mutation.Ints(); ok {
		_spec.SetField(user.FieldInts, field.TypeJSON, value)
		_node.Ints = value
	}
	if value, ok := _c.mutation.Floats(); ok {
		_spec.SetField(user.FieldFloats, field.TypeJSON, value)
		_node.Floats = value
	}
	if value, ok := _c.mutation.Strings(); ok {
		_spec.SetField(user.FieldStrings, field.TypeJSON, value)
		_node.Strings = value
	}
	if value, ok := _c.mutation.IntsValidate(); ok {
		_spec.SetField(user.FieldIntsValidate, field.TypeJSON, value)
		_node.IntsValidate = value
	}
	if value, ok := _c.mutation.FloatsValidate(); ok {
		_spec.SetField(user.FieldFloatsValidate, field.TypeJSON, value)
		_node.FloatsValidate = value
	}
	if value, ok := _c.mutation.StringsValidate(); ok {
		_spec.SetField(user.FieldStringsValidate, field.TypeJSON, value)
		_node.StringsValidate = value
	}
	if value, ok := _c.mutation.Addr(); ok {
		_spec.SetField(user.FieldAddr, field.TypeJSON, value)
		_node.Addr = value
	}
	if value, ok := _c.mutation.Unknown(); ok {
		_spec.SetField(user.FieldUnknown, field.TypeJSON, value)
		_node.Unknown = value
	}
	return _node, _spec
}

// UserCreateBulk is the builder for creating many User entities in bulk.
type UserCreateBulk struct {
	config
	err      error
	builders []*UserCreate
}

// Save creates the User entities in the database.
func (_c *UserCreateBulk) Save(ctx context.Context) ([]*User, error) {
	if _c.err != nil {
		return nil, _c.err
	}
	specs := make([]*sqlgraph.CreateSpec, len(_c.builders))
	nodes := make([]*User, len(_c.builders))
	mutators := make([]Mutator, len(_c.builders))
	for i := range _c.builders {
		func(i int, root context.Context) {
			builder := _c.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*UserMutation)
				if !ok {
					return nil, fmt.Errorf("unexpected mutation type %T", m)
				}
				if err := builder.check(); err != nil {
					return nil, err
				}
				builder.mutation = mutation
				var err error
				nodes[i], specs[i] = builder.createSpec()
				if i < len(mutators)-1 {
					_, err = mutators[i+1].Mutate(root, _c.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, _c.driver, spec); err != nil {
						if sqlgraph.IsConstraintError(err) {
							err = &ConstraintError{msg: err.Error(), wrap: err}
						}
					}
				}
				if err != nil {
					return nil, err
				}
				mutation.id = &nodes[i].ID
				if specs[i].ID.Value != nil {
					id := specs[i].ID.Value.(int64)
					nodes[i].ID = int(id)
				}
				mutation.done = true
				return nodes[i], nil
			})
			for i := len(builder.hooks) - 1; i >= 0; i-- {
				mut = builder.hooks[i](mut)
			}
			mutators[i] = mut
		}(i, ctx)
	}
	if len(mutators) > 0 {
		if _, err := mutators[0].Mutate(ctx, _c.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (_c *UserCreateBulk) SaveX(ctx context.Context) []*User {
	v, err := _c.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (_c *UserCreateBulk) Exec(ctx context.Context) error {
	_, err := _c.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (_c *UserCreateBulk) ExecX(ctx context.Context) {
	if err := _c.Exec(ctx); err != nil {
		panic(err)
	}
}
