package middlewares

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func CheckLogin() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		session, err := ctx.Cookie("session")
		if err != nil{
			ctx.Redirect(http.StatusFound, "/login")
			return
		}
		ctx.Set("userId", session)
		ctx.Next()
	}
}

func CheckNotLogin() gin.HandlerFunc  {
	return func(ctx *gin.Context) {
		_, err := ctx.Cookie("session")
		if err != nil{
			ctx.Next()
			return
		}
		ctx.Redirect(http.StatusFound,"/blog")
		return
	}
}