package rpc

import (
	"github.com/developertom01/post-jsonrpc-server/app/rpc/rpcservices"
	"github.com/developertom01/post-jsonrpc-server/app/service"
)

func registerRpcServices(service service.Service) map[string]any {
	servicesMap := map[string]any{}

	servicesMap["PostService"] = rpcservices.NewPostService(service)

	return servicesMap
}
