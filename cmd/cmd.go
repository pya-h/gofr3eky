package main

import (
	"fmt"
	"gofr3eky/asker"
	"gofr3eky/blocks"
	"gofr3eky/memento"
	"log"
	"strings"
)

func HandleStatement(current_memento *memento.Memento, statement string) {
	terms := strings.Fields(statement)
	count := len(terms)
	switch terms[0] {
	case "#":
		for i := 1; i < count; i++ {

		}
	default:
		// Define new field
		statement = statement[len(terms[0])+1 : len(statement)-1]

		// For test:
		if err := current_memento.DefineField(terms[0], statement); err != nil {
			log.Println("failed defining new field:", terms[0], ";", err)
		} else {
			v, _ := current_memento.Get(terms[0])
			fmt.Println(v, current_memento)
		}
	}
}

func main() {
	main, err := blocks.Main()
	if err != nil {
		panic(err)
	}

	for {
		results, errs := asker.Ask("> ")
		if errs[0] != nil {
			// handle error
			log.Println(err)
		}
		statement := results[0]
		HandleStatement(main.Memento, statement)
	}
}
