package controllers

import (
	. "config"
	"errors"
	"fmt"
	"helpers"
	. "models"
	"net/http"
	"regexp"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/go-pg/pg"
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
	// wordString = strings.ToLower(wordString)
	wordString = strings.TrimSpace(wordString)

	if len(wordString) == 0 {
		return nil, false, errors.New("word is too short")
	} else if len(wordString) > Settings.MaxWordLength {
		return nil, false, errors.New("word is too long")
	}

	match, _ := regexp.MatchString("^([a-zA-Z0-9]+[ -_]?)+$", wordString)

	if !match {
		return nil, false, errors.New("word contains forbidden symbols")
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
		Where("LOWER(word) = LOWER(?word)").
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
		LastUpdateTimestamp:     time.Now().Unix(),
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

func (this *WordController) GetInfoByProvider(c *gin.Context) {
	providerName := c.Param("provider_name")
	word := c.Param("word")

	switch providerName {
	case "pearson.com":
		result, err := getInfoByProviderPearsonCom(word)
		if err == nil {
			c.JSON(http.StatusOK, gin.H{
				"data": result,
			})
		} else {
			_ = c.Error(err)
			c.JSON(http.StatusBadGateway, gin.H{
				"error_message": "something bad happened with provider",
			})
		}
	default:
		c.JSON(http.StatusBadRequest, gin.H{
			"error_message": "bad provider name",
		})
	}
}

func getInfoByProviderPearsonCom(word string) (interface{}, error) {
	url := "https://api.pearson.com/v2/dictionaries/entries?headword=" + word + "&limit=10"
	resp, err := helpers.GetJson(url)

	if err != nil {
		return nil, err
	}

	results := resp.(map[string]interface{})["results"].([]interface{})

	// TODO: convert senses/synonym to synses/synonyms

	for _, result := range results {
		if result, isOk := result.(map[string]interface{}); isOk {
			if senses, isOk := result["senses"]; isOk {
				if senses, isOk := senses.([]interface{}); isOk {
					for _, sense := range senses {
						if sense, isOk := sense.(map[string]interface{}); isOk {
							if definition, isOk := sense["definition"]; isOk {
								if definition, isOk := definition.(string); isOk {
									// convert it
									sense["definition"] = []string{definition}
								}
							}
						}
					}
				}
			}

		}
	}

	return results, nil
}
