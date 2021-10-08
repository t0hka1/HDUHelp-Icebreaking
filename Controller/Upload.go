package Controller

import (
	"C"
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"hduhelp/Model"
	"io/ioutil"
	"net/http"
	"os"
	"path"
	"strings"
	"time"
)

//需求:接收到他人本地的文件，将其上传到服务器文件夹，并生成相应的pic_path，将pic_path更新到数据库

func UploadImage(c *gin.Context) {
	f, err := c.FormFile("upload")
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": 400,
			"msg":  "Upload failed!",
		})
		return
	} else {
		fileExt := strings.ToLower(path.Ext(f.Filename))
		if fileExt != ".png" && fileExt != ".jpg" && fileExt != ".gif" && fileExt != ".jpeg" {
			c.JSON(http.StatusOK, gin.H{
				"code": 400,
				"msg":  "Upload failed!Only png,jpg,gif,jpeg allowed!",
			})
			return
		}
	}


	header := c.GetHeader("Authorization")
	headerList := strings.Split(header, " ")
	content := headerList[1]
	var keyInfo string
	tokenStr := content
	tokenInfo , _ := jwt.Parse(tokenStr, func(token *jwt.Token) (i interface{}, e error) {
		return keyInfo,nil
	})
	finToken := tokenInfo.Claims.(jwt.MapClaims)
	foldStudent :=finToken["StudentId"]
	NowStudent:=fmt.Sprintf("%s",foldStudent)

	fileDir:=fmt.Sprintf("%s%s","C:\\Users\\czk12\\Desktop\\Upload\\", foldStudent)
	_=os.Mkdir(fileDir,os.ModePerm)

	CoverSpace(fileDir)

	filename := Mymd5(fmt.Sprintf("%s%s", f.Filename, time.Now().String()))+path.Ext(f.Filename)
	filePath := fmt.Sprintf("%s%s%s", fileDir, "\\",filename)

	Model.UploadPicpath(filePath,NowStudent)

	err=c.SaveUploadedFile(f, filePath)
	if err != nil {
		c.JSON(http.StatusOK,gin.H{
			"code":400,
			"msg":"save failed!",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"msg":  "Upload success!",
	})
	return
}


func Mymd5(str string) string  {
	h := md5.New()
	h.Write([]byte(str))
	return hex.EncodeToString(h.Sum(nil))
}

func CoverSpace(fileDir string)  {
	dir, err := ioutil.ReadDir(fileDir)
	//该文件夹下无文件时极速返回
	if err!=nil{
		return
	}
	//将文件夹内文件置空
	for _,fi:=range dir{
		_=os.Remove(fmt.Sprintf("%s%s%s",fileDir,"\\",fi.Name()))
	}
	return
}