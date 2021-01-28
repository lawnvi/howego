package main

import (
	"fmt"
	"howego/app/part_x"
	"howego/routers"
)

func main()  {
	//加载多个app的路由配置
	routers.Include(part_x.Routers)
	//初始化路由
	r := routers.Init()
	if err := r.Run(); err != nil {
		fmt.Printf("start service failed, err: %v\n", err)
	}
}
