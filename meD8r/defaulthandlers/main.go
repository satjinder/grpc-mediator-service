package defaulthandlers

import (
	"errors"
	"fmt"

	eh1 "github.com/satjinder/grpc-mediator-service/med8r/handlers/authorisationhandler/v1"
	eh2 "github.com/satjinder/grpc-mediator-service/med8r/handlers/authorisationhandler/v2"
	fh1 "github.com/satjinder/grpc-mediator-service/med8r/handlers/fileservicehandler/v1"
	hh1 "github.com/satjinder/grpc-mediator-service/med8r/handlers/httpservicehandler/v1"
	"github.com/satjinder/grpc-mediator-service/med8r/types"
)

type DefaultProvider struct {
}

func (dp *DefaultProvider) Get(name string) (types.Handler, error) {
	var handler types.Handler
	switch name {
	case "http-backend":
		handler = &hh1.Handler{}
	case "authorisation":
		handler = &eh1.Handler{}
	case "authorisationv2":
		handler = &eh2.Handler{}
	case "file-backend":
		handler = &fh1.Handler{}

	default:
		errMsg := fmt.Errorf("Handler not found %v", name)
		return nil, errors.New(errMsg.Error())
	}

	return handler, nil
}
