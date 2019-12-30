package util

import (
	"encoding/json"
	"reflect"
	"time"

	"github.com/golang/protobuf/protoc-gen-go/generator"
	"github.com/golang/protobuf/ptypes"
	"github.com/golang/protobuf/ptypes/timestamp"
	"github.com/volatiletech/null"
	"github.com/volatiletech/sqlboiler/strmangle"
	"github.com/volatiletech/sqlboiler/types"
)

// Update a from b by paths.
func Update(a interface{}, b interface{}, paths []string) {
	ea := reflect.ValueOf(a).Elem()
	eb := reflect.ValueOf(b).Elem()
	for _, path := range paths {
		fa := ea.FieldByName(strmangle.TitleCase(path))
		fb := eb.FieldByName(generator.CamelCase(path))

		if !fa.CanSet() {
			continue
		}

		if isNumber(fa) && isNumber(fb) {
			setNumber(fa, fb)
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
				continue
			}
		case reflect.TypeOf(time.Time{}):
			if ts, ok := fb.Interface().(*timestamp.Timestamp); ok {
				if t, err := ptypes.Timestamp(ts); err == nil {
					fa.Set(reflect.ValueOf(t))
					continue
				}
			}
		case reflect.TypeOf(types.JSON{}):
			if j, err := json.Marshal(fb.Interface()); err == nil {
				fa.Set(reflect.ValueOf(j))
				continue
			}
		default:
		}
	}
}

func isInt(v reflect.Value) bool {
	switch v.Kind() {
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return true
	default:
		return false
	}
}

func isUint(v reflect.Value) bool {
	switch v.Kind() {
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		return true
	default:
		return false
	}
}

func isNumber(v reflect.Value) bool {
	return isInt(v) || isUint(v)
}

func setNumber(a reflect.Value, b reflect.Value) {
	var i int64
	if isInt(b) {
		i = b.Int()
	}
	if isUint(b) {
		i = int64(b.Uint())
	}
	if isInt(a) {
		a.SetInt(i)
	}
	if isUint(a) {
		a.SetUint(uint64(i))
	}
}
