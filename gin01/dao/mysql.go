package dao

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"haimait/learn/gin01/models"
)
var (
	DB *gorm.DB
)
func InitMySQL()(err error) {
	DB, err := gorm.Open("mysql", "root:c8e772bb8a4ef9e0@/gin01?charset=utf8&parseTime=True&loc=Local")
	if err!=nil{
		fmt.Println("open mysql err:",err)
		defer DB.Close()
		return
	}
	return DB.DB().Ping()
	// 模型绑定
	DB.AutoMigrate(&models.Todo{})
	return
}
func Close(){
	DB.Close()
}