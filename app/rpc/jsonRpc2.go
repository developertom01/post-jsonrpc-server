package rpc

import (
	jsonrpc2 "github.com/developertom01/jsonrpc2"
	"github.com/developertom01/post-jsonrpc-server/app/service"
)

type jsonRpc struct {
	rpc *jsonrpc2.JsonRPC
}

func (rpc *jsonRpc) GetRpc() *jsonrpc2.JsonRPC {
	return rpc.rpc
}

func NewJsonRpc(appServices service.Service) *jsonRpc {
	rpc := jsonrpc2.NewJsonRpc()

	for serviceName, srv := range registerRpcServices(appServices) {
		rpc.RegisterWithName(srv, serviceName)
	}

	return &jsonRpc{
		rpc: &rpc,
	}
}
