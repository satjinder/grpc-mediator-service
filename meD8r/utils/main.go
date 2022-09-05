package utils

import (
	"context"
	"errors"
	"fmt"

	"github.com/satjinder/grpc-mediator-service/med8r/types"
	gpb "go.buf.build/grpc/go/satjinder/med8r/med8rtype/v1"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/reflect/protoreflect"
	"google.golang.org/protobuf/types/dynamicpb"
)

func PopulateDynamicWithJson(dmsg *dynamicpb.Message, jsonBytes []byte) error {
	pm := dmsg.Interface()
	err := protojson.Unmarshal(jsonBytes, pm)
	if err != nil {
		return err
	}
	return nil
}

func ConvertDynamicToJson(dmsg *dynamicpb.Message) ([]byte, error) {
	jsonBytes, err := protojson.Marshal(dmsg)
	if err != nil {
		return nil, err
	}
	return jsonBytes, nil
}

func GetEndpointFromContext(epCtx context.Context) *types.EndpointContext {
	val := epCtx.Value(types.ENDPOINT_CONTEXT_KEY)
	epContext := val.(*types.EndpointContext)
	return epContext
}

func ConvertToOutput(epContext *types.EndpointContext, body []byte) (*dynamicpb.Message, error) {
	outputDesc := epContext.EndpointDescriptor.Output()
	respmsg := dynamicpb.NewMessage(outputDesc)
	err := PopulateDynamicWithJson(respmsg, body)
	if err != nil {
		return nil, err
	}
	return respmsg, nil
}

func SetHandlerContext(method protoreflect.MethodDescriptor, handlerConfig *gpb.Handler, fieldnames []string) (*types.HandlerContext, error) {
	hc := &types.HandlerContext{
		HandlerConfig: handlerConfig,
		Options:       map[string]string{},
		Fields:        map[string]protoreflect.FieldDescriptor{},
	}

	for _, opt := range handlerConfig.Options {
		hc.Options[opt.Key] = opt.Value
	}

	md := method.Input()

	for _, fn := range fieldnames {

		fd := md.Fields().ByName(protoreflect.Name(hc.Options[fn]))
		if fd == nil {
			errM := fmt.Errorf("Required '%v' field not found for '%v'", fn, method.FullName())
			return nil, errors.New(errM.Error())
		}
		hc.Fields[fn] = fd

	}

	return hc, nil
}
