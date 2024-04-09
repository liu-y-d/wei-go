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
	GameOver(user *model.User, request *vo.GameOverRequest) bool
}

type GameRepository struct {
}

func (g GameRepository) GetCurrentGameLevel(user *model.User) uint {
	gameRecord := model.GameRecord{}
	err := common.DB.Model(&gameRecord).Select("game_level").Where("status = 1").Order("game_level DESC").Limit(1).Find(&gameRecord).Error
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
