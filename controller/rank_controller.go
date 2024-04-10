// Package controller @Author yd 2024/3/28 14:43
package controller

import (
	"github.com/gin-gonic/gin"
	"wei/repository"
	"wei/response"
)

type IRankController interface {
	WorldRank(c *gin.Context)
	Test(c *gin.Context)
}
type RankController struct {
	RankRepository repository.IRankRepository
}

func (g RankController) Test(c *gin.Context) {
	//TODO implement me
	response.Success(c, gin.H{"worldRank": 123123}, "获取世界排行榜成功")
}

func (g RankController) WorldRank(c *gin.Context) {
	rank := g.RankRepository.GetWorldRank()
	response.Success(c, gin.H{"worldRank": rank}, "获取世界排行榜成功")
}

func NewRankController() IRankController {
	rankRepository := repository.NewRankRepository()
	controller := RankController{RankRepository: rankRepository}
	return controller
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
