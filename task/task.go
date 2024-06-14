// Package task @Author yd 2024/4/17 10:59
package task

import (
	"fmt"
	"time"
	"wei/common"
	"wei/repository"
)

type ITaskRegister interface {
	RegisterLeafRecover()
}
type TaskRegister struct {
	LeafRepository repository.ILeafRepository
}

func (t TaskRegister) RegisterLeafRecover() {
	// 创建一个定时器，每5分钟执行一次
	ticker := time.NewTicker(10 * time.Minute)

	go func() {
		for range ticker.C {
			common.Log.Info("RegisterLeafRecover")
			t.LeafRepository.RecoveryLeaf()
		}
	}()
}
func (t TaskRegister) RegisterRankCacheCleaner() {
	for {
		now := time.Now()
		tomorrow := now.AddDate(0, 0, 1)                                                               // 获取明天的同一时刻
		target := time.Date(tomorrow.Year(), tomorrow.Month(), tomorrow.Day(), 0, 0, 0, 0, time.Local) // 获取明天的0点时刻

		if now.After(target) { // 如果已经过了今天0点，调整到明天0点
			target = target.AddDate(0, 0, 1)
		}

		sleepDuration := target.Sub(now) // 计算距离明天0点的剩余时间

		fmt.Printf("Sleeping for %v until next execution at %v\n", sleepDuration, target)
		time.Sleep(sleepDuration) // 休眠到明天0点

		repository.WorldRankCache.Delete("worldRank")
	}
}

func NewTaskRegister() ITaskRegister {
	LeafRepository := repository.NewLeafRepository()
	return TaskRegister{LeafRepository: LeafRepository}
}
