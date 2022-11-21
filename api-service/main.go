package main

import (
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type SateliteInfo struct {
	gorm.Model
	Text string `json:"text"`
}

func main() {
	env := os.Getenv("PROD_ENV")

	if env == "" {
		godotenv.Load(".env")
	}

	r := gin.Default()

	dbURL := os.Getenv("DATABASE_URL")

	db, err := gorm.Open(postgres.Open(dbURL), &gorm.Config{})

	if err != nil {
		log.Fatalln(err)
	}
	db.AutoMigrate(&SateliteInfo{})

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	r.Run(":8000")

}
