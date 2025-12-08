package test

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"golang.org/x/crypto/bcrypt"
	"main.go/config"
	"main.go/controllers"
	"main.go/models"
)

func TestUserLoginHandler(t *testing.T) {
	gin.SetMode(gin.TestMode)
	SetUpTestDB()
	r := gin.Default()
	r.POST("/login", controllers.UserLoginHandler)

	password := "MyPassword@123"
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	assert.NoError(t, err)

	config.DB.Create(&models.UserModel{
		Email:    "testmail@gmail.com",
		Password: string(hashedPassword),
	})

	testCases := []struct {
		Name           string
		Input          interface{}
		ExpectedStatus int
		ExpectedBody   string
	}{
		{
			Name: "Valid Login",
			Input: map[string]string{
				"email":    "testmail@gmail.com",
				"password": "MyPassword@123",
			},
			ExpectedStatus: http.StatusOK,
			ExpectedBody:   `"Login successfull"`,
		},
		{
			Name: "Missing Fields",
			Input: map[string]string{
				"email": "",
			},
			ExpectedStatus: http.StatusBadRequest,
			ExpectedBody:   `"Email and Password are required"`,
		},
		{
			Name: "User Not Found",
			Input: map[string]string{
				"email":    "nouser@example.com",
				"password": "somepass",
			},
			ExpectedStatus: http.StatusUnauthorized,
			ExpectedBody:   `"User not found"`,
		},
		{
			Name: "Incorrect Password",
			Input: map[string]string{
				"email":    "testmail@gmail.com",
				"password": "wrongpassword",
			},
			ExpectedStatus: http.StatusUnauthorized,
			ExpectedBody:   `"invalid password"`,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.Name, func(t *testing.T) {
			body, err := json.Marshal(tc.Input)
			assert.NoError(t, err)
			req, err := http.NewRequest(http.MethodPost, "/login", bytes.NewBuffer(body))
			req.Header.Set("Content-Type", "application/json")
			assert.NoError(t, err)
			w := httptest.NewRecorder()
			r.ServeHTTP(w, req)
			assert.Equal(t, tc.ExpectedStatus, w.Code)
			assert.Contains(t, w.Body.String(), tc.ExpectedBody)
		})
	}
}