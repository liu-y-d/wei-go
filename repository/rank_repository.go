// Package repository @Author yd 2024/3/28 14:19
package repository

import (
	"github.com/patrickmn/go-cache"
	"time"
	"wei/common"
	"wei/dto"
)

type IRankRepository interface {

	// GetCurrentGameLevel 获取当前用户游戏关卡
	GetWorldRank() []dto.RankDto
}

var WorldRankCache = cache.New(24*time.Hour, 24*time.Hour)

type RankRepository struct {
}

func (r RankRepository) GetWorldRank() []dto.RankDto {
	var worldRank []dto.RankDto
	// 先获取缓存
	cacheRank, found := WorldRankCache.Get("worldRank")

	if found {
		worldRank = cacheRank.([]dto.RankDto)
	} else {
		const sql = ` select u.openid as player_id,
       game_records.game_level,
       u.created_at,
       u.username,
       u.avatar
from game_records
         join (select userid, MAX(game_level) as game_level from game_records where status = 1 group by userid) as max
on max.userid = game_records.userid and max.game_level = game_records.game_level and status =1
left JOIN users u ON game_records.userid = u.id
order by game_level desc ,created_at limit 50;
`
		dtos := []dto.RankDto{}
		err := common.DB.Raw(sql).Scan(&dtos).Error
		if err != nil {
			return nil
		}
		WorldRankCache.Set("worldRank", dtos, cache.DefaultExpiration)
		worldRank = dtos
	}

	return worldRank
}

func NewRankRepository() IRankRepository {
	return RankRepository{}
}
