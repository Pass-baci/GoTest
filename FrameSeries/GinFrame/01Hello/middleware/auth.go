package middleware

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func Auth() gin.HandlerFunc {
	return func(c *gin.Context) {
		cookie, err := c.Request.Cookie("usercookie")
		if err != nil {
			c.Abort()
			c.String(http.StatusUnauthorized, "error")
		} else {
			c.SetCookie(cookie.Name, cookie.Value, 1000, cookie.Path, cookie.Domain, cookie.Secure, cookie.HttpOnly)
			c.Next()
		}
	}
}
