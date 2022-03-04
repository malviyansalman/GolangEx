package main

import (
	protos "GreeterService/Gen/GoGreeterService"
	"GreeterService/server"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"net"
	"os"
)

func main() {

	grpcServer := grpc.NewServer()
	greeterServer := server.NewGreeter()
	protos.RegisterGreeterServiceServer(grpcServer, greeterServer)
	reflection.Register(grpcServer)
	lis, err := net.Listen("tcp", ":9091")
	if err != nil {
		fmt.Println("Error", err)
		os.Exit(1)
	}
	grpcServer.Serve(lis)
}
