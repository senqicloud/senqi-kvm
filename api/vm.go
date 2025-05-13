package api

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"senqi-kvm/entity/dto"
	"senqi-kvm/internal"
)

// 虚拟机相关 API

// CreateVM 创建虚拟机
func CreateVM(c *gin.Context) {
	var vmCreateRequest dto.VmCreateRequest

	// 解析请求中的 JSON 到结构体
	if err := c.ShouldBindJSON(&vmCreateRequest); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"code": 400, "message": "参数错误", "error": err.Error()})
		return
	}

	// 控制台输出请求信息
	log.Printf("接收到虚拟机创建请求: %+v\n", vmCreateRequest)

	// 调用底层逻辑创建虚拟机
	var vmService internal.VmService = &internal.VmServiceImpl{}

	vm, err := vmService.CreateVm(vmCreateRequest)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":    1,
			"message": "虚拟机创建失败",
			"data":    nil,
		})
	}

	c.JSON(http.StatusOK, gin.H{
		"code":    0,
		"message": "虚拟机创建成功",
		"data": gin.H{
			"vm_id": vm,
		},
	})
}
