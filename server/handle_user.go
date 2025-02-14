package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type user1 struct {
	User_id int    `gorm:"primaryKey" json:"id"`
	Name    string `gorm:"not null" json:"username"`
	Pass    string `gorm:"not null" json:"password"`
}

func registerUser(c *gin.Context) {
	var user user1
	err := c.ShouldBindJSON(&user)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Invalid request"})
		return
	}
	c.JSON(http.StatusAccepted, gin.H{"user": user.User_id, "name": user.Name})
	result := db.Create(&user)

	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Database Error"})
		return
	}
	c.JSON(http.StatusAccepted, gin.H{"message": "User Created"})
}

func handleLogin(c *gin.Context) {

}

func handlePassword(c *gin.Context) {

}
