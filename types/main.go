package types

import (
	"context"

	gpb "go.buf.build/grpc/go/satjinder/schemas/gproto/v1"
	"google.golang.org/protobuf/reflect/protoreflect"
	"google.golang.org/protobuf/types/dynamicpb"
)

const ENDPOINT_CONTEXT_KEY = "endpoint"

type HandlerProvider interface {
	Get(name string) (Handler, error)
}

type SchemaRegistry interface {
	Get(registryName string, protoFile string) (protoreflect.FileDescriptor, error)
}

type Handler interface {
	Init(handlerConfig *gpb.Handler, method protoreflect.MethodDescriptor) error
	Process(epCtx context.Context) error
}

type EndpointContext struct {
	EndpointDescriptor protoreflect.MethodDescriptor // Method descriptor (reflection) for the target endpoint
	Request            *GRequest                     // Contains reference to the request message. Each handler can read, validate or enrich it
	Response           *GResponse                    // Contains reference to the response message. Each handler can read, validate or enrich it
}

type ServerConfig struct {
	Services []ServiceConfig
}
type ServiceConfig struct {
	RegistryName string
	ProtoPath    string
}

type GRequest struct {
	Message *dynamicpb.Message
}

type GResponse struct {
	Message *dynamicpb.Message
}

type HandlerContext struct {
	HandlerConfig *gpb.Handler
	Options       map[string]string
	Fields        map[string]protoreflect.FieldDescriptor
}
