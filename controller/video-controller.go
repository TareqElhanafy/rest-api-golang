package controller

import (
	"RestApi/entity"
	"RestApi/service"
	"RestApi/validators"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator"
)

// VideoController interface
type VideoController interface {
	FindAll() []entity.Video
	Store(ctx *gin.Context) error
	Update(ctx *gin.Context) error
	Delete(ctx *gin.Context) error
}

type controller struct {
	service service.VideoService
}

var validate *validator.Validate

//NewVideoController instance of controller
func NewVideoController(service service.VideoService) VideoController {
	validate = validator.New()
	validate.RegisterValidation("is-cool", validators.TitleValidate)
	return &controller{
		service: service,
	}
}
func (c *controller) FindAll() []entity.Video {
	return c.service.FindAll()
}

func (c *controller) Store(ctx *gin.Context) error {
	var video entity.Video
	err := ctx.ShouldBindJSON(&video)

	if err != nil {
		return err
	}
	err = validate.Struct(video)
	if err != nil {
		return err
	}
	c.service.Store(video)
	return nil

}
func (c *controller) Update(ctx *gin.Context) error {
	var video entity.Video
	err := ctx.ShouldBindJSON(&video)

	if err != nil {
		return err
	}
	id, err := strconv.ParseUint(ctx.Param("id"), 0, 0)
	if err != nil {
		return err
	}
	video.ID = id

	err = validate.Struct(video)
	if err != nil {
		return err
	}
	c.service.Update(video)
	return nil

}

func (c *controller) Delete(ctx *gin.Context) error {
	var video entity.Video

	id, err := strconv.ParseUint(ctx.Param("id"), 0, 0)
	if err != nil {
		return err
	}
	video.ID = id
	if err != nil {
		return err
	}
	c.service.Delete(video)
	return nil

}
