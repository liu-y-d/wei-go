// Package controller @Author yd 2024/3/28 14:43
package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"wei/common"
	"wei/config"
	"wei/repository"
	"wei/response"
	"wei/util"
	"wei/vo"
)

type IGameController interface {
	GetCurrentUserGameLevel(c *gin.Context)
	GameOver(c *gin.Context)
}
type GameController struct {
	UserRepository repository.IUserRepository
	GameRepository repository.IGameRepository
}

func (g GameController) GameOver(c *gin.Context) {

	var req vo.CommonGameRequest

	// 参数绑定
	if err := c.ShouldBind(&req); err != nil {
		response.Fail(c, nil, err.Error())
		return
	}
	if len(req.Param) == 0 {
		response.Fail(c, nil, "参数不正确")
		return
	}
	var req1 vo.GameOverRequest
	//密码通过RSA解密
	// 密码不为空就解密
	if req.Param != "" {
		decodeData, err := util.RSADecrypt(req.Param, config.Conf.System.RSAPrivateBytes)
		if err != nil {
			response.Fail(c, nil, "参数异常")
			return
		}
		param := string(decodeData)
		util.Json2Struct(param, &req1)
	}
	// 参数校验
	if err := common.Validate.Struct(&req1); err != nil {
		errStr := err.(validator.ValidationErrors)[0].Translate(common.Trans)
		response.Fail(c, nil, errStr)
		return
	}
	user, _ := g.UserRepository.GetCurrentUser(c)
	over := g.GameRepository.GameOver(&user, &req1)
	response.Success(c, gin.H{"status": over}, "存储游戏关卡完成状态")
}

func (g GameController) GetCurrentUserGameLevel(c *gin.Context) {
	user, _ := g.UserRepository.GetCurrentUser(c)
	level := g.GameRepository.GetCurrentGameLevel(&user)
	response.Success(c, gin.H{"gameLevel": level}, "获取当前用户所处游戏关卡")

}

func NewGameController() IGameController {
	gameRepository := repository.NewGameRepository()
	userRepository := repository.NewUserRepository()
	controller := GameController{GameRepository: gameRepository, UserRepository: userRepository}
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
