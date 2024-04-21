package memento

import (
	"errors"
	"gofr3eky/fields"
)

type Memento struct {
	Fields map[string]*fields.Field
}

func New(args ...fields.Any) (*Memento, error) {
	var memento Memento
	count := len(args)
	if count%2 != 0 {
		return nil, errors.New("the input should be empty or like name, field pairs respectively")
	}
	i := 0
	for i < count {
		if name, ok := args[i].(string); ok {
			if field, ok := args[i].(*fields.Field); ok {
				memento.Fields[name] = field
			} else if field, ok := args[i].(fields.Field); ok {
				memento.Fields[name] = &field
			} else {
				return nil, errors.New("unsupported init argument list for memento initializing")
			}
		}
		i += 2
	}
	return &memento, nil
}

func (memento Memento) Get(name string) (*fields.Field, error) {
	if field, exists := memento.Fields[name]; exists {
		return field, nil
	}
	return nil, errors.New("no such Field asshole")
}
