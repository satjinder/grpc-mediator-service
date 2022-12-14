// This components parse all the endpoints in the service protos (rpc methods), it validates the request message based on the handler configuration at the endpoint level. It also provides a method to process the request and ensures it calls all the handlers in the correct order
package endpoint

import (
	"context"
	"fmt"

	"github.com/satjinder/grpc-mediator-service/med8r/types"
	gpb "go.buf.build/grpc/go/satjinder/med8r/med8rtype/v1"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/reflect/protoreflect"
	"google.golang.org/protobuf/types/dynamicpb"
)

type Endpoint struct {
	method          protoreflect.MethodDescriptor
	handlers        []types.Handler     // Array of handlers appended in the required order of execution based on the EndpointConfig
	endpointConfig  *gpb.EndpointConfig // Endpoint configuration options defined by the target endpoint
	handlerProvider types.HandlerProvider
}

func NewEndpoint(method protoreflect.MethodDescriptor, handlerProvider types.HandlerProvider) (*Endpoint, error) {
	ep := &Endpoint{}
	ep.method = method
	ep.handlerProvider = handlerProvider
	ep.endpointConfig = parseExtensions(ep.method)
	err := ep.configureHandlers()
	if err != nil {
		return nil, err
	}
	return ep, nil
}

func (ep *Endpoint) Process(ctx context.Context, dec func(interface{}) error) (interface{}, error) {
	epContext := &types.EndpointContext{}

	input := ep.method.Input()
	msg := dynamicpb.NewMessage(input)
	dec(msg)

	epContext.Request = &types.GRequest{Message: msg}
	epContext.Response = &types.GResponse{}
	epContext.EndpointDescriptor = ep.method

	for _, handler := range ep.handlers {
		epCtx := context.WithValue(ctx, types.ENDPOINT_CONTEXT_KEY, epContext)
		err := handler.Process(epCtx)
		if err != nil {
			return nil, err
		}
	}

	return epContext.Response.Message, nil
}

func (ep *Endpoint) configureHandlers() error {
	for _, handlerConfig := range ep.endpointConfig.Handlers {
		handler, err := ep.handlerProvider.Get(handlerConfig.Name)

		if err != nil {
			return err
		}

		err = handler.Init(handlerConfig, ep.method)
		if err != nil {
			return err
		}
		ep.handlers = append(ep.handlers, handler)

	}
	return nil
}

func parseExtensions(method protoreflect.MethodDescriptor) *gpb.EndpointConfig {
	options := method.Options()
	ex := proto.GetExtension(options, gpb.E_EndpointConfig)
	config := ex.(*gpb.EndpointConfig)
	fmt.Println(config.AuthType)
	return config
}
