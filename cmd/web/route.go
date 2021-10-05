package main

import "github.com/gin-gonic/gin"

func (app *application) routes() *gin.Engine{
	router := gin.Default()

	// Signup page
	router.POST("/signup", app.signUp)

	//Login page
	router.POST("/login", app.login)


	return router
}
