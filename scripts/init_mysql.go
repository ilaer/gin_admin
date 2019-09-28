package scripts

import (
	"database/sql"
	_"github.com/go-sql-driver/mysql"
	"log"
)



func InitMysql()(success bool,db *sql.DB){
	const DB_Driver="root:root@tcp(0.0.0.0:3310)/gin_admin"
	var isOpen bool
	db,err:=sql.Open("mysql",DB_Driver)
	if err!=nil{
		isOpen=false
		log.Panicln("err:",err.Error())
	}else{
		isOpen=true
	}

	//db.SetMaxOpenConns(10)
	//db.SetMaxIdleConns(10)
	return isOpen,db

}
