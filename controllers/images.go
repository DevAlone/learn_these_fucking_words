package controllers

import (
	"net/http"

	. "../config"
	"../helpers"
	"github.com/gin-gonic/gin"
)

type ImageController struct{}

func (this *ImageController) Get(c *gin.Context) {
	var searchText = c.Param("search_text")

	url := "https://pixabay.com/api/?key=" + Settings["pixabay_api_key"].(string) + "&q=" + searchText

	resp, err := helpers.GetJson(url)

	if err != nil {
		_ = c.Error(err)
		c.JSON(http.StatusBadGateway, gin.H{"error_message": "some shit happened"})
		return
	}

	data := resp.(map[string]interface{})["hits"]

	c.JSON(http.StatusOK, gin.H{
		"data": data,
	})
}
