package controller

import (
	"net/http"

	"github.com/Abeldlp/fullinfo/api-service/model"
	"github.com/gin-gonic/gin"
)

func SubscribeNewUser() gin.HandlerFunc {
	return func(c *gin.Context) {
		var user model.User
		if err := c.BindJSON(&user); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"message": "Could not serialize payload",
			})
			return
		}

		if err := model.CreateUser(&user).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"message": "Something went wrong",
			})
			return
		}

		c.JSON(http.StatusCreated, user)
	}
}
