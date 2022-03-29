package middleware

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

// InitMiddleware
// @Description: 接收服务实例，并存在gin.Keys中
// @param services
// @return gin.HandlerFunc
func InitMiddleware(service []interface{}) gin.HandlerFunc {
	return func(context *gin.Context) {
		// 将实例存到gin.Keys中
		context.Keys = make(map[string]interface{})
		context.Keys["userService"] = service[0]
		context.Keys["taskService"] = service[1]
		context.Next()
	}
}

// ErrorMiddleware
// @Description: 错误处理中间件
// @return gin.HandlerFunc
func ErrorMiddleware() gin.HandlerFunc {
	return func(context *gin.Context) {
		defer func() {
			if r := recover(); r != nil {
				context.JSON(200, gin.H{
					"code": 404,
					"msg":  fmt.Sprintf("%s", r),
				})
				context.Abort()
			}
		}()
		context.Next()
	}
}
