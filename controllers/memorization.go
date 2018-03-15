package controllers

import (
	"../helpers"
	. "../models"
	"github.com/gin-gonic/gin"
	"net/http"
)

type MemorizationController struct{}

func (this MemorizationController) GetMyMemorizations(context *gin.Context) {
	userId, _ := helpers.JWTGetCurrentUser(context)

	var items []Memorization

	err := DB.Model(&items).Column("User", "Word", "Word.Language").Where("user_id = ?", userId).Select()

	if err != nil {
		panic(err)
	}

	if items == nil {
		items = []Memorization{}
	}

	context.JSON(http.StatusOK, items)
}

func (this MemorizationController) GetAll(context *gin.Context) {
	var items []Memorization

	err := DB.Model(&items).Column("User", "Word").Select()

	if err != nil {
		panic(err)
	}

	if items == nil {
		items = []Memorization{}
	}

	context.JSON(http.StatusOK, items)
}
