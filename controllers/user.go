package controllers

import (
	"github.com/gin-gonic/gin"
	"go-api/repositories"
)

type UserController struct {
	Repository repositories.UserRepositoryInterface
}

func (controller UserController) GetAll(c *gin.Context) {
	users := controller.Repository.GetAll()
	c.JSON(200, users)
}
