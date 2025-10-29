package mid

import (
	"ecoride/auth"
	"ecoride/mode"
	"ecoride/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func CorsMiddleware() gin.HandlerFunc {
    return func(c *gin.Context) {
        if c.GetHeader("Access-Control-Allow-Origin") == "" {
            c.Header("Access-Control-Allow-Origin", "http://localhost:3000") // Allow all origins
        }
        if c.GetHeader("Access-Control-Allow-Methods") == "" {
            c.Header("Access-Control-Allow-Methods", "GET, POST, OPTIONS, PUT, DELETE")
        }
        if c.GetHeader("Access-Control-Allow-Credentials") == "" {
            c.Header("Access-Control-Allow-Credentials", "true")
        }
        if c.GetHeader("Access-Control-Allow-Headers") == "" {
            c.Header("Access-Control-Allow-Headers", "Content-Type, Authorization, Content-Length, Accept-Encoding")
        }

        if c.Request.Method == "OPTIONS" {	
            c.AbortWithStatus(http.StatusNoContent)
            return
        }

        c.Next()
    }
}

func AuthRequired() gin.HandlerFunc {
	return func (c *gin.Context) {
		token, err := c.Cookie("session_token")
        if mode.KnownUuid(token) {
            models.RefreshUser(c)
        } else if auth.CheckSessionExpired(c) {
			if mode.KnownUuid(token) {
	            c.SetCookie("session_token", "", 0, "/", "localhost", false, true)
				c.Status(http.StatusUnauthorized)
			}
            if token == "" {
                if err == http.ErrNoCookie {
                    c.AbortWithStatus(http.StatusUnauthorized)
                }
	            c.SetCookie("session_token", "", 0, "/", "localhost", false, true)
            }
            c.Status(http.StatusForbidden)
        }
        c.Next()
	}
}

//unused...
func Forbidden() gin.HandlerFunc {
    return func (c *gin.Context) {
        if c.GetHeader("Referer") != "http://localhost:3000" {
            c.AbortWithStatus(http.StatusForbidden)
        }
        c.Next()
    }
}