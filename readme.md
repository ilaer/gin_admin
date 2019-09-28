## gin_admin
### 使用gin搭建基础http web基础框架,处理websocket请求.
* web使用golang第三方框架gin
+ token认证使用go-jwt
    - token格式:Authorization:token fgfgdfgfgf
+ websocket使用gorilla/websocket
    - ```github.com/gorilla/websocket``` 
   
* 目录结构如下
    + app 默认app目录
        + index 首页模块
            - index.go 页面
        + user 用户模块
            - info.go 当前用户详情
            - login.go 用户登录
        + socket websocket模块
            front.go 前端页面挂起
            web.go websocket视图
    + conf 项目配置文件目录
        - app.ini app配置文件
    + middleware 中间件
        - jwt token认证中间件模块
    + models 数据库表模块
        - models.go gorm初始化
        - user.go 用户信息表模块
    + pkg 官方定义:存放第三方包
        - e 返回数据部分key value定义
        - logging 日志
        - settings 读取配置文件
        - util 通用功能
    + routes 路由逻辑
        - api api路由
    + scripts 测试脚本
    + unit_test 单元测试
    
    main.go 启动文件
    read.me 项目说明
    