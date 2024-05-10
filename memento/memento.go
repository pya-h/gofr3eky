package memento

import (
	"errors"
	"gofr3eky/fields"
	"strconv"
)

type Memento struct {
	Fields map[string]*fields.Field
}

func New(args ...fields.Any) (*Memento, error) {
	var memento Memento
	var err error = nil
	memento.Fields = make(map[string]*fields.Field)
	if len(args) > 0 {
		err = memento.DefineSomeFields(args)
	}

	return &memento, err
}

func (memento *Memento) DefineSomeFields(fields_data []fields.Any) error {
	count := len(fields_data)
	if count%2 != 0 {
		return errors.New("the input should be empty or like name, field pairs respectively")
	}
	i := 0
	failure_count := 0
	for i < count {
		if name, ok := fields_data[i].(string); ok {
			if err := memento.DefineField(name, fields_data[i+1]); err != nil {
				failure_count++
			}
		}
		i += 2
	}
	return nil
}

func (memento *Memento) DefineField(name string, value fields.Any) error {

	if field, ok := value.(*fields.Field); ok {
		(*memento).Fields[name] = field
	} else if field, ok := value.(fields.Field); ok {
		(*memento).Fields[name] = &field
	} else if raw_data, ok := value.(string); ok {
		field, err := fields.New(raw_data)
		if err != nil {
			return err
		}
		(*memento).Fields[name] = field
	} else {
		return errors.New("unsupported init argument list for memento initializing")
	}
	return nil
}

func (memento *Memento) Get(name string) (*fields.Field, error) {
	if field, exists := memento.Fields[name]; exists {
		return field, nil
	}
	return nil, errors.New("no such Field asshole")
}

func (memento *Memento) Text(name string, value string) error {
	memento.Fields[name] = &fields.Field{Type: fields.Variant(fields.VariantText), Text: value}
	return nil
}

func (memento *Memento) Decimal(name string, value string) error {
	if v, err := strconv.ParseFloat(value, 64); err == nil {
		memento.Fields[name] = &fields.Field{Type: fields.Variant(fields.VariantText), Decimal: v}
	}
	return errors.New("not a decimal")
}
