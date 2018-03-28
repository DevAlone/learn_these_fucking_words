package main

import (
	"./models"
	"./scheduler"
	"./server"
)

func main() {
	if err := models.InitDb(); err != nil {
		panic(err)
	}

	if err := scheduler.Init(); err != nil {
		panic(err)
	}

	if err := server.Run(); err != nil {
		panic(err)
	}

	if err := models.DB.Close(); err != nil {
		panic(err)
	}
}
