// https://stackoverflow.com/questions/71320369/grpc-service-with-generic-proto-request-data-in-golang

package main

import (
	"context"
	"errors"
	"strings"

	//"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"net"

	"github.com/satjinder/med8r/handlers/entitlementshandler"
	"github.com/satjinder/med8r/handlers/fileservicehandler"
	"github.com/satjinder/med8r/handlers/httpservicehandler"
	gpb "github.com/satjinder/med8r/schemas/gprotos"
	"github.com/satjinder/med8r/types"
	"google.golang.org/grpc"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/reflect/protodesc"
	"google.golang.org/protobuf/reflect/protoreflect"
	"google.golang.org/protobuf/reflect/protoregistry"
	"google.golang.org/protobuf/types/descriptorpb"
	"google.golang.org/protobuf/types/dynamicpb"
	"google.golang.org/protobuf/types/known/anypb"
	//"github.com/satjinder/med8r/handlers/httpservicehandler"
)

var (
	port = flag.Int("port", 50051, "The server port")
)

type server struct {
	gpb.UnimplementedGenericServiceServer
	schemaRegister map[string]*protoregistry.Files
}

func (s *server) Call(ctx context.Context, greq *gpb.Request) (*gpb.Response, error) {
	epContext := &types.EndpointContext{}
	serviceDesc, err := GetEndpointDescriptor(s, greq)
	if err != nil {
		return nil, err
	}

	epContext.EndpointDescriptor = serviceDesc
	reqMsg, err := ParseRequest(*epContext, greq)
	if err != nil {
		return nil, err
	}
	epContext.Request = &types.GRequest{Message: reqMsg}
	epContext.Response = &types.GResponse{}

	err = ConfigureHandlers(epContext)
	if err != nil {
		return nil, err
	}

	for _, handler := range epContext.Handlers {
		epCtx := context.WithValue(ctx, types.ENDPOINT_CONTEXT_KEY, epContext)
		err = handler.Process(epCtx)
		if err != nil {
			return nil, err
		}
	}

	resp, _ := anypb.New(epContext.Response.Message)
	return &gpb.Response{Response: resp}, nil
}

func ConfigureHandlers(epContext *types.EndpointContext) error {
	epContext.EndpointConfig = ParseExtensions(*epContext)
	for _, handler := range epContext.EndpointConfig.Handlers {
		switch handler.Name {
		case "http-backend":
			epContext.Handlers = append(epContext.Handlers, httpservicehandler.NewHandler(handler))
		case "entitlements":
			epContext.Handlers = append(epContext.Handlers, entitlementshandler.NewHandler(handler))
		case "file-backend":
			epContext.Handlers = append(epContext.Handlers, fileservicehandler.NewHandler(handler))

		default:
			errMsg := fmt.Errorf("Handler not found %v", handler.Name)
			return errors.New(errMsg.Error())
		}
	}
	return nil
}

func GetEndpointDescriptor(s *server, greq *gpb.Request) (protoreflect.MethodDescriptor, error) {
	endpointName := greq.GetEndpoint()
	schema := greq.GetSchema()
	schemaParts := strings.Split(schema, "/")
	filename := schemaParts[len(schemaParts)-1]
	registry, _ := s.createProtoRegistry(filename)

	desc, err := registry.FindFileByPath(schema)
	if err != nil {
		panic(err)
	}

	fd := desc.Services()

	serviceDescriptor := fd.ByName(protoreflect.Name("StatsAPI")).Methods().ByName(protoreflect.Name(endpointName))
	return serviceDescriptor, err
}

func ParseRequest(epContext types.EndpointContext, greq *gpb.Request) (*dynamicpb.Message, error) {
	fmt.Println("parse request")
	input := epContext.EndpointDescriptor.Input()
	req := greq.GetRequest()
	bytes := req.GetValue()
	msg := dynamicpb.NewMessage(input)
	err := proto.Unmarshal(bytes, msg)
	if err != nil {
		return nil, err
	}
	return msg, nil
}

func ParseExtensions(epContext types.EndpointContext) *gpb.EndpointConfig {
	options := epContext.EndpointDescriptor.Options()
	ex := proto.GetExtension(options, gpb.E_EndpointConfig)
	config := ex.(*gpb.EndpointConfig)
	fmt.Println(config.AuthType)
	return config
}

func main() {
	flag.Parse()
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", *port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	gpb.RegisterGenericServiceServer(s, NewServer())
	log.Printf("server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}

func NewServer() *server {
	return &server{
		schemaRegister: make(map[string]*protoregistry.Files),
	}
}

func (s *server) createProtoRegistry(filename string) (*protoregistry.Files, error) {
	files := s.schemaRegister[filename]
	if files != nil {
		return files, nil
	}

	tmpFile := "../schemas/register/" + filename + "-registry.pb"

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
