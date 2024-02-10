package db

import (
	"context"

	"github.com/developertom01/post-jsonrpc-server/internal/logger"
	"github.com/rs/zerolog"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type mongodb_impl struct {
	ctx    context.Context
	client *mongo.Client
	logger logger.Logger
	doc    *mongo.Database
}

func NewDatabase(url string, dbName string, logger logger.Logger) Database {
	l := logger.GetInstance().(zerolog.Logger)
	ctx := context.TODO()
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(url))
	if err != nil {
		l.Fatal().Msg("Could not connect to database")
	}
	l.Info().Msg("Database connection successful")

	database := client.Database(dbName)

	return &mongodb_impl{
		ctx:    ctx,
		client: client,
		logger: logger,
		doc:    database,
	}
}

func (d *mongodb_impl) GetClient() any {
	return d.client
}

func (d *mongodb_impl) GetDatabase() any {
	return d.doc
}

func (d *mongodb_impl) Disconnect() {
	l := d.logger.GetInstance().(zerolog.Logger)
	if err := d.client.Disconnect(d.ctx); err != nil {
		panic(err)
	}
	l.Info().Msg("Database disconnected")

}
