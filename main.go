package main

import (
	"github.com/gin-gonic/gin"
	"github.com/Kesha005/gin-crud/package/db"
)

func main() {
	url := "root:''@tcp(127.0.0.1:3306)/forgolang"
	database:= db.ConnectToDb(url)


}
