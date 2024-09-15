package calculus

import (
	"errors"
	"gofr3eky/fields"
)

func AddUp(a *fields.Field, b *fields.Field) (*fields.Field, error) {
	if a.Type == fields.VariantDecimal {
		if b.Type == fields.VariantDecimal {
			return fields.NewDecimal(a.Decimal + b.Decimal), nil
		}
		return nil, errors.New("second operand is not numeric")
	}
	return nil, errors.New("first operand is not numeric")
}

func Subtract(a *fields.Field, b *fields.Field) (*fields.Field, error) {
	if a.Type == fields.VariantDecimal {
		if b.Type == fields.VariantDecimal {
			return fields.NewDecimal(a.Decimal - b.Decimal), nil
		}
		return nil, errors.New("second operand is not numeric")
	}
	return nil, errors.New("first operand is not numeric")
}

func Multiply(a *fields.Field, b *fields.Field) (*fields.Field, error) {
	if a.Type == fields.VariantDecimal {
		if b.Type == fields.VariantDecimal {
			return fields.NewDecimal(a.Decimal * b.Decimal), nil
		}
		return nil, errors.New("second operand is not numeric")
	}
	return nil, errors.New("first operand is not numeric")
}

func DivideIn(a *fields.Field, b *fields.Field) (*fields.Field, error) {
	if a.Type == fields.VariantDecimal {
		if b.Type == fields.VariantDecimal {
			return fields.NewDecimal(a.Decimal / b.Decimal), nil
		}
		return nil, errors.New("second operand is not numeric")
	}
	return nil, errors.New("first operand is not numeric")
}
