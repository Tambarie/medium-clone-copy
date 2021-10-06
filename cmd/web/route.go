package main

import "github.com/gin-gonic/gin"

func (app *application) routes() *gin.Engine{
	router := gin.Default()
	router.LoadHTMLGlob("ui/html/**/*")

	//homepage
	router.GET("/",app.homePage)

	// gets the signup page
	router.GET("/signup", app.signup)

	// processes the signup page
	router.POST("/signUp", app.signUp)

	// Gets the Login page
	router.GET("/login", app.login)

	//Logs in the user
	router.POST("/loginUser", app.loginUser)

	//Testing cookies
	router.GET("/loginUser", app.setCookies)
	return router
}
