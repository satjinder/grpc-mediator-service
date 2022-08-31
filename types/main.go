package types

import (
	"context"

	gpb "github.com/satjinder/grpc-mediator-service/gen/gprotos"
	"google.golang.org/protobuf/reflect/protoreflect"
	"google.golang.org/protobuf/types/dynamicpb"
)

const ENDPOINT_CONTEXT_KEY = "endpoint"

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
	DescriptorSetDir *string
	Services         []ServiceConfig
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
