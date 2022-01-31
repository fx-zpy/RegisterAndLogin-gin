package main

import (
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"github.com/spf13/viper"
	"os"
	"registerandlog/common"
)


func main(){
	InitConfig()
	db:=common.InitDb()
	defer db.Close()
	r:=gin.Default()
	r=CollectRoute(r)
	port:=viper.GetString("server.port")
	if port!=""{
		panic(r.Run(":"+port))
	}
	panic(r.Run())
}
func InitConfig(){
	workdir,_:=os.Getwd()
	viper.SetConfigName("application")
	viper.SetConfigType("yml")
	viper.AddConfigPath(workdir+"/config")
	err:=viper.ReadInConfig()
	if err != nil {
		panic(err)
	}
}


