/*
@Time : 2019/9/27 下午3:01
@Author : ila
@File : web_s
@Software: GoLand
*/
package socket
import (
	"fmt"
	_ "fmt"
	"golang.org/x/net/websocket"
	_ "golang.org/x/net/websocket"
	_ "log"
	_ "net/http"
)

func Web_Echo(ws *websocket.Conn) {
	var err error
	for {
		var reply string
		if err = websocket.Message.Receive(ws, &reply); err != nil {
			fmt.Println("Can`t receive")
			break
		}
		fmt.Println("Received back from client:" + reply)
		msg := "Received: " + reply
		fmt.Println("Sending to client:" + msg)

		if err = websocket.Message.Send(ws, msg); err != nil {
			fmt.Println("Cant`t send")
			break
		}
	}
}

