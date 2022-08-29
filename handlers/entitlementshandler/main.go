package entitlementshandler

import (
	"context"
	"fmt"

	gpb "github.com/satjinder/grpc-mediator-service/gen/gprotos"
)

func NewHandler(handlerConfig *gpb.Handler) *Handler {
	return &Handler{HandlerConfig: handlerConfig}
}

type Handler struct {
	HandlerConfig *gpb.Handler
}

func (handler *Handler) Process(ctx context.Context) error {
	fmt.Println("check entitlements")

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
