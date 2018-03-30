package server

import (
	. "../config"
	"../controllers"
	"../middlewares"
	"github.com/gin-contrib/static"
	"github.com/gin-gonic/gin"
)

func Run() error {
	if !Settings.Debug {
		gin.SetMode(gin.ReleaseMode)
	}

	router := gin.New()
	router.Use(gin.Logger())
	router.Use(gin.Recovery())

	if Settings.Debug {
		router.Use(static.Serve("/", static.LocalFile("./frontend/dist", true)))
	}

	// the jwt middleware

	wordController := controllers.WordController{}
	userController := controllers.UserController{}
	languageController := controllers.LanguageController{}
	memorizationController := controllers.MemorizationController{}
	learningController := controllers.LearningController{}
	imageController := controllers.ImageController{}

	router.POST("/api/v1/login", middlewares.AuthMiddleware.LoginHandler)
	router.POST("/api/v1/register", userController.Register)

	api := router.Group("/api/v1")
	api.Use(middlewares.AuthMiddleware.MiddlewareFunc())
	{
		api.GET("/refresh_token", middlewares.AuthMiddleware.RefreshHandler)
		api.GET("/words", wordController.GetAll)
		api.POST("/my/words", wordController.Add)
		api.GET("/word_info_items/:provider_name/:word", wordController.GetInfoByProvider)
		//
		api.GET("/users", userController.GetAll)
		//
		api.GET("/languages", languageController.GetAll)
		api.GET("/memorizations", memorizationController.GetAll)

		api.GET("/my/memorizations", memorizationController.GetMyMemorizations)
		api.PATCH("/my/memorizations/:word_id", memorizationController.UpdateMyMemorization)

		api.GET("/learning/word", learningController.GetWord)

		api.GET("/images/:search_text", imageController.Get)
	}

	return router.Run(Settings.ListenAddress)
}
