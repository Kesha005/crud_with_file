package file_crud

import (

	"log"
	"net/http"
	"github.com/Kesha005/crud_with_file/pkg/common/models"
	"github.com/gin-gonic/gin"
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
	ctx.SaveUploadedFile(up_file)

	save_file := models.FileModel{Name: up_file.Filename, FilePath: ""}
	models.DB.create(&save_file)

	ctx.JSON(200, gin.H{
		"message": "File uploaded",
	})

}

func getFile(ctx *gin.Context) {
	
	id := ctx.Param("id")
	var file models.FileModel
	if err := models.DB.First(&file,id); err !=nil{
		ctx.JSON(http.StatusNotFound,gin.H{
			"message":"Not found",
		})
		return 
	}
	ctx.JSON(200,gin.H{
		"file":file,
	})		
}

func updateFile(ctx *gin.Context) {

}

func deleteFile(ctx *gin.Context) {

}
