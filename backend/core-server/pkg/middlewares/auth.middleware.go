package middlewares

import (
	"context"
	"log"
	"strings"

	"github.com/gin-gonic/gin"
)

type contextKey string

var (
	BearerSchema = "Bearer "
	UserAuthKey  = contextKey("user-auth")
	UserClaims   = contextKey("user-claim")
)

func Middleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		jwt := ""
		if c.Request.Header.Get("Authorization") != "" {
			jwt = c.Request.Header.Get("Authorization")
		} else {
			log.Println("Login or Register request")
			// handler.ServeHTTP(c.Writer, c.Request)
			c.Next()
		}
		jwt = strings.TrimPrefix(jwt, BearerSchema)
		ctx := context.WithValue(c.Request.Context(), UserAuthKey, jwt)
		c.Request = c.Request.WithContext(ctx)
		// handler.ServeHTTP(c.Writer, c.Request)
		c.Next()
	}
}
