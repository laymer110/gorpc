// Code generated by protoc-gen-micro. DO NOT EDIT.
// source: proto/AppAuthServer.proto

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

// Api Endpoints for AppAuthServer service

func NewAppAuthServerEndpoints() []*api.Endpoint {
	return []*api.Endpoint{}
}

// Client API for AppAuthServer service

type AppAuthServerService interface {
	Auth(ctx context.Context, in *AuthRequest, opts ...client.CallOption) (*AuthResponse, error)
	TaskID(ctx context.Context, in *Task, opts ...client.CallOption) (*TaskKey, error)
	TaskUserID(ctx context.Context, in *TaskKey, opts ...client.CallOption) (*Task, error)
}

type appAuthServerService struct {
	c    client.Client
	name string
}

func NewAppAuthServerService(name string, c client.Client) AppAuthServerService {
	return &appAuthServerService{
		c:    c,
		name: name,
	}
}

func (c *appAuthServerService) Auth(ctx context.Context, in *AuthRequest, opts ...client.CallOption) (*AuthResponse, error) {
	req := c.c.NewRequest(c.name, "AppAuthServer.Auth", in)
	out := new(AuthResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *appAuthServerService) TaskID(ctx context.Context, in *Task, opts ...client.CallOption) (*TaskKey, error) {
	req := c.c.NewRequest(c.name, "AppAuthServer.TaskID", in)
	out := new(TaskKey)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *appAuthServerService) TaskUserID(ctx context.Context, in *TaskKey, opts ...client.CallOption) (*Task, error) {
	req := c.c.NewRequest(c.name, "AppAuthServer.TaskUserID", in)
	out := new(Task)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for AppAuthServer service

type AppAuthServerHandler interface {
	Auth(context.Context, *AuthRequest, *AuthResponse) error
	TaskID(context.Context, *Task, *TaskKey) error
	TaskUserID(context.Context, *TaskKey, *Task) error
}

func RegisterAppAuthServerHandler(s server.Server, hdlr AppAuthServerHandler, opts ...server.HandlerOption) error {
	type appAuthServer interface {
		Auth(ctx context.Context, in *AuthRequest, out *AuthResponse) error
		TaskID(ctx context.Context, in *Task, out *TaskKey) error
		TaskUserID(ctx context.Context, in *TaskKey, out *Task) error
	}
	type AppAuthServer struct {
		appAuthServer
	}
	h := &appAuthServerHandler{hdlr}
	return s.Handle(s.NewHandler(&AppAuthServer{h}, opts...))
}

type appAuthServerHandler struct {
	AppAuthServerHandler
}

func (h *appAuthServerHandler) Auth(ctx context.Context, in *AuthRequest, out *AuthResponse) error {
	return h.AppAuthServerHandler.Auth(ctx, in, out)
}

func (h *appAuthServerHandler) TaskID(ctx context.Context, in *Task, out *TaskKey) error {
	return h.AppAuthServerHandler.TaskID(ctx, in, out)
}

func (h *appAuthServerHandler) TaskUserID(ctx context.Context, in *TaskKey, out *Task) error {
	return h.AppAuthServerHandler.TaskUserID(ctx, in, out)
}
