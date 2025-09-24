package auth

import (
	"ecoride/database"
	"ecoride/userstructs"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func CheckSessionExpired(c *gin.Context) bool {
	sessionToken, err := c.Cookie("session_token")
	if err == http.ErrNoCookie {
		return true
	}
	ses := database.GetCookieSessionStruct(sessionToken)
	if ses.Uuid == "" {
		return true
	}
	return ses.Expiry.Before(time.Now())
}

func SetCookieSignin(cred userstructs.Credentials) userstructs.Session {
	var ses userstructs.Session

	sesToken := uuid.New()

	ses.Name = cred.Username
	ses.Uuid = sesToken.String()
	ses.Expiry = time.Now().Add(5 * time.Minute)

	return ses
}