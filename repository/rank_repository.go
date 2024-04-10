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

	const sql = `SELECT t1.openid as player_id,
       t1.game_level,
       t1.created_at,
       t1.username,
       t1.avatar
FROM (
  SELECT gr.userid,
         gr.game_level,
         gr.created_at,
         u.avatar,
         u.openid,
         u.username,
         @rn := IF(@prev_userid = gr.userid, @rn + 1, 1) AS rn,
         @prev_userid := gr.userid
  FROM game_records gr
  JOIN users u ON u.id = gr.userid
  JOIN (SELECT @rn := 0, @prev_userid := NULL) var_init
  WHERE gr.status = 1
  ORDER BY gr.userid, gr.game_level DESC, gr.created_at DESC
) t1
WHERE t1.rn = 1;`
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
