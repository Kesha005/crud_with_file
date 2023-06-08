package main

import (
	"github.com/gin-gonic/gin"
	"github.com/Kesha005/gin-crud/package/db"
	"github.com/Kesha005/gin-crud/package/modules/books"
)

func main() {
	url := "root:''@tcp(127.0.0.1:3306)/forgolang"
	router := gin.Default()
	dbhandler := db.ConnectToDb(url)
	books.RegisterRoutes(router,dbhandler)

	router.Run()



}
