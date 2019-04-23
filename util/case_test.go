package util

import (
	"reflect"
	"testing"
)

func TestToSnakeSlice(t *testing.T) {
	type args struct {
		s []string
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{
			"string slice should be converted to lowercase", args{
				s: []string{"Username", "Password", "IsBot"},
			}, []string{"username", "password", "is_bot"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ToSnakeSlice(tt.args.s); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ToSnakeSlice() = %v, want %v", got, tt.want)
			}
		})
	}
}
