package main

import (
	"RestApi/controller"
	"RestApi/middleware"
	"RestApi/repository"
	"RestApi/service"

	"net/http"

	"github.com/gin-gonic/gin"
)

var (
	databseRepository repository.DatabaseRepository = repository.NewDatabaseRepository()
	userRepository    repository.UserRepository     = repository.NewDatabaseRepository()
	videoRepository   repository.VideoRepository    = repository.NewDatabaseRepository()
	userService       service.UserService           = service.NewUserService(userRepository)
	videoService      service.VideoService          = service.NewVideoService(videoRepository)
	videoController   controller.VideoController    = controller.NewVideoController(videoService)
	userController    controller.UserController     = controller.NewUserController(userService)
)

func main() {

	defer databseRepository.CloseDB()
	server := gin.New()

	server.Use(gin.Recovery())

	apiRoutes := server.Group("/api", middleware.BasicAuth())
	{
		apiRoutes.GET("/videos", func(ctx *gin.Context) {

			ctx.JSON(200, videoController.FindAll())
		})

		apiRoutes.POST("/videos", func(ctx *gin.Context) {
			err := videoController.Store(ctx)
			if err != nil {
				ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			} else {
				ctx.JSON(http.StatusOK, gin.H{"message": "Video input is Invaild"})
			}

		})

		apiRoutes.PUT("/videos/:id", func(ctx *gin.Context) {
			err := videoController.Update(ctx)
			if err != nil {
				ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			} else {
				ctx.JSON(http.StatusOK, gin.H{"message": "video updated successfully"})
			}
		})

		apiRoutes.DELETE("/videos/:id", func(ctx *gin.Context) {
			err := videoController.Delete(ctx)
			if err != nil {
				ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			} else {
				ctx.JSON(http.StatusOK, gin.H{"message": "video deleted successfully"})
			}
		})

	}
	userRoutes := server.Group("/api")
	{
		userRoutes.GET("/users/", func(ctx *gin.Context) {
			ctx.JSON(200, userController.FindAllUsers())
		})

		userRoutes.POST("/users/", func(ctx *gin.Context) {
			err := userController.StoreUser(ctx)
			if err != nil {
				ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			} else {
				ctx.JSON(http.StatusOK, gin.H{"message": "user created successfully"})
			}
		})
	}
	server.Run(":8090")
}
