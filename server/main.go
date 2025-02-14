package main

import (
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var db *gorm.DB

func main() {
	var err error
	err = godotenv.Load(".env")
	if err != nil {
		log.Fatal("env err")
	}
	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbName := os.Getenv("DB_NAME")
	connStr := "user=" + dbUser + " password=" + dbPassword + " dbname=" + dbName + " sslmode=disable"
	db, err = gorm.Open(postgres.Open(connStr), &gorm.Config{})
	if err != nil {
		log.Fatal("Database connection error", err)
	}
	db.AutoMigrate(&user1{})
	r := gin.Default()
	r.POST("/register", registerUser)
	r.POST("/login", handleLogin)
	r.POST("/forgot-password", handlePassword)
	r.Run(":8080")
}
