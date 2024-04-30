package main

import (
	"fmt"
	"gofr3eky/asker"
	"gofr3eky/memento"
	"log"
	"strings"
)

func HandleStatement(current_memento *memento.Memento, statement string) {
	terms := strings.Fields(statement)

	switch terms[0] {
	default:
		// Define new field
		statement = statement[len(terms[0])+1 : len(statement)-1]

		// For test:
		if err := current_memento.DefineField(terms[0], statement); err != nil {
			log.Println("failed defining new field:", terms[0], ";", err)
		} else {
			v, _ := current_memento.Get(terms[0])
			fmt.Println(v)
		}
	}
}

func main() {
	global, err := memento.New()
	if err != nil {
		log.Fatalln(err)
	}

	for {
		results, errs := asker.Ask("> ")
		if errs[0] != nil {
			// handle error
			log.Println(err)
		}
		statement := results[0]
		HandleStatement(global, statement)
	}
}
