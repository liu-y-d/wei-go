// Package model @Author yd 2024/3/28 14:10
package model

type GamePropsConfig struct {
	Id                uint `gorm:"type:bigint;autoIncrement:false" json:"id"`
	Type              uint `gorm:"type:tinyint(1);comment:'1固定道具, 2地图随机道具'" json:"type"`
	FixedPropsNum     uint `gorm:"type:tinyint UNSIGNED;comment:'固定道具数量'" json:"fixedPropsNum"`
	RandomPropsWeight uint `gorm:"type:tinyint UNSIGNED;comment:'随机道具权重'" json:"randomPropsWeight"`
	Show              uint `gorm:"type:tinyint(1);comment:'0 显示, 1 不显示'" json:"show"`
}
