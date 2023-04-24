package main

import (
	"github.com/gin-gonic/gin"
	"github.com/Kesha005/crud_with_file/pkg/db"
	"github.com/Kesha005/crud_with_file/pkg/controller"
)

func main() {
	url := "root:s@tcp(127.0.0.1:3306)/forgolang"
	db.ConnectDb(url)

	app := gin.Default()
	app.GET("/",usercontrol.getUsers)
	app.Run(":80")

}
