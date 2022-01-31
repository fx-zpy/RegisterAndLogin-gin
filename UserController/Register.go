package UserController

import (
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"golang.org/x/crypto/bcrypt"
	"log"
	"net/http"
	"registerandlog/common"
	"registerandlog/dto"
	"registerandlog/model"
	"registerandlog/response"
	"registerandlog/util"
)

func Register(ctx *gin.Context){
	db:=common.GetDB()
	name:=ctx.PostForm("name")
	telephone:=ctx.PostForm("telephone")
	password:=ctx.PostForm("password")
	if len(telephone)!=11{
		response.Response(ctx,http.StatusUnprocessableEntity,422,nil,"手机号应该为11位")
		return
	}
	if len(password)<6{
		response.Response(ctx,http.StatusUnprocessableEntity,422,nil,"密码不能少于六位")
		return
	}
	if len(name)==0{
		name=util.RandomName(8)
	}
	if IsTelephoneExist(db,telephone){
		response.Response(ctx,http.StatusUnprocessableEntity,422,nil,"该手机号已经注册过了")
		return
	}
	hashpassword,err:=bcrypt.GenerateFromPassword([]byte(password),bcrypt.DefaultCost)
	if err!=nil{
		response.Response(ctx,http.StatusInternalServerError,500,nil,"加密错误")
		return
	}
	newuser:=model.User{
		Name:      name,
		Telephone: telephone,
		Password:  string(hashpassword),
	}
	db.Create(&newuser)
	log.Println(name,password,telephone)
	response.Success(ctx,nil,"注册成功")
}
func Login(ctx *gin.Context){
	telephone:=ctx.PostForm("telephone")
	password:=ctx.PostForm("password")
	db:=common.GetDB()
	if len(telephone)!=11{
		response.Response(ctx,http.StatusUnprocessableEntity,422,nil,"手机号应该为11位")
		return
	}
	if len(password)<6{
		response.Response(ctx,http.StatusUnprocessableEntity,422,nil,"密码不能少于六位")
		return
	}
	var user model.User
	db.Where("telephone=?",telephone).First(&user)
	if user.ID==0{
		response.Response(ctx,http.StatusUnprocessableEntity,422,nil,"用户不存在，请先注册")
		return
	}
	if err:=bcrypt.CompareHashAndPassword([]byte(user.Password),[]byte(password));err!=nil{
		response.Fail(ctx,nil,"输入的密码不正确")
		return
	}
	token,err:=common.ReleaseToken(user)
	if err!=nil{
		response.Response(ctx,http.StatusInternalServerError,500,nil,"系统异常")
		return
	}
	response.Success(ctx,gin.H{"token":token},"登录成功")
}
func Info(ctx *gin.Context){
	user,_:=ctx.Get("user")
	response.Success(ctx,gin.H{"user":dto.ToUserDto(user.(model.User))},"")
}
func IsTelephoneExist(db *gorm.DB, telephone string) bool {
	var user model.User
	db.Where("telephone=?",telephone).First(&user)
	if user.ID!=0{
		return true
	}
	return false
}
