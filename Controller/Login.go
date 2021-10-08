package Controller

import (
	_ "hduhelp/Model"

	"hduhelp/Middlewires"

	"hduhelp/Model"

	"github.com/gin-gonic/gin"

	"net/http"

)

func Login(c *gin.Context)  {
	StudentId:=c.PostForm("StudentId")
	Password:=c.PostForm("Password")
	//对登录传递的参数进行校验，为空则failed
	if StudentId==""||Password==""{
			c.JSON(http.StatusOK,gin.H{
				"err":"Please enter your StudentId and Password!",
				"msg":"failed",
				"data":"",
				"redirect":"/iceBreaking/login",
			})
		return
	}
	//对登录传递的参数进行校验，与数据库中的数据不符则failed
	err := Model.TryLogin(StudentId,Password)
	if err!=nil{
		c.JSON(http.StatusOK,gin.H{
			"err":err,
			"msg":"Wrong StudentId or Password",
			"data":"",
			"redirect":"/iceBreaking/login",
		})
		return
	}
	//对登录传递的参数进行校验，与数据库中的数据相符则success,同时返回token
	tokenString:= Middlewires.GenToken(StudentId)
	c.JSON(http.StatusOK,gin.H{
		"err":"",
		"msg":"success",
		"data":gin.H{"token":tokenString},
		"redirect":"/iceBreaking/guess",
	})
	return
}