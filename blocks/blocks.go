package blocks

import (
	"errors"
	"fmt"
	"gofr3eky/fields"
	"gofr3eky/memento"
	"log"
	"strings"
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
	Content    []*Liny
	Memento    *memento.Memento
	Variant    Variant
	Identifier string
}

func Main() (*Block, error) {
	mem, err := memento.New()
	if err != nil {
		return nil, err
	}
	return &Block{
		Memento:    mem,
		Content:    make([]*Liny, 0),
		Variant:    Variant(main),
		Identifier: "main",
	}, nil
}

func (block *Block) NextLiny(statement string) *Liny {
	liny := NewLiny(statement)
	block.Content = append(block.Content, liny)
	return liny
}

func (block *Block) ExtractStatements(statements []fields.Any) error {
	// this function extracts the list of initial Linies provided for a block;
	// arguments passed in any argument-pass format supported.
	if !fields.IsCollection(statements) {
		// as: func(arg1, arg2, arg3, ...)
		for _, statement := range statements {
			if v, ok := statement.(string); ok { // as: func("line1", "line2", "line3", ...)
				block.NextLiny(v)
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
				block.NextLiny(string_list[i])
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
	return &block, block.ExtractStatements(statements)
}

func (block *Block) Process(liny *Liny) {
	liny.Parse(block.Memento)
	terms := strings.Fields(liny.Statement)
	count := len(terms)
	switch terms[0] {
	case "#":
		for i := 1; i < count; i++ {
			// TODO: Check the next term is operator, evaluate until next one isnt operator.
			if field, err := block.Memento.Get(terms[i]); err == nil {
				fmt.Print(field.Value, " ")
			} else {
				fmt.Print("wtf ")
			}
		}
		fmt.Print("\n")
	default:
		// Define new field
		liny.Statement = liny.Statement[len(terms[0])+1 : len(liny.Statement)-1]

		// For test:
		if err := block.Memento.DefineField(terms[0], liny.Statement); err != nil {
			log.Println("failed defining new field:", terms[0], ";", err)
		} else {
			v, _ := block.Memento.Get(terms[0])
			fmt.Println(v, block.Memento)
		}
	}
}
