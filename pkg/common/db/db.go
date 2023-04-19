package db

import (
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"

	"github.com/Kesha005/crud_with_file/pkg/common/models"

)

func ConnectToDb(url string) *gorm.DB {

	db, err := gorm.Open(mysql.Open(url), *gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}
	db.AutoMigrate(&models.FileModel{})
	return db
}
