package main

import (
	"github.com/Tambarie/medium-clone/pkg/helpers"
	"github.com/Tambarie/medium-clone/pkg/models"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"net/http"
	"time"
)

func (app *application) signUp(ctx *gin.Context) {
	var user models.User
	err := ctx.BindJSON(&user)
	if err != nil{
		panic(err)
	}

	if user.Password == "" || user.Email == "" || user.FirstName== "" || user.LastName == ""{
		ctx.String(400,"Please fill in the forms")
		return
	}
	userPassword := user.Password
	hasedPass, err := helpers.HashPassword(userPassword)
	if err != nil{
		panic(err)
	}
	user.Password = hasedPass


	isEmail := helpers.IsEmailValid(user.Email)
	if !isEmail {
		ctx.String(http.StatusBadRequest, "Please enter a valid email")
		return
	}

	_, ok := app.user.IfEmailExists(&user, user.Email)

	if ok {
		ctx.String(http.StatusBadRequest, "Email already exists, signup with a new email address")
		return
	}

	user.Id = uuid.New().String()
	user.CreatedAt = time.Now().Format("2006-January-02")
	user.UpdatedAt = time.Now().Format("2006-January-02")


	if err != nil {
		panic(err)
	}
	err = app.user.SignUp(user)
	if err != nil {
		panic(err)
	}
	ctx.String(200, "success")
}

func (app *application) login(ctx *gin.Context)  {
	var user = &models.User{}
	err := ctx.BindJSON(&user)
	if err != nil{
		panic(err)
	}

	if user.Password == "" || user.Email == "" || user.FirstName== "" || user.LastName == ""{
		ctx.String(400,"Please fill in the forms")
		return
	}

	isEmail := helpers.IsEmailValid(user.Email)
	if !isEmail {
		ctx.String(http.StatusBadRequest, "Please enter a valid email")
		return
	}

	_,ok := app.user.IfEmailExists(user,user.Email)
	if !ok{
		ctx.String(http.StatusNotFound,"sorry, you don't have an account ...")
		return
	}
}








