package main

import (
	"log"
	"net/http"

	"github.com/developertom01/post-jsonrpc-server/app/rpc"
	"github.com/developertom01/post-jsonrpc-server/app/service"
	"github.com/developertom01/post-jsonrpc-server/config"
	"github.com/developertom01/post-jsonrpc-server/internal/db"
	"github.com/developertom01/post-jsonrpc-server/internal/logger"
)

func main() {
	s := http.NewServeMux()

	logger := logger.NewLogger()

	database := db.NewDatabase(config.DATABASE_URL, config.DATABASE_NAME, logger)

	appServices := service.NewService(database, logger)

	rpc := rpc.NewJsonRpc(appServices)
	s.Handle("/rpc", *rpc.GetRpc())

	if err := http.ListenAndServe(":8000", s); err != nil {
		log.Fatal("Server failed to start", err)
	}
}
