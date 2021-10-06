package main

import (
	"github.com/Tambarie/medium-clone/pkg/helpers/bycrypt"
	"github.com/Tambarie/medium-clone/pkg/helpers/emailValidator"
	"github.com/Tambarie/medium-clone/pkg/models"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"net/http"
	"time"
)



//Homepage
func (app *application) homePage(ctx *gin.Context)  {
	ctx.HTML(http.StatusOK,"home.page.gohtml",nil)
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
		CreatedAt: time.Now().Format("2006-January-02"),
		UpdatedAt: time.Now().Format("2006-January-02"),
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
	ctx.String(200, "success")
}

//Login page
func (app *application) login(ctx *gin.Context)  {
	ctx.HTML(http.StatusOK,"login.page.gohtml",nil)
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

	//compare user's password
	ok := bycrypt.CheckPasswordHash(password,hashPassword.Password)
	if ok{
		ctx.Redirect(http.StatusFound,"/")
		return
	}else{
		ctx.String(200,"wrong password")
	}

}

func (app *application) setCookies(ctx *gin.Context)  {
	ctx.SetCookie("email","some-cookies",12,"/setCookies",)
}








