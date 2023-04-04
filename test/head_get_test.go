package test

import (
	"context"
	"dcard_2023_bk/handler"
	"dcard_2023_bk/pkg/redis"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestGetHeadExist(t *testing.T) {
	r := setupRouter()
	r.GET("/head/:id", handler.GetHead)
	defer teardown()

	headId := "exist"

	_, err := redis.RC.SetNX(context.Background(), "head_"+headId, "abcdefghijklmno", time.Minute).Result()
	if err != nil {
		panic(err)
	}

	req, _ := http.NewRequest("GET", "/head/"+headId, nil)

	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)
	assert.Equal(t, http.StatusOK, w.Code)
}

func TestGetHeadNotExist(t *testing.T) {
	r := setupRouter()
	r.GET("/head/:id", handler.GetHead)
	defer teardown()

	headId := "notexist"

	req, _ := http.NewRequest("GET", "/head/"+headId, nil)

	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)
	assert.Equal(t, http.StatusNotFound, w.Code)
}
