package entitlementshandler

import (
	"context"
	"fmt"

	"github.com/satjinder/grpc-mediator-service/med8r/types"
	"github.com/satjinder/grpc-mediator-service/med8r/utils"
	gpb "go.buf.build/grpc/go/satjinder/med8r/med8rtype/v1"
	"google.golang.org/protobuf/reflect/protoreflect"
)

const (
	role     = "role"
	resource = "resource"
)

func (handler *Handler) Init(handlerConfig *gpb.Handler, method protoreflect.MethodDescriptor) error {
	hctx, err := utils.SetHandlerContext(method, handlerConfig, []string{resource})
	if err != nil {
		return err
	}
	handler.Context = hctx
	return nil
}

type Handler struct {
	Context *types.HandlerContext
}

func (handler *Handler) Process(ctx context.Context) error {
	fmt.Println("check roles")

	val := ctx.Value(types.ENDPOINT_CONTEXT_KEY)
	epContext := val.(*types.EndpointContext)

	resource := epContext.Request.Message.Get(handler.Context.Fields[resource]).String()

	fmt.Printf("check %v role for resource %v", handler.Context.Options[role], resource)
	if resource != "Nation" {
		return fmt.Errorf("Incorrect Resource %v, expected %v", resource, "Nation")
	}

	return nil
}
