package user

import (
	"EduCRM/tools/middleware"
	"github.com/gin-gonic/gin"
)

func UserRouter(api *gin.Engine, handler *UserHandler) {
	user := api.Group("/api/v1/user")
	{
		user.POST("/sign-in", handler.UserEndPoint.SignInUser)
		user.POST("/sign-up", handler.UserEndPoint.SignUpUser)
	}
	userAuth := api.Group("/api/v1/user", middleware.AuthRequestHandler)
	{
		userAuth.POST("/create", handler.UserEndPoint.CreateUser)
		userAuth.POST("/update/:id", handler.UserEndPoint.UpdateUser)
		userAuth.DELETE("/delete/:id", handler.UserEndPoint.DeleteUser)
		userAuth.GET("/list", handler.UserEndPoint.GetUserList)
		userAuth.GET("/:id", handler.UserEndPoint.GetUserByID)
	}
	usersAuth := api.Group("/api/v1/users", middleware.AuthRequestHandler)
	{
		usersAuth.POST("/create", handler.UserEndPoint.CreateUsers)
		usersAuth.PUT("/update", handler.UserEndPoint.UpdateUsers)
		usersAuth.DELETE("/delete", handler.UserEndPoint.DeleteUsers)
	}

}
