package memento


import (
	"fields"
)

type Memento struct {
	fields	map[string]*fields.Field

}


func New(init_fields ...*fields.Field) *Memento {
	var memento Memento
	for _, field := range init_fields {
		memento
	}
}
