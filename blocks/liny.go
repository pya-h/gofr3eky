package blocks

import (
	"gofr3eky/fields"
)

type Liny struct {
	Statement string
	Draft     *fields.Field // last run result of the line
	Tag       string
}

func (line Liny) Do() (*fields.Field, error) {

	return line.Draft, nil
}

func NextLiny(statement string) *Liny {
	return &Liny{
		Statement: statement,
		Draft:     nil,
		Tag:       "",
	}
}
