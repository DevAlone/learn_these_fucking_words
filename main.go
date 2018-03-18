package main

//import . "./models"

import (
	"./controllers"
	"./middlewares"
	"./models"
	"github.com/gin-contrib/static"
	"github.com/gin-gonic/gin"
)

func main() {
	err := models.InitDb()
	if err != nil {
		panic(err)
	}

	//gin.SetMode(gin.ReleaseMode)
	router := gin.New()
	router.Use(gin.Logger())
	router.Use(gin.Recovery())

	router.Use(static.Serve("/", static.LocalFile("./frontend/dist", true))) // static.LocalFile("./frontend/dist/", false)))
	//router.LoadHTMLGlob("templates/*")

	// the jwt middleware

	wordController := controllers.WordController{}
	userController := controllers.UserController{}
	languageController := controllers.LanguageController{}
	memorizationController := controllers.MemorizationController{}
	learningController := controllers.LearningController{}

	router.POST("/api/v1/login", middlewares.AuthMiddleware.LoginHandler)
	router.POST("/api/v1/register", userController.Register)

	api := router.Group("/api/v1")
	api.Use(middlewares.AuthMiddleware.MiddlewareFunc())
	{
		api.GET("/refresh_token", middlewares.AuthMiddleware.RefreshHandler)
		api.GET("/words", wordController.GetAll)
		api.POST("/words", wordController.Add)
		//
		api.GET("/users", userController.GetAll)
		//
		api.GET("/languages", languageController.GetAll)
		api.GET("/memorizations", memorizationController.GetAll)

		api.GET("/my/memorizations", memorizationController.GetMyMemorizations)
		api.PATCH("/my/memorizations/:word_id", memorizationController.UpdateMyMemorization)

		api.GET("/learning/word", learningController.GetWord)
	}

	err = router.Run()

	if err != nil {
		panic(err)
	}

	err = models.DB.Close()

	if err != nil {
		panic(err)
	}
}
