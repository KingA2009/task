package user

import (
	"EduCRM/package/service"
	"EduCRM/util/logrus_log"
	"github.com/gin-gonic/gin"
)

type UserHandler struct {
	UserEndPoint
}
type UserEndPoint interface {
	SignInUser(ctx *gin.Context)
	SignUpUser(ctx *gin.Context)
	CreateUser(ctx *gin.Context)
	CreateUsers(ctx *gin.Context)
	UpdateUser(ctx *gin.Context)
	UpdateUsers(ctx *gin.Context)
	DeleteUser(ctx *gin.Context)
	DeleteUsers(ctx *gin.Context)
	GetUserList(ctx *gin.Context)
	GetUserByID(ctx *gin.Context)
}

func NewUserHandler(service *service.Service,
	loggers *logrus_log.Logger) *UserHandler {
	return &UserHandler{
		UserEndPoint: NewUserEndPointHandler(service, loggers),
	}
}
