package main

import (
	"github.com/gin-gonic/gin"
	"howego/app/part_x/controller"
	"howego/config/database"
	"howego/config/log"
	"howego/config/routers"
)

func main() {
	//加载log
	//log.Init(log.NoWrite, "./howego.log")

	gin.SetMode(gin.DebugMode)
	//加载数据库配置
	if err := database.Init(); err != nil {
		log.E("init", "database init error")
		return
	}
	defer database.DB.Close()

	//加载多个app的路由配置
	routers.Include(controller.UserRouters)
	//初始化路由
	r := routers.Init()
	if err := r.Run(":8088"); err != nil {
		log.E("init", "start service failed")
	}
	log.I("init", "web start success")
}
