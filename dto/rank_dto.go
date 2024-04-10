package dto

type RankDto struct {
	PlayerId  string `json:"playerId"`
	Username  string `json:"username"`
	Avatar    string `json:"avatar"`
	GameLevel uint   `json:"gameLevel"`
}
