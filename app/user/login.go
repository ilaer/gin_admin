package user

import (
	"../../models"
	"../../pkg/e"
	_ "fmt"
	"github.com/gin-gonic/gin"
	_ "github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"log"
	_ "log"
	"net/http"
	_ "net/http"
)

//登录接口
func Login_Api(c *gin.Context) {
	email_ := c.Request.FormValue("email")
	password_ := c.DefaultPostForm("password", "gin123")
	log.Println("email=====>" + email_)
	log.Println("password======>" + password_)

	msg_data:= make(map[string]interface{})

	msg_data["data"], msg_data["status"] = models.Email_Login(email_, password_)
	if (msg_data["status"] == 1) {
		msg_data["code"]=e.NORMAL_CODE
		msg_data["message"]=e.GetMsg(e.NORMAL_CODE)
	}else{
		msg_data["code"]=e.PASSWORD_ERROR_CODE
		msg_data["message"]=e.GetMsg(e.PASSWORD_ERROR_CODE)
	}

	c.JSON(http.StatusOK, &msg_data)
}
