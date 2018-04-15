package controllers

import . "models"
import . "config"

import (
	"github.com/gin-gonic/gin"
	"github.com/go-pg/pg"
	"helpers"
	"net/http"
	"time"
)

type LearningController struct{}

func (this *LearningController) GetWord(c *gin.Context) {
	userId, _ := helpers.JWTGetCurrentUser(c)

	memorization := Memorization{}

	err := DB.Model(&memorization).
		Column("Word").
		Where("user_id = ?", userId).
		Where("next_show_timestamp < ?", time.Now().Unix()).
		Order("memorization_coefficient").
		First()

	if err == pg.ErrNoRows {
		c.JSON(http.StatusNotFound, gin.H{
			"error_message": "there is no any word to learn, try to add one",
		})
		return
	}

	if err != nil {
		_ = c.Error(err)
		c.JSON(http.StatusBadRequest, gin.H{
			"error_message": "some shit happened",
		})
		return
	}

	memorization.NextShowTimestamp = time.Now().Unix() + Settings.LearningNextShowMinTime
	_, err = DB.Model(&memorization).
		Column("next_show_timestamp").
		Update()

	if err != nil {
		c.Error(err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"error_message": "some shit happened",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": memorization,
	})
}
