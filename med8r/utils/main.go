package utils

import (
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
