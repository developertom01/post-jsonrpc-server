package rpc

import (
	"github.com/developertom01/post-jsonrpc-server/app/service"
)

func registerRpcServices(service service.Service) map[string]any {
	servicesMap := map[string]any{}

	servicesMap["PostService"] = NewPostService(service)

	return servicesMap
}
