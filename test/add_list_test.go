package test

import (
	"bytes"
	"dcard_2023_bk/handler"
	"dcard_2023_bk/model"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAddList(t *testing.T) {
	r := setupRouter()
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
