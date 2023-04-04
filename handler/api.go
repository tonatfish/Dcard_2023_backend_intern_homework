package handler

import (
	"context"
	"dcard_2023_bk/model"
	"dcard_2023_bk/pkg/postgres"
	"dcard_2023_bk/pkg/redis"
	"encoding/json"
	"errors"
	"fmt"
	"math/rand"
	"net/http"

	_ "github.com/lib/pq"

	"github.com/gin-gonic/gin"
	redis_glob "github.com/redis/go-redis/v9"
)

const letterBytes = "abcdefghijklmnopqrstuvwxyz1234567890"
const idLen = 15
const redisPageLimit = 5

// create page id
func getRandomStringId() string {
	idBytes := make([]byte, idLen)
	for i := range idBytes {
		idBytes[i] = letterBytes[rand.Intn(len(letterBytes))]
	}
	return string(idBytes)
}

func errorCheck(err error) {
	if err != nil {
		fmt.Println(err)
		panic(err)
	}
}

// add list of page into db
// create both head and page
func AddList(c *gin.Context) {
	var inputList model.List
	err := c.Bind(&inputList)
	errorCheck(err)
	// fmt.Printf("%v\n", inputList)

	// add head
	headId := c.Param("id")
	headKey := "head_" + headId
	pageId := getRandomStringId()
	err = redis.AddData(headKey, pageId)
	errorCheck(err)

	// add page
	for i := range inputList.Data {
		var newPage model.Page
		newPage.Articles = inputList.Data[i]
		if i != len(inputList.Data)-1 {
			newPage.NextPageKey = getRandomStringId()
		}
		pageKey := "page_" + pageId
		pageData, err := json.Marshal(newPage)
		errorCheck(err)
		if i < redisPageLimit { // insert into redis
			err := redis.AddData(pageKey, string(pageData))
			errorCheck(err)
		} else { // insert into postgreSQL
			err := postgres.InsertPage(pageId, string(pageData))
			errorCheck(err)
		}
		pageId = newPage.NextPageKey
	}
	c.JSON(http.StatusAccepted, gin.H{"status": "OK"})
}

// get head information from redis
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
	c.JSON(http.StatusOK, gin.H{"nextPageKey": val})
}

// get page information from redis or postgreSQL
func GetPage(c *gin.Context) {
	id := c.Param("id")
	val, err := redis.GetData("page_" + id)

	// go postgreSQL when redis not work
	if errors.Is(err, redis_glob.Nil) {
		dbData, err := postgres.GetDBPage(id)
		errorCheck(err)
		if len(dbData) == 0 {
			c.JSON(http.StatusNotFound, gin.H{})
			return
		}
		val = dbData
	} else {
		errorCheck(err)
	}

	var pageInfo model.Page
	getPageErr := json.Unmarshal(val, &pageInfo)
	if getPageErr != nil {
		panic(getPageErr)
	}

	if len(pageInfo.NextPageKey) > 0 {
		c.JSON(http.StatusOK, gin.H{"articles": pageInfo.Articles, "nextPageKey": pageInfo.NextPageKey})
	} else {
		c.JSON(http.StatusOK, gin.H{"articles": pageInfo.Articles})
	}
}
