package routes

import (
	_ "fmt"
	"gin_admin/app/index"
	socket_ "gin_admin/app/socket"
	"gin_admin/app/user"
	"gin_admin/middleware/jwt"
	"gin_admin/pkg/settings"
	"gin_admin/routes/api"
	"github.com/gin-gonic/gin"
	_ "github.com/gin-gonic/gin"
	"net/http"
	_ "net/http"
)

//初始化gin default得到实例,加载路由等
func InitRouter() *gin.Engine {
	// http.Handle("/static/",http.StripPrefix("/static/",http.FileServer(http.Dir("static"))))

	//r := gin.Default()

	r:=gin.New()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())

	gin.SetMode(settings.RunMode)



	//处理templates目录里多个子目录
	r.LoadHTMLGlob("app/templates/**/*")
	//r.LoadHTMLFiles("templates/**/*")
	r.Static("/static", "app/static")


	r.GET("/auth",api.GetAuth)


	//index模块分组
	i1 := r.Group("")
	{
		i1.GET("/", index.Index_Api)
	}

	//user模块分组
	u1 := r.Group("/user")
	u1.Use(jwt.JWT())
	{
		u1.GET("/login", func(c *gin.Context) {
			c.HTML(http.StatusOK, "index/login.html",
				gin.H{})
		})
		u1.POST("/login", user.Login_Api)

		u1.GET("/list", func(c *gin.Context) {
			c.HTML(http.StatusOK, "user/list.html",
				gin.H{
					"title": "user list",
				})
		})
		u1.GET("/info", user.User_Info_Api)
	}


	//socket分组
	s1:=r.Group("/socket")
	{
		s1.GET("/blockchain",socket_.Ping)
		s1.GET("/front",socket_.Front)
	}

	return r
}
