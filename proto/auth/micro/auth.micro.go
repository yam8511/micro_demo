// Code generated by protoc-gen-micro. DO NOT EDIT.
// source: proto/auth/auth.proto

/*
Package pb is a generated protocol buffer package.

It is generated from these files:
	proto/auth/auth.proto

It has these top-level messages:
	PingRequest
	PongResponse
	User
	RegisterRequest
	RegisterResponse
	LoginRequest
	LoginResponse
	CheckRequest
	CheckResponse
	LogoutRequest
	LogoutResponse
*/
package pb

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

import (
	client "github.com/micro/go-micro/client"
	server "github.com/micro/go-micro/server"
	context "context"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ client.Option
var _ server.Option

// Client API for Auth service

type AuthClient interface {
	// Ping 測試服務連線
	Ping(ctx context.Context, in *PingRequest, opts ...client.CallOption) (*PongResponse, error)
	// Register 註冊
	Register(ctx context.Context, in *RegisterRequest, opts ...client.CallOption) (*RegisterResponse, error)
	// Login 登入
	Login(ctx context.Context, in *LoginRequest, opts ...client.CallOption) (*LoginResponse, error)
	// Check 確認是否登入
	Check(ctx context.Context, in *CheckRequest, opts ...client.CallOption) (*CheckResponse, error)
	// Logout 登出
	Logout(ctx context.Context, in *LogoutRequest, opts ...client.CallOption) (*LogoutResponse, error)
}

type authClient struct {
	c           client.Client
	serviceName string
}

func NewAuthClient(serviceName string, c client.Client) AuthClient {
	if c == nil {
		c = client.NewClient()
	}
	if len(serviceName) == 0 {
		serviceName = "pb"
	}
	return &authClient{
		c:           c,
		serviceName: serviceName,
	}
}

func (c *authClient) Ping(ctx context.Context, in *PingRequest, opts ...client.CallOption) (*PongResponse, error) {
	req := c.c.NewRequest(c.serviceName, "Auth.Ping", in)
	out := new(PongResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authClient) Register(ctx context.Context, in *RegisterRequest, opts ...client.CallOption) (*RegisterResponse, error) {
	req := c.c.NewRequest(c.serviceName, "Auth.Register", in)
	out := new(RegisterResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authClient) Login(ctx context.Context, in *LoginRequest, opts ...client.CallOption) (*LoginResponse, error) {
	req := c.c.NewRequest(c.serviceName, "Auth.Login", in)
	out := new(LoginResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authClient) Check(ctx context.Context, in *CheckRequest, opts ...client.CallOption) (*CheckResponse, error) {
	req := c.c.NewRequest(c.serviceName, "Auth.Check", in)
	out := new(CheckResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authClient) Logout(ctx context.Context, in *LogoutRequest, opts ...client.CallOption) (*LogoutResponse, error) {
	req := c.c.NewRequest(c.serviceName, "Auth.Logout", in)
	out := new(LogoutResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Auth service

type AuthHandler interface {
	// Ping 測試服務連線
	Ping(context.Context, *PingRequest, *PongResponse) error
	// Register 註冊
	Register(context.Context, *RegisterRequest, *RegisterResponse) error
	// Login 登入
	Login(context.Context, *LoginRequest, *LoginResponse) error
	// Check 確認是否登入
	Check(context.Context, *CheckRequest, *CheckResponse) error
	// Logout 登出
	Logout(context.Context, *LogoutRequest, *LogoutResponse) error
}

func RegisterAuthHandler(s server.Server, hdlr AuthHandler, opts ...server.HandlerOption) {
	s.Handle(s.NewHandler(&Auth{hdlr}, opts...))
}

type Auth struct {
	AuthHandler
}

func (h *Auth) Ping(ctx context.Context, in *PingRequest, out *PongResponse) error {
	return h.AuthHandler.Ping(ctx, in, out)
}

func (h *Auth) Register(ctx context.Context, in *RegisterRequest, out *RegisterResponse) error {
	return h.AuthHandler.Register(ctx, in, out)
}

func (h *Auth) Login(ctx context.Context, in *LoginRequest, out *LoginResponse) error {
	return h.AuthHandler.Login(ctx, in, out)
}

func (h *Auth) Check(ctx context.Context, in *CheckRequest, out *CheckResponse) error {
	return h.AuthHandler.Check(ctx, in, out)
}

func (h *Auth) Logout(ctx context.Context, in *LogoutRequest, out *LogoutResponse) error {
	return h.AuthHandler.Logout(ctx, in, out)
}