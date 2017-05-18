package main

import (
	"os"
	"testing"

	"net/http"
	"net/http/httptest"

	"github.com/gin-gonic/gin"
)

var tmpArticles []article

func TestMain(m *testing.M) {
	gin.SetMode(gin.TestMode)
	os.Exit(m.Run())
}

func getRouter(withTmplt bool) *gin.Engine {
	r := gin.Default()
	if withTmplt {
		r.LoadHTMLGlob("template/*")
	}
	return r
}

func testHttpResponze(t *testing.T, r *gin.Engine, req *http.Request, f func(w *httptest.ResponseRecorder) bool) {
	w := httptest.NewRecorder()

	r.ServeHTTP(w, req)

	if !f(w) {
		t.Fail()
	}
}

func saveArticles() {
	tmpArticles = articles
}
func restoreArticles() {
	articles = tmpArticles
}
