package cache

import (
	"context"
	"github.com/sirupsen/logrus"
	"go-portfolio-service/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"sync"
)

type Cache struct {
	NewData     []model.Templates
	All         []model.Templates
	templateCol *mongo.Collection
}

func NewCache(templateCol *mongo.Collection) *Cache {
	return &Cache{
		templateCol: templateCol,
	}
}

func (c *Cache) Init() {
	c.All = c.getAllData()
	c.NewData = c.getNewData()
	c.watchingNewData()
}

func (c *Cache) watchingNewData() {
	var waitGroup sync.WaitGroup
	episodesStream, err := c.templateCol.Watch(context.TODO(), mongo.Pipeline{})
	if err != nil {
		panic(err)
	}
	waitGroup.Add(1)

	routineCtx, _ := context.WithCancel(context.Background())
	go c.iterateChangeStream(routineCtx, waitGroup, episodesStream)
	go waitGroup.Wait()
}

func (c *Cache) iterateChangeStream(routineCtx context.Context, waitGroup sync.WaitGroup, stream *mongo.ChangeStream) {
	defer stream.Close(routineCtx)
	defer waitGroup.Done()
	for stream.Next(routineCtx) {
		c.Init()
	}
}

func (c *Cache) getNewData() []model.Templates {
	findOptions := options.Find()
	findOptions.SetSort(bson.D{{"_id", -1}})
	result, err := c.templateCol.Find(context.TODO(), bson.M{}, findOptions.SetLimit(10))
	if err != nil {
		logrus.Error(err)
	}

	var response []model.Templates
	for result.Next(context.TODO()) {
		var temp model.Templates
		if e := result.Decode(&temp); e != nil {
			logrus.Error(e)
		}
		response = append(response, temp)
	}
	return response
}

func (c *Cache) getAllData() []model.Templates {
	result, err := c.templateCol.Find(context.TODO(), bson.M{})
	if err != nil {
		logrus.Error(err)
	}

	var response []model.Templates
	for result.Next(context.TODO()) {
		var temp model.Templates
		if e := result.Decode(&temp); e != nil {
			logrus.Error(e)
		}
		response = append(response, temp)
	}

	return response
}
