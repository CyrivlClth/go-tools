package main

import (
	"fmt"

	"github.com/CyrivlClth/snowflake/snowflake"
)

func main() {
	generator, _ := snowflake.New(0, 0)
	id, _ := generator.NextID()
	fmt.Println(id)
}
