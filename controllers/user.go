package controllers

import . "../models"
import . "../config"

import (
	"../middlewares"
	"fmt"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"net/http"
	"regexp"
	"strings"
)

type UserController struct{}

func (this *UserController) GetAll(context *gin.Context) {
	var users []User

	err := DB.Model(&users).Select()

	if err != nil {
		panic(err)
	}

	context.JSON(http.StatusOK, users)
}

func (this *UserController) Register(c *gin.Context) {
	var user struct {
		Username string `json:"username"`
		Password string `json:"password"`
		Token    string `json:"token"`
	}

	if err := c.ShouldBindJSON(&user); err != nil {
		err = c.Error(err)
		if err != nil {
			panic(err)
		}
		c.JSON(http.StatusBadRequest, gin.H{"error_message": "Some error happened"})
		return
	}

	if user.Token != Settings.RegisterToken {
		c.JSON(http.StatusBadRequest, gin.H{"error_message": "you don't have a token, don't you?"})
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
		_ = c.Error(err)
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
		panic(err)
	}

	resultUser := User{
		Username: user.Username,
		Password: hashedPassword,
	}

	_, err = DB.Model(&resultUser).
		Returning("id").
		Insert()

	if err != nil {
		c.Error(err)
		c.JSON(http.StatusInternalServerError, gin.H{"error_message": "some shit happened"})
		return
	}

	token := fmt.Sprintf("%d,%s", resultUser.Id, resultUser.Username)
	token, time, err := middlewares.AuthMiddleware.TokenGenerator(token)

	if err != nil {
		_ = c.Error(err)
		c.JSON(
			http.StatusInternalServerError,
			gin.H{"error_message": "some shit happened, try to login with new login and password"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code":   200,
		"expire": time,
		"token":  token,
	})
}
