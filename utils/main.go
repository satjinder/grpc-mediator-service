package utils

import (
	"context"

	"github.com/satjinder/grpc-mediator-service/types"
	"google.golang.org/protobuf/encoding/protojson"
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
