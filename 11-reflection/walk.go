package main

import "reflect"

func walk(x interface{}, fn func(input string)) {
	val := getValue(x)

	// previous code handled slice type first and then struct
	// better to see if it is a slice type, then struct type because we can send slice's item's into walk again and get the required fields
	switch val.Kind() {
	case reflect.Struct:
		for i := 0; i < val.NumField(); i++ {
			walk(val.Field(i).Interface(), fn)
		}
	case reflect.Slice:
		for i := 0; i < val.Len(); i++ {
			walk(val.Index(i).Interface(), fn)
		}
	case reflect.String:
		fn(val.String())
	}
}

func getValue(x interface{}) reflect.Value {
	val := reflect.ValueOf(x)

	if val.Kind() == reflect.Ptr {
		val = val.Elem()
	}

	return val
}
