// Code generated by protoc-gen-micro. DO NOT EDIT.
// source: proto/v1/user.proto

package v1

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

// Api Endpoints for UserService service

func NewUserServiceEndpoints() []*api.Endpoint {
	return []*api.Endpoint{}
}

// Client API for UserService service

type UserService interface {
	GetMapping(ctx context.Context, in *GetMappingRequest, opts ...client.CallOption) (*GetMappingResponse, error)
}

type userService struct {
	c    client.Client
	name string
}

func NewUserService(name string, c client.Client) UserService {
	return &userService{
		c:    c,
		name: name,
	}
}

func (c *userService) GetMapping(ctx context.Context, in *GetMappingRequest, opts ...client.CallOption) (*GetMappingResponse, error) {
	req := c.c.NewRequest(c.name, "UserService.GetMapping", in)
	out := new(GetMappingResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for UserService service

type UserServiceHandler interface {
	GetMapping(context.Context, *GetMappingRequest, *GetMappingResponse) error
}

func RegisterUserServiceHandler(s server.Server, hdlr UserServiceHandler, opts ...server.HandlerOption) error {
	type userService interface {
		GetMapping(ctx context.Context, in *GetMappingRequest, out *GetMappingResponse) error
	}
	type UserService struct {
		userService
	}
	h := &userServiceHandler{hdlr}
	return s.Handle(s.NewHandler(&UserService{h}, opts...))
}

type userServiceHandler struct {
	UserServiceHandler
}

func (h *userServiceHandler) GetMapping(ctx context.Context, in *GetMappingRequest, out *GetMappingResponse) error {
	return h.UserServiceHandler.GetMapping(ctx, in, out)
}