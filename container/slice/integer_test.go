// package slice
package slice

import (
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
)

type _suiteForInteger struct {
	suite.Suite
}

func (s *_suiteForInteger) TestDistinct() {
	tests := []struct {
		name string
		args []int
		want []int
	}{
		{"no duplicate elem", []int{1, 2, 3}, []int{1, 2, 3}},
		{"duplicate elem many times", []int{1, 3, 3, 2, 5, 3}, []int{1, 3, 2, 5}},
		{"start with duplicate elem", []int{3, 3, 2, 1}, []int{3, 2, 1}},
		{"empty", []int{}, []int{}},
		{"nil", nil, nil},
		{"only one", []int{1}, []int{1}},
	}
	for _, tt := range tests {
		s.Run(tt.name, func() {
			got := Integer(tt.args).Distinct()
			require.EqualValues(s.T(), tt.want, got)
		})
	}
}

func TestInteger(t *testing.T) {
	suite.Run(t, &_suiteForInteger{})
}
