package blocks

import (
	"gofr3eky/calculus"
	"gofr3eky/fields"
	"gofr3eky/memento"
	"strings"
)

type Liny struct {
	Statement string
	Terms     []*fields.Field
	Draft     *fields.Field // last run result of the line
	Tag       string
	Cursor    int16
}

func NewLiny(statement string) *Liny {
	return &Liny{
		Statement: statement,
		Draft:     nil,
		Tag:       "",
		Cursor:    0,
	}
}

func (liny *Liny) Parse(source *memento.Memento) {
	// for now assume its only math line
	terms := strings.Fields(liny.Statement)
	liny.Terms = make([]*fields.Field, 0)
	for _, v := range terms[1:] {
		if field, err := source.Get(v); err != nil {
			liny.Terms = append(liny.Terms, field)
		} else if field, err := fields.New(field); err == nil {
			liny.Terms = append(liny.Terms, field)
		}
	}
}

// TODO: Best way to check string is numeric
func (liny *Liny) Do() {
	count := len(liny.Terms)

	for i := 1; i < count-1; i++ {
		var result *fields.Field = nil
		var err error
		if liny.Terms[i+1].Type == fields.VariantText {
			switch liny.Terms[i].Text {
			case "+":
				// Todo: handle i+2 doesnt exist
				result, err = calculus.AddUp(liny.Terms[i], liny.Terms[i+2])
			case "-":
				result, err = calculus.Substract(liny.Terms[i], liny.Terms[i+2])
			case "*":
				result, err = calculus.Multiply(liny.Terms[i], liny.Terms[i+2])
			case "%":
				result, err = calculus.DevideIn(liny.Terms[i], liny.Terms[i+2])

			}
		}

		if result != nil && err == nil {
			liny.Terms[i] = result
			if i+3 < count {
				liny.Terms = append(liny.Terms[:i+1], liny.Terms[i+3:]...)
				count -= 2
			} else {
				liny.Terms = liny.Terms[:i+1]
			}
			i--
		}
	}
}
