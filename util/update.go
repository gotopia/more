package util

import (
	"github.com/volatiletech/null"
	"reflect"
)

// Update a from b by paths.
func Update(a interface{}, b interface{}, paths []string) {
	ea := reflect.ValueOf(a).Elem()
	eb := reflect.ValueOf(b).Elem()
	for _, path := range paths {
		fa := ea.FieldByName(path)
		fb := eb.FieldByName(path)
		if !fa.CanSet() {
			continue
		}
		if fa.Type() == fb.Type() {
			fa.Set(fb)
			continue
		}
		switch fa.Type() {
		case reflect.TypeOf(null.String{}):
			if s, ok := fb.Interface().(string); ok {
				fa.Set(reflect.ValueOf(null.NewString(s, s != "")))
			}
		default:
		}
	}
}
