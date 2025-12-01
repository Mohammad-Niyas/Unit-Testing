package test

import (
	"bytes"
	"encoding/json"
	"net/http"
	"testing"
	"net/http/httptest"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"github.com/stretchr/testify/assert"
	"github.com/glebarez/sqlite"
	"main.go/config"
	"main.go/controllers"
	"main.go/models"
)

func SetUpTestDB() {
	db, _ := gorm.Open(sqlite.Open(":memory:"), &gorm.Config{})
	db.AutoMigrate(&models.UserModel{})
	config.DB = db
}

func TestSignUpHandler(t *testing.T) {
	gin.SetMode(gin.TestMode)
	SetUpTestDB()
	r := gin.Default()
	r.POST("/signup", controllers.SignUp)
	t.Run("Valid Signup", func(t *testing.T) {
		user := models.UserModel{
			Name:     "Niyas",
			Email:    "niyas@gmail.com",
			Password: "Niyas@786",
		}
		body, err := json.Marshal(user)
		assert.NoError(t, err)
		req, err := http.NewRequest(http.MethodPost, "/signup", bytes.NewBuffer(body))
		assert.NoError(t, err)
		req.Header.Set("Content-Type", "application/json")
		w := httptest.NewRecorder()
		r.ServeHTTP(w, req)
		assert.Equal(t, http.StatusOK, w.Code)
		assert.JSONEq(t, `{"message":"User created successfully"}`, w.Body.String())
	})

	t.Run("Missing Fields", func(t *testing.T) {
		user := models.UserModel{
			Email:    "no-name@gmail.com",
			Password: "NoName",
		}
		body, _ := json.Marshal(user)
		req, _ := http.NewRequest(http.MethodPost, "/signup", bytes.NewBuffer(body))
		req.Header.Set("Content-Type", "application/json")
		w := httptest.NewRecorder()
		r.ServeHTTP(w, req)
		assert.Equal(t, http.StatusBadRequest, w.Code)
		assert.Contains(t, w.Body.String(), "All fields are required")
	})

	t.Run("Invalid JSON", func(t *testing.T) {
		invalidJson := []byte(`{"email":"fake@gmail.com",}`)

		req, _ := http.NewRequest(http.MethodPost, "/signup", bytes.NewBuffer(invalidJson))
		req.Header.Set("Content-Type", "application/json")
		w := httptest.NewRecorder()
		r.ServeHTTP(w, req)

		assert.Equal(t, http.StatusBadRequest, w.Code)
		assert.Contains(t, w.Body.String(), "Binding the data")
	})

	t.Run("Dupicate User", func(t *testing.T) {
		user := models.UserModel{
			Name:     "Niyas",
			Email:    "niyas@gmail.com",
			Password: "Niyas@786",
		}
		body, _ := json.Marshal(user)
		req, _ := http.NewRequest(http.MethodPost, "/signup", bytes.NewBuffer(body))
		req.Header.Set("Content-Type", "application/json")
		w := httptest.NewRecorder()

		r.ServeHTTP(w, req)
		assert.Equal(t, http.StatusInternalServerError, w.Code)
		assert.Contains(t, w.Body.String(), "User already exist")
	})
}