package main

import (
	"github.com/Tambarie/medium-clone/pkg/helpers/bycrypt"
	"github.com/Tambarie/medium-clone/pkg/helpers/emailValidator"
	"github.com/Tambarie/medium-clone/pkg/models"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"log"
	"net/http"
	"time"
)


////////////////////////////////USERS-CONTROLLERS///////////////////////////////////////
//Homepage
func (app *application) homePage(ctx *gin.Context)  {
	posts, err := app.post.GetAllPosts()
	if err != nil{
		panic(err)
	}
	ctx.HTML(http.StatusOK,"home.page.gohtml",posts)

}

// Read blog post
func (app *application) readPostPage(ctx *gin.Context)  {
	ctx.HTML(200,"blog.readPost.gohtml",nil)
}

// Personalized page
func (app application) myArticles(ctx *gin.Context)  {
	authorId, err := ctx.Cookie("session")
	if err != nil{
		panic(err)
	}

	userPost := &models.Post{}
	posts, err := app.post.GetAllPostOfAUser(userPost,authorId)
	if err != nil{
		panic(err)
	}
	ctx.HTML(http.StatusOK,"blog.myArticles.page.gohtml",posts)
}

func (app *application) blogPage(ctx *gin.Context)  {

	posts, err := app.post.GetAllPosts()

	if err != nil{
		panic(err)
	}
	ctx.HTML(http.StatusOK,"blog.page.gohtml", posts)
}
//Signup page
func (app *application) signup(ctx *gin.Context)  {
	ctx.HTML(http.StatusOK,"signup.page.gohtml",nil)
}



//Processing the signup page
func (app *application) signUp(ctx *gin.Context) {

	lastname :=  ctx.PostForm("lastname")
	firstname := ctx.PostForm("firstname")
	password := ctx.PostForm("password")
	email := ctx.PostForm("email")

	// checking for empty fields
	if lastname == "" || firstname == "" || password== "" || password == ""{
		ctx.String(400,"Please fill in the forms")
		return
	}

	//hash the user's password
	hashedPass, err := bycrypt.HashPassword(password)
	if err != nil{
		panic(err)
	}

	//check if the email is valid
	isEmail := emailValidator.IsEmailValid(email)
	if !isEmail {
		ctx.String(http.StatusBadRequest, "Please enter a valid email")
		return
	}

	//Initializing the user
	user := &models.User{
		Id:        uuid.New().String(),
		FirstName: firstname,
		LastName:  lastname,
		Email:     email,
		Password:  hashedPass,
		CreatedAt: time.Now().Format(time.RFC822Z),
		UpdatedAt: time.Now().Format(time.RFC822Z),
	}

	// checking if email exists
	_, ok := app.user.IfEmailExists(user, email)

	if ok {
		ctx.String(http.StatusBadRequest, "Email already exists, signup with a new email address")
		return
	}

	if err != nil {
		panic(err)
	}
	err = app.user.SignUp(*user)
	if err != nil {
		panic(err)
	}
	ctx.Redirect(302,"/login")
}

//Login page
func (app *application) login(ctx *gin.Context)  {
	ctx.HTML(http.StatusOK,"login.page.gohtml",nil)
}

// handles  logout
func (app *application) logout(ctx *gin.Context)  {
	ctx.SetCookie("session","",-1,"/","localhost",true,true)
	ctx.Redirect(http.StatusFound,"/")
	//ctx.HTML(http.StatusOK,"home.page.gohtml",nil)

}


// logs in the user
func (app *application) loginUser(ctx *gin.Context)  {
	user := &models.User{}

	password := ctx.PostForm("password")
	email := ctx.PostForm("email")

	if password == "" || email == ""{
		ctx.String(400,"Please fill in the forms")
		return
	}
	// Verify if the email entered is valid
	isEmail := emailValidator.IsEmailValid(email)
	if !isEmail {
		ctx.String(http.StatusBadRequest, "Please enter a valid email")
		return
	}

	// Get user's password based on the email inputted
	hashPassword, _ := app.user.IfPasswordExists(user, email)

	// Get user's id
	userId, err := app.user.GetUserId(user,email)
	if err != nil{
		ctx.String(http.StatusSeeOther,"User does not exist")
		ctx.Redirect(http.StatusSeeOther,"/login")
		return
	}
	id := userId.Id

	//Setting the cookies
	ctx.SetCookie("session",id,60*30,"/","localhost",true,true)
	//compare user's password
	ok := bycrypt.CheckPasswordHash(password,hashPassword.Password)
	if ok{
		ctx.Redirect(http.StatusFound,"/blog")
		return
	}

	ctx.String(200,"wrong password")

}

// Gets the form page
func (app *application) getFormPage(ctx *gin.Context)  {
	ctx.HTML(http.StatusOK,"blog.form.gohtml",nil)
}


////////////////////////////////BLOG-CONTROLLERS///////////////////////////////////////

//posts the blog to the database
func (app *application) postForm(ctx *gin.Context)  {

	userId, err := ctx.Cookie("session")
	if err != nil{
		panic(err)
	}

	blogTitle := ctx.PostForm("title")
	blogArticle := ctx.PostForm("article")

	if blogTitle == "" || blogArticle == ""{
		ctx.String(400,"Please fill in the forms")
		return
	}

	post := &models.Post{
		Id:        uuid.New().String(),
		Content:   blogArticle,
		Title:     blogTitle,
		CreatedAt: time.Now().Format("2006-January-02"),
		UpdatedAt: time.Now().Format("2006-January-02"),
		Author: userId,
	}

	err = app.post.CreatePost(post)
	if err != nil{
		panic(err)
	}
	//ctx.Redirect(200,"/blog")
	ctx.Redirect(302,"/blog")
}


// EditPage
func (app *application) editPostPage(ctx *gin.Context)  {
	post := &models.Post{}
	postId := ctx.Param("id")
	sessionId, err := ctx.Cookie("session")
	if err != nil{
		log.Printf("error: %s",err.Error())
	}
	postStruct := app.post.GetAPost(post, postId)
	if sessionId == postStruct.Author{
		postData := app.post.GetAPost(post,postId)
		ctx.HTML(http.StatusOK,"blog.page.edit.gohtml",postData)
		return
	}
	//ctx.String(200,"sorry you can't edit this post")
	ctx.Redirect(302,"/blog")
}

//Edits the blog post
func (app *application) updatePost(ctx *gin.Context)  {
	id := ctx.Param("id")
	title := ctx.PostForm("title")
	content := ctx.PostForm("content")

	var post = &models.Post{
		Id: id,
		Content:   content,
		Title:     title,
	}

	err := app.post.UpdatePost(post)
	if err != nil{
		log.Printf("%v", err.Error())

	}
	ctx.Redirect(302, "/blog")
}


// Deletes the blog post
func (app *application) deletePost(ctx *gin.Context)  {
	post := &models.Post{}
	postId := ctx.Param("id")

	sessionId, err :=  ctx.Cookie("session")
	if err != nil{
		log.Printf("error :%s",err.Error())
	}

	postStruct := app.post.GetAPost(post, postId)
	if sessionId == postStruct.Author{
		err = app.post.DeletePost(postId)
		if err != nil{
			panic(err)

		}
		ctx.Redirect(http.StatusFound,"/blog")
		return

	}
	ctx.Redirect(302,"/blog")
}


////////////////////////////////COMMENTS-CONTROLLERS///////////////////////////////////////

func (app *application) addComment(ctx *gin.Context)  {
	//comment := models.Comments{}
}







