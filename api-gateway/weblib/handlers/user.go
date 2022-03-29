package handlers

import (
	"api-gateway/pkg/utils"
	"api-gateway/services"
	"context"
	"github.com/gin-gonic/gin"
	"net/http"
)

// UserRegister
// @Description: 用户注册
// @param ginCtx
func UserRegister(ginCtx *gin.Context) {
	var userReq services.UserRequest
	PanicIfUserError(ginCtx.Bind(&userReq))
	// 从gin.Key中取出服务实例
	userService := ginCtx.Keys["userService"].(services.UserService)
	userResp, err := userService.UserRegister(context.Background(), &userReq)
	PanicIfUserError(err)
	ginCtx.JSON(http.StatusOK, gin.H{
		"data": userResp,
	})
}

// UserLogin
// @Description: 用户登录
// @param ginCtx
func UserLogin(ginCtx *gin.Context) {
	var userReq services.UserRequest
	PanicIfUserError(ginCtx.Bind(&userReq))
	// 从gin.Key中取出服务实例
	userService := ginCtx.Keys["userService"].(services.UserService)
	userResp, err := userService.UserLogin(context.Background(), &userReq)
	PanicIfUserError(err)
	token, err := utils.GenerateToken(uint(userResp.UserDetail.ID))
	ginCtx.JSON(http.StatusOK, gin.H{
		"code": userResp.Code,
		"msg":  "success",
		"data": gin.H{
			"user":  userResp.UserDetail,
			"token": token,
		},
	})
}
