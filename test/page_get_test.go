package test

import (
	"context"
	"dcard_2023_bk/handler"
	"dcard_2023_bk/model"
	"dcard_2023_bk/pkg/redis"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestGetPageExist(t *testing.T) {
	r := setupRouter()
	r.GET("/page/:id", handler.GetPage)

	pageId := "exist"

	var newPage model.Page
	newPage.Articles = []string{"1", "2", "3"}
	newPage.NextPageKey = "nextexist"
	pageData, err := json.Marshal(newPage)
	if err != nil {
		panic(err)
	}

	_, err = redis.RC.SetNX(context.Background(), "page_"+pageId, pageData, time.Minute).Result()
	if err != nil {
		panic(err)
	}

	req, _ := http.NewRequest("GET", "/page/"+pageId, nil)

	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)
	assert.Equal(t, http.StatusOK, w.Code)
}

func TestGetPageNotExist(t *testing.T) {
	r := setupRouter()
	r.GET("/page/:id", handler.GetPage)

	pageId := "notexist"

	req, _ := http.NewRequest("GET", "/page/"+pageId, nil)

	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)
	assert.Equal(t, http.StatusNotFound, w.Code)
}
