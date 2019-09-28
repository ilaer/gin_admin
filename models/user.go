package models

import (
	_ "github.com/jinzhu/gorm"
	"time"
)

//定义User模型结构体
type User struct {
	ID          int        `gorm:"primary_key" json:"id"`
	Username    string     `gorm:"type:char(255)" json:"username"`
	Password    string     `gorm:"type:char(255)" json:"password"`
	Email       string     `gorm:"type:char(255)" json:"email"`
	Active      int        `gorm:"type:tinyint" json:"active"`
	Create_ip   string     `gorm:"type:varchar(50)" json:"create_ip"`
	Login_ip    string     `gorm:"type:varchar(50)"  json:"login_ip"`
	Login_time  int        `gorm:"type:int" json:"login_time"`
	Create_time *time.Time `gorm:"type:datetime" json:"create_time"`
	Update_time *time.Time `gorm:"type:datetime" json:"update_time"`
}

//检查当前email的用户row记录
func CheckAuth(email,password string)bool{
	var user User
	db.Select("id").Where(User{Email:email,Password:password}).First(&user)
	if user.ID>0{
		return true
	}
	return false
}

//generate token时更新login_ip和login_time
func UpdateUser(email,remote_addr string ) bool{
	update_user:=User{
		Login_ip:remote_addr,
		Login_time:int(time.Now().Unix()),
		}

	err:=db.Model(User{}).Where("email=?",email).Update(&update_user)
	if err!=nil{
		return false
	}
	return true

}


//email查询用户信息
func GetUser(email string) (user []User, count int) {
	//User模型的查表方法<====email
	datas := db.Where("email=?", email).First(&user)
	datas = datas.Count(&count) //计算本次查询结果row数量
	//log.Println("counts=====>",count)
	//log.Println("username======>",user[0].Username)
	return
}

//用户登录验证方法
func Email_Login(email, password string)   (user []User,vaild int) {

	datas := db.Where("email=? AND password=?", email,password).First(&user)
	var (
		count int

	)
	datas=datas.Count(&count)

	if(count==1){
		vaild=1
	}else{
		vaild=0
	}

	return

}
