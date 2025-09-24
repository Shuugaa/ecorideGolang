package models

import (
	"ecoride/auth"
	"ecoride/database"
	"ecoride/userstructs"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func LoginUser(c *gin.Context) {
	var user userstructs.Credentials
	user.Username = c.Request.PostFormValue("name")
	user.Password = c.Request.PostFormValue("password")
	if user.Username == "" {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}
	if result, err := database.LoginUser(user); err != nil {
		c.AbortWithError(result, err)
		return
	}
	buildCookie := auth.SetCookieSignin(user)
	c.SetCookie("session_token", buildCookie.Uuid, 0, "/", "localhost", false, true)
	c.Status(database.StoreSessionWithCookie(buildCookie))
	c.Redirect(http.StatusMovedPermanently, "/users")
}

func CreateUserLogic(c *gin.Context) {
	var newusr userstructs.User
	newusr.Name = c.Request.PostFormValue("name")
	newusr.Password = c.Request.PostFormValue("password")
	newusr.Email = c.Request.PostFormValue("mail")
	checkDup := database.CheckUserExist(newusr)
	if !checkDup {
		database.InsertUser(newusr)
		c.Redirect(http.StatusMovedPermanently, "/users")
		return
	}
	c.AbortWithStatus(http.StatusConflict)
}

func ReadAllUsers(c *gin.Context) {
	users := database.ReadAllUsers()
	c.JSON(http.StatusOK, users)
}

func LogoutUser(c *gin.Context) {
	uuid, err := c.Cookie("session_token")
	if err == http.ErrNoCookie || uuid == "" {
		c.Status(http.StatusInternalServerError)
		return
	}
	sessionToken := database.GetCookieSessionStruct(uuid)
	if sessionToken.Name == "" {
		c.Status(http.StatusBadRequest)
		return
	}
	database.LogoutUser(uuid)
	c.SetCookie("session_token", "", 0, "/", "localhost", false, true)
	c.Redirect(http.StatusMovedPermanently, "/")
}

func RefreshUser(c *gin.Context) {
	cookie, err := c.Cookie("session_token")
	if err == http.ErrNoCookie || cookie == "" {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}
	session := database.GetCookieSessionStruct(cookie)
	session.Expiry = time.Now().Add(5 * time.Minute)
	database.UpdateSession(session)
	c.Status(http.StatusOK)
}

func CheckUuid(c *gin.Context) {
	uuid, err := c.Cookie("session_token")
	if err == http.ErrNoCookie || uuid == "" {
		c.Status(http.StatusBadRequest)
	}
	session, _ := database.CheckUuidExists(uuid)
	c.JSON(http.StatusAccepted, session)
}