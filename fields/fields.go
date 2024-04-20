package gofr3eky

import "errors"

type FieldType uint8

type Numerix interface{}

const (
	Numeric FieldType = iota
	Letters
	WhatTheFux
	// ...
)

type Field struct {
	Type  FieldType
	Value interface{}
}

var fields map[string]*Field

func Get(name string) (*Field, error) {
	if field, exists := fields[name]; exists {
		return field, nil
	}
	return nil, errors.New("no such Field asshole")
}

func NewField(x Numerix) (*Field, error) {
	if _, ok := x.(int); ok {
		return &Field{Type: FieldType(Numeric), Value: x}, nil
	}
	if _, ok := x.(float64); ok {
		return &Field{Type: FieldType(Numeric), Value: x}, nil
	}
	if _, ok := x.(string); ok {
		return &Field{Type: FieldType(Letters), Value: x}, nil
	}

	return &Field{Type: FieldType(WhatTheFux), Value: x}, nil

}
