package main

import (
	"context"
	"log"
	"net"

	"google.golang.org/grpc"
	"istio.io/api/mixer/adapter/model/v1beta1"
	"istio.io/istio/ditto"
	"istio.io/istio/ditto/config"
)

type adapter struct{}

func (adapter) HandleDitto(ctx context.Context, req *ditto.HandleDittoRequest) (*v1beta1.CheckResult, error) {
	cfg := &config.Params{}
	if err := cfg.Unmarshal(req.AdapterConfig.Value); err != nil || cfg.ApiRequirements[0].Endpoints[0].Method == "" {
		log.Printf("error unmarshalling adapter config: %v %#v", err, cfg)
		return nil, err
	}

	return &v1beta1.CheckResult{}, nil
}

func main() {
	listener, err := net.Listen("tcp", ":9090")
	if err != nil {
		panic(err)
	}
	server := grpc.NewServer()
	ditto.RegisterHandleDittoServiceServer(server, adapter{})
	server.Serve(listener)
}
