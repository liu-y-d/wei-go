package model

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Openid       string  `gorm:"type:varchar(200)" json:"openid"`
	Username     string  `gorm:"type:varchar(20);not null" json:"username"`
	Password     string  `gorm:"size:255;not null" json:"password"`
	Mobile       string  `gorm:"type:varchar(11);" json:"mobile"`
	Avatar       string  `gorm:"type:varchar(255)" json:"avatar"`
	Nickname     *string `gorm:"type:varchar(20)" json:"nickname"`
	Introduction *string `gorm:"type:varchar(255)" json:"introduction"`
	RoleId       uint    `gorm:"type:bigint" json:"roleId"`
	Status       uint    `gorm:"type:tinyint(1);default:1;comment:'1正常, 2禁用'" json:"status"`
	Creator      string  `gorm:"type:varchar(20);" json:"creator"`
}
