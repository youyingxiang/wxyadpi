package main

import (
	"wxyapi/conf"
	"wxyapi/server"
)

func main() {
	// 从配置文件读取配置
	conf.Init()

	// 装载路由
	r := server.NewRouter()
	r.Run(":9000")
}
