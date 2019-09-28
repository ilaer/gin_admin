package settings
import (
	"github.com/go-ini/ini"
	"log"
	"time"
)

var(
	Cfg *ini.File

	RunMode string

	HTTPPort int
	ReadTimeout time.Duration
	WriteTimeout time.Duration

	PageSize int
	JwtSecret string

	TYPE string
	USER string
	PASSWORD string
	HOST string
	NAME string
	TABLE_PREFIX string
)

func init(){
	var err error
	Cfg, err=ini.Load("conf/app.ini")
	if err!=nil{
		log.Fatalf("加载app.ini配置文件失败:%v",err)
	}
	LoadBase()
	LoadServer()
	LoadApp()
	LoadDataBase()
}

func LoadBase(){
	RunMode=Cfg.Section("").Key("RUN_MODE").MustString("debug")
}

func LoadServer(){
	sec,err:=Cfg.GetSection("server")
	if err!=nil{
		log.Fatalf("Fail to get section 'server':%v",err)
	}
	RunMode=Cfg.Section("").Key("RUN_MODE").MustString("debug")

	HTTPPort=sec.Key("HTTP_PORT").MustInt(8765)
	ReadTimeout=time.Duration(sec.Key("READ_TIMEOUT").MustInt(60))*time.Second
	WriteTimeout=time.Duration(sec.Key("WRITE_TIMEOUT").MustInt(60))*time.Second

}

func LoadApp(){
	sec,err:=Cfg.GetSection("app")
	if err!=nil{
		log.Fatalf("fail to get section 'app':%v",err)
	}
	JwtSecret=sec.Key("JWT_SECRET").MustString("!@)*#)!@U#@*!@!)")
	PageSize=sec.Key("PAGE_SIZE").MustInt(10)
}



func LoadDataBase(){
	sec,err:=Cfg.GetSection("database")
	if err!=nil{
		log.Fatalf("fail to get section 'database':%v",err)
	}
	TYPE=sec.Key("TYPE").MustString("mysql")
	USER=sec.Key("USER").MustString("root")
	PASSWORD=sec.Key("PASSWORD").MustString("123")
	HOST=sec.Key("HOST").MustString("0.0.0.0:3306")
	NAME=sec.Key("NAME").MustString("gin_admin")
	TABLE_PREFIX=sec.Key("TABLE_PREFIX").MustString("ga_")
}





















