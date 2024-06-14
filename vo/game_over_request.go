// Package vo @Author yd 2024/3/28 16:36
package vo

type GameOverRequest struct {
	GameLevel uint `json:"gameLevel" form:"gameLevel" validate:"required"`
	// 1 win 2 lose
	Status uint `json:"status" form:"status" validate:"required,oneof=1 2"`
}

type CommonGameRequest struct {
	Param string `json:"param" form:"param"`
}

type SavePropsGuideRequest struct {
	PropsId *uint `json:"propsId" form:"propsId" validate:"required"`
	ShowTip uint  `json:"showTip" form:"showTip" validate:"required,oneof=0 1"`
}
