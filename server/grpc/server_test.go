// package main
package main

import (
	"context"
	"fmt"
	"log"

	"google.golang.org/grpc"

	snowflakepb "github.com/CyrivlClth/snowflake/server/grpc/proto"
)

func ExampleClient() {
	s, err := grpc.Dial("localhost:5550", grpc.WithInsecure())
	if err != nil {
		fmt.Println(err)
	}

	client := snowflakepb.NewSnowflakeServiceClient(s)
	id, err := client.NextID(context.TODO(), &snowflakepb.NextIDReq{})
	if err != nil {
		fmt.Println(err)
	}
	log.Println(id.GetId())
	//Output:
	//
}
