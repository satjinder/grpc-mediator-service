// Code generated by protoc-gen-connect-go. DO NOT EDIT.
//
// Source: usstats/v2/usstats.proto

package usstatsv2connect

import (
	context "context"
	errors "errors"
	connect_go "github.com/bufbuild/connect-go"
	v2 "go.buf.build/grpc/go/satjinder/schemas/usstats/v2"
	http "net/http"
	strings "strings"
)

// This is a compile-time assertion to ensure that this generated file and the connect package are
// compatible. If you get a compiler error that this constant is not defined, this code was
// generated with a version of connect newer than the one compiled into your binary. You can fix the
// problem by either regenerating this code with an older version of connect or updating the connect
// version compiled into your binary.
const _ = connect_go.IsAtLeastVersion0_1_0

const (
	// StatsAPIName is the fully-qualified name of the StatsAPI service.
	StatsAPIName = "med8r.schemas.samples.usstats.v2.StatsAPI"
)

// StatsAPIClient is a client for the med8r.schemas.samples.usstats.v2.StatsAPI service.
type StatsAPIClient interface {
	GetStatsData(context.Context, *connect_go.Request[v2.GetStatsRequest]) (*connect_go.Response[v2.GetStatsResponse], error)
}

// NewStatsAPIClient constructs a client for the med8r.schemas.samples.usstats.v2.StatsAPI service.
// By default, it uses the Connect protocol with the binary Protobuf Codec, asks for gzipped
// responses, and sends uncompressed requests. To use the gRPC or gRPC-Web protocols, supply the
// connect.WithGRPC() or connect.WithGRPCWeb() options.
//
// The URL supplied here should be the base URL for the Connect or gRPC server (for example,
// http://api.acme.com or https://acme.com/grpc).
func NewStatsAPIClient(httpClient connect_go.HTTPClient, baseURL string, opts ...connect_go.ClientOption) StatsAPIClient {
	baseURL = strings.TrimRight(baseURL, "/")
	return &statsAPIClient{
		getStatsData: connect_go.NewClient[v2.GetStatsRequest, v2.GetStatsResponse](
			httpClient,
			baseURL+"/med8r.schemas.samples.usstats.v2.StatsAPI/GetStatsData",
			opts...,
		),
	}
}

// statsAPIClient implements StatsAPIClient.
type statsAPIClient struct {
	getStatsData *connect_go.Client[v2.GetStatsRequest, v2.GetStatsResponse]
}

// GetStatsData calls med8r.schemas.samples.usstats.v2.StatsAPI.GetStatsData.
func (c *statsAPIClient) GetStatsData(ctx context.Context, req *connect_go.Request[v2.GetStatsRequest]) (*connect_go.Response[v2.GetStatsResponse], error) {
	return c.getStatsData.CallUnary(ctx, req)
}

// StatsAPIHandler is an implementation of the med8r.schemas.samples.usstats.v2.StatsAPI service.
type StatsAPIHandler interface {
	GetStatsData(context.Context, *connect_go.Request[v2.GetStatsRequest]) (*connect_go.Response[v2.GetStatsResponse], error)
}

// NewStatsAPIHandler builds an HTTP handler from the service implementation. It returns the path on
// which to mount the handler and the handler itself.
//
// By default, handlers support the Connect, gRPC, and gRPC-Web protocols with the binary Protobuf
// and JSON codecs. They also support gzip compression.
func NewStatsAPIHandler(svc StatsAPIHandler, opts ...connect_go.HandlerOption) (string, http.Handler) {
	mux := http.NewServeMux()
	mux.Handle("/med8r.schemas.samples.usstats.v2.StatsAPI/GetStatsData", connect_go.NewUnaryHandler(
		"/med8r.schemas.samples.usstats.v2.StatsAPI/GetStatsData",
		svc.GetStatsData,
		opts...,
	))
	return "/med8r.schemas.samples.usstats.v2.StatsAPI/", mux
}

// UnimplementedStatsAPIHandler returns CodeUnimplemented from all methods.
type UnimplementedStatsAPIHandler struct{}

func (UnimplementedStatsAPIHandler) GetStatsData(context.Context, *connect_go.Request[v2.GetStatsRequest]) (*connect_go.Response[v2.GetStatsResponse], error) {
	return nil, connect_go.NewError(connect_go.CodeUnimplemented, errors.New("med8r.schemas.samples.usstats.v2.StatsAPI.GetStatsData is not implemented"))
}