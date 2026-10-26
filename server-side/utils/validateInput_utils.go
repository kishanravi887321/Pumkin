package utils

import (
	"reflect"
	"strings"
)

func CleanInput(data any) {
	v := reflect.ValueOf(data)

	// Must be a non-nil pointer to a struct.
	if v.Kind() != reflect.Ptr || v.IsNil() {
		return
	}

	v = v.Elem()

	if v.Kind() != reflect.Struct {
		return
	}

	for i := 0; i < v.NumField(); i++ {
		field := v.Field(i)

		// Only modify fields that can be set.
		if field.Kind() == reflect.String && field.CanSet() {
			field.SetString(strings.TrimSpace(field.String()))
		}
	}
}
