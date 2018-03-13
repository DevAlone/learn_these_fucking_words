package controllers

import . "../models"

import (
	"../middlewares"
	"github.com/gin-gonic/gin"
	"net/http"
	"regexp"
	"golang.org/x/crypto/bcrypt"
	"strings"
)

func Register(c *gin.Context) {
	var user struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}

	if err := c.ShouldBindJSON(&user); err != nil {
		c.Error(err)
		c.JSON(http.StatusBadRequest, gin.H{"error_message": "Some error happened"})
		return
	}

	user.Username = strings.TrimSpace(user.Username)
	// user.Username = strings.ToLower(user.Username)

	match, _ := regexp.MatchString("^[a-zA-Z0-9_]{4,9}$", user.Username)
	if !match {
		c.JSON(http.StatusBadRequest, gin.H{"error_message": "bad username"})
		return
	}

	match, _ = regexp.MatchString("^.{6,64}$", user.Password)
	if !match {
		c.JSON(http.StatusBadRequest, gin.H{"error_message": "bad password"})
		return
	}

	count, err := DB.Model(&User{}).Where("LOWER(username) = ?", strings.ToLower(user.Username)).Count()

	if err != nil {
		c.Error(err)
		c.JSON(
			http.StatusInternalServerError,
			gin.H{"error_message": "some shit happened"})
		return
	}

	if count > 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error_message": "this username is already taken"})
		return
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)

	if err != nil {
		panic(err);
	}

	_, err = DB.Model(&User{}).
		Exec(`
			INSERT INTO users (username, password) 
			VALUES (?, ?)`, user.Username, hashedPassword)

	if err != nil {
		panic(err)
	}

	token, time, err := middlewares.AuthMiddleware.TokenGenerator("1,admin")

	if err != nil {
		c.Error(err)
		c.JSON(
			http.StatusInternalServerError,
			gin.H{"error_message": "some shit happened, try to login with new login and password"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"expire": time,
		"token": token,
	})
}
