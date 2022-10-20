package main

import (
	"ginapp/controller"
	"ginapp/tool"
	"github.com/gin-gonic/gin"
)

func main() {
	cfg, err := tool.ParseConfig("./conf/app.json")
	if err != nil {
		panic(err)
	}
	_, err = tool.OrmEngine(cfg)

	if err != nil {
		panic(err)
	}
	app := gin.Default()
	registerRouter(app)
	app.Run(cfg.AppHost + ":" + cfg.AppPort)
}

//路由设置
func registerRouter(router *gin.Engine) {
	new(controller.HelloController).Router(router)
	new(controller.MemberController).Router(router)
}
