package controller

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"go-portfolio-service/cache"
	"go-portfolio-service/model"
	"math/rand"
	"net/http"
)

type TemplateService struct {
	_cache *cache.Cache
}

func NewTemplateController(_cache *cache.Cache) *TemplateService {
	return &TemplateService{
		_cache: _cache,
	}
}

func (t *TemplateService) GetNewTemplates(ctx echo.Context) error {
	response := make(map[string][]model.Templates)
	response["template"] = t._cache.NewData
	return ctx.JSON(http.StatusOK, response)
}

func (t *TemplateService) GetTemplateByColors(ctx echo.Context) error {
	request := struct {
		Color string `json:"color"`
	}{}

	if err := ctx.Bind(&request); err != nil {
		return err
	}

	response := make(map[string][]model.Templates)
	var data []model.Templates
	for _, value := range t._cache.All {
		if value.Color == request.Color {
			data = append(data, model.Templates{
				Code:    value.Code,
				Preview: value.Preview,
				Color:   value.Color,
			})
		}

		for _, child := range value.Child {
			if child.Color == request.Color {
				data = append(data, model.Templates{
					Color:   child.Color,
					Preview: child.Preview,
					Code:    child.Code,
				})
			}
		}
	}
	response["template"] = data
	return ctx.JSON(http.StatusOK, response)
}

func (t *TemplateService) GetTemplateByTheme(ctx echo.Context) error {
	request := struct {
		Theme string `json:"theme"`
	}{}

	if err := ctx.Bind(&request); err != nil {
		return err
	}

	fmt.Println(request.Theme)

	response := make(map[string][]model.Templates)
	var data []model.Templates
	for _, value := range t._cache.All {
		if value.Style == request.Theme {
			data = append(data, model.Templates{
				Code:    value.Code,
				Preview: value.Preview,
				Color:   value.Color,
				Style:   value.Style,
			})
		}
	}
	response["template"] = data
	return ctx.JSON(http.StatusOK, response)
}

func (t *TemplateService) GetAllTemplate(ctx echo.Context) error {
	response := make(map[string][]model.Templates)
	var data []model.Templates

	for _, template := range t._cache.All {
		data = append(data, model.Templates{
			Code:    template.Code,
			Preview: template.Preview,
			//Paper:   template.Paper,
			Style: template.Style,
			Price: template.Price,
		})
	}
	response["template"] = data
	return ctx.JSON(http.StatusOK, response)
}

func (t *TemplateService) GetTemplateID(ctx echo.Context) error {
	request := struct {
		Code string `json:"code"`
	}{}

	if err := ctx.Bind(&request); err == nil {
		for _, template := range t._cache.All {
			if request.Code == template.Code {
				response := make(map[string]model.Templates)
				response["template"] = model.Templates{
					Code:    template.Code,
					Preview: template.Preview,
					Color:   template.Color,
					Price:   template.Price,
					Paper:   template.Paper,
					Style:   template.Style,
					Child:   template.Child,
				}
				return ctx.JSON(http.StatusOK, response)
			}
		}
	}
	response := make(map[string]string)
	response["template"] = "not found"
	return ctx.JSON(http.StatusOK, response)
}

func (t *TemplateService) GetRecommendTemplate(ctx echo.Context) error {
	size := 4
	var data []model.Templates
	response := make(map[string][]model.Templates)
	var checkNum []int
	for true {
		if len(data) == size {
			break
		}
		iRand := rand.Intn(len(t._cache.All) - 1)
		found := false
		for _, n := range checkNum {
			if n == iRand {
				found = true
			}
		}
		if found {
			continue
		}
		checkNum = append(checkNum, iRand)
		template := t._cache.All[iRand]
		data = append(data, model.Templates{
			Code:    template.Code,
			Preview: template.Preview,
			Paper:   template.Paper,
			Style:   template.Style,
			Price:   template.Price,
		})
	}

	response["template"] = data

	return ctx.JSON(http.StatusOK, response)
}
