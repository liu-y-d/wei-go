// Package routes @Author yd 2024/3/28 14:52
package routes

import (
	jwt "github.com/appleboy/gin-jwt/v2"
	"github.com/gin-gonic/gin"
	"wei/controller"
)

// InitGameRoutes 注册游戏路由
func InitGameCoreRoutes(r *gin.RouterGroup, authMiddleware *jwt.GinJWTMiddleware) gin.IRoutes {
	gameController := controller.NewGameController()
	rankController := controller.NewRankController()
	leafController := controller.NewLeafController()
	router := r.Group("/game")
	// 开启jwt认证中间件
	router.Use(authMiddleware.MiddlewareFunc())
	// 开启casbin鉴权中间件
	//router.Use(middleware.CasbinMiddleware())
	{
		router.GET("/currentGameLevel", gameController.GetCurrentUserGameLevel)
		router.POST("/gameOver", gameController.GameOver)
		router.GET("/worldRank", rankController.WorldRank)
		router.GET("/getLeaf", leafController.GetLeaf)
		router.GET("/consumeLeaf", leafController.ConsumeLeaf)
		router.GET("/infinityLeaf", leafController.InfinityLeaf)
		router.GET("/propsConfig", gameController.GetGamePropsConfig)
		router.GET("/getAllPropsGuide", gameController.GetUserAllPropsGuide)
		router.POST("/savePropsGuide", gameController.SaveUserPropsGuide)
	}
	return r
}
