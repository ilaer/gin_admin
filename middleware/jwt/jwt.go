/*
@Time : 2019/9/27 上午8:46
@Author : ila
@File : jwt
@Software: GoLand
*/
package jwt

import (
	"../../pkg/e"
	"../../pkg/util"
	"github.com/gin-gonic/gin"
	_ "github.com/gin-gonic/gin"
	"net/http"
	_ "net/http"
	"strings"
	_ "strings"
	"time"
	_ "time"
)
func JWT() gin.HandlerFunc{
	return func(c *gin.Context){
		var code int
		var data interface{}
		var token string
		code=e.NORMAL_CODE
		//token:=c.Query("token")//从get参数里获取


		//把token修改成使用headers传递
		Authorization:=c.Request.Header.Get("Authorization")
		splits:=strings.Split(Authorization," ")
		if len(splits)<=1{
			token=""
		}else{
			token=splits[1]
		}


		if token ==""{
			code=e.VALIDATE_PARAM_CODE
		}else{
			claims,err:=util.ParseToken(token)
			if err!=nil{
				code=e.NON_AUTHORIZED_CODE
			}else if time.Now().Unix()>claims.ExpiresAt{
				code=e.NON_AUTHORIZED_CODE
			}
		}


		if code!=e.NORMAL_CODE{
			c.JSON(http.StatusUnauthorized,gin.H{
				"code":code,
				"message":e.GetMsg(code),
				"data":data,
			})
			c.Abort()
			return
		}
		c.Next()
	}
}





















