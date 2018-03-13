package controllers

import (
	. "../models"
	"../helpers"
	"github.com/gin-gonic/gin"
	"net/http"
)

//func GetUserMemorizations(context *gin.Context) {
//	username := context.Param("username")
//
//	var user User
//	err := DB.Model(&user).Where("username = ?", username).Select()
//
//	//if err != nil {
//	//	panic(err)
//	//}
//
//	var items []Memorization
//
//	err = DB.Model(&items).Where("user_id = ?", user.Id).Select()
//
//	if err != nil {
//		panic(err)
//	}
//
//	context.JSON(http.StatusOK, items)
//}

func GetMyMemorizations(context *gin.Context) {
	userId, _ := helpers.JWTGetCurrentUser(context)

	var items []Memorization

	err := DB.Model(&items).Column("User", "Word", "Word.Language").Where("user_id = ?", userId).Select()

	if err != nil {
		panic(err)
	}

	if items == nil {
		items = []Memorization{};
	}

	context.JSON(http.StatusOK, items)
}

func GetMemorizations(context *gin.Context) {
	var items []Memorization

	err := DB.Model(&items).Column("User", "Word").Select()

	if err != nil {
		panic(err)
	}

	if items == nil {
		items = []Memorization{};
	}

	context.JSON(http.StatusOK, items)
}
