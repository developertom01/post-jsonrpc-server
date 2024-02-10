package rpcservices

import (
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
