package Middlewires

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

// AuthJWT 验证 JWT 的中间件
func AuthJWT() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		header := ctx.GetHeader("Authorization")
		headerList := strings.Split(header, " ")
		if len(headerList) != 2 {
			ctx.JSON(http.StatusOK,gin.H{
				"err":"Could not parse Authorization!",
				"msg":"failed",
				"data":gin.H{},
				"redirect":"/iceBreaking/login",
			})
			ctx.Abort()
			return
		}
		t := headerList[0]
		content := headerList[1]
		if t != "Bearer" {
			ctx.JSON(http.StatusOK,gin.H{
				"err":"Only Bearer supported!",
				"msg":"failed",
				"data":gin.H{},
				"redirect":"/iceBreaking/login",
			})
			ctx.Abort()
			return
		}
		if content!=token{
			ctx.JSON(http.StatusOK,gin.H{
				"err":"Wrong token!",
				"msg":"failed",
				"data":gin.H{},
				"redirect":"/iceBreaking/login",
			})
			ctx.Abort()
			return
		}
		ctx.Next()
	}
}

