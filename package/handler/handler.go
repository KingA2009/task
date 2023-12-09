package handler

import (
	"EduCRM/config"
	"EduCRM/docs"
	"EduCRM/package/handler/user"
	"EduCRM/package/service"
	"EduCRM/tools/middleware"
	"EduCRM/util/logrus_log"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

//	type Handler struct {
//		service *service.Service
//		loggers *loggers_log.Logger
//		config  *config.Configuration
//	}
//
//	handler_func NewHandler(service *service.Service, loggers *loggers_log.Logger, config *config.Configuration) *Handler {
//		return &Handler{service: service, loggers: loggers, config: config}
//	}
type Handler struct {
	User *user.UserHandler
}

func NewHandler(service *service.Service,
	loggers *logrus_log.Logger) *Handler {
	return &Handler{
		User: user.NewUserHandler(service, loggers),
	}
}
func (handler *Handler) InitRoutes() (route *gin.Engine) {
	cfg := config.Config()
	route = gin.New()
	//gin.SetMode(gin.ReleaseMode)
	route.HandleMethodNotAllowed = true
	middleware.GinMiddleware(route)
	//swagger settings
	docs.SwaggerInfo.Title = cfg.AppName
	docs.SwaggerInfo.Version = cfg.AppVersion
	//docs.SwaggerInfo.Host = cfg.AppURL
	//docs.SwaggerInfo.BasePath = cfg.AppURL + "/api/v1"
	docs.SwaggerInfo.Schemes = []string{"http", "https"}
	serverHost := ""
	route.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler),
		func(ctx *gin.Context) {
			serverHost = ctx.Request.Host
		})
	docs.SwaggerInfo.Host = serverHost
	route.Static("/public", "./public/")
	//routers
	user.UserRouter(route, handler.User)
	return
}
