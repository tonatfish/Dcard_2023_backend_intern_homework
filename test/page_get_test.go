package test

import (
	"dcard_2023_bk/handler"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetPage(t *testing.T) {
	r := SetUpRouter()
	r.POST("/page/:id", handler.GetPage)

	req, _ := http.NewRequest("GET", "/page/0", nil)

	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)
	assert.Equal(t, http.StatusNotFound, w.Code)
}
