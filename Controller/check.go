package Controller

import (
	"hduhelp/Model"

	"github.com/gin-gonic/gin"

	"net/http"
)

func Check(c *gin.Context)  {
	chooseName:=c.Query("chooseName")
	flag:= Model.CheckName(chooseName)
	c.JSON(http.StatusOK,gin.H{
		"err":"",
		"msg":"success",
		"data":gin.H{
			"result":flag,
		},
		"redirect":"",
	})
}