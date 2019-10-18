package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"web-app/models"
	"web-app/user/handler"
	"web-app/user/repository"
	"web-app/user/usecase"
)

func gormConnect() *gorm.DB {

	DBMS := "mysql"
	USER := "root"
	PASS := "test"
	PROTOCOL := "(localhost:3306)"
	DBNAME := "example"
	CONNECT := USER + ":" + PASS + "@" + PROTOCOL + "/" + DBNAME + "?charset=utf8&parseTime=True&loc=Local"
	db, error := gorm.Open(DBMS, CONNECT)

	if error != nil {
		panic(error.Error())
	}

	fmt.Print("db connected; ", &db)
	return db
}

type JsonErrorResponse struct {
	message string
}

func main() {
	db := gormConnect()
	defer db.Close()
	db.LogMode(true)
	db.AutoMigrate(&models.User{})
	rep := repository.NewUserRepository(db)
	use := usecase.NewUserUsecase(rep)
	route := gin.Default()
	handler.NewUserHandler(route, use)
	route.Run(":5000")
}
