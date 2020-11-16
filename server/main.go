package main

import (
	"fmt"
	"net"
	"os"

	"github.com/hashicorp/go-hclog"
	"github.com/vaibhav/try-assignment2/server/handlers"
	"github.com/vaibhav/try-assignment2/server/protos/user"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func main() {
	// creating a logger instance

	opts := []grpc.ServerOption{}
	grpcServer := grpc.NewServer(opts...)
	reflection.Register(grpcServer)

	// getting service server
	log := hclog.Default()
	userServer := handlers.NewUserServer(log)

	user.RegisterUserServiceServer(grpcServer, userServer)

	listener, err := net.Listen("tcp", ":4000")
	if err != nil {
		log.Error("Unable to listen. ERROR: ", err)
		os.Exit(1)
	}

	//  starting go routine for server server
	go func() {
		err := grpcServer.Serve(listener)
		if err != nil {
			log.Error("Unable to serve. ERROR: ", err)
			os.Exit(1)
		}
	}()
	fmt.Println("server is running...")
}
