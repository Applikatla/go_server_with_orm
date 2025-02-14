package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
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
	var log user1
	var store user1
	err := c.ShouldBindJSON(&log)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Internal Error"})
		return
	}
	result := db.Where("id = ?", log.User_id).First(&store)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Database Error"})
		return
	}
	if log.Pass == store.Pass {
		c.JSON(http.StatusAccepted, gin.H{"message": "Password Verified"})
		return
	} else {
		c.JSON(http.StatusForbidden, gin.H{"massage": "Password not match"})
		return
	}
}

func handlePassword(c *gin.Context) {
	var log user1
	var store user1
	err := c.ShouldBindJSON(&log)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "internal error"})
		return
	}
	// Fetch user by ID using GORM
	result := db.First(&store, log.User_id)
	// // featch with where
	// x := db.Where("id = ?", log.User_id).First(&store)

	// if x.Error != nil {
	// 	c.JSON(http.StatusInternalServerError, gin.H{"message": "internal server error"})
	// 	return
	// }
	// c.JSON(http.StatusAccepted, gin.H{"message": "reset link will send you soon"})
	if result.Error == gorm.ErrRecordNotFound {
		c.JSON(http.StatusNotFound, gin.H{"message": "User not found"})
		return
	}
	c.JSON(http.StatusAccepted, gin.H{"message": "reset link will send you soon"})

}
