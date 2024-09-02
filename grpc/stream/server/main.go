package server

import (
	"log"

	"github.com/kl7sn/demo/grpc/pb"
)

type Hello struct {
	pb.UnimplementedServer
}

// UnaryEcho 一个普通的UnaryAPI
func (e *Echo) UnaryEcho(ctx context.Context, req *pb.EchoRequest) (*pb.EchoResponse, error) {
	log.Printf("Recved: %v", req.GetMessage())
	resp := &pb.EchoResponse{Message: req.GetMessage()}
	return resp, nil
}
