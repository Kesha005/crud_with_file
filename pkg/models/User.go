package models

import "gorm.io/gorm"





type Users struct {
	gorm.Model
	Name string
	Surname string
	PassFile string
}

