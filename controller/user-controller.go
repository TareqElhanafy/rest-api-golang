package controller

import (
	"RestApi/entity"
	"RestApi/service"

	"github.com/gin-gonic/gin"
)

//UserController interface
type UserController interface {
	StoreUser(ctx *gin.Context) error
	FindAllUsers() []entity.Person
}

type userController struct {
	service service.UserService
}

//NewUserController construction function
func NewUserController(service service.UserService) UserController {
	return &userController{
		service: service,
	}
}

func (c *userController) StoreUser(ctx *gin.Context) error {
	var user entity.Person
	err := ctx.ShouldBindJSON(&user)
	if err != nil {
		return err
	}
	c.service.StoreUser(user)
	return nil
}

func (c *userController) FindAllUsers() []entity.Person {
	return c.service.FindAllUsers()

}
