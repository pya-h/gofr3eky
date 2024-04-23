package fields

type Variant uint8

const (
	numeric Variant = iota
	text
	whatTheFux
	// ...
)

type Field struct {
	Type  Variant
	Value interface{}
}

func New(x interface{}) (*Field, error) {
	if _, ok := x.(int); ok {
		return &Field{Type: Variant(numeric), Value: x}, nil
	}
	if _, ok := x.(float64); ok {
		return &Field{Type: Variant(numeric), Value: x}, nil
	}
	if _, ok := x.(string); ok {
		return &Field{Type: Variant(text), Value: x}, nil
	}

	return &Field{Type: Variant(whatTheFux), Value: x}, nil

}

func Text(name string, value string) *Field {
	return &Field{Type: Variant(text), Value: value}
}

func Numeric(name string, value string) *Field {
	return &Field{Type: Variant(text), Value: value}
}
