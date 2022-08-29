package httpservicehandler

import (
	"context"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"regexp"
	"strings"

	gpb "github.com/satjinder/grpc-mediator-service/gen/gprotos"
	"github.com/satjinder/grpc-mediator-service/utils"
)

const (
	HttpMethod    = "method"
	UrlPattern    = "url_pattern"
	HostConfigKey = "host_config_key"
	AuthType      = "auth_type"
	Body          = "body"
)

func NewHandler(handlerConfig *gpb.Handler) *Handler {
	handler := &Handler{HandlerConfig: handlerConfig, Options: map[string]string{}}
	for _, opt := range handlerConfig.Options {
		handler.Options[opt.Key] = opt.Value
	}
	return handler
}

type Handler struct {
	HandlerConfig *gpb.Handler
	Options       map[string]string
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
	regX := regexp.MustCompile(`\{(.*?)\}`)
	for _, m := range regX.FindAllStringSubmatch(`api/data?drilldowns={drilldowns}&measures={measures}`, -1) {
		replace := jsonData[m[1]].(string)
		url = strings.Replace(url, m[0], replace, -1)
	}
	return url
}
