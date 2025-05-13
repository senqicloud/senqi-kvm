package middlewares

import "github.com/gin-gonic/gin"

// 全局中间件

// LoadGlobalMiddlewares 加载全局中间件
func LoadGlobalMiddlewares(r *gin.Engine) {

	// Logger 记录每个 HTTP 请求的日志
	r.Use(gin.Logger())

	// Recovery 从任何 panic 中恢复，保证服务器不会因为单个请求崩溃
	r.Use(gin.Recovery())
}
