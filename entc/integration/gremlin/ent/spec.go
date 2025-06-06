// Copyright 2019-present Facebook Inc. All rights reserved.
// This source code is licensed under the Apache 2.0 license found
// in the LICENSE file in the root directory of this source tree.

// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"

	"entgo.io/ent/dialect/gremlin"
)

// Spec is the model entity for the Spec schema.
type Spec struct {
	config
	// ID of the ent.
	ID string `json:"id,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the SpecQuery when eager-loading is set.
	Edges SpecEdges `json:"edges"`
}

// SpecEdges holds the relations/edges for other nodes in the graph.
type SpecEdges struct {
	// Card holds the value of the card edge.
	Card []*Card `json:"card,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [1]bool
}

// CardOrErr returns the Card value or an error if the edge
// was not loaded in eager-loading.
func (e SpecEdges) CardOrErr() ([]*Card, error) {
	if e.loadedTypes[0] {
		return e.Card, nil
	}
	return nil, &NotLoadedError{edge: "card"}
}

// FromResponse scans the gremlin response data into Spec.
func (_m *Spec) FromResponse(res *gremlin.Response) error {
	vmap, err := res.ReadValueMap()
	if err != nil {
		return err
	}
	var scan_m struct {
		ID string `json:"id,omitempty"`
	}
	if err := vmap.Decode(&scan_m); err != nil {
		return err
	}
	_m.ID = scan_m.ID
	return nil
}

// QueryCard queries the "card" edge of the Spec entity.
func (_m *Spec) QueryCard() *CardQuery {
	return NewSpecClient(_m.config).QueryCard(_m)
}

// Update returns a builder for updating this Spec.
// Note that you need to call Spec.Unwrap() before calling this method if this Spec
// was returned from a transaction, and the transaction was committed or rolled back.
func (_m *Spec) Update() *SpecUpdateOne {
	return NewSpecClient(_m.config).UpdateOne(_m)
}

// Unwrap unwraps the Spec entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (_m *Spec) Unwrap() *Spec {
	_tx, ok := _m.config.driver.(*txDriver)
	if !ok {
		panic("ent: Spec is not a transactional entity")
	}
	_m.config.driver = _tx.drv
	return _m
}

// String implements the fmt.Stringer.
func (_m *Spec) String() string {
	var builder strings.Builder
	builder.WriteString("Spec(")
	builder.WriteString(fmt.Sprintf("id=%v", _m.ID))
	builder.WriteByte(')')
	return builder.String()
}

// Specs is a parsable slice of Spec.
type Specs []*Spec

// FromResponse scans the gremlin response data into Specs.
func (_m *Specs) FromResponse(res *gremlin.Response) error {
	vmap, err := res.ReadValueMap()
	if err != nil {
		return err
	}
	var scan_m []struct {
		ID string `json:"id,omitempty"`
	}
	if err := vmap.Decode(&scan_m); err != nil {
		return err
	}
	for _, v := range scan_m {
		node := &Spec{ID: v.ID}
		*_m = append(*_m, node)
	}
	return nil
}
