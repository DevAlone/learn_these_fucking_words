package main

import (
	"./models"
	"./views"
	"github.com/gin-gonic/gin"
	"github.com/gin-contrib/static"
	"github.com/appleboy/gin-jwt"
	"time"
)

func main() {
	defer models.DB.Close()

	err := models.InitDb()
	if err != nil {
		panic(err)
	}

	//gin.SetMode(gin.ReleaseMode)
	router := gin.Default()
	router.Use(static.Serve("/", static.LocalFile("./frontend/dist", true)))  // static.LocalFile("./frontend/dist/", false)))
	//router.LoadHTMLGlob("templates/*")

	// the jwt middleware
	authMiddleware := &jwt.GinJWTMiddleware{
		Realm:      "test zone",
		Key:        []byte("secret key"),
		Timeout:    time.Hour,
		MaxRefresh: time.Hour,
		Authenticator: func(userId string, password string, c *gin.Context) (string, bool) {
			if (userId == "admin" && password == "admin") || (userId == "test" && password == "test") {
				return userId, true
			}

			return userId, false
		},
		Authorizator: func(userId string, c *gin.Context) bool {
			if userId == "admin" {
				return true
			}

			return false
		},
		Unauthorized: func(c *gin.Context, code int, message string) {
			c.JSON(code, gin.H {
				"code":    code,
				"message": message,
			})
		},
		// TokenLookup is a string in the form of "<source>:<name>" that is used
		// to extract token from the request.
		// Optional. Default value "header:Authorization".
		// Possible values:
		// - "header:<name>"
		// - "query:<name>"
		// - "cookie:<name>"
		TokenLookup: "header:Authorization",
		// TokenLookup: "query:token",
		// TokenLookup: "cookie:token",

		// TokenHeadName is a string in the header. Default value is "Bearer"
		TokenHeadName: "Bearer",

		// TimeFunc provides the current time. You can override it to use another time value. This is useful for
		// testing or if your server uses a different time zone than your tokens.
		TimeFunc: time.Now,
	}

	router.POST("/api/v1/login", authMiddleware.LoginHandler)

	api := router.Group("/api/v1")
	api.Use(authMiddleware.MiddlewareFunc())
	{
		//api.GET("/refresh_token", authMiddleware.RefreshHandler)
		api.GET("/words", views.GetWords)
		api.POST("/words", views.AddWord)
		//
		api.GET("/users", views.GetUsers)
		//
		api.GET("/languages", views.GetLanguages)
		api.GET("/user_memorizations/:username", views.GetUserMemorizations)
	}

	router.Run()
}
