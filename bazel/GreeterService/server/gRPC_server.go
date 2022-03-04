package server

import (
	greeter "GreeterService/Gen/GoGreeterService"
	"context"
)

type GreeterServer struct {
	//greeter.UnimplementedGreeterServiceServer
	//l *log.Logger
	// constructor
}

func NewGreeter() *GreeterServer {
	return &GreeterServer{}
}
func (grt *GreeterServer) SayHello(ctx context.Context, req *greeter.HelloRequest) (*greeter.HelloResponse, error) {
	response := &greeter.HelloResponse{Name: "Hello " + req.GetName()}
	return response, nil
}
