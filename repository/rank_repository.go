// Package repository @Author yd 2024/3/28 14:19
package repository

import (
	"wei/common"
	"wei/dto"
)

type IRankRepository interface {

	// GetCurrentGameLevel 获取当前用户游戏关卡
	GetWorldRank() []dto.RankDto
}

type RankRepository struct {
}

func (r RankRepository) GetWorldRank() []dto.RankDto {

	const sql = `WITH RankedPlayerLevels AS (
  SELECT
    gr.userid,
    gr.game_level ,
    gr.created_at,
    u.avatar,
    u.openid,
    u.username,
    ROW_NUMBER() OVER (PARTITION BY gr.userid ORDER BY gr.game_level DESC, gr.created_at) AS rn
  FROM game_records gr
  JOIN users u ON u.id = gr.userid
  where gr.status = 1
)
SELECT
  openid player_id,
  game_level,
  created_at,
  username,
  avatar
FROM RankedPlayerLevels
WHERE rn = 1;`
	dtos := []dto.RankDto{}
	err := common.DB.Raw(sql).Scan(&dtos).Error
	if err != nil {
		return nil
	}
	return dtos
}

func NewRankRepository() IRankRepository {
	return RankRepository{}
}
