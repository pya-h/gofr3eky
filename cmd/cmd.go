package main

import (
	"fmt"

	fields "gofr3eky/fields"
)

func main() {
	x, _ := fields.NewField("10s")
	fmt.Println(x)
}
