package httpservicehandler

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"regexp"
	"strings"

	"github.com/satjinder/grpc-mediator-service/utils"
	gpb "go.buf.build/grpc/go/satjinder/schemas/gproto/v1"
	"google.golang.org/protobuf/reflect/protoreflect"
)

const (
	HttpMethod    = "method"
	UrlPattern    = "url_pattern"
	HostConfigKey = "host_config_key"
	AuthType      = "auth_type"
	Body          = "body"
)

func (handler *Handler) Init(handlerConfig *gpb.Handler, method protoreflect.MethodDescriptor) error {
	handler.HandlerConfig = handlerConfig
	handler.Options = map[string]string{}
	handler.UrlFields = map[string]string{}

	for _, opt := range handlerConfig.Options {
		handler.Options[opt.Key] = opt.Value
	}

	url := handler.Options[UrlPattern]
	regX := regexp.MustCompile(`\{(.*?)\}`)
	for _, m := range regX.FindAllStringSubmatch(url, -1) {
		md := method.Input()
		fd := md.Fields().ByName(protoreflect.Name(m[1]))
		if fd == nil {
			errM := fmt.Errorf("Required '%v' field not found for '%v' by httpservicehandler", m[1], method.FullName())
			return errors.New(errM.Error())
		}
		handler.UrlFields[m[1]] = m[0]
	}

	return nil
}

type Handler struct {
	HandlerConfig *gpb.Handler
	Options       map[string]string
	UrlFields     map[string]string
}

func (handler *Handler) Process(epCtx context.Context) error {
	epContext := utils.GetEndpointFromContext(epCtx)

	requestJson, err := utils.ConvertDynamicToJson(epContext.Request.Message)
	if err != nil {
		return nil
	}

	var jsonData map[string]interface{}
	json.Unmarshal(requestJson, &jsonData)

	url := getUrl(handler, jsonData)

	resp, err := http.Get("https://datausa.io/" + url)
	if err != nil {
		return err
	}

	// We Read the response body on the line below.
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	respmsg, err := utils.ConvertToOutput(epContext, body)
	if err != nil {
		return err
	}
	epContext.Response.Message = respmsg
	return nil
}

func getUrl(handler *Handler, jsonData map[string]interface{}) string {
	url := handler.Options[UrlPattern]
	for field, val := range handler.UrlFields {
		replace := jsonData[field].(string)
		url = strings.Replace(url, val, replace, -1)
	}
	return url
}
