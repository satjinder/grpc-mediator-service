package main

import (
	"context"
	"flag"
	"log"
	"time"

	spb "github.com/satjinder/med8r/schemas/statsservice"
	upb "github.com/satjinder/med8r/schemas/usstats"
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
	// Contact the server and print out its response.
	WithHttp(conn)
	WithFile(conn)
}

func WithHttp(conn *grpc.ClientConn) {
	c := upb.NewStatsAPIClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	r, err := c.GetStats(ctx, &upb.GetStatsRequest{Drilldowns: *name, Measures: "Population"})
	if err != nil {
		log.Fatalf("failed: %v", err)
	}

	log.Println(r)
}

func WithFile(conn *grpc.ClientConn) {
	c := spb.NewStatsAPIClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	r, err := c.GetStats(ctx, &spb.GetStatsRequest{Drilldowns: *name})
	if err != nil {
		log.Fatalf("failed: %v", err)
	}

	log.Println(r)
}
