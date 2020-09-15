// package convert
package convert

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_camelToSnake(t *testing.T) {
	tests := []struct {
		name string
		args string
		want string
	}{
		{"small camel", "nameField", "name_field"},
		{"big camel", "NameField", "name_field"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := CamelToSnake(tt.args)
			assert.Equal(t, tt.want, got)
		})
	}
}

func Test_snakeToCamel(t *testing.T) {
	type args struct {
		s     string
		first bool
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{"snake", args{"name_field", false}, "nameField"},
		{"snake", args{"name_field", true}, "NameField"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := SnakeToCamel(tt.args.s, tt.args.first)
			assert.Equal(t, tt.want, got)
		})
	}
}
