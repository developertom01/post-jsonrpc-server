package service

import (
	"context"
	"errors"

	"github.com/developertom01/post-jsonrpc-server/config"
	"github.com/developertom01/post-jsonrpc-server/internal/db"
)

func (srv service) HandleAuthorization(ctx context.Context) (*db.User, error) {

	userId, ok := ctx.Value(config.AUTH_CONTEXT_KEY).(string)
	if !ok {
		srv.logger.Error("Decoded auth token is not of type `string`")

		return nil, errors.New("Internal server error")
	}

	return srv.db.GetUserById(context.TODO(), userId)
}
