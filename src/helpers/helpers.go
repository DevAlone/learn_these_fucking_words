package helpers

import (
	"github.com/appleboy/gin-jwt"
	"github.com/gin-gonic/gin"
	"strconv"
	"strings"
)

func JWTGetCurrentUser(context *gin.Context) (uint64, string) {
	claims := jwt.ExtractClaims(context)

	result := strings.Split(claims["id"].(string), ",")
	userId, err := strconv.ParseUint(result[0], 10, 64)

	if err != nil {
		panic(err)
	}

	return userId, result[1]
}

func IsStringInSlice(s string, list []string) bool {
	for _, item := range list {
		if item == s {
			return true
		}
	}
	return false
}
