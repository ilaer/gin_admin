package index

import (
	_ "fmt"
	"gin_admin/pkg/e"
	"gin_admin/pkg/logging"
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
	logging.Warn("welcome to index")
	c.JSON(http.StatusOK,data)
}
