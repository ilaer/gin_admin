package main

import (
	"fmt"
	"gin_admin/pkg/settings"
	"gin_admin/routes"
	"github.com/fvbock/endless"
	"log"
	//"net/http"
	_ "github.com/fvbock/endless"
	"syscall"

	//"../gin_admin/scripts"
)
func main(){

	endless.DefaultReadTimeOut=settings.ReadTimeout
	endless.DefaultWriteTimeOut=settings.WriteTimeout
	endless.DefaultMaxHeaderBytes=1<<20
	endPoint:=fmt.Sprintf(":%d",settings.HTTPPort)

	server:=endless.NewServer(endPoint,routes.InitRouter())
	server.BeforeBegin=func(add string){
		log.Printf("Actual pid is %d",syscall.Getpid())
	}

	err:=server.ListenAndServe()
	if err!=nil{
		log.Print("Server err: %v",err)
	}

	//scripts.Select_One()

	//r:=routes.InitRouter()
	////_=r.Run("0.0.0.0:8080")
	//s:=&http.Server{
	//	Addr:           fmt.Sprintf(":%d", settings.HTTPPort),
	//	Handler:        r,
	//	ReadTimeout:    settings.ReadTimeout,
	//	WriteTimeout:   settings.WriteTimeout,
	//	MaxHeaderBytes: 1<<20,
	//}
	//s.ListenAndServe()
}




















