package models

import (
	"gorm.io/gorm"
)

var DB *gorm.DB

type FileModel struct {
	gorm.Model
	Name     string `json:"name"`
	FilePath string `json:"filepath"`
}
