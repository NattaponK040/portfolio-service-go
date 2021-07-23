package context

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/sirupsen/logrus"
	"go-portfolio-service/config"
	"go-portfolio-service/logger"
	"os"
	"strconv"
)

type ServerContext struct {
	Serv   *echo.Echo
	Config config.ServerConfig
}

func CreateServer() *ServerContext {
	serv := &ServerContext{
		Serv: echo.New(),
		Config: config.LoadConfig(
			"",
			"resource",
			os.Getenv("ENV"),
			"application"),
	}

	serv.Serv.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowHeaders: []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept},
	}))
	logger.Logger = logrus.New()
	serv.Serv.Logger = logger.GetEchoLogger()
	serv.Serv.Use(logger.Hook())

	return serv
}

func (s *ServerContext) GetPort() string {
	var port = os.Getenv("PORT") // ----> (A)
	if port == "" {
		port = strconv.Itoa(s.Config.Server.Port)
	}
	return ":" + port // ----> (B)
}
