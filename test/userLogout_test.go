package test

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"main.go/controllers"
)

func TestUserLogoutHandler(t *testing.T) {
	gin.SetMode(gin.TestMode)

	r := gin.Default()
	r.GET("/logout", controllers.UserLogoutHandler)
	req, err := http.NewRequest(http.MethodGet, "/logout", nil)
	assert.NoError(t, err)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)
	assert.Equal(t, http.StatusOK, w.Code)
	var responce map[string]string
	err = json.Unmarshal(w.Body.Bytes(), &responce)
	assert.NoError(t, err)
	assert.Equal(t, "Logout successful", responce["message"])

}