// generic server accepts configurations as input to make it testable. It reads the config and loads all the services into the grpC server.

// reference
// https://stackoverflow.com/questions/71320369/grpc-service-with-generic-proto-request-data-in-golang
// https://stackoverflow.com/questions/65561125/grpc-go-single-generic-service-handler

package main

import (

	//"encoding/json"
	"flag"
	"fmt"
	"log"
	"net"

	"github.com/satjinder/grpc-mediator-service/defaulthandlers"
	"github.com/satjinder/grpc-mediator-service/genericserver"
	"github.com/satjinder/grpc-mediator-service/types"
	//"github.com/satjinder/grpc-mediator-service/handlers/httpservicehandler"
)

var (
	port         = flag.Int("port", 50051, "The server port")
	ServerConfig = types.ServerConfig{
		DescriptorSetDir: flag.String("descriptor-sets", "gen/descriptor-sets", "directory containing all descriptor sets to load"),
		Services: []types.ServiceConfig{
			{RegistryName: "usstats.v1.StatsAPI.fds", ProtoPath: "usstats/v1/usstats.proto"},
			{RegistryName: "fileservice.v1.FileAPI.fds", ProtoPath: "fileservice/v1/fileservice.proto"},
			{RegistryName: "usstats.v2.StatsAPI.fds", ProtoPath: "usstats/v2/usstats.proto"},
		},
	}
)

func main() {
	flag.Parse()
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", *port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	gs, err := genericserver.NewServer(ServerConfig, &defaulthandlers.DefaultProvider{})
	if err != nil {
		panic(err)
	}

	log.Printf("server listening at %v", lis.Addr())
	if err := gs.GrpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
