package main

import (
	"fmt"

	fields "gofr3eky/fields"
)

func test(x interface{}) interface{} {
	if v, ok := x.(*fields.Field); ok {
		return v
	}
	return nil
}

func main() {
	x, _ := fields.New("10s")

	fmt.Println(x)
	fmt.Println(test(x))
}
