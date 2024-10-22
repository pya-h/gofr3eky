package calculus

import (
	"errors"
	"gofr3eky/fields"
	"math"
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

func RaiseTo(a *fields.Field, b *fields.Field) (*fields.Field, error) {
	if a.Type == fields.VariantDecimal {
		if b.Type == fields.VariantDecimal {
			return fields.NewDecimal(math.Pow(a.Decimal, b.Decimal)), nil
		}
		return nil, errors.New("second operand is not numeric")
	}
	return nil, errors.New("first operand is not numeric")
}

func MultiplyIn2ndRootOf(a *fields.Field, b *fields.Field) (*fields.Field, error) {
	if a.Type == fields.VariantDecimal {
		if b.Type == fields.VariantDecimal {
			return fields.NewDecimal(a.Decimal * math.Pow(b.Decimal, 0.5)), nil
		}
		return nil, errors.New("there must be a number inside root operator")
	}
	return nil, errors.New("multiplier operand is not numeric")
}

func MultiplyInNthRootOf(a *fields.Field, b *fields.Field, n float64) (*fields.Field, error) {
	if a.Type == fields.VariantDecimal {
		if b.Type == fields.VariantDecimal {
			if n > 0 {
				return fields.NewDecimal(a.Decimal * math.Pow(b.Decimal, 1/n)), nil
			}
			return nil, errors.New("the root order must be a positive number")
		}
		return nil, errors.New("there must be a number inside root operator")
	}
	return nil, errors.New("multiplier operand is not numeric")
}
