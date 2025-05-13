package routers

import (
	"github.com/gin-gonic/gin"
	"senqi-kvm/middlewares"
)

// InitRouter 初始化路由
func InitRouter() *gin.Engine {
	// 设置运行模式
	gin.SetMode(gin.DebugMode)

	r := gin.New()

	// 加载全局中间件
	middlewares.LoadGlobalMiddlewares(r)

	//r.Use(jwt.JWT())
	router := r.Group("/api/v1")

	// 加载路由
	LoadRouter(router)

	r.GET("/test", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "test",
		})
	})

	return r
}

// LoadRouter 初始化和注册所有路由 API
func LoadRouter(router *gin.RouterGroup) {
	RegisterVmRoutes(router)
}
