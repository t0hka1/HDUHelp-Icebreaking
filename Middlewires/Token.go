package Middlewires
//放在中间件下的伪中间件package

import (
	"github.com/dgrijalva/jwt-go"
	"time"
)
var token string
type MyClaims struct {
	StudentId string
	jwt.StandardClaims
}
type UserInfo struct {
	StudentId string
	Password string
}

//GenToken生成JWT TOKEN

func GenToken(StudentId string) (s string)  {
	mySigningKey := []byte("t0hka")
	d := MyClaims{
		StudentId: StudentId,
		StandardClaims:jwt.StandardClaims{
			NotBefore: time.Now().Unix() - 60,      //生效时间
			ExpiresAt: time.Now().Unix() + 60*60*2,  //过期时间
			Issuer: "admin",                        //签发人
		},
	}
	t := jwt.NewWithClaims(jwt.SigningMethodHS256,d)
	s, _= t.SignedString(mySigningKey)          //对token进行加密
	token=s
	return s
}


