package user

import (
	"../../../gin_admin/models"
	_ "fmt"
	"github.com/gin-gonic/gin"
	_ "github.com/gin-gonic/gin"
	_ "log"
	"net/http"
	_ "net/http"
)
func User_Info_Api(c *gin.Context){
	email:=c.Request.FormValue("email")
	data := make(map[string]interface{})
	data["code"]=200
	data["message"]="success"
	pagination:=make(map[string]int)
	pagination["current_page"]=1
	pagination["perpage"]=10
	pagination["pages"]=1
	data["data"],pagination["total"]=models.GetUser(email)
	data["pagination"]=pagination
	//log.Println("=======>",count)
	//log.Println("data[\"data\"] type======>",reflect.TypeOf(data["data"]))

	c.JSON(http.StatusOK,data)
}