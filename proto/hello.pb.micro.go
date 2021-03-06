// Code generated by protoc-gen-micro. DO NOT EDIT.
// source: hello.proto

package proto

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	math "math"
)

import (
	context "context"
	api "github.com/micro/go-micro/v2/api"
	client "github.com/micro/go-micro/v2/client"
	server "github.com/micro/go-micro/v2/server"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

// Reference imports to suppress errors if they are not otherwise used.
var _ api.Endpoint
var _ context.Context
var _ client.Option
var _ server.Option

// Api Endpoints for Hello service

func NewHelloEndpoints() []*api.Endpoint {
	return []*api.Endpoint{}
}

// Client API for Hello service

type HelloService interface {
	Hello(ctx context.Context, in *Req, opts ...client.CallOption) (*Reply, error)
}

type helloService struct {
	c    client.Client
	name string
}

func NewHelloService(name string, c client.Client) HelloService {
	return &helloService{
		c:    c,
		name: name,
	}
}

func (c *helloService) Hello(ctx context.Context, in *Req, opts ...client.CallOption) (*Reply, error) {
	req := c.c.NewRequest(c.name, "Hello.Hello", in)
	out := new(Reply)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Hello service

type HelloHandler interface {
	Hello(context.Context, *Req, *Reply) error
}

func RegisterHelloHandler(s server.Server, hdlr HelloHandler, opts ...server.HandlerOption) error {
	type hello interface {
		Hello(ctx context.Context, in *Req, out *Reply) error
	}
	type Hello struct {
		hello
	}
	h := &helloHandler{hdlr}
	return s.Handle(s.NewHandler(&Hello{h}, opts...))
}

type helloHandler struct {
	HelloHandler
}

func (h *helloHandler) Hello(ctx context.Context, in *Req, out *Reply) error {
	return h.HelloHandler.Hello(ctx, in, out)
}
