package routers

import (
	"github.com/gin-gonic/gin"
	"go-api/controllers"
	"go-api/repositories"
)

func init() {

}

func InitRoutes() *gin.Engine {
	router := gin.New()
	v1 := router.Group("api/v1")
	{
		AddRoutesUser(v1)
	}

	return router
}

func AddRoutesUser(r *gin.RouterGroup) {
	controller := controllers.UserController{Repository: repositories.NewUserRepository()}
	routes := r.Group("/users")
	{
		routes.GET("/", controller.GetAll)
	}
}
