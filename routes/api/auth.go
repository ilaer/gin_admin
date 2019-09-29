/*
@Time : 2019/9/27 上午8:57
@Author : ila
@File : auth
@Software: GoLand
*/
package api

import (
	"github.com/astaxie/beego/validation"
	"github.com/gin-gonic/gin"

	"net/http"
	_ "net/http"

	_ "github.com/gin-gonic/gin"

	"gin_admin/models"
	"gin_admin/pkg/e"
	"gin_admin/pkg/logging"
	"gin_admin/pkg/util"
)

type auth struct {
	Email string `valid:"Required;MaxSize(150)"`
	Password string `valid:"Required;MaxSize(50)"`
}

func GetAuth(c *gin.Context) {
	email := c.Query("email")
	password := c.Query("password")
	remote_addr:=c.ClientIP()


	valid := validation.Validation{}
	a := auth{Email: email, Password: password}
	ok, _ := valid.Valid(&a)
	data := make(map[string]interface{})
	code := e.VALIDATE_PARAM_CODE
	if ok {
		isExist := models.CheckAuth(email, password)
		if isExist {
			token, err := util.GenerateToken(email, password)
			if err != nil {
				code = e.NON_AUTHORIZED_CODE
			} else {
				data["token"] = token
				code = e.NORMAL_CODE
				//成功生成token,准备update user表
				models.UpdateUser(email,remote_addr)
			}
		} else {
			code = e.NON_AUTHORIZED_CODE
		}

	} else {
		for _, err := range valid.Errors {
			logging.Warn(err.Key, err.Message)
		}
	}
	c.JSON(http.StatusOK, gin.H{
		"code":    code,
		"message": e.GetMsg(code),
		"data":    data,
	})
}
