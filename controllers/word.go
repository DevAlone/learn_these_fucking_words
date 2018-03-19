package controllers

import . "../models"

import (
	"../helpers"
	"../settings"
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/go-pg/pg"
	"net/http"
	"strings"
	"time"
)

type WordController struct{}

func (this *WordController) Add(context *gin.Context) {
	userId, _ := helpers.JWTGetCurrentUser(context)

	var userWord struct {
		Word         string `json:"word"`
		LanguageCode string `json:"languageCode"`
	}

	if err := context.ShouldBindJSON(&userWord); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error_message": err.Error()})
		return
	}

	if len(userWord.LanguageCode) < 2 {
		context.JSON(http.StatusBadRequest, gin.H{"error_message": "bad language"})
		return
	}

	memorization, wasAdded, err := addWord(userWord.LanguageCode, userWord.Word, userId, context)

	if err != nil {
		context.JSON(
			http.StatusBadRequest,
			gin.H{"status": "error", "error_message": err.Error()},
		)
		return
	}

	context.JSON(http.StatusOK, gin.H{
		"data":     *memorization,
		"wasAdded": wasAdded,
	})
}

func addWord(languageCode string, wordString string, userId uint64, c *gin.Context) (*Memorization, bool, error) {
	wordString = strings.ToLower(wordString)
	wordString = strings.TrimSpace(wordString)

	if len(wordString) == 0 {
		return nil, false, errors.New("wordString is too short")
	} else if len(wordString) > settings.MAX_WORD_LENGTH {
		return nil, false, errors.New("wordString is too long")
	}

	// TODO: detect language

	languageCode = strings.ToLower(languageCode)
	languageCode = strings.TrimSpace(languageCode)

	language := Language{}
	fmt.Println(language)
	err := DB.Model(&language).Where("code = ?", languageCode).Select()

	if err == pg.ErrNoRows {
		return nil, false, errors.New("language does not exist")
	}

	if err != nil {
		_ = c.Error(err)
		return nil, false, errors.New("some shit happened")
	}

	word := Word{Word: wordString, LanguageId: language.Id}

	err = DB.Model(&word).
		Column("word.id", "language_id", "Language").
		Where("word = ?word").
		Where("language_id = ?language_id").
		Select()

	if err == pg.ErrNoRows {
		_, err = DB.Model(&word).
			Returning("id").
			Insert()

		if err != nil {
			_ = c.Error(err)
			return nil, false, errors.New("some shit happened")
		}

		// TODO: select language
		err = DB.Model(&word).
			Column("Language").
			Where("word.id = ?id").
			Select()

		if err != nil {
			_ = c.Error(err)
			return nil, false, errors.New("some shit happened")
		}
	} else if err != nil {
		_ = c.Error(err)
		return nil, false, errors.New("some shit happened")
	}

	memorization := Memorization{
		UserId:                  userId,
		WordId:                  word.Id,
		MemorizationCoefficient: 0,
		LastUpdateTimestamp:     uint64(time.Now().Unix()),
	}
	result, err := DB.Model(&memorization).
		OnConflict("(word_id, user_id) DO NOTHING").
		Returning("memorization_coefficient", "last_update_timestamp").
		Insert()

	wasAdded := result.RowsAffected() > 0

	memorization.Word = &word

	if err != nil {
		_ = c.Error(err)
		return nil, false, errors.New("some shit happened")
	}

	return &memorization, wasAdded, nil
}

func (this *WordController) GetAll(context *gin.Context) {
	var words []Word

	err := DB.Model(&words).
		Column("word.id", "word", "Language.id", "Language.code").Select()

	if err != nil {
		panic(err)
	}

	context.JSON(http.StatusOK, words)
}
