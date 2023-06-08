package books

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type bookhandler struct{
	DB *gorm.DB
}

func RegisterRoutes(router *gin.Engine,db *gorm.DB){
	h := &bookhandler{
		DB:db,
	}
	routes:=router.Group("/books")
	routes.POST("/",h.storeBook)
	
}