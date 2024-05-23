package db

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)




func ConnectDb()(*gorm.DB,error){
	db, err := gorm.Open(mysql.Open("there must be url"), &gorm.Config{})
	if err!=nil{
		panic(err)
		return nil, err
	}
	//there must be automigrate 
	return db,nil
}