package main

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
)

func TestPingRoute(t *testing.T) {
	gin.SetMode(gin.TestMode)
	router := gin.Default()

	router.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	req, _ := http.NewRequest(http.MethodGet, "/ping", nil)
	resp := httptest.NewRecorder()
	router.ServeHTTP(resp, req)

	if resp.Code != http.StatusOK {
		t.Errorf("期待されるステータスコード200ではありませんでした。得られたステータスコード: %v", resp.Code)
	}

	expected := `{"message":"pong"}`
	if resp.Body.String() != expected {
		t.Errorf("期待されるボディ'%v'ではありませんでした。得られたボディ: %v", expected, resp.Body.String())
	}
}
