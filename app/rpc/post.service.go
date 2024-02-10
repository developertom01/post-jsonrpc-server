package rpc

import (
	"context"

	jsonrpc2 "github.com/developertom01/json-rpc2"
	"github.com/developertom01/post-jsonrpc-server/app/service"
)

type postService struct {
	service service.Service
}

func NewPostService(service service.Service) *postService {
	return &postService{
		service: service,
	}
}

func (postService) CreatePost(ctx context.Context) (any, error, *jsonrpc2.RpcErrorCode) {
	return nil, nil, nil
}
