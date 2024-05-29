// Code generated by protoc-gen-micro. DO NOT EDIT.
// source: BackendAPI/BackendAPI.proto

package gorpc

import (
	fmt "fmt"
	proto "google.golang.org/protobuf/proto"
	math "math"
)

import (
	context "context"
	api "go-micro.dev/v4/api"
	client "go-micro.dev/v4/client"
	server "go-micro.dev/v4/server"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// Reference imports to suppress errors if they are not otherwise used.
var _ api.Endpoint
var _ context.Context
var _ client.Option
var _ server.Option

// Api Endpoints for BackendAPI service

func NewBackendAPIEndpoints() []*api.Endpoint {
	return []*api.Endpoint{}
}

// Client API for BackendAPI service

type BackendAPIService interface {
	CallUnionDataBackend(ctx context.Context, in *BackendRequest, opts ...client.CallOption) (*BackendResponse, error)
}

type backendAPIService struct {
	c    client.Client
	name string
}

func NewBackendAPIService(name string, c client.Client) BackendAPIService {
	return &backendAPIService{
		c:    c,
		name: name,
	}
}

func (c *backendAPIService) CallUnionDataBackend(ctx context.Context, in *BackendRequest, opts ...client.CallOption) (*BackendResponse, error) {
	req := c.c.NewRequest(c.name, "BackendAPI.CallUnionDataBackend", in)
	out := new(BackendResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for BackendAPI service

type BackendAPIHandler interface {
	CallUnionDataBackend(context.Context, *BackendRequest, *BackendResponse) error
}

func RegisterBackendAPIHandler(s server.Server, hdlr BackendAPIHandler, opts ...server.HandlerOption) error {
	type backendAPI interface {
		CallUnionDataBackend(ctx context.Context, in *BackendRequest, out *BackendResponse) error
	}
	type BackendAPI struct {
		backendAPI
	}
	h := &backendAPIHandler{hdlr}
	return s.Handle(s.NewHandler(&BackendAPI{h}, opts...))
}

type backendAPIHandler struct {
	BackendAPIHandler
}

func (h *backendAPIHandler) CallUnionDataBackend(ctx context.Context, in *BackendRequest, out *BackendResponse) error {
	return h.BackendAPIHandler.CallUnionDataBackend(ctx, in, out)
}
