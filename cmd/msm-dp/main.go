package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net"

	pb "github.com/media-streaming-mesh/msm-dp/api/v1alpha1/msm_dp"
	"google.golang.org/grpc"
)

var (
	port = flag.Int("port", 9000, "The server port")
)

// server is used to implement msm_dp.server.
type server struct {
	pb.UnimplementedMsmDataPlaneServer
}

func (s *server) StreamAddDel(ctx context.Context, in *pb.StreamData) (*pb.StreamResult, error) {
	log.Printf("Received: message from client Endpoint = %v", in.Endpoint)
	log.Printf("Received: message from client Enable = %v", in.Enable)
	log.Printf("Received: message from client Protocol = %v", in.Protocol)
	log.Printf("Received: message from client Id = %v", in.Id)
	log.Printf("Received: message from client Operation = %v", in.Operation)
	log.Printf("Received: message from client Context = %v", ctx)

	return &pb.StreamResult{
		Success: true,
	}, nil
}

func main() {
	flag.Parse()
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", *port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterMsmDataPlaneServer(s, &server{})
	log.Printf("server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
