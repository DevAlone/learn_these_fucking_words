package controllers

import (
	. "config"
	"helpers"
	. "models"

	"github.com/gin-gonic/gin"
	"github.com/go-pg/pg"
	"net/http"
	"strconv"
	"time"
)

type MemorizationController struct{}

func (this *MemorizationController) GetMyMemorizations(context *gin.Context) {
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

func (this *MemorizationController) GetAll(context *gin.Context) {
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

func (this *MemorizationController) UpdateMyMemorization(c *gin.Context) {
	userId, _ := helpers.JWTGetCurrentUser(c)

	wordId, err := strconv.ParseInt(c.Param("word_id"), 10, 0)

	if err != nil || wordId <= 0 {
		c.JSON(http.StatusBadRequest, gin.H{
			"error_message": "Bad word",
		})
		return
	}

	var userData struct {
		MemorizationCoefficient float64 `json:"memorizationCoefficient"`
	}

	if err := c.ShouldBindJSON(&userData); err != nil {
		c.Error(err)
		c.JSON(http.StatusBadRequest, gin.H{"error_message": "bad json"})
		return
	}

	if userData.MemorizationCoefficient < 0 || userData.MemorizationCoefficient > 1 {
		c.JSON(http.StatusBadRequest, gin.H{"error_message": "bad coefficient"})
		return
	}

	memorization := Memorization{}

	err = DB.Model(&memorization).Column("Word").
		Where("user_id = ?", userId).Where("word_id = ?", wordId).First()

	if err == pg.ErrNoRows {
		c.JSON(http.StatusBadRequest, gin.H{
			"error_message": "you don't learn this word",
		})
		return
	}

	if err != nil {
		_ = c.Error(err)
		c.JSON(http.StatusBadRequest, gin.H{
			"error_message": "Some shit happened",
		})
		return
	}

	memorization.MemorizationCoefficient = userData.MemorizationCoefficient

	dt := float64(Settings.LearningNextShowMaxTime-Settings.LearningNextShowMinTime)*memorization.MemorizationCoefficient + float64(Settings.LearningNextShowMinTime)

	memorization.NextShowTimestamp = time.Now().Unix() + int64(dt)
	_, err = DB.Model(&memorization).
		Column("memorization_coefficient", "next_show_timestamp").
		Update()

	if err != nil {
		_ = c.Error(err)
		c.JSON(http.StatusBadRequest, gin.H{
			"error_message": "Some shit happened",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": memorization,
	})
}
