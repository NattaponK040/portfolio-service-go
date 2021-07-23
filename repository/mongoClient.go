package repository

import (
	"context"
	"github.com/sirupsen/logrus"
	"go-portfolio-service/config"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func CreateMongoClient(cf *config.ServerConfig) (*mongo.Client, error) {
	option := options.Client().ApplyURI(cf.MongoDb.Uri).SetMinPoolSize(cf.MongoDb.MinPool).SetMinPoolSize(cf.MongoDb.MaxPool)
	mongoClient, err := mongo.Connect(context.TODO(), option)
	if err != nil {
		return nil, err
	}

	err = mongoClient.Ping(context.TODO(), nil)
	if err != nil {
		logrus.Error(err)
	}
	logrus.Info("MongoDB Connected!")
	return mongoClient, nil
}
