package file_crud

import (
	"fmt"
	"log"
	"net/http"
	"github.com/gin-gonic/gin"
	"github.com/Kesha005/crud_with_file/pkg/common/models"
)

type File struct {
	Name     string `json:"name"`
	FilePath string `json:"file_path"`
}

func getFiles(ctx *gin.Context) {
	var files []models.FileModel

	models.DB.Find(&files)

	ctx.JSON(http.StatusOK, gin.H{
		"files": files,
	})
	return
}

func addFile(ctx *gin.Context) {
	var file File

	if err := ctx.ShouldBindJSON(&file); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": err,
		})
		return
	}
	up_file, _ := ctx.FormFile("file")
	log.Println(up_file.Filename)
	ctx.SaveUploadedFile(up_file, ../files)

	save_file := models.FileModel{Name: up_file.Filename, FilePath: ""}
	models.DB.create(&save_file)

	ctx.JSON(200, gin.H{
		"message": "File uploaded",
	})

}

func getFile(ctx *gin.Context) {

}

func updateFile(ctx *gin.Context) {

}

func deleteFile(ctx *gin.Context) {

}
