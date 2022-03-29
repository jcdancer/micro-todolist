package handlers

import (
	"api-gateway/pkg/utils"
	"api-gateway/services"
	"context"
	"github.com/gin-gonic/gin"
	"strconv"
)

func GetTaskList(ginCtx *gin.Context) {
	var taskReq services.TaskRequest
	PanicIfTaskError(ginCtx.Bind(&taskReq))
	taskService := ginCtx.Keys["taskService"].(services.TaskService)
	claim, _ := utils.ParseToken(ginCtx.GetHeader("Authorization")) // 拿到当前访问用户的Id，才能拿到自己的备忘录信息
	taskReq.Uid = uint64(claim.Id)
	// 调用服务端的函数
	taskResp, err := taskService.GetTasksList(context.Background(), &taskReq)
	if err != nil {
		PanicIfTaskError(err)
	}
	ginCtx.JSON(200, gin.H{
		"data": gin.H{
			"task":  taskResp.TaskList,
			"count": taskResp.Count,
		},
	})
}

func CreateTask(ginCtx *gin.Context) {
	var taskReq services.TaskRequest
	PanicIfTaskError(ginCtx.Bind(&taskReq))
	// 从ginCtx.Keys中取出服务实例
	claim, _ := utils.ParseToken(ginCtx.GetHeader("Authorization"))
	taskReq.Uid = uint64(claim.Id)
	taskService := ginCtx.Keys["taskService"].(services.TaskService)
	taskResp, err := taskService.CreateTask(context.Background(), &taskReq)
	PanicIfTaskError(err)
	ginCtx.JSON(200, gin.H{
		"data": taskResp.TaskDetail,
	})
}

func UpdateTask(ginCtx *gin.Context) {
	var taskReq services.TaskRequest
	PanicIfTaskError(ginCtx.Bind(&taskReq))
	// 从ginCtx.Keys中取出服务实例
	claim, _ := utils.ParseToken(ginCtx.GetHeader("Authorization"))
	taskReq.Uid = uint64(claim.Id)
	id, _ := strconv.Atoi(ginCtx.Param("id")) // 获取task_id
	taskReq.Id = uint64(id)
	taskService := ginCtx.Keys["taskService"].(services.TaskService)
	taskResp, err := taskService.UpdateTask(context.Background(), &taskReq)
	PanicIfTaskError(err)
	ginCtx.JSON(200, gin.H{
		"data": taskResp.TaskDetail,
	})
}

func DeleteTask(ginCtx *gin.Context) {
	var taskReq services.TaskRequest
	PanicIfTaskError(ginCtx.Bind(&taskReq))
	// 从ginCtx.Keys中取出服务实例
	claim, _ := utils.ParseToken(ginCtx.GetHeader("Authorization"))
	taskReq.Uid = uint64(claim.Id)
	id, _ := strconv.Atoi(ginCtx.Param("id")) // 获取task_id
	taskReq.Id = uint64(id)
	taskService := ginCtx.Keys["taskService"].(services.TaskService)
	taskResp, err := taskService.DeleteTask(context.Background(), &taskReq)
	PanicIfTaskError(err)
	ginCtx.JSON(200, gin.H{
		"data": taskResp.TaskDetail,
	})
}

func GetTaskDetail(ginCtx *gin.Context) {
	var taskReq services.TaskRequest
	PanicIfTaskError(ginCtx.Bind(&taskReq))
	// 从ginCtx.Keys中取出服务实例
	claim, _ := utils.ParseToken(ginCtx.GetHeader("Authorization"))
	taskReq.Uid = uint64(claim.Id)
	id, _ := strconv.Atoi(ginCtx.Param("id")) // 获取task_id
	taskReq.Id = uint64(id)
	taskService := ginCtx.Keys["taskService"].(services.TaskService)
	taskResp, err := taskService.GetTask(context.Background(), &taskReq)
	PanicIfTaskError(err)
	ginCtx.JSON(200, gin.H{
		"data": taskResp.TaskDetail,
	})
}
