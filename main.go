package main

import . "./models"

import (
	"./settings"
	"./models"
	"github.com/gin-gonic/gin"
	// "net/http"
	//"io/ioutil"
	"strings"
	"errors"
	"io/ioutil"
	"net/http"
	//"time"
	"time"
)

func index(context *gin.Context) {
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

	var languages[]Language

	err = DB.Model(&languages).Select()
	if err != nil {
		panic(err)
	}

	var memorizations[]Memorization

	err = DB.Model(&memorizations).Select()
	if err != nil {
		panic(err)
	}

	context.HTML(http.StatusOK, "index.html", gin.H{
		"title": "test",
		"words": words,
		"users": users,
		"languages": languages,
		"memorizations": memorizations,
	})
}

func _addWord(wordString string) error {
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
		UserId: 1,
		WordId: word.Id,
		MemorizationCoefficient: 0,
		LastUpdateTimestamp: uint64(time.Now().Unix()),
	}).OnConflict("(word_id, user_id) DO NOTHING").Insert()

	return err
}

func addWord(context *gin.Context) {
	_body, _ := ioutil.ReadAll(context.Request.Body)
	body := string(_body)

	words := strings.Split(body, " ")

	var errorsList []string

	for _, word := range words {
		err := _addWord(word)
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

func main() {
	defer DB.Close()

	err := models.InitDb()
	if err != nil {
		panic(err)
	}

	//gin.SetMode(gin.ReleaseMode)
	server := gin.Default()
	server.LoadHTMLGlob("templates/*")
	server.GET("/", index)
	server.POST("/api/words/", addWord)
	server.Run()
}
