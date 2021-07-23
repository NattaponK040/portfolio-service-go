package main

import (
	ctx "context"
	"github.com/sirupsen/logrus"
	"go-portfolio-service/cache"
	"go-portfolio-service/context"
	"go-portfolio-service/controller"
	"go-portfolio-service/repository"
	"go-portfolio-service/service"
)

func main() {
	server := context.CreateServer()
	client, err := repository.CreateMongoClient(&server.Config)

	if err != nil {
		logrus.Fatal(err)
	}
	mongo := repository.NewMongoRepository(client, &server.Config)
	defer client.Disconnect(ctx.TODO())
	//data := model.GetEducationTemplateModel()
	//for _, d := range data {
	//	mongo.TemplateCollection.InsertOne(mongo.CTX, d)
	//}
	_cache := cache.NewCache(mongo.TemplateCollection)
	_cache.Init()
	control := controller.NewTemplateController(_cache)
	cOut := service.NewCheckOut()

	server.Serv.POST("/getTemplate/id", control.GetTemplateID)
	server.Serv.POST("/getTemplate/new", control.GetNewTemplates)
	server.Serv.POST("/getTemplate/color", control.GetTemplateByColors)
	server.Serv.POST("/getTemplate/theme", control.GetTemplateByTheme)
	server.Serv.POST("/getTemplate/all", control.GetAllTemplate)
	server.Serv.POST("/getTemplate/recommend", control.GetRecommendTemplate)
	server.Serv.POST("/user/checkout/alert", cOut.Checkout)

	logrus.Fatal(server.Serv.Start(server.GetPort()))
}
