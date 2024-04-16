// Package repository @Author yd 2024/3/28 14:19
package repository

import (
	"time"
	"wei/common"
	"wei/model"
)

type ILeafRepository interface {
	GetLeaf(user *model.User) model.Leaf
	ConsumeLeaf(user *model.User) bool

	InfinityLeaf(user *model.User) bool
	CleanInfinityLeaf(user *model.User) bool
}

type LeafRepository struct {
}

func (l LeafRepository) CleanInfinityLeaf(user *model.User) bool {
	var leaf model.Leaf
	err := common.DB.Where("userid", user.Model.ID).First(&leaf).Error
	if err != nil {
		return false
	}
	err = common.DB.Model(&leaf).Where("userid", user.Model.ID).Update("infinity", nil).Error
	if err != nil {
		return false
	}
	return true
}

func (l LeafRepository) InfinityLeaf(user *model.User) bool {
	var leaf model.Leaf
	err := common.DB.Where("userid", user.Model.ID).First(&leaf).Error
	now := time.Now()

	// 计算今晚12点的时间点
	tonight := time.Date(now.Year(), now.Month(), now.Day()+1, 0, 0, 0, 0, now.Location())
	yenight := time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, now.Location())

	if err != nil || (leaf.Infinity != nil && leaf.Infinity.Before(tonight) && leaf.Infinity.After(yenight)) {
		return false
	}
	err = common.DB.Model(&leaf).Where("userid", user.Model.ID).Update("infinity", time.Now()).Error
	if err != nil {
		return false
	}
	return true
}

func (l LeafRepository) ConsumeLeaf(user *model.User) bool {

	var leaf model.Leaf
	err := common.DB.Where("userid", user.Model.ID).First(&leaf).Error
	if err != nil || leaf.Remaining-5 < 0 {
		return false
	}
	err = common.DB.Model(&leaf).Where("userid", user.Model.ID).Update("remaining", leaf.Remaining-5).Error
	if err != nil {
		return false
	}
	return true
}

func (l LeafRepository) GetLeaf(user *model.User) model.Leaf {
	var leaf model.Leaf
	err := common.DB.Where("userid", user.Model.ID).First(&leaf).Error
	if err != nil {
		leaf = model.Leaf{
			Userid:    user.Model.ID,
			Remaining: 100,
		}
		common.DB.Create(&leaf)

	}
	return leaf
}

func NewLeafRepository() ILeafRepository {
	return LeafRepository{}
}
