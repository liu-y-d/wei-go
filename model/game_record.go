// Package model @Author yd 2024/3/28 14:10
package model

import "time"

type GameRecord struct {
	Userid    uint `gorm:"type:bigint" json:"userid"`
	GameLevel uint `gorm:"type:bigint" json:"gameLevel"`
	CreatedAt time.Time
	Status    uint `gorm:"type:tinyint(1);comment:'1win, 2lose'" json:"status"`
}
