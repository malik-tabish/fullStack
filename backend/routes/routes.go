package routes

import (
	"gin-example/controllers"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes(routes *gin.Engine) {
	routes.GET("/", controllers.Home)
	routes.POST("user", controllers.CreateUserController)
	routes.GET("/users", controllers.GetAllUsersController)
	routes.GET("/users/:id", controllers.GetByIdController)
	routes.DELETE("/users/:id", controllers.DeleteByIdController)
	routes.PUT("/users/:id", controllers.UpdateUserController)
}
