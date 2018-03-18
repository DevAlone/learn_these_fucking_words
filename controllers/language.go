package controllers

import . "../models"

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type LanguageController struct{}

func (this *LanguageController) GetAll(context *gin.Context) {
	var languages []Language

	err := DB.Model(&languages).Select()

	if err != nil {
		panic(err)
	}

	context.JSON(http.StatusOK, languages)
}
