/*
@Time : 2019/9/27 下午3:11
@Author : ila
@File : web
@Software: GoLand
*/
package socket
import (
	"github.com/gin-gonic/gin"
	_"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	_"github.com/gorilla/websocket"
	"net/http"
	_"net/http"
)

var upGrader =websocket.Upgrader{
	CheckOrigin:func(r *http.Request) bool{
		return true
	},
}

func Ping(c *gin.Context){
	ws,err:=upGrader.Upgrade(c.Writer,c.Request,nil)
	if err!=nil{
		return
	}
	defer ws.Close()
	for {
		//读取ws中的数据
		mt,message,err:=ws.ReadMessage()
		if err!=nil{
			break
		}
		if string(message)=="ping"{
			message=[]byte("pong")
		}

		//写入ws数据
		err=ws.WriteMessage(mt,message)
		if err!=nil{
			break
		}
	}
}


/*
<script>
var ws = new WebSocket("ws://0.0.0.0:8765/socket/blockchain");
//连接打开时触发
ws.onopen = function(evt) {
    console.log("Connection open ...");
    ws.send("Hello WebSockets!");
};
//接收到消息时触发
ws.onmessage = function(evt) {
    console.log("Received Message: " + evt.data);
};
//连接关闭时触发
ws.onclose = function(evt) {
    console.log("Connection closed.");
};

</script>

*/




















