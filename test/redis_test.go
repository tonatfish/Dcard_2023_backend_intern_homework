package test

import (
	"context"
	"dcard_2023_bk/pkg/redis"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestGetData(t *testing.T) {
	r := setupRouter()
	r.GET("/mock", nil)
	defer teardown()

	existId := "exist"
	expectData := "abcdefghijklmno"

	_, err := redis.RC.SetNX(context.Background(), existId, expectData, time.Minute).Result()
	if err != nil {
		panic(err)
	}

	data, err := redis.GetData(existId)
	if err != nil {
		panic(err)
	}

	assert.Equal(t, string(data), expectData)
}

func TestAddData(t *testing.T) {
	r := setupRouter()
	r.GET("/mock", nil)
	defer teardown()

	existId := "exist"
	expectData := "abcdefghijklmno"

	err := redis.AddData(existId, expectData)
	if err != nil {
		panic(err)
	}

	data, err := redis.RC.Get(context.Background(), existId).Bytes()
	if err != nil {
		panic(err)
	}

	assert.Equal(t, string(data), expectData)
}
