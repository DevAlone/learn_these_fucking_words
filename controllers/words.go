package controllers

import . "../models"

import (
	"../settings"
	"../helpers"
	"strings"
	"errors"
	"time"
	"github.com/gin-gonic/gin"
	"net/http"
)

func addWord(languageId uint16, wordString string, userId uint64) (Word, error) {
	wordString = strings.ToLower(wordString)
	wordString = strings.TrimSpace(wordString)

	if len(wordString) == 0 {
		return Word{}, errors.New("wordString is too short")
	} else if len(wordString) > settings.MAX_WORD_LENGTH {
		return Word{}, errors.New("wordString is too long")
	}

	// TODO: detect language and user

	word := Word{Word: wordString, LanguageId: languageId}

	_, err := DB.Model(&word).Column("id").
		Where("word = ?word").Where("language_id = ?language_id").
		OnConflict("DO NOTHING").Returning("id").SelectOrInsert()

	if err != nil {
		// TODO: do not return error to user
		return Word{}, err
	}

	_, err = DB.Model(&Memorization {
		UserId:                  userId,
		WordId:                  word.Id,
		MemorizationCoefficient: 0,
		LastUpdateTimestamp:     uint64(time.Now().Unix()),
	}).OnConflict("(word_id, user_id) DO NOTHING").Insert()

	return word, err
}

func AddWord(context *gin.Context) {
	userId, _ := helpers.JWTGetCurrentUser(context)

	var word Word

	if err := context.ShouldBindJSON(&word); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	word, err := addWord(word.LanguageId, word.Word, userId)

	if err != nil {
		context.JSON(
			403,
			gin.H{"status": "error", "error_message": err.Error()},
		)
		return
	}

	context.JSON(200, word)
}

func GetWords(context *gin.Context) {
	var words []Word

	err := DB.Model(&words).
		Column("word.id", "word", "Language.id", "Language.code").Select()

	if err != nil {
		panic(err)
	}

	context.JSON(http.StatusOK, words)
}
