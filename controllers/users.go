package controllers

import . "../models"

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func GetUsers(context *gin.Context) {
	var users []User

	err := DB.Model(&users).Select()

	if err != nil {
		panic(err)
	}

	context.JSON(http.StatusOK, users)
}

