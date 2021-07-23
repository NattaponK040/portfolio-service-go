package router

import (
	"github.com/labstack/echo/v4"
	"go-portfolio-service/controller"
	"go-portfolio-service/repository"
	"go-portfolio-service/service"
)

type Routes struct {
	ctx *echo.Echo
	mg  *repository.MongoRepository
}

func NewRoutes(ctx *echo.Echo, repository *repository.MongoRepository) *Routes {
	return &Routes{
		ctx: ctx,
		mg:  repository,
	}
}
func (r *Routes) InitRoute(tps *controller.TemplateService, out *service.CheckOut) {
	r.ctx.POST("/getTemplate/id",tps.GetTemplateID)
	r.ctx.POST("/getTemplate/new",tps.GetNewTemplates)
	r.ctx.POST("/getTemplate/color",tps.GetTemplateByColors)
	r.ctx.POST("/getTemplate/theme",tps.GetTemplateByTheme)
	r.ctx.POST("/getTemplate/all",tps.GetAllTemplate)
	r.ctx.POST("/getTemplate/recommend",tps.GetRecommendTemplate)
	r.ctx.POST("/user/checkout/alert",out.Checkout)
}
