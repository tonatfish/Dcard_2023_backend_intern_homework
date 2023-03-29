package main

import (
	"dcard_2023_bk/handler"
	"dcard_2023_bk/pkg/redis"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

var balance = 1000

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
}

func ownJson(c *gin.Context) {
	input := c.DefaultQuery("input", "haha")
	id := c.Param("id")
	type jsonReply struct {
		Name string `json:"name"`
		Unit int    `json:"unit"`
		Id   string `json:"id"`
	}
	var rs = jsonReply{
		Name: input,
		Unit: balance,
		Id:   id,
	}
	c.JSON(http.StatusAccepted, rs)
}

func changeBalance(c *gin.Context) {
	inputJson := make(map[string]interface{})
	c.BindJSON(&inputJson)
	input := fmt.Sprint(inputJson["check"])
	fmt.Println(input)
	amount, err := strconv.Atoi(input)
	if err == nil {
		balance = amount
		c.JSON(http.StatusOK, gin.H{
			"action": "OK",
		})
	} else {
		c.JSON(http.StatusBadRequest, gin.H{
			"action": "not OK",
		})
	}
}

func test(c *gin.Context) {
	str := []byte("ok")                      // 對於[]byte感到疑惑嗎？ 因為網頁傳輸沒有string的概念，都是要轉成byte字節方式進行傳輸
	c.Data(http.StatusOK, "text/plain", str) // 指定contentType為 text/plain，就是傳輸格式為純文字啦～
}
