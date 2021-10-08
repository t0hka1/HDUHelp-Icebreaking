package Controller

import (
	"github.com/gin-gonic/gin"
	"hduhelp/Model"
	"net/http"
)
//生成破冰游戏的主数据返回给前端

func Guess(c *gin.Context)  {
	picPath,m2:= Model.Fetch()
	c.JSON(http.StatusOK,gin.H{
		"err":"",
		"msg":"success",
		"data": gin.H{
			"picPath": picPath,
			"name1": m2[1],
			"name2": m2[2],
			"name3": m2[3],
			"name4": m2[4],
		},
		"redirect":"",
	})
	return
}

