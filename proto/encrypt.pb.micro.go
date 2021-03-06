// Code generated by protoc-gen-micro. DO NOT EDIT.
// source: encrypt.proto

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

// Api Endpoints for Encrypt service

func NewEncryptEndpoints() []*api.Endpoint {
	return []*api.Endpoint{}
}

// Client API for Encrypt service

type EncryptService interface {
	GetPublicKey(ctx context.Context, in *GetPublicKeyRequest, opts ...client.CallOption) (*GetPublicKeyResponse, error)
	Encrypt(ctx context.Context, in *EncryptRequest, opts ...client.CallOption) (*EncryptResponse, error)
	Decrypt(ctx context.Context, in *DecryptRequest, opts ...client.CallOption) (*DecryptResponse, error)
}

type encryptService struct {
	c    client.Client
	name string
}

func NewEncryptService(name string, c client.Client) EncryptService {
	return &encryptService{
		c:    c,
		name: name,
	}
}

func (c *encryptService) GetPublicKey(ctx context.Context, in *GetPublicKeyRequest, opts ...client.CallOption) (*GetPublicKeyResponse, error) {
	req := c.c.NewRequest(c.name, "Encrypt.GetPublicKey", in)
	out := new(GetPublicKeyResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *encryptService) Encrypt(ctx context.Context, in *EncryptRequest, opts ...client.CallOption) (*EncryptResponse, error) {
	req := c.c.NewRequest(c.name, "Encrypt.Encrypt", in)
	out := new(EncryptResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *encryptService) Decrypt(ctx context.Context, in *DecryptRequest, opts ...client.CallOption) (*DecryptResponse, error) {
	req := c.c.NewRequest(c.name, "Encrypt.Decrypt", in)
	out := new(DecryptResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Encrypt service

type EncryptHandler interface {
	GetPublicKey(context.Context, *GetPublicKeyRequest, *GetPublicKeyResponse) error
	Encrypt(context.Context, *EncryptRequest, *EncryptResponse) error
	Decrypt(context.Context, *DecryptRequest, *DecryptResponse) error
}

func RegisterEncryptHandler(s server.Server, hdlr EncryptHandler, opts ...server.HandlerOption) error {
	type encrypt interface {
		GetPublicKey(ctx context.Context, in *GetPublicKeyRequest, out *GetPublicKeyResponse) error
		Encrypt(ctx context.Context, in *EncryptRequest, out *EncryptResponse) error
		Decrypt(ctx context.Context, in *DecryptRequest, out *DecryptResponse) error
	}
	type Encrypt struct {
		encrypt
	}
	h := &encryptHandler{hdlr}
	return s.Handle(s.NewHandler(&Encrypt{h}, opts...))
}

type encryptHandler struct {
	EncryptHandler
}

func (h *encryptHandler) GetPublicKey(ctx context.Context, in *GetPublicKeyRequest, out *GetPublicKeyResponse) error {
	return h.EncryptHandler.GetPublicKey(ctx, in, out)
}

func (h *encryptHandler) Encrypt(ctx context.Context, in *EncryptRequest, out *EncryptResponse) error {
	return h.EncryptHandler.Encrypt(ctx, in, out)
}

func (h *encryptHandler) Decrypt(ctx context.Context, in *DecryptRequest, out *DecryptResponse) error {
	return h.EncryptHandler.Decrypt(ctx, in, out)
}
