package usercontrol

import (
	"mime/multipart"
	"net/http"
	"path/filepath"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"golang.org/x/text/message"
	"gorm.io/gorm"

	"github.com/Kesha005/crud_with_file/pkg/models"
)

var DBASE *gorm.DB

type Users struct {
	Name     string
	Surname  string
	PassFile string
}

type UserAdd struct {
	Name     string `json:"name"`
	Surname  string `json:"surname"`
	PassFile multipart.Form
}

type UserUpdate struct {
}

func getUsers(ctx *gin.Context) {
	var users []models.Users
	if res := DBASE.Find(&users); res.Error != nil {
		ctx.JSON(404, gin.H{
			"message": "Not found",
		})
		return
	}
	ctx.JSON(200, gin.H{
		users,
	})

}

func addUser(ctx *gin.Context) {
	user := UserAdd{}
	if res := ctx.ShouldBindJSON(&user); res.Error !=nil{
		ctx.JSON(400,gin.H{
			"message":res.Error,
		})
		return 
	}

	var newUser models.Users
	file, err := ctx.FormFile("file")

	// The file cannot be received.
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"message": "No file is received",
		})
		return
	}

	// Retrieve file information
	extension := filepath.Ext(file.Filename)
	// Generate random file name for the new uploaded file so it doesn't override the old file with same name
	newFileName := uuid.New().String() + extension

	// The file is received, so let's save it
	if err := ctx.SaveUploadedFile(file, "/some/path/on/server/"+newFileName); err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"message": "Unable to save the file",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Your file has been successfully uploaded.",
	})
	newUser.Name = user.Name
	newUser.Surname = user.Surname
	newUser.PassFile = newFileName
	if res := DBASE.Create(&newUser); res.Error !=nil {
		ctx.JSON(400,gin.H{
			"message":res.Error,
		})
		return 
	}
	
	ctx.JSON(200,gin.H{
		"message":"User created successfully",
	})

}

func getUser(ctx *gin.Context) {
	id := ctx.Param("id")
	var user models.Users
	if res := DBASE.First(&user, id); res.Error != nil {
		ctx.JSON(404, gin.H{
			"message": res.Error,
		})
		return
	}
	ctx.JSON(200, user)
}

func updateUser(ctx *gin.Context) {

}

func deleteUser(ctx *gin.Context) {

}

func uploadSingleFile(c *gin.Context) {
	
}
