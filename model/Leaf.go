// Package model @Author yd 2024/4/12 10:49
package model

import "time"

type Leaf struct {
	Userid    uint       `gorm:"type:bigint" json:"userid"`
	Remaining int        `json:"remaining"`
	Infinity  *time.Time `json:"infinity"`
}
