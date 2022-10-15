package main

import "reflect"

func walk(x interface{}, fn func(input string)) {
	val := reflect.ValueOf(x)

	// NumField can't be used on a pointer value, will need to extract underlying value
	if val.Kind() == reflect.Ptr {
		val = val.Elem()
	}

	for i := 0; i < val.NumField(); i++ {
		field := val.Field(i)

		switch field.Kind() {
		case reflect.String:
			fn(field.String())
		case reflect.Struct:
			walk(field.Interface(), fn)
		}
	}
}
