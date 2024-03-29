// Code generated by protoc-gen-go-http. DO NOT EDIT.
// versions:
// - protoc-gen-go-http v2.7.2
// - protoc             v4.23.4
// source: idgen/v1/sequence.proto

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

const OperationSequenceNext = "/api.idgen.v1.Sequence/Next"

type SequenceHTTPServer interface {
	Next(context.Context, *NextRequest) (*NextReply, error)
}

func RegisterSequenceHTTPServer(s *http.Server, srv SequenceHTTPServer) {
	r := s.Route("/")
	r.GET("/api/idgen/v1/sequence/next", _Sequence_Next0_HTTP_Handler(srv))
}

func _Sequence_Next0_HTTP_Handler(srv SequenceHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in NextRequest
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationSequenceNext)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.Next(ctx, req.(*NextRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*NextReply)
		return ctx.Result(200, reply)
	}
}

type SequenceHTTPClient interface {
	Next(ctx context.Context, req *NextRequest, opts ...http.CallOption) (rsp *NextReply, err error)
}

type SequenceHTTPClientImpl struct {
	cc *http.Client
}

func NewSequenceHTTPClient(client *http.Client) SequenceHTTPClient {
	return &SequenceHTTPClientImpl{client}
}

func (c *SequenceHTTPClientImpl) Next(ctx context.Context, in *NextRequest, opts ...http.CallOption) (*NextReply, error) {
	var out NextReply
	pattern := "/api/idgen/v1/sequence/next"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation(OperationSequenceNext))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "GET", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}
