package main

import (
	"github.com/gin-gonic/gin"
	"github.com/nelsonomoi/gin_gonic_poc/controllers"
	"github.com/nelsonomoi/gin_gonic_poc/service"
)

var (
	videoService service.VideoService = service.New()
	videoController controllers.VideoController = controllers.New(videoService)
)

func main() {

	server := gin.Default()

	server.GET("/",func(ctx *gin.Context) {
		ctx.JSON(200,gin.H{
			"message": "ok",
		})
	})

	server.GET("/videos",func(ctx *gin.Context) {
		ctx.JSON(200,videoController.FindAll())
	})

	server.POST("/video",func(ctx *gin.Context) {
		ctx.JSON(200,videoController.Save(ctx))
	})

	server.Run(":3000")
}