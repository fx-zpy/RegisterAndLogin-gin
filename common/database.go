package common

import (
	"fmt"
	"github.com/jinzhu/gorm"
	"github.com/spf13/viper"
	"registerandlog/model"
)
var DB *gorm.DB
func InitDb() *gorm.DB{
	driveName:=viper.GetString("datasource.driveName")
	host:=viper.GetString("datasource.host")
	port:=viper.GetString("datasource.port")
	database:=viper.GetString("datasource.database")
	username:=viper.GetString("datasource.username")
	password:=viper.GetString("datasource.password")
	chasret:=viper.GetString("datasource.charset")
	args:=fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=%s&parseTime=true",
		username,
		password,
		host,
		port,
		database,
		chasret)
	db,err:=gorm.Open(driveName,args)
	if err!=nil{
		panic("failed to connect database,err:"+err.Error())
	}
	db.AutoMigrate(&model.User{})
	DB=db
	return db
}
func GetDB() *gorm.DB{
	return DB
}