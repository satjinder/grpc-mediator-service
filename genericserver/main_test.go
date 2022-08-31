package genericserver

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net"
	"testing"

	fpb "github.com/satjinder/grpc-mediator-service/gen/fileservice"
	hpb "github.com/satjinder/grpc-mediator-service/gen/usstats"
	"github.com/satjinder/grpc-mediator-service/types"

	//"google.golang.org/protobuf/types/dynamicpb"
	//"google.golang.org/protobuf/encoding/protojson"

	"google.golang.org/grpc"
	"google.golang.org/grpc/test/bufconn"
	//"time"
)

const bufSize = 1024 * 1024

var lis *bufconn.Listener

func init() {

	lis = bufconn.Listen(bufSize)

	gs, err := NewServer(types.ServerConfig{
		DescriptorSetDir: flag.String("descriptor-sets", "/Users/baths/src/mediator-service/gen/descriptor-sets", "directory containing all descriptor sets to load"),
		Services: []types.ServiceConfig{
			{RegistryName: "usstats.StatsAPI.fds", ProtoPath: "usstats/usstats.proto"},
			{RegistryName: "fileservice.FileAPI.fds", ProtoPath: "fileservice/fileservice.proto"},
		},
	})

	if err != nil {
		panic(err)
	}

	log.Printf("server listening at %v", lis.Addr())

	go func() {
		if err := gs.GrpcServer.Serve(lis); err != nil {
			log.Fatalf("failed to serve: %v", err)
		}
	}()
}

func bufDialer(context.Context, string) (net.Conn, error) {
	return lis.Dial()
}

func Test_ServerReturnsCorrectResponseForFile(t *testing.T) {
	ctx := context.Background()
	conn, err := grpc.DialContext(ctx, "bufnet", grpc.WithContextDialer(bufDialer), grpc.WithInsecure())
	if err != nil {
		t.Fatalf("Failed to dial bufnet: %v", err)
	}
	defer conn.Close()
	client := fpb.NewFileAPIClient(conn)
	req := &fpb.GetJsonRequest{Filename: "response.json"}
	resp, err := client.GetJson(ctx, req)
	if err != nil {
		t.Fatalf("SayHello failed: %v", err)
	}
	log.Printf("Response: %+v", resp)
	fmt.Println(&resp)
	if resp.Data[len(resp.Data)-1].Year != "2021" {
		t.Fatalf("Incorrect year ")
	}
	// Test for output here.
}

func Test_ServerReturnsCorrectResponseForHttp(t *testing.T) {
	ctx := context.Background()
	conn, err := grpc.DialContext(ctx, "bufnet", grpc.WithContextDialer(bufDialer), grpc.WithInsecure())
	if err != nil {
		t.Fatalf("Failed to dial bufnet: %v", err)
	}
	defer conn.Close()
	client := hpb.NewStatsAPIClient(conn)
	req := &hpb.GetStatsRequest{Drilldowns: "Nation", Measures: "Population"}
	resp, err := client.GetStats(ctx, req)
	if err != nil {
		t.Fatalf("SayHello failed: %v", err)
	}
	log.Printf("Response: %+v", resp)
	fmt.Println(&resp)
	if resp.Data[len(resp.Data)-1].Nation != "United States" {
		t.Fatalf("Incorrect year ")
	}
}
