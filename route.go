package main

import (
	"github.com/gin-gonic/gin"
	"registerandlog/UserController"
	"registerandlog/middleware"
)

func CollectRoute(r *gin.Engine) *gin.Engine{
	r.POST("/api/auth/register",UserController.Register)
	r.POST("/api/auth/login",UserController.Login)
	r.GET("/api/auth/info",middleware.AuthMiddleWare(),UserController.Info)
	return r
}
