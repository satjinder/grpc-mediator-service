// https://stackoverflow.com/questions/71320369/grpc-service-with-generic-proto-request-data-in-golang

package main

import (
	"context"

	//"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"net"
	"os"

	gpb "github.com/satjinder/med8r/gprotos"
	"google.golang.org/grpc"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/reflect/protodesc"
	"google.golang.org/protobuf/reflect/protoreflect"
	"google.golang.org/protobuf/reflect/protoregistry"
	"google.golang.org/protobuf/types/descriptorpb"
	"google.golang.org/protobuf/types/dynamicpb"
	"google.golang.org/protobuf/types/known/anypb"
)

var (
	port = flag.Int("port", 50051, "The server port")
)

var registry protoregistry.Files

// server is used to implement helloworld.GreeterServer.
type server struct {
	gpb.UnimplementedGenericServiceServer
}

// SayHello implements helloworld.GreeterServer
func (s *server) Call(ctx context.Context, in *gpb.Request) (*anypb.Any, error) {
	json := ReturnJson(in)
	log.Printf("Received: %v", json)

	return GetResponse(in)
	//resp := &pb.GetStatsResponse{Test: "abc"}
	//log.Printf("returning %v", resp.GetTest())
	//return anypb.New(resp)
}

func ReturnJson(in *gpb.Request) string {
	msg := GetMsg(in, in.GetRequestname())

	req := in.GetRequest()
	bytes := req.GetValue()

	err := proto.Unmarshal(bytes, msg)
	if err != nil {
		fmt.Println("here3")
		panic(err)
	}

	jsonBytes, err := protojson.Marshal(msg)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(jsonBytes))
	return string(jsonBytes)
}

func GetMsg(in *gpb.Request, msgName string) *dynamicpb.Message {
	srvname := in.GetServicename()

	file := fmt.Sprintf("%v.proto", srvname)
	registry, _ := createProtoRegistry(file)

	desc, err := registry.FindFileByPath(file)
	if err != nil {
		panic(err)
	}
	fd := desc.Messages()
	data := fd.ByName(protoreflect.Name(msgName))
	msg := dynamicpb.NewMessage(data)
	return msg
}

func GetResponse(in *gpb.Request) (*anypb.Any, error) {
	msg := GetMsg(in, "GetStatsResponse")

	jsonFile, err := os.Open("genericserver/response.json")
	if err != nil {
		fmt.Println(err)
	}
	defer jsonFile.Close()

	byteValue, err2 := ioutil.ReadAll(jsonFile)
	if err2 != nil {
		fmt.Println(err2)
	}
	jsonBytes := []byte(byteValue)

	pm := msg.Interface()
	err = protojson.Unmarshal(jsonBytes, pm)
	if err != nil {
		fmt.Println(err)
	}

	jsonBytes, err3 := protojson.Marshal(msg)
	if err3 != nil {
		fmt.Println(err3)
	}
	fmt.Println(string(jsonBytes))

	return anypb.New(msg)
}

func main() {
	flag.Parse()
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", *port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	gpb.RegisterGenericServiceServer(s, &server{})
	log.Printf("server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}

func main2() {
	in := &gpb.Request{Servicename: "stats", Requestname: "GetStatsRequest"}
	GetResponse(in)
}

func createProtoRegistry(filename string) (*protoregistry.Files, error) {
	// Create descriptors using the protoc binary.
	// Imported dependencies are included so that the descriptors are self-contained.
	tmpFile := "schemaregister/" + filename + "-tmp.pb"

	marshalledDescriptorSet, err := ioutil.ReadFile(tmpFile)
	if err != nil {
		return nil, err
	}
	descriptorSet := descriptorpb.FileDescriptorSet{}
	err = proto.Unmarshal(marshalledDescriptorSet, &descriptorSet)
	if err != nil {
		return nil, err
	}

	//fmt.Println(descriptorSet)
	files, err := protodesc.NewFiles(&descriptorSet)
	if err != nil {
		return nil, err
	}
	return files, nil
}
