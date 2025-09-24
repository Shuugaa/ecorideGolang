package models

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Csrf_token (id string) string {
	IdBack := id
	return IdBack
}

func Welcome (c *gin.Context) {

	c.HTML(http.StatusOK, "home.html", gin.H{
		"title": "EcoRide",
	})
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