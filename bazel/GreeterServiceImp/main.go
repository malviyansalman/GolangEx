package main

import (
	greeter "../../GreeterService/Gen/GoGreeterService"
	ImplementedServer "GreeterServiceImp/server"

	"flag"
	"fmt"
	"google.golang.org/grpc"
	"net"
)

func main() {
	port := flag.Int("Port", 8181, "Running Server")
	grpcServer, lis := runGrpcServer(*port)
	err := grpcServer.Serve(lis)
	if err != nil {
		fmt.Println("Something went wrong")
	}
}

func runGrpcServer(port int) (*grpc.Server, net.Listener) {
	lis, err := net.Listen("tcp", fmt.Sprintf("0.0.0.0:%d", port))
	if err != nil {
		fmt.Println("Error")
	}
	//
	grpcServer := grpc.NewServer()
	greeter.RegisterGreeterServiceServer(grpcServer, &ImplementedServer.GreeterServer{})
	return grpcServer, lis
}
