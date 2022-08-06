package main

import (
	gpb "github.com/satjinder/med8r/schemas/gprotos"
	pb "github.com/satjinder/med8r/schemas/statsservice"

	//"google.golang.org/protobuf/types/dynamicpb"
	"google.golang.org/protobuf/types/known/anypb"
	"testing"
	//"time"
	"fmt"
	"os"
)

func Test_ServerReturnsCorrectResponse(t *testing.T) {
	fmt.Println(os.Getwd())
	s := &server{}
	req, _ := anypb.New(&pb.GetStatsRequest{Drilldowns: "Nation"})
	m, err := s.ConfigureEndpoint(&gpb.Request{Endpoint: "GetStats", Schema: "statsservice/stats.proto", Request: req})
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(m)
}
