package main

//import . "./models"

import (
	"./models"
	"./controllers"
	"./middlewares"
	"github.com/gin-gonic/gin"
	"github.com/gin-contrib/static"
)


func main() {
	defer models.DB.Close()

	err := models.InitDb()
	if err != nil {
		panic(err)
	}

	//gin.SetMode(gin.ReleaseMode)
	router := gin.New()
	router.Use(gin.Logger())
	router.Use(gin.Recovery())

	router.Use(static.Serve("/", static.LocalFile("./frontend/dist", true)))  // static.LocalFile("./frontend/dist/", false)))
	//router.LoadHTMLGlob("templates/*")

	// the jwt middleware

	router.POST("/api/v1/login", middlewares.AuthMiddleware.LoginHandler)
	router.POST("/api/v1/register", controllers.Register)

	api := router.Group("/api/v1")
	api.Use(middlewares.AuthMiddleware.MiddlewareFunc())
	{
		api.GET("/refresh_token", middlewares.AuthMiddleware.RefreshHandler)
		api.GET("/words", controllers.GetWords)
		api.POST("/words", controllers.AddWord)
		//
		api.GET("/users", controllers.GetUsers)
		//
		api.GET("/languages", controllers.GetLanguages)
		api.GET("/memorizations", controllers.GetMemorizations)

		api.GET("/my/memorizations", controllers.GetMyMemorizations)
	}

	router.Run()
}
