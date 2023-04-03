package test

import (
	"dcard_2023_bk/handler"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetHead(t *testing.T) {
	r := SetUpRouter()
	r.POST("/head/:id", handler.GetHead)

	req, _ := http.NewRequest("GET", "/head/0", nil)

	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)
	assert.Equal(t, http.StatusNotFound, w.Code)
}
