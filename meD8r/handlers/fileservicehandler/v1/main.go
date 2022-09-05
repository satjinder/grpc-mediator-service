package fileservicehandler

import (
	"context"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/satjinder/grpc-mediator-service/med8r/types"
	"github.com/satjinder/grpc-mediator-service/med8r/utils"
	gpb "go.buf.build/grpc/go/satjinder/med8r/med8rtype/v1"
	"google.golang.org/protobuf/reflect/protoreflect"
	"google.golang.org/protobuf/types/dynamicpb"
)

const FileName = "filename"

func (handler *Handler) Init(handlerConfig *gpb.Handler, method protoreflect.MethodDescriptor) error {
	hctx, err := utils.SetHandlerContext(method, handlerConfig, []string{FileName})
	if err != nil {
		return err
	}
	handler.Context = hctx
	return nil
}

type Handler struct {
	Context *types.HandlerContext
}

func (handler *Handler) Process(epCtx context.Context) error {
	fmt.Println(handler.Context.HandlerConfig)
	val := epCtx.Value(types.ENDPOINT_CONTEXT_KEY)
	epContext := val.(*types.EndpointContext)

	filename := epContext.Request.Message.Get(handler.Context.Fields[FileName]).String()
	fmt.Println(filename)
	jsonBytes, _ := CallExternalAPI(filename)
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
