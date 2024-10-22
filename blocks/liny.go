package blocks

import (
	"gofr3eky/calculus"
	"gofr3eky/fields"
	"gofr3eky/memento"
	"strconv"
	"strings"
)

type Liny struct {
	Statement string
	Terms     []*fields.Field
	Tag       string
}

func NewLiny(statement string) *Liny {
	return &Liny{
		Statement: statement,
		Tag:       "",
	}
}

func (liny *Liny) Parse(source *memento.Memento) {
	// for now assume its only math line
	terms := strings.Fields(liny.Statement)
	liny.Terms = make([]*fields.Field, 0)
	for _, v := range terms {
		if field, err := source.Get(v); err == nil {
			liny.Terms = append(liny.Terms, field)
		} else if field, err := fields.New(v); err == nil {
			liny.Terms = append(liny.Terms, field)
		}
	}
}

// TODO: Best way to check string is numeric
func (liny *Liny) Do(start uint, end uint) {
	for i := start; i < end-1; i++ {
		var result *fields.Field = nil
		var err error
		if liny.Terms[i+1].Type == fields.VariantText {
			switch liny.Terms[i+1].Text {
			case "+":
				// Todo: handle i+2 doesn't exist
				result, err = calculus.AddUp(liny.Terms[i], liny.Terms[i+2])
			case "-":
				result, err = calculus.Subtract(liny.Terms[i], liny.Terms[i+2])
			case "*":
				result, err = calculus.Multiply(liny.Terms[i], liny.Terms[i+2])
			case "%":
				result, err = calculus.DivideIn(liny.Terms[i], liny.Terms[i+2])
			case "^":
				result, err = calculus.RaiseTo(liny.Terms[i], liny.Terms[i+2])
			case "!/":
				result, err = calculus.MultiplyIn2ndRootOf(liny.Terms[i], liny.Terms[i+2])
			default:
				middle_param := liny.Terms[i+1].Text
				param_length := len(middle_param)
				if middle_param[0] == '!' && middle_param[param_length-1] == '/' {
					if n, conversion_err := strconv.ParseFloat(middle_param[1:param_length-1], 64); conversion_err == nil {
						result, err = calculus.MultiplyInNthRootOf(liny.Terms[i], liny.Terms[i+2], n)
					}
				}
			}
		}

		if result != nil && err == nil {
			liny.Terms[i] = result
			if i+3 < end {
				liny.Terms = append(liny.Terms[:i+1], liny.Terms[i+3:]...)
				end -= 2
			} else {
				liny.Terms = liny.Terms[:i+1]
				break
			}
			i--
		}
	}
}
