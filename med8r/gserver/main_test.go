package main

import (
	"context"
	gpb "github.com/satjinder/med8r/schemas/gprotos"
	pb "github.com/satjinder/med8r/schemas/statsservice"
	hpb "github.com/satjinder/med8r/schemas/usstats"

	//"google.golang.org/protobuf/types/dynamicpb"
	//"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/types/known/anypb"
	"testing"
	//"time"
	"fmt"
	"os"
)

func Test_ServerReturnsCorrectResponseForFile(t *testing.T) {
	fmt.Println(os.Getwd())
	s := NewServer()
	req, _ := &pb.GetStatsRequest{Drilldowns: "response.json"}
	ctx := context.Background()
	m, err := s.Handler(ctx, &gpb.Request{Endpoint: "GetStats", Schema: "statsservice/stats.proto", Request: req})
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("----Response----")
	var resp pb.GetStatsResponse
	m.Response.UnmarshalTo(&resp)

	fmt.Println(&resp)
}

func Test_ServerReturnsCorrectResponseForHttp(t *testing.T) {
	fmt.Println(os.Getwd())
	s := NewServer()

	req, _ := anypb.New(&hpb.GetStatsRequest{Drilldowns: "Nation"})
	ctx := context.Background()
	m, err := s.Call(ctx, &gpb.Request{Endpoint: "GetStats", Schema: "usstats/usstats.proto", Request: req})
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("----Response----")

	var resp hpb.GetStatsResponse
	m.Response.UnmarshalTo(&resp)

	fmt.Println(&resp)
}
