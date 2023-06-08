package db


import(
	"gorm.io/gorm"
	"gorm.io/driver/mysql"
)


func ConnectToDb(url string)*gorm.DB{
	database,err := gorm.Open(mysql.Open(url),&gorm.Config{})
	if err!=nil{
		panic(err)
	}
	return database
}