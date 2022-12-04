package model

import (
	"github.com/Abeldlp/iss/api-service/config"
	"gorm.io/gorm"
)

type User struct {
	Email    string `json:"email"`
	Timezone string `json:"timezone"`
}

func CreateUser(user *User) *gorm.DB {
	return config.DB.Create(&user)
}
