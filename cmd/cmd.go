package main

import (
	"gofr3eky/asker"
	"gofr3eky/blocks"
	"log"
)

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
		main.HandleStatement(main.NextLiny(statement))
	}
}
