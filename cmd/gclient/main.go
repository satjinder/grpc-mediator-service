package main

import (
	"context"
	"flag"
	"log"
	"time"

	fpb "go.buf.build/grpc/go/satjinder/schemas/fileservice/v1"
	hpb "go.buf.build/grpc/go/satjinder/schemas/usstats/v1"
	hpb2 "go.buf.build/grpc/go/satjinder/schemas/usstats/v2"
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
	WithHttp2(conn)
	WithFile(conn)
}

func WithHttp(conn *grpc.ClientConn) {
	c := hpb.NewStatsAPIClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	r, err := c.GetStats(ctx, &hpb.GetStatsRequest{Drilldowns: "Nation", Measures: "Population"})
	if err != nil {
		log.Fatalf("failed: %v", err)
	}

	log.Println(r)
}

func WithHttp2(conn *grpc.ClientConn) {
	c := hpb2.NewStatsAPIClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	r, err := c.GetStatsData(ctx, &hpb2.GetStatsRequest{Drilldowns: "Nation", Measures: "Population"})
	if err != nil {
		log.Fatalf("failed: %v", err)
	}

	log.Println(r)
}

func WithFile(conn *grpc.ClientConn) {
	c := fpb.NewFileAPIClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	r, err := c.GetJson(ctx, &fpb.GetJsonRequest{Filename: "response.json"})
	if err != nil {
		log.Fatalf("failed: %v", err)
	}

	log.Println("worked....")
	log.Println(r)
}
