package fileservicehandler

import (
	"context"
	"errors"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/satjinder/grpc-mediator-service/types"
	"github.com/satjinder/grpc-mediator-service/utils"
	gpb "go.buf.build/grpc/go/satjinder/schemas/gproto/v1"
	"google.golang.org/protobuf/reflect/protoreflect"
	"google.golang.org/protobuf/types/dynamicpb"
)

const FileName = "filename"

func (handler *Handler) Init(handlerConfig *gpb.Handler, method protoreflect.MethodDescriptor) error {
	handler.HandlerConfig = handlerConfig
	handler.Options = map[string]string{}

	for _, opt := range handlerConfig.Options {
		handler.Options[opt.Key] = opt.Value
	}

	md := method.Input()
	fd := md.Fields().ByName(protoreflect.Name(handler.Options[FileName]))
	if fd == nil {
		errM := fmt.Errorf("Required '%v' field not found for '%v' by fileservicehandler", handler.Options[FileName], method.FullName())
		return errors.New(errM.Error())
	}
	handler.FileNameField = fd

	return nil
}

type Handler struct {
	HandlerConfig *gpb.Handler
	Options       map[string]string
	FileNameField protoreflect.FieldDescriptor
}

func (handler *Handler) Process(epCtx context.Context) error {
	fmt.Println(handler.HandlerConfig)
	val := epCtx.Value(types.ENDPOINT_CONTEXT_KEY)
	epContext := val.(*types.EndpointContext)

	filename := epContext.Request.Message.Get(handler.FileNameField).String()
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
