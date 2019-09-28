/*
@Time : 2019/9/26 下午5:39
@Author : ila
@File : jwt
@Software: GoLand
*/
package util
import (
	"../settings"
	"github.com/dgrijalva/jwt-go"
	"time"
)

var jwtSecret=[]byte(settings.JwtSecret)
type Claims struct{
	Username string `json:"username"`
	Password string `json:"password"`
	jwt.StandardClaims
}

//根据用户名和密码生成token
func GenerateToken(email,password string)(string,error){
	nowTime:=time.Now()
	//有效时间7天
	expireTime:=nowTime.Add(24*7*time.Hour)

	claims:=Claims{
		email,//使用email生成token
		password,
		jwt.StandardClaims{
			ExpiresAt:expireTime.Unix(),
			Issuer:"gin-admin",
		},
	}
	tokenClaims:=jwt.NewWithClaims(jwt.SigningMethodHS256,claims)
	token,err:=tokenClaims.SignedString(jwtSecret)

	return token,err
}

//解析token,返回认证对比结果
func ParseToken(token string)(*Claims,error){
	tokenClaims,err:=jwt.ParseWithClaims(token,&Claims{},func(token *jwt.Token)(interface{},error){
		return jwtSecret,nil
	})
	if tokenClaims!=nil{
		if claims,ok:=tokenClaims.Claims.(*Claims);ok&&tokenClaims.Valid{
			return claims,nil
		}
	}
	return nil,err
}

















