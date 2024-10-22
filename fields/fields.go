package fields

import (
	"strconv"
)

type Variant uint8

const (
	VariantDecimal Variant = iota
	VariantText
	VariantWhatTheFux
	// ...
)

type Field struct {
	Type    Variant
	Decimal float64
	Text    string
	Wtf     Any
	Name    string
}

func New(x interface{}) (*Field, error) {
	// if v, ok := x.(int64); ok {
	// 	return &Field{Type: Variant(decimal), Decimal: (float64)v}, nil
	// }
	if v, ok := x.(float64); ok {
		return &Field{Type: Variant(VariantDecimal), Decimal: v}, nil
	}
	if str, ok := x.(string); ok {
		num, err := strconv.ParseFloat(str, 64) // for now numbers only support float
		if err != nil {
			return &Field{Type: Variant(VariantText), Text: str}, nil
		}
		return &Field{Type: Variant(VariantDecimal), Decimal: num}, nil

	}

	return &Field{Type: Variant(VariantWhatTheFux), Wtf: x}, nil

}

func (field *Field) Value() Any {
	if field.Type == VariantDecimal {
		return field.Decimal
	}

	if field.Type == VariantText {
		return field.Text
	}

	return field.Wtf
}

func NewDecimal(value float64) *Field {
	return &Field{Type: VariantDecimal, Decimal: value}
}

func NewText(value string) *Field {
	return &Field{Type: VariantDecimal, Text: value}
}

func (lhs *Field) Assign(rhs *Field) {
	lhs.Type = rhs.Type
	switch lhs.Type {
	case VariantDecimal:
		lhs.Decimal = rhs.Decimal
	case VariantText:
		lhs.Text = rhs.Text
	case VariantWhatTheFux:
		lhs.Wtf = rhs.Wtf
	}

	// ** Important: This way when a field changes type, it holds its previous type value.
}
