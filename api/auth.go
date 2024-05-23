package auth

import (
	"github.com/Kesha005/go_encryptor"
	jwt "github.com/Kesha005/go_encryptor/token"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)



type User struct{
	*gorm.Model
	Name string 
	Username string
	Password string
}

type UserRequest struct{
	Name string 	`json::"name"`
	Username string	`json:"username"`
	Password string	`json:"password"`
}

type LoginBody struct{
	Username string	`json:"username"`
	Password string	`json:"password"`
}

type DbDial struct{
	Db *gorm.DB
}

func (db DbDial)Register(ctx *gin.Context){

	body := UserRequest{}

	if err := ctx.ShouldBind(&body);err!=nil{
		ctx.JSON(401,gin.H{
			"success":false,
			"error":err,
		})
		return 
	}

	var user User

	user.Name= body.Name
	user.Username = body.Username
	user.Password= body.Password


	if saverr:=db.Db.Create(&user).Error; saverr!=nil{
		ctx.JSON(401,gin.H{
			"success":false,
			"error":saverr,
		})
		return 
	}


	token := jwt.UserToken{Id:int(user.ID),Username:user.Username}

	tokenstring, tokerr := token.GenerateToken()
	if tokerr!=nil{
		ctx.JSON(401,gin.H{
			"success":false,
			"error":tokerr,
		})
		return 
	}


	ctx.JSON(200,gin.H{
		"success":true,
		"token":tokenstring,
		"username":user.Username,
	})


}

func (db DbDial)Login(ctx *gin.Context){
	body := LoginBody{}

	if err := ctx.ShouldBind(&body); err!=nil{
		ctx.JSON(401,gin.H{
			"success":false,
			"error":err,
		})
		return 
	}


	var user User


	if getuser := db.Db.Where("username = ?", body.Username).First(&user).Error; getuser!=nil{
		ctx.JSON(401,gin.H{
			"success":false,
			"error":"Username or password is incorrect",
		})
		return
	}

	passhash ,passerr:= go_encryptor.Decrypt(body.Password)
	if passerr!=nil{
		ctx.JSON(401,gin.H{
			"success":false,
			"error":"server error",
		})
		return 
	}
	if passhash !=user.Password{
		ctx.JSON(401,gin.H{
			"success":false,
			"error":"Username or password is incorrect",
		})
		return 
	} 


	token := jwt.UserToken{Id: int(user.ID), Username:user.Username}
	tokenstring , tokerr := token.GenerateToken()
	if tokerr!=nil{
		ctx.JSON(500,gin.H{
			"success":false,
			"error":"Server error",
		})

		return 
	}


	ctx.JSON(200,gin.H{
		"success":true,
		"username":user.Username,
		"token":tokenstring,

	})
	return 
}