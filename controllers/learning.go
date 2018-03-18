package controllers

import . "../models"

import (
	"../helpers"
	"github.com/gin-gonic/gin"
	"github.com/go-pg/pg"
	"net/http"
)

type LearningController struct{}

func (this *LearningController) GetWord(c *gin.Context) {
	userId, _ := helpers.JWTGetCurrentUser(c)

	memorization := Memorization{}

	err := DB.Model(&memorization).
		Column("Word").
		Where("user_id = ?", userId).Order("memorization_coefficient").
		First()

	if err == pg.ErrNoRows {
		c.JSON(http.StatusNotFound, gin.H{
			"error_message": "there is no any word to learn, try to add one",
		})
		return
	}

	if err != nil {
		c.Error(err)
		c.JSON(http.StatusBadRequest, gin.H{
			"error_message": "some shit happened",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": memorization,
	})
}
