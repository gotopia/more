package util

import (
	"reflect"
)

// InterfaceSlice converts any data slice to interface slice.
func InterfaceSlice(dataSlice interface{}) []interface{} {
	s := reflect.ValueOf(dataSlice)
	if s.Kind() != reflect.Slice {
		return []interface{}{}
	}
	interfaceSlice := make([]interface{}, s.Len())
	for i := 0; i < s.Len(); i++ {
		interfaceSlice[i] = s.Index(i).Interface()
	}
	return interfaceSlice
}
