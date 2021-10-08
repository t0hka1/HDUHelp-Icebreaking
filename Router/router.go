package Router

import (
	api "hduhelp/Controller"
	 "hduhelp/Middlewires"

	"github.com/gin-gonic/gin"
)

//注册的路由包括登录、注册、破冰、上传照片...

func InitRouter()  {
	gin.SetMode(gin.ReleaseMode)
	router:=gin.Default()
	router.NoRoute(api.NotFound)
	v1:=router.Group("iceBreaking")
	{
		v1.GET("/register",api.Register)
		v1.POST("/login",api.Login)
	}

	v1.Use(Middlewires.AuthJWT())
	{

		v1.GET("/guess",api.Guess)
		v1.GET("/check",api.Check)
		v1.POST("/upload",api.UploadImage)
	}
	_=router.Run(":2020")
}
