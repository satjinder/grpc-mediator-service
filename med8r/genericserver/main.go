// https://stackoverflow.com/questions/71320369/grpc-service-with-generic-proto-request-data-in-golang

package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net"

	gpb "github.com/satjinder/med8r/gprotos"
	pb "github.com/satjinder/med8r/protos"
	"google.golang.org/grpc"
	"google.golang.org/protobuf/types/known/anypb"
)

var (
	port = flag.Int("port", 50051, "The server port")
)

// server is used to implement helloworld.GreeterServer.
type server struct {
	gpb.UnimplementedGenericServiceServer
}

// SayHello implements helloworld.GreeterServer
func (s *server) Call(ctx context.Context, in *anypb.Any) (*anypb.Any, error) {
	var req pb.GetStatsRequest
	in.UnmarshalTo(&req)
	log.Printf("Received: %v", req.GetDrilldowns())
	resp := &pb.GetStatsResponse{Test: "abc"}
	log.Printf("returning %v", resp.GetTest())
	return anypb.New(resp)
}

func main() {
	flag.Parse()
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", *port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	gpb.RegisterGenericServiceServer(s, &server{})
	log.Printf("server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
