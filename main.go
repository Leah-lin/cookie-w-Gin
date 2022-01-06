package main

import (
	"GO-GIN_REST_API/article"
	"GO-GIN_REST_API/middleware"
	"GO-GIN_REST_API/user"
	"fmt"

	"github.com/gin-gonic/gin"
)

func main() {
	fmt.Println("Hello World")

	r := gin.Default()
	r.LoadHTMLGlob("templates/*.html")

	r.Use(middleware.SetUserStatus())

	r.GET("/", article.ShowIndexPage)

	AuthRoutes := r.Group("/user")
	{
		AuthRoutes.GET("/register", middleware.EnsureNotLoggedIn(), user.ShowRegistrationPage)

		AuthRoutes.POST("/register", middleware.EnsureNotLoggedIn(), user.Register)
		AuthRoutes.GET("/login", middleware.EnsureNotLoggedIn(), user.ShowLoginPage)
		AuthRoutes.POST("/login", middleware.EnsureNotLoggedIn(), user.PerformLogin)
		AuthRoutes.GET("/logout", middleware.EnsureLoggedIn(), user.Logout)
	}

	articleRoutes := r.Group("/article")
	{

		articleRoutes.GET("/view/:article_id", article.GetArticle)

		articleRoutes.GET("/create", middleware.EnsureLoggedIn(), article.ShowArticleCreationPage)

		articleRoutes.POST("/create", middleware.EnsureLoggedIn(), article.CreateArticle)
	}

	r.Run(":1106")
}
