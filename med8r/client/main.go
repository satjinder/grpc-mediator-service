package main

import (
	"context"
	"flag"
	"log"
	"time"

	pb "github.com/satjinder/med8r/protos"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
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
	c := pb.NewStatsAPIClient(conn)

	// Contact the server and print out its response.
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	r, err := c.GetStats(ctx, &pb.GetStatsRequest{Drilldowns: *name})
	if err != nil {
		log.Fatalf("failed: %v", err)
	}
	log.Printf("Result: %s", r.GetTest())
}
