package main

import (
	"dcard_2023_bk/handler"
	"dcard_2023_bk/pkg/postgres"
	"dcard_2023_bk/pkg/redis"

	"github.com/gin-gonic/gin"
)

func main() {
	load()

	router := gin.Default()
	router.RedirectFixedPath = true
	router.GET("/head/:id", handler.GetHead)
	router.GET("/page/:id", handler.GetPage)
	router.POST("/list/:id", handler.AddList)
	router.Run(":9876")
}

func load() {
	redis.Init()
	postgres.Init()
}
