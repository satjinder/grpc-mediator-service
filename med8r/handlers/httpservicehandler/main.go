package httpservicehandler

import (
	"context"
	"io/ioutil"

	"net/http"

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
	resp, err := http.Get("https://datausa.io/api/data?drilldowns=Nation&measures=Population")
	if err != nil {
		return err
	}

	//We Read the response body on the line below.
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	//Convert the body to type string
	//fmt.Println(string(body))
	val := epCtx.Value(types.ENDPOINT_CONTEXT_KEY)
	epContext := val.(*types.EndpointContext)

	outputDesc := epContext.EndpointDescriptor.Output()
	respmsg := dynamicpb.NewMessage(outputDesc)
	err = utils.PopulateDynamicWithJson(respmsg, body)
	if err != nil {
		return err
	}
	epContext.Response.Message = respmsg
	return nil
}
