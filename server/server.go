package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"grpc-example/auth"
	"log"
	"net"
)

const port = 5001

type authServer struct{}

func (s *authServer) Login(ctx context.Context, request *rpc_auth.LoginRequest) (*rpc_auth.Token, error) {
	return &rpc_auth.Token{Token: "this is a demo token"}, nil
}

func (s *authServer) GetUserInfo(ctx context.Context, request *rpc_auth.Token) (*rpc_auth.User, error) {
	return &rpc_auth.User{Name: "test:" + request.Token, Age: 11, Friends: []string{"a", "b", "c"}}, nil
}

func main() {
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", port))
	if err != nil {
		panic(err)
	}
	s := grpc.NewServer()
	rpc_auth.RegisterAuthServer(s, &authServer{})
	//reflection.Register(s)  这个是为了外网cli调用的时候,能够提供反射服务

	//like $ ./grpc_cli call localhost:50051 SayHello "name: 'gRPC CLI'"
	log.Printf("auth server run in :%d", port)
	if err := s.Serve(lis); err != nil {
		panic(err)
	}
}
