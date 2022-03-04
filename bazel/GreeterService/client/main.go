package main

import (
	greeter "GreeterService/Gen/GoGreeterService"
	"context"
	"fmt"
	"google.golang.org/grpc"
)

func main() {
	serverAddress := "localhost:9091"
	conn, err := grpc.Dial(serverAddress, grpc.WithInsecure())
	if err != nil {
		fmt.Println("Unable to connect")
	}
	defer conn.Close()
	client := greeter.NewGreeterServiceClient(conn)
	doSomeWork(&client)
}

func doSomeWork(client *greeter.GreeterServiceClient) {
	request := &greeter.HelloRequest{Name: "Hello My Name is Salman"}
	resp, err := (*client).SayHello(context.Background(), request)
	if err != nil {
		fmt.Println("Something went wrong")
	} else {
		fmt.Println("Respnse is : ", resp)
	}
}
