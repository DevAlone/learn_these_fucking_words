package views

import . "../models"

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func Index(context *gin.Context) {
	context.JSON(200, gin.H{"status":"ok"})
	return

	var words []Word

	err := DB.Model(&words).Select()
	if err != nil {
		panic(err)
	}

	var users []User

	err = DB.Model(&users).Select()
	if err != nil {
		panic(err)
	}

	var languages []Language

	err = DB.Model(&languages).Select()
	if err != nil {
		panic(err)
	}

	var memorizations []Memorization

	err = DB.Model(&memorizations).Select()
	if err != nil {
		panic(err)
	}

	context.HTML(http.StatusOK, "index.html", gin.H{
		"title":         "test",
		"words":         words,
		"users":         users,
		"languages":     languages,
		"memorizations": memorizations,
	})
}
