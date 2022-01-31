package middleware

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"registerandlog/common"
	"registerandlog/model"
	"strings"
)

func AuthMiddleWare() gin.HandlerFunc{
	return func(ctx *gin.Context){
		tokenString:=ctx.GetHeader("Authorization")
		if tokenString==""||!strings.HasPrefix(tokenString,"Bearer"){
			ctx.JSON(http.StatusUnauthorized,gin.H{"code":401,"msg":"权限不足"})
			ctx.Abort()
			return
		}
		tokenString=tokenString[7:]
		token,claims,err:=common.ParseToken(tokenString)
		if err != nil || !token.Valid{
			ctx.JSON(http.StatusUnauthorized,gin.H{"code":401,"msg":"权限不足"})
			ctx.Abort()
			return
		}
		userid:=claims.Userid
		DB:=common.GetDB()
		var user model.User
		DB.First(&user,userid)

		if user.ID==0{
			ctx.JSON(http.StatusUnauthorized,gin.H{"code":401,"msg":"权限不足"})
			ctx.Abort()
			return
		}
		ctx.Set("user",user)
		ctx.Next()
	}
}
