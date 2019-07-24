/*
   @Time : 2019-07-12 14:51
   @Author : frozenChen
   @File : task
   @Software: studio-tcc
*/
package task

import (
	"studio-tcc/model"
	"studio-tcc/service"
	"time"
)

type Task struct {
	svc *service.Service
}

func New(s *service.Service) *Task {
	return &Task{
		svc: s,
	}
}

func Run() {
	go func() {
		for {
			timer := time.NewTimer(1 * time.Minute)
			<-timer.C

		}

	}()
}

func (t *Task) getExTransactionList() []*model.Transaction {
//t.svc
}
