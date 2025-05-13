package routers

import (
	"github.com/gin-gonic/gin"
	"senqi-kvm/api"
)

// 虚拟机相关的 API

// RegisterVmRoutes 注册虚拟机相关路由
func RegisterVmRoutes(r *gin.RouterGroup) {
	vmGroup := r.Group("/vm")
	{
		vmGroup.POST("/create", api.CreateVM)
		//vmGroup.POST("/update", api.UpdateKvm)
		//vmGroup.POST("/remove", api.RemoveKvm)
		//vmGroup.POST("/reinstall", api.ReinstallKvm)
		//vmGroup.POST("/status", api.GetStatusKvm)
		//vmGroup.POST("/boot_order", api.BootOrderKvm)
	}
}
