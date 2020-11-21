package main

import (
	"reflect"
	"fmt"
)

func walk(x interface{}, fn func(input string)) {
	val := reflect.ValueOf(x)

	for i:=0; i<val.NumField(); i++ {
		field := val.Field(i)
			fmt.Println(field.Kind())
		if field.Kind() == reflect.String {
			fn(field.String())
		}
		if field.Kind() == reflect.Struct {
			walk(field.Interface(), fn)
		}
	}
}

func main() {
}
