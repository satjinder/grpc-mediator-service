package genericserver

import (
	"context"
	"fmt"
	"log"
	"net"
	"testing"

	"github.com/satjinder/grpc-mediator-service/med8r/defaulthandlers"
	bsr "github.com/satjinder/grpc-mediator-service/med8r/schemaregistry/bsr"
	"github.com/satjinder/grpc-mediator-service/med8r/types"

	pb1 "go.buf.build/grpc/go/satjinder/med8r/samples/fakes/v1"
	pb2 "go.buf.build/grpc/go/satjinder/med8r/samples/fakes/v2"

	"google.golang.org/grpc"
	"google.golang.org/grpc/test/bufconn"
	//"time"
)

const bufSize = 1024 * 1024

var lis *bufconn.Listener

func init() {

	lis = bufconn.Listen(bufSize)

	//descriptorSetDir := "../gen/descriptor-sets"
	gs, err := NewServer(types.ServerConfig{

		Services: []types.ServiceConfig{
			{RegistryName: "usstats.v1.StatsAPI.fds", ProtoPath: "samples/fakes/v1/service.proto"},
			{RegistryName: "usstats.v2.StatsAPI.fds", ProtoPath: "samples/fakes/v2/service.proto"},
		},
	}, &defaulthandlers.DefaultProvider{}, bsr.New("satjinder", "med8r", "main"))

	if err != nil {
		fmt.Println(err)
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
	client := pb2.NewFileAPIClient(conn)
	req := &pb2.GetJsonRequest{Filename: "response.json"}
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
	client := pb1.NewStatsAPIClient(conn)
	req := &pb1.GetStatsRequest{Drilldowns: "Nation", Measures: "Population"}
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
	client := pb2.NewStatsAPIClient(conn)
	req := &pb2.GetStatsRequest{Drilldowns: "Nation", Measures: "Population"}
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
	client := pb2.NewFileAPIClient(conn)
	req := &pb2.GetJsonRequest{Filename: "response.json"}
	resp, err := client.UnsupportedHandler(ctx, req)
	if err == nil || resp != nil {
		t.Fatalf("no error detected: %v", err)
	}

}
