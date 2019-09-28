package index

import (
	"../../pkg/e"
	_ "fmt"
	"github.com/gin-gonic/gin"
	_ "github.com/gin-gonic/gin"
	_ "log"
	"net/http"
	_ "net/http"
	"time"
	_ "time"
)

func Index_Api(c *gin.Context){
	data:=make(map[string] interface{})
	data["code"]=e.NORMAL_CODE
	data["message"]=e.GetMsg(e.NORMAL_CODE)
	data["data"]=""
	data["time"]=int(time.Now().Unix())//时间戳
	c.JSON(http.StatusOK,data)
}
