// Package controller @Author yd 2024/3/28 14:43
package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"golang.org/x/net/context"
	"time"
	"wei/repository"
	"wei/response"
)

type ILeafController interface {
	GetLeaf(c *gin.Context)
	ConsumeLeaf(c *gin.Context)
	InfinityLeaf(c *gin.Context)
}
type LeafController struct {
	LeafRepository repository.ILeafRepository
	UserRepository repository.IUserRepository
}

func (l LeafController) InfinityLeaf(c *gin.Context) {
	user, _ := l.UserRepository.GetCurrentUser(c)
	status := l.LeafRepository.InfinityLeaf(&user)
	response.Success(c, gin.H{"status": status}, "获得无限叶子成功")
}

func (l LeafController) ConsumeLeaf(c *gin.Context) {
	user, _ := l.UserRepository.GetCurrentUser(c)
	status := l.LeafRepository.ConsumeLeaf(&user)
	response.Success(c, gin.H{"status": status}, "消费叶子成功")
}

func (l LeafController) GetLeaf(c *gin.Context) {
	user, _ := l.UserRepository.GetCurrentUser(c)
	leaf := l.LeafRepository.GetLeaf(&user)
	response.Success(c, gin.H{"leaf": leaf}, "获取叶子详情成功")
}

func NewLeafController() ILeafController {
	LeafRepository := repository.NewLeafRepository()
	userRepository := repository.NewUserRepository()
	controller := LeafController{LeafRepository: LeafRepository, UserRepository: userRepository}
	return controller
}

func runScheduledTask(c *gin.Context, f func()) {
	ctx := c.Request.Context()
	// 使用 WithDeadline 或 WithTimeout 从父上下文中派生出一个带超时限制的新上下文
	taskCtx, cancel := context.WithTimeout(ctx, 10*time.Second) // 设置任务的超时时间
	defer cancel()

	// 设置定时任务
	timer := time.NewTimer(5 * time.Second) // 假设任务将在5秒后执行

	// 等待定时器触发或上下文被取消
	select {
	case <-timer.C:
		fmt.Println("Task executed after 5 seconds")
		// 在这里执行你的定时任务
		f()

	case <-taskCtx.Done():
		fmt.Println("Task cancelled due to request timeout or cancellation")
		return
	}
}

//	GetUserInfo(c *gin.Context)          // 获取当前登录用户信息
//	GetUsers(c *gin.Context)             // 获取用户列表
//	ChangePwd(c *gin.Context)            // 更新用户登录密码
//	CreateUser(c *gin.Context)           // 创建用户
//	UpdateUserById(c *gin.Context)       // 更新用户
//	BatchDeleteUserByIds(c *gin.Context) // 批量删除用户
//}
//
//type UserController struct {
//	UserRepository repository.IUserRepository
//}
//
//// 构造函数
//func NewUserController() IUserController {
//	userRepository := repository.NewUserRepository()
//	userController := UserController{UserRepository: userRepository}
//	return userController
//}
