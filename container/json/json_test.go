package json

import (
	"fmt"
)

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
