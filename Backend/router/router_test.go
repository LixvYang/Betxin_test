package router

import (
	v1 "betxin/api/v1"
	"betxin/api/v1/topic"
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/zeebo/assert"
)

func SetupRouter() *gin.Engine {
	return gin.Default()
}

func TestListTopicsHandler(t *testing.T) {
	r := SetupRouter()
	r.POST("/api/v1/topic/list", topic.ListTopics)

	// values := map[string]interface{}{
	// 		"offset": 0,
	// 		"limit": 10,
	// 		"title": "",
	// 		"content": "",
	// }
	postBody := []byte(`{"offset": 0,"limit": 10,"title": "","content": "",}`)
	req, _ := http.NewRequest("POST", "http://localhost:3000/api/v1/topic/list", bytes.NewBuffer(postBody))
	req.Header.Set("Content-Type", "application/json; charset=UTF-8")
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	var res []v1.Response
	json.Unmarshal(w.Body.Bytes(), &res)
	fmt.Println(res)

	assert.Equal(t, http.StatusOK, w.Code)
}
