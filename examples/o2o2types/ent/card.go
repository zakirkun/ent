// Copyright 2019-present Facebook Inc. All rights reserved.
// This source code is licensed under the Apache 2.0 license found
// in the LICENSE file in the root directory of this source tree.

// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/examples/o2o2types/ent/card"
	"entgo.io/ent/examples/o2o2types/ent/user"
)

// Card is the model entity for the Card schema.
type Card struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// Expired holds the value of the "expired" field.
	Expired time.Time `json:"expired,omitempty"`
	// Number holds the value of the "number" field.
	Number string `json:"number,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the CardQuery when eager-loading is set.
	Edges        CardEdges `json:"edges"`
	user_card    *int
	selectValues sql.SelectValues
}

// CardEdges holds the relations/edges for other nodes in the graph.
type CardEdges struct {
	// Owner holds the value of the owner edge.
	Owner *User `json:"owner,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [1]bool
}

// OwnerOrErr returns the Owner value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e CardEdges) OwnerOrErr() (*User, error) {
	if e.Owner != nil {
		return e.Owner, nil
	} else if e.loadedTypes[0] {
		return nil, &NotFoundError{label: user.Label}
	}
	return nil, &NotLoadedError{edge: "owner"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Card) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case card.FieldID:
			values[i] = new(sql.NullInt64)
		case card.FieldNumber:
			values[i] = new(sql.NullString)
		case card.FieldExpired:
			values[i] = new(sql.NullTime)
		case card.ForeignKeys[0]: // user_card
			values[i] = new(sql.NullInt64)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Card fields.
func (_m *Card) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case card.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			_m.ID = int(value.Int64)
		case card.FieldExpired:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field expired", values[i])
			} else if value.Valid {
				_m.Expired = value.Time
			}
		case card.FieldNumber:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field number", values[i])
			} else if value.Valid {
				_m.Number = value.String
			}
		case card.ForeignKeys[0]:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for edge-field user_card", value)
			} else if value.Valid {
				_m.user_card = new(int)
				*_m.user_card = int(value.Int64)
			}
		default:
			_m.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the Card.
// This includes values selected through modifiers, order, etc.
func (_m *Card) Value(name string) (ent.Value, error) {
	return _m.selectValues.Get(name)
}

// QueryOwner queries the "owner" edge of the Card entity.
func (_m *Card) QueryOwner() *UserQuery {
	return NewCardClient(_m.config).QueryOwner(_m)
}

// Update returns a builder for updating this Card.
// Note that you need to call Card.Unwrap() before calling this method if this Card
// was returned from a transaction, and the transaction was committed or rolled back.
func (_m *Card) Update() *CardUpdateOne {
	return NewCardClient(_m.config).UpdateOne(_m)
}

// Unwrap unwraps the Card entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (_m *Card) Unwrap() *Card {
	_tx, ok := _m.config.driver.(*txDriver)
	if !ok {
		panic("ent: Card is not a transactional entity")
	}
	_m.config.driver = _tx.drv
	return _m
}

// String implements the fmt.Stringer.
func (_m *Card) String() string {
	var builder strings.Builder
	builder.WriteString("Card(")
	builder.WriteString(fmt.Sprintf("id=%v, ", _m.ID))
	builder.WriteString("expired=")
	builder.WriteString(_m.Expired.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("number=")
	builder.WriteString(_m.Number)
	builder.WriteByte(')')
	return builder.String()
}

// Cards is a parsable slice of Card.
type Cards []*Card
