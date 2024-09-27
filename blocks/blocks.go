package blocks

import (
	"errors"
	"fmt"
	"gofr3eky/fields"
	"gofr3eky/memento"
	"log"
	"runtime"
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

func showMemoryStats() {
	var memStats runtime.MemStats
	runtime.ReadMemStats(&memStats)

	// Print total memory allocated (in bytes)
	fmt.Printf("Memory Alloc = %v kb\n", memStats.Alloc/1024.0)
	fmt.Printf("Total Alloc = %v kb\n", memStats.TotalAlloc/1024.0)
	fmt.Printf("Heap Alloc = %v kb\n", memStats.HeapAlloc/1024.0)
	fmt.Printf("Sys = %v kb\n", memStats.Sys)
}

func (block *Block) Process(liny *Liny) {
	liny.Parse(block.Memento)
	count := len(liny.Terms)

	switch liny.Terms[0].Value() {
	case "#":
		liny.Do(1, uint(count))
		count = len(liny.Terms)
		for i := 1; i < count; i++ {
			fmt.Print(liny.Terms[i].Value(), " ")
		}
		fmt.Print("\n")
	case "#M":
		showMemoryStats()
	default:
		// Define new field
		liny.Do(1, uint(count))

		if liny.Terms[0].Type == fields.VariantText {
			count = len(liny.Terms)
			if count > 2 {
				// TODO: It's probably a string or array
				if liny.Terms[0].Type == fields.VariantText && liny.Terms[1].Text == ":" {
					str := fmt.Sprint(liny.Terms[2].Value())
					for i := 3; i < count; i++ {
						str += fmt.Sprint(" ", liny.Terms[i].Value())
					}
					if _, err := block.Memento.DefineField(liny.Terms[0].Text, str); err != nil {
						log.Println("failed defining new field:", liny.Terms[0].Text, ";", err)
					}
				}
			} else if count == 2 {
				if _, err := block.Memento.DefineField(liny.Terms[0].Text, liny.Terms[1]); err != nil {
					log.Println("failed defining new field:", liny.Terms[0].Text, ";", err)
				}
			}
		} else {
			// Its a simple math statement
			// TODO: Calculate it and save in drafts
		}

	}
	// FIXME: After running each command, used memory increase as +[4-9]
}
