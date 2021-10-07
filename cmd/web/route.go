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

	//router.GET("/router",middlewares.CheckLogin(), app.logout)


	// Blog Pages

	blogRoutes := router.Group("/blog")
	{

		blogRoutes.Use(middlewares.CheckLogin())

		blogRoutes.GET("/logout",app.logout)
		blogRoutes.GET("/",app.blogPage)
		blogRoutes.GET("/myArticles",app.myArticles)
		blogRoutes.GET("readPost",app.readPostPage)
		blogRoutes.GET("/form", app.getFormPage)
		blogRoutes.POST("/postForm",app.postForm)
		blogRoutes.GET("/edit/:id", app.editPostPage)
		blogRoutes.POST("/updatePost/:id",app.updatePost)
		blogRoutes.GET("/delete/:id",app.deletePost)


	}

	return router
}
