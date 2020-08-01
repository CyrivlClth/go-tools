package jsonobj

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestObject_Get(t *testing.T) {
	b := `{"a":"1","b":2,"c":{"d":{"e":"e"},"f":[1,2,3]}}`
	type args struct {
		data string
		key  string
	}
	type want struct {
		result string
		err    bool
	}
	tests := []struct {
		name string
		args args
		want want
	}{
		{"get str from dict", args{b, "a"}, want{`"1"`, false}},
		{"get int from dict", args{b, "b"}, want{`2`, false}},
		{"get dict from dict", args{b, "c"}, want{`{"d":{"e":"e"},"f":[1,2,3]}`, false}},
		{"get sub dict from dict", args{b, "c.d"}, want{`{"e":"e"}`, false}},
		{"get sub list from dict", args{b, "c.f"}, want{`[1,2,3]`, false}},
		{"get string in sub list from dict", args{b, "c.d.e"}, want{`"e"`, false}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := NewObject([]byte(tt.args.data)).Get(tt.args.key)
			assert.Equal(t, tt.want.err, err != nil)
			if err == nil {
				assert.Equal(t, tt.want.result, got.String())
			}
		})
	}
}

func ExampleObject_Get() {
	b := []byte(`{"a":"1","b":2,"c":{"d":{"e":"e"},"f":[1,2,3]}}`)
	o := NewObject(b)
	fmt.Println(o.Get("a"))
	fmt.Println(o.Get("b"))
	fmt.Println(o.Get("c"))
	fmt.Println(o.Get("c.f"))
	fmt.Println(o.GetString("a"))
	fmt.Println(o.GetString("c.d.e"))
	fmt.Println(o.GetInt("b"))
	fmt.Println(o.GetString("c"))
	fmt.Println(o.GetString("c.f"))

	//Output:
	//"1" <nil>
	//2 <nil>
	//{"d":{"e":"e"},"f":[1,2,3]} <nil>
	//[1,2,3] <nil>
	//1
	//e
	//2
}
