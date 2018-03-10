package main

//import . "./models"

import (
	"./settings"
	"./models"
	"fmt"
	"github.com/gin-gonic/gin"
	// "net/http"
	"io/ioutil"
	"strings"
)

func index(context *gin.Context) {
	context.String(200, "OK\n")
}

func addWord(context *gin.Context) {
	_body, _ := ioutil.ReadAll(context.Request.Body)
	body := string(_body)
	body = strings.TrimSpace(body)
	body = strings.ToLower(body)

	if len(body) == 0 {
		context.JSON(403, gin.H{"status": "error", "error_message": "word is too short"})
		return
	}

	if len(body) > settings.MAX_WORD_LENGTH {
		context.JSON(403, gin.H{"status": "error", "error_message": "word is too long"})
		return
	}

	//word := Word{Word: body}
	//DB.Insert(word)

	context.JSON(200, gin.H{"status": "ok"})
}

func main() {
	models.InitDb()
	//gin.SetMode(gin.ReleaseMode)
	server := gin.Default()
	server.GET("/", index)
	server.POST("/api/words/", addWord)
	server.Run()
	fmt.Println("It works!")
}
