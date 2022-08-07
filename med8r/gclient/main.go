package main

import (
	"context"
	"flag"
	"log"
	"time"

	gpb "github.com/satjinder/med8r/schemas/gprotos"
	pb "github.com/satjinder/med8r/schemas/usstats"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/protobuf/types/known/anypb"
)

const (
	defaultDrilldown = "Nation"
)

var (
	addr = flag.String("addr", "localhost:50051", "the address to connect to")
	name = flag.String("Drillsdown", defaultDrilldown, "Stats at the level of Nation or State")
)

func main() {
	flag.Parse()
	// Set up a connection to the server.
	conn, err := grpc.Dial(*addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := gpb.NewGenericServiceClient(conn)

	// Contact the server and print out its response.
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	req, _ := anypb.New(&pb.GetStatsRequest{Drilldowns: *name})
	r, err := c.Call(ctx, &gpb.Request{Endpoint: "GetStats", Schema: "usstats/usstats.proto", Request: req})
	if err != nil {
		log.Fatalf("failed: %v", err)
	}

	var resp pb.GetStatsResponse
	r.Response.UnmarshalTo(&resp)
	log.Printf("Result: %s", resp.GetData())
}
