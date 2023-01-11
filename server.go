package main

import (
	"github.com/gin-gonic/gin"
	"github.com/leke/goland-gin-lol/controller"
	"github.com/leke/goland-gin-lol/service"
)

var (
	videoService    service.VideoService       = service.New()
	videoController controller.VideoController = controller.New(videoService)
)

func main() {
	server := gin.Default()

	server.GET("/posts", func(ctx *gin.Context) {
		ctx.JSON(200, videoController.FindAll())
	})

	server.POST("/posts", func(ctx *gin.Context) {
		ctx.JSON(200, videoController.Save(ctx))

	})
	server.Run(":8080")
}
