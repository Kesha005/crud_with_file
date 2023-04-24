package db

import (
	"log"

	"github.com/go-sql-driver/mysql"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"

	"github.com/Kesha005/crud_with_file/pkg/models"
)

func ConnectDb(url string) *gorm.DB {
	db, err := gorm.Open(mysql.Open(url), *gorm.Config{})
	if err != nil {
		log.Fatal(err)
		return
	}
	db.AutoMigrate(models.Users{})
	return db
}
