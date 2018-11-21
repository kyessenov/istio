package main

import (
	"net"

	"google.golang.org/grpc"
	"istio.io/istio/userkey"
)

func main() {
	listener, err := net.Listen("tcp", ":9070")
	if err != nil {
		panic(err)
	}
	server := grpc.NewServer()
	userkey.RegisterHandleUserkeyServiceServer(server, userkey.Userkey{})
	server.Serve(listener)
}
