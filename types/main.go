package types

import (
	"context"

	gpb "github.com/satjinder/grpc-mediator-service/gen/gprotos"
	"google.golang.org/protobuf/reflect/protoreflect"
	"google.golang.org/protobuf/types/dynamicpb"
)

const ENDPOINT_CONTEXT_KEY = "endpoint"

type Handler interface {
	Process(epCtx context.Context) error
}

type EndpointContext struct {
	EndpointDescriptor protoreflect.MethodDescriptor // Method descriptor (reflection) for the target endpoint
	EndpointConfig     *gpb.EndpointConfig           // Endpoint configuration options defined by the target endpoint
	Handlers           []Handler                     // Array of handlers appended in the required order of execution based on the EndpointConfig
	Request            *GRequest                     // Contains reference to the request message. Each handler can read, validate or enrich it
	Response           *GResponse                    // Contains reference to the response message. Each handler can read, validate or enrich it
}

type GRequest struct {
	Message *dynamicpb.Message
}

type GResponse struct {
	Message *dynamicpb.Message
}
