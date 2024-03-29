// Code generated by protoc-gen-go-http. DO NOT EDIT.
// versions:
// - protoc-gen-go-http v2.7.2
// - protoc             v4.23.4
// source: idgen/v1/snowflake.proto

package v1

import (
	context "context"
	http "github.com/go-kratos/kratos/v2/transport/http"
	binding "github.com/go-kratos/kratos/v2/transport/http/binding"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the kratos package it is being compiled against.
var _ = new(context.Context)
var _ = binding.EncodeURL

const _ = http.SupportPackageIsVersion1

const OperationSnowflakeGenerate = "/api.idgen.v1.Snowflake/Generate"

type SnowflakeHTTPServer interface {
	Generate(context.Context, *GenerateRequest) (*GenerateReply, error)
}

func RegisterSnowflakeHTTPServer(s *http.Server, srv SnowflakeHTTPServer) {
	r := s.Route("/")
	r.GET("/api/idgen/v1/snowflake/generate", _Snowflake_Generate0_HTTP_Handler(srv))
}

func _Snowflake_Generate0_HTTP_Handler(srv SnowflakeHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in GenerateRequest
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationSnowflakeGenerate)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.Generate(ctx, req.(*GenerateRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*GenerateReply)
		return ctx.Result(200, reply)
	}
}

type SnowflakeHTTPClient interface {
	Generate(ctx context.Context, req *GenerateRequest, opts ...http.CallOption) (rsp *GenerateReply, err error)
}

type SnowflakeHTTPClientImpl struct {
	cc *http.Client
}

func NewSnowflakeHTTPClient(client *http.Client) SnowflakeHTTPClient {
	return &SnowflakeHTTPClientImpl{client}
}

func (c *SnowflakeHTTPClientImpl) Generate(ctx context.Context, in *GenerateRequest, opts ...http.CallOption) (*GenerateReply, error) {
	var out GenerateReply
	pattern := "/api/idgen/v1/snowflake/generate"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation(OperationSnowflakeGenerate))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "GET", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}
