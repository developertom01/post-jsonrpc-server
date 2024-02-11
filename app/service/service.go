package service

import (
	"context"

	"github.com/developertom01/post-jsonrpc-server/internal/db"
	"github.com/developertom01/post-jsonrpc-server/internal/logger"
)

// Application services
type Service interface {
	HandleAuthorization(ctx context.Context) (*db.User, error)

	//User services

	UserService

	//Posts services
	PostsServices
}

type service struct {
	db     db.Database
	logger logger.Logger
}

func NewService(db db.Database, logger logger.Logger) Service {
	return &service{
		db:     db,
		logger: logger,
	}
}
