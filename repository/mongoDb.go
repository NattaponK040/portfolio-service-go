package repository

import (
	"context"
	"go-portfolio-service/config"
	"go.mongodb.org/mongo-driver/mongo"
)

type MongoRepository struct {
	CTX                context.Context
	TemplateCollection *mongo.Collection
	Db                 *mongo.Client
}

func NewMongoRepository(db *mongo.Client, cfg *config.ServerConfig) *MongoRepository {
	return &MongoRepository{
		CTX:                context.Background(),
		TemplateCollection: db.Database(cfg.MongoDb.Database).Collection("template"),
		Db:                 db,
	}
}
