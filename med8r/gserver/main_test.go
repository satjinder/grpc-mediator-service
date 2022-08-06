package main

import (
	gpb "github.com/satjinder/med8r/gprotos"
	pb "github.com/satjinder/med8r/statsservice"

	//"google.golang.org/protobuf/types/dynamicpb"
	"google.golang.org/protobuf/types/known/anypb"
	"testing"
	//"time"
	"fmt"
	"os"
)

func Test_NewAggregate_StartsWithNoLimitRecords(t *testing.T) {
	fmt.Println(os.Getwd())
	s := &server{}
	req, _ := anypb.New(&pb.GetStatsRequest{Drilldowns: "Nation"})
	m, err := s.ConfigureEndpoint(&gpb.Request{Endpoint: "GetStats", Request: req})
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("gggg")
	fmt.Println(m)
}
