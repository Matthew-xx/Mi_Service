package main

import (
    "Mi_house/ihomeWeb/handler"
    "github.com/julienschmidt/httprouter"
    "github.com/micro/go-micro/registry"
    "github.com/micro/go-micro/util/log"
    "github.com/micro/go-micro/web"
    "github.com/micro/go-plugins/registry/etcdv3"
    "net/http"
)

func main() {
    reg := etcdv3.NewRegistry(func(op *registry.Options){
        op.Addrs = []string{
            "127.0.0.1:2379",
        }
    })

	// create new web service
    service := web.NewService(
            web.Registry(reg),
            web.Name("go.micro.web.ihomeWeb"),
            web.Version("latest"),
            web.Address(":8999"),
    )

	// initialise service
    if err := service.Init(); err != nil {
            log.Fatal(err)
    }

    rou := httprouter.New()
    //映射静态页面
    rou.NotFound = http.FileServer(http.Dir("html"))
    //获取地区请求
    rou.GET("/api/v1.0/areas", handler.GetArea)
    // 获取session请求
    rou.GET("/api/v1.0/session", handler.GetSession)
    // 获取首页轮播请求
    rou.GET("/api/v1.0/house/index", handler.GetIndex)
    // 获取验证码图片
    rou.GET("/api/v1.0/imagecode/:uuid", handler.GetImageCode)
    // 获取短信验证码
    rou.GET("/api/v1.0/smscode/:mobile", handler.GetSmsCode)
    // 提交注册表单
    rou.POST("/api/v1.0/users", handler.PostReg)
    // 提交登录请求
    rou.POST("/api/v1.0/session", handler.PostSession)
    // 退出登录请求
    rou.DELETE("/api/v1.0/session", handler.DeleteSession)
    // 获取用户信息
    rou.GET("/api/v1.0/user", handler.GetUserInfo)
	// 上传用户头像
	rou.POST("/api/v1.0/user/avatar", handler.PostAvatar)
    // 获取用户实名状态
    rou.GET("/api/v1.0/user/auth", handler.GetUserAuth)
    // 发送进行实名认证请求
    rou.POST("/api/v1.0/user/auth", handler.PostUserAuth)
    // 获取用户房源
    rou.GET("/api/v1.0/user/houses", handler.GetUserHouses)
	// register html handler
	//service.Handle("/", http.FileServer(http.Dir("html")))

	// register call handler
	//service.HandleFunc("/ihomeWeb/call", handler.IhomeWebCall)

    // 注册服务
    service.Handle("/", rou)
	// run service
        if err := service.Run(); err != nil {
                log.Fatal(err)
        }
}
