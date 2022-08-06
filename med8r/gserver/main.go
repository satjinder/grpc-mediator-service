// https://stackoverflow.com/questions/71320369/grpc-service-with-generic-proto-request-data-in-golang

package main

import (
	"context"
	"os"

	//"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"net"

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

type middleware struct {
}

type server struct {
	gpb.UnimplementedGenericServiceServer
	schemaRegister map[string]*protoregistry.Files
	middlewareList []middleware
}

func (s *server) Call(ctx context.Context, in *gpb.Request) (*gpb.Response, error) {
	return s.ConfigureEndpoint(in)
}

func (s *server) ConfigureEndpoint(greq *gpb.Request) (*gpb.Response, error) {
	servicedata, err := GetServiceDescriptor(s, greq)

	// check extensions
	CheckExtensions(servicedata)

	// get request json
	jsonBytes := ParseRequest(servicedata, greq, err)

	// get response
	return ProcessRequest(jsonBytes, servicedata)
}

func GetServiceDescriptor(s *server, greq *gpb.Request) (protoreflect.MethodDescriptor, error) {
	endpointName := greq.GetEndpoint()

	filename := "stats.proto"
	registry, _ := s.createProtoRegistry(filename)

	desc, err := registry.FindFileByPath(filename)
	if err != nil {
		panic(err)
	}

	fd := desc.Services()

	serviceDescriptor := fd.ByName(protoreflect.Name("StatsAPI")).Methods().ByName(protoreflect.Name(endpointName))
	return serviceDescriptor, err
}

func ProcessRequest(requestJson []byte, servicedata protoreflect.MethodDescriptor) (*gpb.Response, error) {
	fmt.Println("get response")

	jsonBytes, _ := CallExternalAPI("response.json")
	output := servicedata.Output()
	respmsg := dynamicpb.NewMessage(output)
	pm := respmsg.Interface()
	err := protojson.Unmarshal(jsonBytes, pm)
	if err != nil {
		fmt.Println(err)
	}

	resp, _ := anypb.New(respmsg)
	return &gpb.Response{Response: resp}, nil
}

func ParseRequest(servicedata protoreflect.MethodDescriptor, greq *gpb.Request, err error) []byte {
	fmt.Println("parse request")
	input := servicedata.Input()
	req := greq.GetRequest()
	bytes := req.GetValue()
	msg := dynamicpb.NewMessage(input)
	err = proto.Unmarshal(bytes, msg)
	if err != nil {
		panic(err)
	}

	jsonBytes, err := protojson.Marshal(msg)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(jsonBytes))
	return jsonBytes
}

func CheckExtensions(servicedata protoreflect.MethodDescriptor) {
	options := servicedata.Options()
	ex := proto.GetExtension(options, gpb.E_Med8RConfig)
	config := ex.(*gpb.Med8RConfig)
	fmt.Println(config.AuthType)

	if config.EntitlementOperations != nil {
		fmt.Println("check entitlements")
		fmt.Println(config.EntitlementOperations)
	}
}

func CallExternalAPI(filename string) ([]byte, error) {
	jsonFile, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer jsonFile.Close()

	byteValue, err2 := ioutil.ReadAll(jsonFile)
	if err2 != nil {
		return nil, err2
	}
	jsonBytes := []byte(byteValue)

	return jsonBytes, nil
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

func (s *server) createProtoRegistry(filename string) (*protoregistry.Files, error) {
	if s.schemaRegister == nil {
		s.schemaRegister = make(map[string]*protoregistry.Files)
	}
	files := s.schemaRegister[filename]
	if files != nil {
		return files, nil
	}

	tmpFile := "../schemaregister/" + filename + "-tmp.pb"

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
	files, err = protodesc.NewFiles(&descriptorSet)
	if err != nil {
		return nil, err
	}

	s.schemaRegister[filename] = files
	return files, nil
}
