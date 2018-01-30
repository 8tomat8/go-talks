package main

import (
	"context"
	"log"
	"net"
	"strconv"
	"time"

	"fmt"

	"github.com/8tomat8/go-talks/code-from-class/gRPC/api_proto"
	"google.golang.org/grpc"
)

func main() {
	srv := service{"ServiceName", "BIAUygv3if76abvkuVAEU^Tvb64uvckuvyb"}
	gServer := grpc.NewServer()
	api.RegisterServServer(gServer, &srv)

	listener, err := net.Listen("tcp", ":5555")
	if err != nil {
		log.Fatal(err)
	}
	if err := gServer.Serve(listener); err != nil {
		log.Fatal(err)
	}
}

type service struct {
	Name     string
	CheckSum string
}

func (s *service) GetName(ctx context.Context, in *api.Empty) (*api.GetNameResponse, error) {
	fmt.Println("GetName was triggred")
	return &api.GetNameResponse{s.Name}, nil
}

func (s *service) SayHello(ctx context.Context, user *api.User) (*api.SayHelloResponse, error) {
	fmt.Println("SayHello was triggred")
	return &api.SayHelloResponse{
		"Hello " + user.Name + ". You are " + strconv.Itoa(int(user.Age)),
	}, nil
}

func (s *service) Mul(ctx context.Context, in *api.MulRequest) (*api.MulResponse, error) {
	fmt.Println("Mul was triggred")
	time.Sleep(time.Second)
	return &api.MulResponse{in.A * in.B}, nil
}
