package middleware

import (
	"github.com/gin-gonic/gin"
)

//BasicAuth middleware authorization
func BasicAuth() gin.HandlerFunc {
	return gin.BasicAuth(gin.Accounts{
		"test": "password",
	})
}
