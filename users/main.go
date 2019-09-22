//go:generate protoc --go_out=plugins=grpc:./pb  -I ../proto ../proto/users.proto

package main

import (
	"log"
	"net"

	"github.com/petuhovskiy/grpc-hydra-bench/users/impl"
	"github.com/petuhovskiy/grpc-hydra-bench/users/pb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func main() {
	users := &impl.UsersServer{}

	server := grpc.NewServer()
	pb.RegisterUsersServer(server, users)
	reflection.Register(server)

	listener, err := net.Listen("tcp", ":3000")
	if err != nil {
		log.Fatal(err)
	}

	log.Println("gRPC server started")
	defer log.Println("gRPC server exited")
	if err := server.Serve(listener); err != nil {
		log.Fatal(err)
	}
}
