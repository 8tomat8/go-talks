package main

import (
	"log"

	"context"

	"fmt"

	"github.com/8tomat8/go-talks/code-from-class/gRPC/api_proto"
	"google.golang.org/grpc"
)

func main() {
	ctx := context.Background()
	cliConn, err := grpc.Dial("127.0.0.1:5555", grpc.WithInsecure())
	if err != nil {
		log.Fatal(err)
	}
	defer func() {
		err := cliConn.Close()
		if err != nil {
			log.Println(err)
		}
	}()
	cli := api.NewServClient(cliConn)

	fmt.Println("######################")
	mulResp, err := cli.Mul(ctx, &api.MulRequest{4, 8})
	if err != nil {
		log.Println(err)
	}
	fmt.Printf("#%v\n", mulResp)

	fmt.Println("######################")
	nameResp, err := cli.GetName(ctx, &api.Empty{})
	if err != nil {
		log.Println(err)
	}
	fmt.Printf("#%v\n", nameResp)

	fmt.Println("######################")
	sayHelloResp, err := cli.SayHello(ctx, &api.User{"UserName", 98})
	if err != nil {
		log.Println(err)
	}
	fmt.Printf("#%v\n", sayHelloResp)
}
