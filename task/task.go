// Package task @Author yd 2024/4/17 10:59
package task

import (
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
	ticker := time.NewTicker(5 * time.Minute)

	go func() {
		for range ticker.C {
			common.Log.Info("RegisterLeafRecover")
			t.LeafRepository.RecoveryLeaf()
		}
	}()
}

func NewTaskRegister() ITaskRegister {
	LeafRepository := repository.NewLeafRepository()
	return TaskRegister{LeafRepository: LeafRepository}
}
