package handler

import (
	"context"
	"dcard_2023_bk/model"
	"dcard_2023_bk/pkg/redis"
	"encoding/json"
	"errors"
	"fmt"
	"math/rand"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	redis_glob "github.com/redis/go-redis/v9"
)

const letterBytes = "abcdefghijklmnopqrstuvwxyz1234567890"
const expireTime = 30 * time.Second

func getRandomStringId() string {
	idLen := 15
	idBytes := make([]byte, idLen)
	for i := range idBytes {
		idBytes[i] = letterBytes[rand.Intn(len(letterBytes))]
	}
	return string(idBytes)
}

func AddList(c *gin.Context) {
	var inputList model.List
	err := c.Bind(&inputList)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("%v\n", inputList)
	headId := c.Param("id")
	headKey := "head_" + headId
	pageId := getRandomStringId()
	setHead, err := redis.RC.SetNX(context.Background(), headKey, pageId, expireTime).Result()
	if err != nil {
		fmt.Println(setHead)
		fmt.Println(err)
	}

	for i := range inputList.Data {
		var newPage model.Page
		newPage.Articles = inputList.Data[i]
		if i != len(inputList.Data)-1 {
			newPage.NextPageKey = getRandomStringId()
		}
		pageKey := "page_" + pageId
		pageData, err := json.Marshal(newPage)
		if err != nil {
			fmt.Println(err)
		}
		setPage, err := redis.RC.SetNX(context.Background(), pageKey, pageData, expireTime).Result()
		if err != nil {
			fmt.Println(setPage)
			fmt.Println(err)
		}
		pageId = newPage.NextPageKey
	}
	c.JSON(http.StatusAccepted, gin.H{"status": "OK"})
}

func GetHead(c *gin.Context) {
	id := c.Param("id")
	val, err := redis.RC.Get(context.Background(), "head_"+id).Result()
	if errors.Is(err, redis_glob.Nil) {
		c.JSON(http.StatusNotFound, gin.H{})
		return
	} else if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{})
		panic(err)
	}
	fmt.Println("key", val)
	c.JSON(http.StatusOK, gin.H{"nextPageKey": val})
}

func GetPage(c *gin.Context) {
	id := c.Param("id")
	val, err := redis.RC.Get(context.Background(), "page_"+id).Bytes()
	if errors.Is(err, redis_glob.Nil) {
		c.JSON(http.StatusNotFound, gin.H{})
		return
	} else if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{})
		panic(err)
	}
	var pageInfo model.Page
	getPageErr := json.Unmarshal(val, &pageInfo)
	if getPageErr != nil {
		panic(getPageErr)
	}
	fmt.Println("key", pageInfo)
	if len(pageInfo.NextPageKey) > 0 {
		c.JSON(http.StatusOK, gin.H{"articles": pageInfo.Articles, "nextPageKey": pageInfo.NextPageKey})
	} else {
		c.JSON(http.StatusOK, gin.H{"articles": pageInfo.Articles})
	}
}