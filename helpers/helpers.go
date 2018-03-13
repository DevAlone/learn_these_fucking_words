package helpers

import (
	"github.com/appleboy/gin-jwt"
	"github.com/gin-gonic/gin"
	"strings"
	"strconv"
)

func JWTGetCurrentUser(context *gin.Context) (uint64, string)  {
	claims := jwt.ExtractClaims(context)

	result := strings.Split(claims["id"].(string), ",")
	username, err := strconv.ParseUint(result[0], 10, 64)

	if err != nil {
		panic(err)
	}

	return username, result[1]
}
