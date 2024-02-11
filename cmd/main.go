package main

import (
	"log"
	"net/http"

	"github.com/developertom01/post-jsonrpc-server/app/middleware"
	"github.com/developertom01/post-jsonrpc-server/app/rpc"
	"github.com/developertom01/post-jsonrpc-server/app/service"
	"github.com/developertom01/post-jsonrpc-server/config"
	"github.com/developertom01/post-jsonrpc-server/internal/db"
	"github.com/developertom01/post-jsonrpc-server/internal/logger"
	"github.com/go-chi/chi/v5"
)

func main() {

	logger := logger.NewLogger()

	database := db.NewDatabase(config.DATABASE_URL, config.DATABASE_NAME, logger)

	appServices := service.NewService(database, logger)

	rpc := rpc.NewJsonRpc(appServices)
	r := chi.NewRouter()

	r.Use(middleware.AuthenticationMiddleware())
	r.Handle("/rpc", *rpc.GetRpc())

	if err := http.ListenAndServe(":8000", r); err != nil {
		log.Fatal("Server failed to start", err)
	}
}
