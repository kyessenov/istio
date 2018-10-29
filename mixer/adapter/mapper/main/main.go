package main

import (
	"net"

	"google.golang.org/grpc"
	"istio.io/istio/mixer/adapter/mapper"
)

func main() {
	listener, err := net.Listen("tcp", ":9070")
	if err != nil {
		panic(err)
	}
	server := grpc.NewServer()
	mapper.RegisterHandleMapperServiceServer(server, mapper.MyAdapter{})
	server.Serve(listener)
}
