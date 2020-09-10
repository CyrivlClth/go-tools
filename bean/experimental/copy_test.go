// package experimental
package experimental

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/suite"
)

type Alpha struct {
	Name string `json:"name"`
	Code uint32 `json:"code"`
}
type Beta struct {
	Name string `json:"name"`
	Code uint32 `json:"code"`
}

type Charlie struct {
	Name   string `json:"name"`
	Status uint32 `json:"code"`
}

type Delta struct {
	First  string `json:"first"`
	Status uint32 `json:"status"`
}

func ExampleCopyProperties() {
	a := Alpha{
		Name: "alpha",
		Code: 2,
	}
	b := Beta{}
	err := CopyProperties(&a, &b)
	fmt.Println(err)
	fmt.Println(b)

	//Output:
	//<nil>
	//{alpha 2}
}

type beanTestSuite struct {
	suite.Suite
}

func (s beanTestSuite) TestCopyProperties() {
	type args struct {
		src    interface{}
		target interface{}
	}
	type want struct {
		target interface{}
	}
	tests := []struct {
		name    string
		args    args
		want    want
		wantErr bool
	}{
		{"simple struct: same", args{Alpha{"alpha", 2}, &Alpha{}}, want{&Alpha{"alpha", 2}}, false},
		{"simple struct: different but same", args{Alpha{"alpha", 2}, &Beta{}}, want{&Beta{"alpha", 2}}, false},
		{"simple struct: different one same", args{Alpha{"alpha", 2}, &Charlie{}}, want{&Charlie{"alpha", 0}}, false},
		{"simple struct: different no same", args{Alpha{"alpha", 2}, &Delta{}}, want{&Delta{"", 0}}, false},
		{"simple map: to map", args{map[string]interface{}{"name": "1", "code": 2}, map[string]interface{}{}}, want{map[string]interface{}{"name": "1", "code": 2}}, false},
		{"simple map: to struct", args{map[string]interface{}{"name": "1", "code": 2}, &Alpha{}}, want{&Alpha{"1", 2}}, false},
	}

	for _, tt := range tests {
		s.Run(tt.name, func() {
			var err error
			s.NotPanics(func() {
				err = CopyProperties(&tt.args.src, &tt.args.target)
			})
			s.Equal(tt.wantErr, err != nil)
			if !tt.wantErr {
				s.Equal(tt.want.target, tt.args.target)
			}
		})
	}
}

func TestBeanSuite(t *testing.T) {
	suite.Run(t, new(beanTestSuite))
}
