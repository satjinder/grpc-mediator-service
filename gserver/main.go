// reference
// https://stackoverflow.com/questions/71320369/grpc-service-with-generic-proto-request-data-in-golang
// https://stackoverflow.com/questions/65561125/grpc-go-single-generic-service-handler

package main

import (
	"context"
	"errors"
	"path/filepath"
	"strings"

	//"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"net"

	gpb "github.com/satjinder/grpc-mediator-service/gen/gprotos"
	"github.com/satjinder/grpc-mediator-service/handlers/entitlementshandler"
	"github.com/satjinder/grpc-mediator-service/handlers/fileservicehandler"
	"github.com/satjinder/grpc-mediator-service/handlers/httpservicehandler"
	"github.com/satjinder/grpc-mediator-service/types"
	"google.golang.org/grpc"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/reflect/protodesc"
	"google.golang.org/protobuf/reflect/protoreflect"
	"google.golang.org/protobuf/reflect/protoregistry"
	"google.golang.org/protobuf/types/descriptorpb"
	"google.golang.org/protobuf/types/dynamicpb"
	//"github.com/satjinder/grpc-mediator-service/handlers/httpservicehandler"
)

var (
	port              = flag.Int("port", 50051, "The server port")
	descriptorSetsDir = flag.String("descriptor-sets", "gen/descriptor-sets", "directory containing all descriptor sets to load")
	servicesConfig    = []ServiceConfig{
		{registryName: "usstats.fds", protoPath: "usstats/usstats.proto"},
		{registryName: "stats.fds", protoPath: "statsservice/stats.proto"},
	}
)

type ServiceConfig struct {
	registryName string
	protoPath    string
}

type GRPCService struct {
	grpcServer           *grpc.Server
	serviceDescriptorMap map[string]protoreflect.ServiceDescriptor
}

func NewServer() *GRPCService {
	return &GRPCService{
		grpcServer:           grpc.NewServer(),
		serviceDescriptorMap: make(map[string]protoreflect.ServiceDescriptor),
	}
}

func main() {
	flag.Parse()
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", *port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	gs := NewServer()
	if err := gs.LoadServices(); err != nil {
		panic(err)
	}

	log.Printf("server listening at %v", lis.Addr())
	if err := gs.grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}

func (s *GRPCService) LoadServices() error {
	for _, f := range servicesConfig {
		fdsFile := filepath.Join(*descriptorSetsDir, f.registryName)
		registry, err := s.createProtoRegistry(fdsFile)
		if err != nil {
			return err
		}
		if err := s.LoadService(f.protoPath, registry); err != nil {
			return err
		}
	}
	return nil
}

func (s *GRPCService) LoadService(serviceProtoName string, registry *protoregistry.Files) error {
	fd, err := registry.FindFileByPath(serviceProtoName)
	if err != nil {
		return err
	}
	services := fd.Services()

	for i := 0; i < services.Len(); i++ {
		rsd := services.Get(i)
		srvName := string(rsd.FullName())
		fmt.Println("For service ", srvName)
		s.serviceDescriptorMap[string(srvName)] = rsd
		gsd := grpc.ServiceDesc{ServiceName: srvName, HandlerType: (*interface{})(nil)}

		methods := rsd.Methods()
		for m := 0; m < methods.Len(); m++ {
			method := methods.Get(m)
			methodName := string(method.Name())
			fmt.Println(" For method ", methodName)
			gsd.Methods = append(gsd.Methods, grpc.MethodDesc{MethodName: methodName, Handler: s.Handler})
		}
		s.grpcServer.RegisterService(&gsd, s)
	}
	return nil
}

func (s *GRPCService) createProtoRegistry(path string) (*protoregistry.Files, error) {
	marshalledDescriptorSet, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, err
	}
	descriptorSet := descriptorpb.FileDescriptorSet{}
	err = proto.Unmarshal(marshalledDescriptorSet, &descriptorSet)
	if err != nil {
		return nil, err
	}

	files, err := protodesc.NewFiles(&descriptorSet)
	if err != nil {
		return nil, err
	}

	return files, nil
}

func (s *GRPCService) Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	fmt.Print("handle")

	epContext, err := s.ConfigureEndpointContext(ctx, dec)
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

	return epContext.Response.Message, nil
}

func (s *GRPCService) ConfigureEndpointContext(ctx context.Context, dec func(interface{}) error) (*types.EndpointContext, error) {
	stream := grpc.ServerTransportStreamFromContext(ctx)
	arr := strings.Split(stream.Method(), "/")
	fmt.Println(stream.Method())

	serviceName := arr[1]
	methodName := arr[2]

	fmt.Printf(" for service %v and method %v", serviceName, methodName)

	service := s.serviceDescriptorMap[serviceName]

	epContext := &types.EndpointContext{}
	epContext.EndpointDescriptor = service.Methods().ByName(protoreflect.Name(methodName))

	input := epContext.EndpointDescriptor.Input()
	msg := dynamicpb.NewMessage(input)
	dec(msg)

	epContext.Request = &types.GRequest{Message: msg}
	epContext.Response = &types.GResponse{}

	err := ConfigureHandlers(epContext)
	if err != nil {
		return nil, err
	}

	return epContext, nil
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

func ParseExtensions(epContext types.EndpointContext) *gpb.EndpointConfig {
	options := epContext.EndpointDescriptor.Options()
	ex := proto.GetExtension(options, gpb.E_EndpointConfig)
	config := ex.(*gpb.EndpointConfig)
	fmt.Println(config.AuthType)
	return config
}
