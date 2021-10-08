package Controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

//设置默认路由当访问一个错误网站时返回

func NotFound(c *gin.Context) {
	c.JSON(http.StatusNotFound, gin.H{
		"error":  "404 ,page not exists!",
		"msg":"failed",
		"data":"",
		"redirect":"",
	})
}