package Model

import (
	"math/rand"
	"time"
)

//在HDUHelp这个项目中，主要使用的语句为查询语句和增添语句，其中查询语句用于返回给用户作于判断的选项
//对于查询语句而言，要求返回一个与学生图片所匹配的一个username和三个不匹配的username
//初步考虑先随机获取一张图片路径，再返回该图片路径所在的一条记录中的username，将其放置于四个选项中的随机一个
//其他三个选项将随机获取username，并置于空位

var correctName string

func Fetch()(picPath string,name[5]string){
	var count int
	DB.Model(&User{}).Count(&count)
	a:=[100]int{0}
	rand.Seed(time.Now().UnixNano())
	randId:=rand.Intn(count)+1  //选择生成图片的ID

	var randRecord User
	DB.Where("id=?",randId).Find(&randRecord)   //选出生成图片的该记录

	picPath=randRecord.PicPath
	correctName="\""+randRecord.Name+"\""

	time.Sleep(time.Microsecond)
	rand.Seed(time.Now().UnixNano())
	correctPosition:=rand.Intn(4)+1


	for i:=1;i<=4;i++{
		var outputRecord User
		if i==correctPosition{
			DB.Where("id=?",randId).Find(&outputRecord)
			a[randId]=1
			name[i]=outputRecord.Name
			continue
		}else {
			time.Sleep(time.Microsecond)
			rand.Seed(time.Now().UnixNano())
			randOutPutId:=rand.Intn(count)+1
			for {
				if a[randOutPutId]==1{
					time.Sleep(time.Microsecond)
					rand.Seed(time.Now().UnixNano())
					randOutPutId=rand.Intn(count)+1
				}else {
					break
				}
			}
			DB.Where("id=?",randOutPutId).Find(&outputRecord)
			a[randOutPutId]=1
			name[i]=outputRecord.Name
		}
	}
	return picPath, name
}
