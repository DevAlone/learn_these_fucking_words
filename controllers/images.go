package controllers

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"net/http"
)

type ImageController struct{}

func (this *ImageController) Get(c *gin.Context) {
	var searchText = c.Param("search_text")

	url := "https://pixabay.com/api/?key=YOUR_KEY_GOES_HERE&q=" + searchText

	resp, err := http.Get(url)

	if err != nil {
		c.Error(err)
		c.JSON(http.StatusInternalServerError, gin.H{"error_message": "some shit happened"})
		return
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		c.Error(err)
		c.JSON(http.StatusInternalServerError, gin.H{"error_message": "some shit happened"})
		return
	}

	var jsonResult = make(map[string]interface{})

	if err := json.Unmarshal(body, &jsonResult); err != nil {
		c.Error(err)
		c.JSON(http.StatusInternalServerError, gin.H{"error_message": "some shit happened"})
		return
	}

	c.JSON(http.StatusOK, body)
}
