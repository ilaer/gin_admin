/*
@Time : 2019/9/27 下午3:30
@Author : ila
@File : front
@Software: GoLand
*/
package socket

import (
	"github.com/gin-gonic/gin"
	_ "github.com/gin-gonic/gin"
	"net/http"
	_ "net/http"
)

func Front(c *gin.Context) {
	c.HTML(http.StatusOK, "front.html", gin.H{})
}
