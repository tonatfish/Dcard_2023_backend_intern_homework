package test

import (
	"context"
	"dcard_2023_bk/pkg/postgres"
	"dcard_2023_bk/pkg/redis"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/alicebob/miniredis/v2"
	"github.com/gin-gonic/gin"
	redis_glob "github.com/redis/go-redis/v9"
	"github.com/stretchr/testify/assert"
)

var redisServer *miniredis.Miniredis
var mockPostgresServer sqlmock.Sqlmock

func setupRouter() *gin.Engine {
	redisServer = mockRedis()
	redis.RC = redis_glob.NewClient(&redis_glob.Options{
		Addr: redisServer.Addr(),
	})
	mockPostgres()

	router := gin.Default()
	router.RedirectFixedPath = true
	return router
}

func mockPostgres() {
	var err error
	postgres.PC, mockPostgresServer, err = sqlmock.New()
	if err != nil {
		panic(err)
	}
}

func mockRedis() *miniredis.Miniredis {
	s, err := miniredis.Run()
	if err != nil {
		panic(err)
	}
	return s
}

func teardown() {
	redisServer.Close()
}

func TestRedisConnection(t *testing.T) {
	redis.Init()
	_, err := redis.RC.Ping(context.Background()).Result()
	assert.Equal(t, err, nil)
}

func TestPostgreSQLConnection(t *testing.T) {
	postgres.Init()
	err := postgres.PC.Ping()
	assert.Equal(t, err, nil)
}
