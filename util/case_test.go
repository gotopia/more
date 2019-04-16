package util

import (
	"reflect"
	"testing"
)

func TestToLowerSlice(t *testing.T) {
	type args struct {
		s []string
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{"string slice should be converted to lowercase", args{s: []string{"Username", "Password"}}, []string{"username", "password"}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ToLowerSlice(tt.args.s); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ToLowerSlice() = %v, want %v", got, tt.want)
			}
		})
	}
}
