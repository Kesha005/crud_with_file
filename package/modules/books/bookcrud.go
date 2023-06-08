package books

import (
	"net/http"

	"github.com/Kesha005/gin-crud/package/models"
	"github.com/gin-gonic/gin"
)

type AddBookBody struct {
	Title       string `json:"title"`
	Author      string `json:"author"`
	Description string `json:"description"`
}

type getBookBody struct {
	Title       string `json:"title"`
	Author      string `json:"author"`
	Description string `json:"description"`
}

type UpdateBookBody struct {
	Title       string `json:"title"`
	Author      string `json:"author"`
	Description string `json:"description"`
}

func (h bookhandler)getBooks(ctx *gin.Context) {
	var books []models.Book
	if result := h.DB.Find(&books);result.Error!=nil{
		ctx.AbortWithError(http.StatusNotFound,result.Error)
		return 
	}
	ctx.JSON(http.StatusOK,books)
}

func (h bookhandler) storeBook(ctx *gin.Context) {
	body := AddBookBody{}
	if err := ctx.BindJSON(&body); err!=nil{
		ctx.AbortWithError(http.StatusBadRequest,err)
		return
	}
	var book models.Book
	book.Title=body.Title
	book.Author=body.Author
	book.Description=body.Description
	if result := h.DB.Create(&book);result.Error!=nil{
		ctx.AbortWithError(http.StatusNotFound,result.Error)
		return 
	}
	ctx.JSON(http.StatusCreated,book)
}

func (h bookhandler) getBook(ctx *gin.Context) {
	id := ctx.Param("id")
	var book models.Book
	if result := h.DB.First(&book,id);result.Error!=nil{
		ctx.AbortWithError(http.StatusNotFound,result.Error)
		return 
	}
	ctx.JSON(http.StatusOK,book)
}

func (h bookhandler) updateBook(ctx *gin.Context) {
	id := ctx.Param("id")
	body := UpdateBookBody{}
	if err := ctx.BindJSON(&body);err!=nil{
		ctx.AbortWithError(http.StatusBadRequest,err)
		return 
	}
	var book models.Book
	if result := h.DB.First(&book,id);result.Error!=nil{
		ctx.AbortWithError(http.StatusNotFound,result.Error)
		return
	}
	book.Title = body.Title
	book.Author=body.Author
	book.Description=body.Description
	h.DB.Save(&body)
	ctx.JSON(http.StatusOK,&book)
}
func (h bookhandler) DeleteBook(ctx *gin.Context) {
	id := ctx.Param("id")
	var book models.Book
	if result := h.DB.First(&book,id);result.Error!=nil{
		ctx.AbortWithError(http.StatusNotFound,result.Error)
		return
	}
	h.DB.Delete(&book)

	ctx.JSON(http.StatusOK,"All ok")
}
