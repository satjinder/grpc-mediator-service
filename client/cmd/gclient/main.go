package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"time"

	pb1 "go.buf.build/grpc/go/satjinder/med8rtestservices/usstats/v1"
	pb2 "go.buf.build/grpc/go/satjinder/med8rtestservices/usstats/v2"

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
	fmt.Println("------------------------")
	WithHttp2(conn)
	fmt.Println("------------------------")
	WithFile(conn)
}

func WithHttp(conn *grpc.ClientConn) {
	c := pb1.NewStatsAPIClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	r, err := c.GetStats(ctx, &pb1.GetStatsRequest{Drilldowns: "Nation", Measures: "Population"})
	if err != nil {
		log.Fatalf("failed: %v", err)
	}

	log.Println(r)
}

func WithHttp2(conn *grpc.ClientConn) {
	c := pb2.NewStatsAPIClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	r, err := c.GetStatsData(ctx, &pb2.GetStatsRequest{Drilldowns: "Nation", Measures: "Population"})
	if err != nil {
		log.Fatalf("failed: %v", err)
	}

	log.Println(r)
}

func WithFile(conn *grpc.ClientConn) {
	c := pb2.NewStatsAPIClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	r, err := c.GetJson(ctx, &pb2.GetJsonRequest{Filename: "response.json"})
	if err != nil {
		log.Fatalf("failed: %v", err)
	}

	log.Println("worked....")
	log.Println(r)
}
