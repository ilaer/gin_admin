package scripts

import (
	"../app/models"
	"fmt"
)
func Select_One(){
	email_:="admin@git_admin.com"
	isopen,db:=models.InitMysql()
	if isopen==true{
		row:=db.QueryRow("select email,password,username from user where email=?",email_)


		if row!=nil{

				var email string
				var user string
				var pwd string
				row.Scan(&email,&user,&pwd)
				fmt.Println("email",email)
				fmt.Println("user",user)
				fmt.Println("pwd",pwd)

		}
	}
	db.Close()
}

