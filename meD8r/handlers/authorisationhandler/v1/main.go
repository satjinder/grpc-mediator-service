package entitlementshandler

import (
	"context"
	"fmt"

	gpb "go.buf.build/grpc/go/satjinder/med8r/med8rtype/v1"
	"google.golang.org/protobuf/reflect/protoreflect"
)

func (handler *Handler) Init(handlerConfig *gpb.Handler, method protoreflect.MethodDescriptor) error {
	handler.HandlerConfig = handlerConfig
	return nil
}

type Handler struct {
	HandlerConfig *gpb.Handler
}

func (handler *Handler) Process(ctx context.Context) error {
	fmt.Println("check auth v1")

	// val := ctx.Value(types.ENDPOINT_CONTEXT_KEY)
	// epContext := val.(*types.EndpointContext)

	if handler.HandlerConfig != nil {
		for _, op := range handler.HandlerConfig.Options {
			fmt.Printf("operation:%v", op.Value)
			fmt.Println("")
		}
	}
	return nil
}
