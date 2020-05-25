// Code generated by protoc-gen-micro. DO NOT EDIT.
// source: captcha_verfify.proto

package captcha_srv

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

// Api Endpoints for Captcha service

func NewCaptchaEndpoints() []*api.Endpoint {
	return []*api.Endpoint{}
}

// Client API for Captcha service

type CaptchaService interface {
	CaptchaVerfify(ctx context.Context, in *CaptchaVerfifyRequest, opts ...client.CallOption) (*CaptchaVerfifyResponse, error)
}

type captchaService struct {
	c    client.Client
	name string
}

func NewCaptchaService(name string, c client.Client) CaptchaService {
	return &captchaService{
		c:    c,
		name: name,
	}
}

func (c *captchaService) CaptchaVerfify(ctx context.Context, in *CaptchaVerfifyRequest, opts ...client.CallOption) (*CaptchaVerfifyResponse, error) {
	req := c.c.NewRequest(c.name, "Captcha.CaptchaVerfify", in)
	out := new(CaptchaVerfifyResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Captcha service

type CaptchaHandler interface {
	CaptchaVerfify(context.Context, *CaptchaVerfifyRequest, *CaptchaVerfifyResponse) error
}

func RegisterCaptchaHandler(s server.Server, hdlr CaptchaHandler, opts ...server.HandlerOption) error {
	type captcha interface {
		CaptchaVerfify(ctx context.Context, in *CaptchaVerfifyRequest, out *CaptchaVerfifyResponse) error
	}
	type Captcha struct {
		captcha
	}
	h := &captchaHandler{hdlr}
	return s.Handle(s.NewHandler(&Captcha{h}, opts...))
}

type captchaHandler struct {
	CaptchaHandler
}

func (h *captchaHandler) CaptchaVerfify(ctx context.Context, in *CaptchaVerfifyRequest, out *CaptchaVerfifyResponse) error {
	return h.CaptchaHandler.CaptchaVerfify(ctx, in, out)
}
