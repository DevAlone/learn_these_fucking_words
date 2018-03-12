package views

import . "../models"

import (
	"../settings"
	"strings"
	"errors"
	"time"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"net/http"
)

func addWord(wordString string) error {
	wordString = strings.ToLower(wordString)
	wordString = strings.TrimSpace(wordString)

	if len(wordString) == 0 {
		return errors.New("wordString is too short")
	} else if len(wordString) > settings.MAX_WORD_LENGTH {
		return errors.New("wordString is too long")
	}

	// TODO: detect language and user
	word := Word{Word: wordString, LanguageId: 1}

	_, err := DB.Model(&word).Column("id").
		Where("word = ?word").Where("language_id = ?language_id").
		OnConflict("DO NOTHING").Returning("id").SelectOrInsert()

	if err != nil {
		return err
	}

	_, err = DB.Model(&Memorization{
		UserId:                  1,
		WordId:                  word.Id,
		MemorizationCoefficient: 0,
		LastUpdateTimestamp:     uint64(time.Now().Unix()),
	}).OnConflict("(word_id, user_id) DO NOTHING").Insert()

	return err
}

func AddWord(context *gin.Context) {
	_body, _ := ioutil.ReadAll(context.Request.Body)
	body := string(_body)

	words := strings.Split(body, " ")

	var errorsList []string

	for _, word := range words {
		err := addWord(word)
		if err != nil {
			// TODO: log error and return something else to user
			errorsList = append(errorsList, err.Error())
		}
	}

	if len(errorsList) > 0 {
		context.JSON(
			403,
			gin.H{"status": "error", "error_message": "multiple errors happened(see errors)", "errors": errorsList},
		)
		return
	}

	context.JSON(200, gin.H{"status": "ok"})
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
