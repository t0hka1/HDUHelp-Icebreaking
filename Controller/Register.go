package Controller

import (
	"net/http"

	"hduhelp/Model"

	"github.com/gin-gonic/gin"
)

func Register(c *gin.Context) {
	StudentId:=c.PostForm("StudentId")
	Name:=c.PostForm("Name")
	Password:=c.PostForm("Password")
	Department:=c.PostForm("Department")
	var exist bool
	exist,err:= Model.CheckUserExist(StudentId)
	if err!=nil {
		c.JSON(http.StatusOK,gin.H{
			"error":err,
			"msg":"failed",
			"data":"",
			"redirect":"/iceBreaking/register",
		})
	}
	if exist{
		c.JSON(http.StatusOK,gin.H{
			"error":"用户已存在",
			"msg":"failed",
			"data":"",
			"redirect":"/iceBreaking/register",
		})
		return
	}
	err= Model.InsertUser(StudentId,Name,Password,Department)
	if err!=nil {
		c.JSON(http.StatusOK,gin.H{
			"error":err,
			"msg":"failed",
			"data":"",
			"redirect":"/iceBreaking/register",
		})
	}
	if exist==false {
		c.JSON(http.StatusOK,gin.H{
			"message":"register success!",
			"msg":"success",
			"data":"",
			"redirect":"/iceBreaking/login",
		})
	}
}
