package router

import (
	"ecoride/mid"
	"ecoride/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func ServeRouter() *gin.Engine {
	r := gin.Default()

	r.Use(mid.CorsMiddleware())

	r.LoadHTMLGlob("./assets/html/*")
	home := r.Group("")
	assets := home.Group("/assets")
	assets.StaticFS("css", http.Dir("./assets/css"))
	assets.StaticFS("js", http.Dir("./assets/js"))
	home.GET("", models.Welcome)
	home.GET("home", models.Welcome)
	home.POST("check", models.CheckUuid)
	home.GET("covoitPage", models.CovoitPage)
	home.GET("contact", models.ContactPage)

	user := r.Group("/users")
	user.StaticFS("/assets/css", http.Dir("./assets/css"))
	user.StaticFS("/assets/js", http.Dir("./assets/js"))
	user.POST("/login", models.LoginUser)
	user.GET("/create", models.Create)
	user.POST("/create/new", models.CreateUserLogic)
	user.Use(mid.AuthRequired())
	user.GET("", models.Profile)
	user.POST("", models.RefreshUser)
	user.POST("/logout", models.LogoutUser)
	user.POST("/refresh", models.RefreshUser)

	return r
}