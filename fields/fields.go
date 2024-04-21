package fields

type FieldType uint8

type Numerix interface{}

const (
	numeric FieldType = iota
	text
	whatTheFux
	// ...
)

type Field struct {
	Type  FieldType
	Value interface{}
}

func New(x interface{}) (*Field, error) {
	if _, ok := x.(int); ok {
		return &Field{Type: FieldType(numeric), Value: x}, nil
	}
	if _, ok := x.(float64); ok {
		return &Field{Type: FieldType(numeric), Value: x}, nil
	}
	if _, ok := x.(string); ok {
		return &Field{Type: FieldType(text), Value: x}, nil
	}

	return &Field{Type: FieldType(whatTheFux), Value: x}, nil

}

func Text(name string, value string) *Field {
	return &Field{Type: FieldType(text), Value: value}
}

func Numeric(name string, value string) *Field {
	return &Field{Type: FieldType(text), Value: value}
}
