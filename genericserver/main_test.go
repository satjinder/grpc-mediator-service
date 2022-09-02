package genericserver

import (
	"context"
	"fmt"
	"log"
	"net"
	"testing"

	"github.com/satjinder/grpc-mediator-service/defaulthandlers"
	lsm "github.com/satjinder/grpc-mediator-service/schemaregistry/local"
	"github.com/satjinder/grpc-mediator-service/types"
	fpb "go.buf.build/grpc/go/satjinder/schemas/fileservice/v1"
	hpb "go.buf.build/grpc/go/satjinder/schemas/usstats/v1"
	hpb2 "go.buf.build/grpc/go/satjinder/schemas/usstats/v2"

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

	descriptorSetDir := "../gen/descriptor-sets"
	gs, err := NewServer(types.ServerConfig{

		Services: []types.ServiceConfig{
			{RegistryName: "usstats.v1.StatsAPI.fds", ProtoPath: "usstats/v1/usstats.proto"},
			{RegistryName: "usstats.v2.StatsAPI.fds", ProtoPath: "usstats/v2/usstats.proto"},
			{RegistryName: "fileservice.v1.FileAPI.fds", ProtoPath: "fileservice/v1/fileservice.proto"},
		},
	}, &defaulthandlers.DefaultProvider{}, lsm.New(descriptorSetDir))

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
		t.Fatalf(" failed: %v", err)
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
		t.Fatalf(" failed: %v", err)
	}
	log.Printf("Response: %+v", resp)
	fmt.Println(&resp)
	if resp.Data[len(resp.Data)-1].Nation != "United States" {
		t.Fatalf("Incorrect year ")
	}
}

func Test_ServerReturnsCorrectResponseForHttp2(t *testing.T) {
	ctx := context.Background()
	conn, err := grpc.DialContext(ctx, "bufnet", grpc.WithContextDialer(bufDialer), grpc.WithInsecure())
	if err != nil {
		t.Fatalf("Failed to dial bufnet: %v", err)
	}
	defer conn.Close()
	client := hpb2.NewStatsAPIClient(conn)
	req := &hpb2.GetStatsRequest{Drilldowns: "Nation", Measures: "Population"}
	resp, err := client.GetStatsData(ctx, req)
	if err != nil {
		t.Fatalf(" failed: %v", err)
	}
	log.Printf("Response: %+v", resp)
	fmt.Println(&resp)
	if resp.Data[len(resp.Data)-1].Nation != "United States" {
		t.Fatalf("Incorrect year ")
	}
}

func Test_DoesntLoadServiceWhenHandlerIsAvailable(t *testing.T) {
	ctx := context.Background()
	conn, err := grpc.DialContext(ctx, "bufnet", grpc.WithContextDialer(bufDialer), grpc.WithInsecure())
	if err != nil {
		t.Fatalf("Failed to dial bufnet: %v", err)
	}
	defer conn.Close()
	client := fpb.NewFileAPIClient(conn)
	req := &fpb.GetJsonRequest{Filename: "response.json"}
	resp, err := client.UnsupportedHandler(ctx, req)
	if err == nil || resp != nil {
		t.Fatalf("no error detected: %v", err)
	}

}
