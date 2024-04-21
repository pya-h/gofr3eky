package fields

import (
	"fmt"
	"reflect"
)

type Any interface{}

func TypeName(any Any) string {
	typeOf := reflect.TypeOf(any)
	if typeOf.Kind().String() == "ptr" {
		return fmt.Sprint(typeOf.Elem().Name(), "ptr")
	}
	return reflect.TypeOf(&any).Elem().Name()
}

func TypeOf(any Any, type_name string) bool {
	return reflect.TypeOf(&any).Elem().Name() == type_name
}