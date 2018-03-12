package views

import . "../models"

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func GetUserMemorizations(context *gin.Context) {
	username := context.Param("username")

	var user User
	err := DB.Model(&user).Where("username = ?", username).Select()

	//if err != nil {
	//	panic(err)
	//}

	var items []Memorization

	err = DB.Model(&items).Where("user_id = ?", user.Id).Select()

	if err != nil {
		panic(err)
	}

	context.JSON(http.StatusOK, items)
}
