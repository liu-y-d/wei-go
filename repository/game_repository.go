// Package repository @Author yd 2024/3/28 14:19
package repository

import (
	"wei/common"
	"wei/model"
	"wei/vo"
)

type IGameRepository interface {

	// GetCurrentGameLevel 获取当前用户游戏关卡
	GetCurrentGameLevel(user *model.User) uint
	// GameOver 保存游戏结束记录
	GameOver(user *model.User, request *vo.GameOverRequest) bool
	// GetGamePropsConfig 获取游戏道具配置
	GetGamePropsConfig() []model.GamePropsConfig
	// GetUserAllPropsGuide 获取用户所有道具引导配置
	GetUserAllPropsGuide(user *model.User) []model.GamePropsGuide
	// SaveUserPropsGuide 保存用户道具引导配置
	SaveUserPropsGuide(user *model.User, propsId uint, show uint) bool
}

type GameRepository struct {
}

func (g GameRepository) GetUserAllPropsGuide(user *model.User) []model.GamePropsGuide {
	var models []model.GamePropsGuide
	err := common.DB.Where("userid = ?", user.Model.ID).Find(&models).Error
	if err != nil {
		return nil
	} else {
		return models
	}
}

func (g GameRepository) SaveUserPropsGuide(user *model.User, propsId uint, show uint) bool {

	record := model.GamePropsGuide{
		Userid:  user.Model.ID,
		PropsId: propsId,
		Show:    show,
	}
	err := common.DB.Create(&record).Error
	if err != nil {
		return false
	}
	return true
}

func (g GameRepository) GetGamePropsConfig() []model.GamePropsConfig {
	var models []model.GamePropsConfig
	err := common.DB.Where("`show` = 0").Find(&models).Error
	if err != nil {
		return nil
	} else {
		return models
	}

}

func (g GameRepository) GetCurrentGameLevel(user *model.User) uint {
	gameRecord := model.GameRecord{}
	err := common.DB.Model(&gameRecord).Select("game_level").Where("status = ? and userid = ?", 1, user.Model.ID).Order("game_level DESC").Limit(1).Find(&gameRecord).Error
	if err != nil {
		return 0
	}
	return gameRecord.GameLevel
}
func (g GameRepository) GameOver(user *model.User, request *vo.GameOverRequest) bool {
	gameRecord := model.GameRecord{}

	err := common.DB.Model(&gameRecord).Select("game_level").Where("status = 1").Order("game_level DESC").Limit(1).Find(&gameRecord).Error
	level := gameRecord.GameLevel
	if err != nil {
		level = 0
	}

	if request.GameLevel == level+1 {
		record := model.GameRecord{
			Userid:    user.Model.ID,
			GameLevel: request.GameLevel,
			Status:    request.Status,
		}
		err := common.DB.Create(&record).Error
		if err != nil {
			return false
		}
		return true
	}
	return false
}

func NewGameRepository() IGameRepository {
	return GameRepository{}
}
