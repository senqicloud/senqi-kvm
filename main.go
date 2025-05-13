package main

import (
	"senqi-kvm/internal"
	"senqi-kvm/routers"
)

// 初始化钩子函数
func init() {
	internal.Setup()
}

// 项目启动方法
func main() {

	// 初始化路由
	router := routers.InitRouter()

	// 启动
	err := router.Run(":8999")
	if err != nil {
		return
	}

}
