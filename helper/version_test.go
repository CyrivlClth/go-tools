package helper

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestVersionCompare(t *testing.T) {
	tests := []struct {
		name string
		v1   string
		v2   string
		want int
	}{
		{"v1>v2", "v1", "v0", 1},
		{"v1>v2", "v2", "v1", 1},
		{"v1>v2", "1", "0", 1},
		{"v1>v2", "2", "1", 1},
		{"v1>v2", "v1.0.0", "v0.0.0", 1},
		{"v1>v2", "v2.0.0", "v1.0.0", 1},
		{"v1>v2", "1.0.0", "0.0.0", 1},
		{"v1>v2", "2.0.0", "1.0.0", 1},
		{"v1>v2", "v0.0.1", "v0.0.0", 1},
		{"v1>v2", "v0.0.2", "v0.0.1", 1},
		{"v1>v2", "0.0.1", "0.0.0", 1},
		{"v1>v2", "0.0.2", "0.0.1", 1},
		{"v1>v2", "v0.0.11", "v0.0.10", 1},
		{"v1>v2", "v0.0.12", "v0.0.11", 1},
		{"v1>v2", "0.0.11", "0.0.10", 1},
		{"v1>v2", "0.0.12", "0.0.11", 1},
		{"v1>v2", "v0.0.11", "v0.0.1", 1},
		{"v1>v2", "v0.0.12", "v0.0.9", 1},
		{"v1>v2", "0.0.11", "0.0.1", 1},
		{"v1>v2", "0.0.12", "0.0.9", 1},
		{"v1>v2",  "v0.1.0","v0.0.1", 1},
		{"v1>v2",  "v0.1.1","v0.0.2", 1},
		{"v1>v2", "0.1.0","0.0.1",  1},
		{"v1>v2", "0.1.1","0.0.2",  1},
		{"v1>v2",  "v0.1.10","v0.0.11", 1},
		{"v1>v2",  "v0.1.11","v0.0.12", 1},
		{"v1>v2", "0.1.10","0.0.11",  1},
		{"v1>v2", "0.1.11","0.0.12",  1},
		{"v1>v2",  "v0.1","v0.0.1", 1},
		{"v1>v2",  "v0.1","v0.0.2", 1},
		{"v1>v2", "0.1","0.0.1",  1},
		{"v1>v2", "0.1","0.0.2",  1},
		{"v1>v2",  "v0.1","v0.0.11", 1},
		{"v1>v2",  "v0.1","v0.0.12", 1},
		{"v1>v2", "0.1","0.0.11",  1},
		{"v1>v2", "0.1","0.0.12",  1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := VersionCompare(tt.v1, tt.v2)
			assert.Equal(t, tt.want, got)
		})
	}
}
