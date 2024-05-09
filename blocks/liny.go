package blocks

import (
	"gofr3eky/fields"
	"gofr3eky/memento"
	"strings"
)

type Liny struct {
	Statement string
	Terms     []fields.Any
	Cursor    int16
	Draft     *fields.Field // last run result of the line
	Tag       string
}

func NewLiny(statement string) *Liny {
	return &Liny{
		Statement: statement,
		Draft:     nil,
		Tag:       "",
	}
}

func (liny *Liny) Parse(source *memento.Memento) {
	// for now assume its only math lin6
	terms := strings.Fields(liny.Statement)
	liny.Cursor = 0
	liny.Terms = make([]fields.Any, 0)
	for _, v := range terms {
		if field, err := source.Get(v); err != nil {
			liny.Terms = append(liny.Terms, v)
		} else {
			liny.Terms = append(liny.Terms, field)
		}
	}
}
func (liny *Liny) Do() {

}
