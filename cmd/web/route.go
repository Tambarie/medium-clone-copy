package main

import (
	"github.com/Tambarie/medium-clone/pkg/middlewares"
	"github.com/gin-gonic/gin"
)

func (app *application) routes() *gin.Engine{
	router := gin.Default()
	router.LoadHTMLGlob("ui/html/**/*")

	//homepage



	router.GET("/",middlewares.CheckNotLogin(),app.homePage)

	// gets the signup page
	router.GET("/signup",middlewares.CheckNotLogin() ,app.signup)

	// processes the signup page
	router.POST("/signUp",middlewares.CheckNotLogin() ,app.signUp)

	// Gets the Login page
	router.GET("/login",middlewares.CheckNotLogin(), app.login)

	//Logs in the user
	router.POST("/loginUser", middlewares.CheckNotLogin(),app.loginUser)

	//Testing cookies'
	//router.GET("/cookies", app.setCookies)


	// Blog Pages

	blogRoutes := router.Group("/blog")
	{
		blogRoutes.GET("/",app.blogPage)
		blogRoutes.GET("/form", app.getFormPage)
	}

	return router
}
