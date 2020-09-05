package main

import (
	"fmt"

	"github.com/CyrivlClth/go-tools/idgen/snowflake"
)

func main() {
	// return error when workID or dataCenterID < 0
	generator, _ := snowflake.New(0, 0)
	// return error when system clock moved backwards
	id, _ := generator.NextID()
	fmt.Println(id)
}
