package main

import (
	"fmt"
	"github.com/grpc-ecosystem/go-grpc-middleware"
	"github.com/soichisumi/grpc-api-setting/app"
	"github.com/soichisumi/grpc-api-setting/app/proto"
	"google.golang.org/grpc"
	"log"
	"net"
)

const (
	grpcPort = ":3000"
)

func main() {
	lis, err := net.Listen("tcp", grpcPort)
	if err != nil {
		log.Fatalf("failed to listen: %+v\n", err)
	}

	grpcServer := grpc.NewServer(grpc.UnaryInterceptor(
			grpc_middleware.ChainUnaryServer(
				app.AuthInterceptor(app.APIPolicy),
			)))

	server := app.NewServer()
	proto.RegisterServerServer(grpcServer, server)

	methods, err := app.GetMethodDescriptor(grpcServer)
	if err != nil {
		log.Fatalf("failed to retrieve method descriptor. err: %+v", err)
	}

	if err := app.VerifyPolicy(methods, app.APIPolicy); err != nil {
		log.Fatalf("failed to retrieve method descriptor: %+v", err)
	}

	fmt.Printf("api is running on port: %s\n", grpcPort)
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalln(err)
	}
}