package server

import (
	greeter "/Users/salman/Desktop/me/GO/bazel/GreeterService/Gen/GoGreeterService"
	"context"
)

type GreeterServer struct {
	greeter.UnimplementedGreeterServiceServer
}

func (grt *GreeterServer) SayHello(ctx context.Context, req *greeter.HelloRequest) (*greeter.HelloResponse, error) {
	response := &greeter.HelloResponse{Name: "Hello " + req.GetName()}
	return response, nil
}
