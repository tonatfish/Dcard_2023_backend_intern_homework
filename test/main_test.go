package test

import (
	"bytes"
	"dcard_2023_bk/handler"
	"dcard_2023_bk/model"
	"dcard_2023_bk/pkg/redis"
	"encoding/json"
	"net/http"
	"net/http/httptest"
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

func TestAddList(t *testing.T) {
	r := SetUpRouter()
	r.POST("/list/:id", handler.AddList)

	row1 := []string{"1", "2", "3"}
	row2 := []string{"4", "5", "6"}
	listData := [][]string{row1, row2}
	list := model.List{
		Data: listData,
	}
	jsonValue, _ := json.Marshal(list)
	req, _ := http.NewRequest("POST", "/list/0", bytes.NewBuffer(jsonValue))

	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)
	assert.Equal(t, http.StatusAccepted, w.Code)
}

func TestGetHead(t *testing.T) {
	r := SetUpRouter()
	r.POST("/head/:id", handler.GetHead)

	req, _ := http.NewRequest("GET", "/head/0", nil)

	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)
	assert.Equal(t, http.StatusNotFound, w.Code)
}

func TestGetPage(t *testing.T) {
	r := SetUpRouter()
	r.POST("/page/:id", handler.GetPage)

	req, _ := http.NewRequest("GET", "/page/0", nil)

	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)
	assert.Equal(t, http.StatusNotFound, w.Code)
}
