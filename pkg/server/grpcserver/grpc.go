package grpcserver

import (
	"net"

	"github.com/labstack/gommon/log"
	"gitlab.com/indie-developers/go-api-echo-template/pb"
	"gitlab.com/indie-developers/go-api-echo-template/pkg/server"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

type grpcServer struct {
	client      *grpc.Server
	userService pb.UserServiceServer
	pb.UnimplementedUserServiceServer
}

func NewServer(server *grpc.Server, userService pb.UserServiceServer) server.RpcServer {
	return &grpcServer{
		client:      server,
		userService: userService,
	}
}

func (s *grpcServer) Run(address string) {
	pb.RegisterUserServiceServer(s.client, s.userService)
	reflection.Register(s.client)

	listener, err := net.Listen("tcp", address)
	if err != nil {
		log.Fatalf("cannot create listener: %s", err.Error())
	}

	log.Printf("start gRPC server at %s", listener.Addr().String())
	err = s.client.Serve(listener)
	if err != nil {
		log.Fatalf("cannot start gRPC server: %s", err.Error())
	}
}
