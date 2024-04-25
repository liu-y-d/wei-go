// Package model @Author yd 2024/3/28 14:10
package model

type GamePropsGuide struct {
	Userid  uint `gorm:"type:bigint;primaryKey;autoIncrement:false" json:"userid"`
	PropsId uint `gorm:"type:bigint;primaryKey;autoIncrement:false" json:"propsId"`
	Show    uint `gorm:"type:tinyint(1);comment:'0 显示, 1 不显示'" json:"show"`
}
