package models

import (
	"ecoride/mode"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Csrf_token (id string) string {
	IdBack := id
	return IdBack
}

func Welcome (c *gin.Context) {

	cookiie, err := c.Cookie("session_token")
	if (err == nil) {
	}
	if (cookiie != "") {
		c.HTML(http.StatusOK, "homeAuth.html", gin.H{
			"title": "EcoRide",
			"username": mode.GetUsernameFromUuid(cookiie),
			"Csrf_token": Csrf_token(cookiie),
		})
	} else {
		c.HTML(http.StatusOK, "home.html", gin.H{
		"title": "EcoRide",
		})
	}
}

func Create (c *gin.Context) {
	c.HTML(http.StatusOK, "create.html", gin.H{
		"title": "Création..",
	})
}

func Profile (c *gin.Context) {
	c.HTML(http.StatusOK, "profile.html", gin.H{
		"title": "Profile",
	})
}

func CovoitPage (c *gin.Context) {
	c.HTML(http.StatusOK, "covoitPage.html", gin.H{
		"title": "Covoiturages",
	})
}

func ContactPage (c *gin.Context) {
	c.HTML(http.StatusOK, "contact.html", gin.H{
		"title": "Contact",
	})
}