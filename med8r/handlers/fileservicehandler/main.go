package fileservicehandler

import (
	"context"
	"fmt"
	"io/ioutil"
	"os"

	gpb "github.com/satjinder/med8r/schemas/gprotos"
	"github.com/satjinder/med8r/types"
	"github.com/satjinder/med8r/utils"
	"google.golang.org/protobuf/types/dynamicpb"
)

func NewHandler(handlerConfig *gpb.Handler) *Handler {
	return &Handler{HandlerConfig: handlerConfig}
}

type Handler struct {
	HandlerConfig *gpb.Handler
}

func (handler *Handler) Process(epCtx context.Context) error {
	fmt.Println(handler.HandlerConfig)
	val := epCtx.Value(types.ENDPOINT_CONTEXT_KEY)
	epContext := val.(*types.EndpointContext)

	requestJson, err := utils.ConvertDynamicToJson(epContext.Request.Message)
	if err != nil {
		return nil
	}
	fmt.Println()

	filename := string(requestJson)
	fmt.Println(filename)
	jsonBytes, _ := CallExternalAPI("response.json")
	outputDesc := epContext.EndpointDescriptor.Output()
	respmsg := dynamicpb.NewMessage(outputDesc)
	utils.PopulateDynamicWithJson(respmsg, jsonBytes)
	epContext.Response.Message = respmsg

	return nil
}

func CallExternalAPI(filename string) ([]byte, error) {
	jsonFile, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer jsonFile.Close()

	byteValue, err2 := ioutil.ReadAll(jsonFile)
	if err2 != nil {
		return nil, err2
	}
	jsonBytes := []byte(byteValue)

	return jsonBytes, nil
}
