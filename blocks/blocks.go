package blocks

import (
	"errors"
	"fmt"
	"gofr3eky/fields"
)

type Variant uint8

const (
	main Variant = iota
	check
	loop
	method
	concept

	// ...
)

type Block struct {
	Content []*Liny

	Variant    Variant
	Identifier string
}

func (block *Block) extractStatements(statements []fields.Any) error {
	// this function extracts the list of initial Linies provided for a block;
	// arguments passed in any argument-pass format supported.
	if !fields.IsCollection(statements) {
		// as: func(arg1, arg2, arg3, ...)
		for _, statement := range statements {
			if v, ok := statement.(string); ok { // as: func("line1", "line2", "line3", ...)
				block.Content = append(block.Content, NextLiny(v))
			} else if liny, ok := statement.(Liny); ok { // as: func(LinyObject1, LinyObject2, LinyObject3, ...)
				block.Content = append(block.Content, &liny)
			} else if liny_ptr, ok := statement.(*Liny); ok { // as: func(*LinyObjectPtr1, *LinyObjectPtr2, *LinyObjectPtr3, ...)
				block.Content = append(block.Content, liny_ptr)
			} else {
				return errors.New(fmt.Sprint("unsupported argument data type:", statement))
			}
		}
	} else {
		// as: func([arg1, arg2, ...])
		if len(statements) != 1 {
			return errors.New("invalid argument list")
		}
		if string_list, ok := statements[0].([]string); ok { // as: func(["line1", "line2", ...])
			for i := range string_list[0] {
				block.Content = append(block.Content, NextLiny(string_list[i]))
			}
		} else if linies, ok := statements[0].([]*Liny); ok { // as: func([*linyPtr1, *linyPtr2, ...])
			block.Content = linies
		} else if linies, ok := statements[0].([]Liny); ok { // as: func([*linyObj1, *linyObj2, ...])
			for i := range linies {
				block.Content = append(block.Content, &linies[i])
			}
		} else {
			return errors.New(fmt.Sprint("unsupported argument data type:", statements))
		}
	}
	return nil
}

func NextMethod(identifier string, statements ...fields.Any) (*Block, error) {
	// TODO: Test these sections and packages until now.
	block := Block{Variant: Variant(method), Identifier: identifier}
	return &block, block.extractStatements(statements)
}
