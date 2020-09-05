package main

import (
	"flag"
	"fmt"
	"log"
	"net"

	"google.golang.org/grpc"

	snowflakepb "github.com/CyrivlClth/go-tools/idgen/server/grpc/proto"
	"github.com/CyrivlClth/go-tools/idgen/server/grpc/service"
	"github.com/CyrivlClth/go-tools/idgen/snowflake"
)

var port = flag.Int("port", 5550, "server port")
var worker = flag.Int64("worker", 0, "worker id")
var dataCenter = flag.Int64("data", 0, "data center id")

func main() {
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", *port))
	if err != nil {
		log.Fatal(err)
	}

	s := grpc.NewServer()

	svc, err := snowflake.New(*worker, *dataCenter)
	if err != nil {
		log.Fatal(err)
	}

	snowflakepb.RegisterSnowflakeServiceServer(s, service.NewSnowflake(svc))
	if err := s.Serve(lis); err != nil {
		log.Fatal(err)
	}
}
