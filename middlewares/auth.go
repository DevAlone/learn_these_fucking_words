package middlewares

import . "../models"

import (
	"fmt"
	"github.com/appleboy/gin-jwt"
	"github.com/gin-gonic/gin"
	"github.com/go-pg/pg"
	"golang.org/x/crypto/bcrypt"
	"strings"
	"time"
)

var AuthMiddleware *jwt.GinJWTMiddleware

func init() {
	AuthMiddleware = &jwt.GinJWTMiddleware{
		Realm:      "test zone",
		Key:        []byte("secret key"),
		Timeout:    time.Hour,
		MaxRefresh: time.Hour,
		Authenticator: func(username string, password string, c *gin.Context) (string, bool) {
			var resultString string

			user := User{}

			username = strings.ToLower(username)

			err := DB.Model(&user).
				Where("LOWER(username) = ?", username).
				Select()

			if err == pg.ErrNoRows {
				return resultString, false
			}

			err = bcrypt.CompareHashAndPassword(user.Password, []byte(password))

			if err != nil {
				return resultString, false
			}

			resultString += fmt.Sprintf("%d,%s", user.Id, user.Username)

			return resultString, true
		},
		Authorizator: func(userId string, c *gin.Context) bool {
			return true
			//if userId == "admin" {
			//	return true
			//}
			//
			//return false
		},
		Unauthorized: func(c *gin.Context, code int, message string) {
			c.JSON(code, gin.H{
				"code":          code,
				"error_message": message,
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
}
