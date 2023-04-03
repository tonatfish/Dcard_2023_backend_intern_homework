package test

import (
	"context"
	"dcard_2023_bk/pkg/redis"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func SetUpRouter() *gin.Engine {
	redis.Init()
	router := gin.Default()
	router.RedirectFixedPath = true
	return router
}

func TestRedisConnection(t *testing.T) {
	redis.Init()
	_, err := redis.RC.Ping(context.Background()).Result()
	assert.Equal(t, err, nil)
}
