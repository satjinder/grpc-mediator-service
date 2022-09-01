// reference
// https://stackoverflow.com/questions/71320369/grpc-service-with-generic-proto-request-data-in-golang
// https://stackoverflow.com/questions/65561125/grpc-go-single-generic-service-handler

package genericserver

import (
	"context"
	"path/filepath"
	"strings"

	//"encoding/json"

	"fmt"
	"io/ioutil"

	endpoint "github.com/satjinder/grpc-mediator-service/endpoint"
	"github.com/satjinder/grpc-mediator-service/types"
	"google.golang.org/grpc"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/reflect/protodesc"
	"google.golang.org/protobuf/reflect/protoreflect"
	"google.golang.org/protobuf/reflect/protoregistry"
	"google.golang.org/protobuf/types/descriptorpb"
	//"github.com/satjinder/grpc-mediator-service/handlers/httpservicehandler"
)

type GenericServer struct {
	GrpcServer           *grpc.Server
	serviceDescriptorMap map[string]*GService
	config               types.ServerConfig
}

type GService struct {
	descriptor protoreflect.ServiceDescriptor
	endpoints  map[string]*endpoint.Endpoint
}

func NewServer(config types.ServerConfig) (*GenericServer, error) {
	gs := &GenericServer{
		GrpcServer:           grpc.NewServer(),
		serviceDescriptorMap: make(map[string]*GService),
		config:               config,
	}

	if err := gs.loadServices(); err != nil {
		return nil, err
	}

	return gs, nil
}

func (s *GenericServer) loadServices() error {
	for _, f := range s.config.Services {
		fdsFile := filepath.Join(*s.config.DescriptorSetDir, f.RegistryName)
		registry, err := s.createProtoRegistry(fdsFile)
		if err != nil {
			return err
		}
		if err := s.loadService(f.ProtoPath, registry); err != nil {
			return err
		}
	}
	return nil
}

func (s *GenericServer) loadService(serviceProtoName string, registry *protoregistry.Files) error {
	fd, err := registry.FindFileByPath(serviceProtoName)
	if err != nil {
		return err
	}
	services := fd.Services()

	for i := 0; i < services.Len(); i++ {
		rsd := &GService{descriptor: services.Get(i), endpoints: make(map[string]*endpoint.Endpoint)}
		srvName := string(rsd.descriptor.FullName())
		fmt.Println("For service ", srvName)
		s.serviceDescriptorMap[string(srvName)] = rsd
		gsd := grpc.ServiceDesc{ServiceName: srvName, HandlerType: (*interface{})(nil)}

		methods := rsd.descriptor.Methods()
		for m := 0; m < methods.Len(); m++ {
			method := methods.Get(m)
			methodName := string(method.Name())
			fmt.Println(" For method ", methodName)
			ep, err := endpoint.NewEndpoint(method)
			if err != nil {
				fmt.Errorf("Could not load endpoint %v for service %v", methodName, srvName)
				break
			}
			rsd.endpoints[string(methodName)] = ep
			gsd.Methods = append(gsd.Methods, grpc.MethodDesc{MethodName: methodName, Handler: s.Handler})
		}
		s.GrpcServer.RegisterService(&gsd, s)
	}
	return nil
}

func (s *GenericServer) Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	stream := grpc.ServerTransportStreamFromContext(ctx)
	arr := strings.Split(stream.Method(), "/")
	fmt.Println(stream.Method())

	serviceName := arr[1]
	methodName := arr[2]

	fmt.Printf(" for service %v and method %v", serviceName, methodName)

	service := s.serviceDescriptorMap[serviceName]
	ep := service.endpoints[methodName]
	return ep.Process(ctx, dec)
}

func (s *GenericServer) createProtoRegistry(path string) (*protoregistry.Files, error) {
	marshalledDescriptorSet, err := ioutil.ReadFile(path)
	if err != nil {
		fmt.Println(err)
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
