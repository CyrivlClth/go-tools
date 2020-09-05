// package serialnum
package serialnum

import (
	"fmt"
)

func ExampleGenerator_NextIDNew() {
	idG := New(9)
	for i := 0; i < 12; i++ {
		fmt.Println(idG.NextID())
	}

	// Output: 0 <nil>
	// 1 <nil>
	// 2 <nil>
	// 3 <nil>
	// 4 <nil>
	// 5 <nil>
	// 6 <nil>
	// 7 <nil>
	// 8 <nil>
	// 9 <nil>
	// 0 <nil>
	// 1 <nil>
}
