package test

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/Abeldlp/fullinfo/api-service/config"
	"github.com/Abeldlp/fullinfo/api-service/model"
	"github.com/Abeldlp/fullinfo/api-service/route"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/assert/v2"
	"gorm.io/gorm"
)

func beforeEachUserTest() (*gorm.DB, *gin.Engine) {
	test_db = config.TestDBInit()
	test_db.AutoMigrate(&model.User{})

	r := config.InitializeServer()
	route.AppendUserRoute(r)

	return test_db, r
}

func TestCreateUserEmail(t *testing.T) {
	test_db, r := beforeEachUserTest()

	body := []byte(`{
		"email": "test@mail.com",
		"timezone": "Europe/Amsterdam"
	}`)

	req, _ := http.NewRequest("POST", "/users", bytes.NewBuffer(body))
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	var user model.User

	test_db.First(&user)

	assert.Equal(t, user.Email, "test@mail.com")

	config.TestDBFree(test_db)
}

func TestCreateUserTimezone(t *testing.T) {
	test_db, r := beforeEachUserTest()

	body := []byte(`{
		"email": "test@mail.com",
		"timezone": "Europe/Amsterdam"
	}`)

	req, _ := http.NewRequest("POST", "/users", bytes.NewBuffer(body))
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	var user model.User

	test_db.First(&user)

	assert.Equal(t, user.Timezone, "Europe/Amsterdam")

	config.TestDBFree(test_db)
}
