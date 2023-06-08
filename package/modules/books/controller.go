package books


import("gorm.io/gorm")

type bookhandler struct{
	DB *gorm.DB
}