/*
   @Time : 2019-07-12 14:51
   @Author : frozenChen
   @File : task
   @Software: studio-tcc
*/
package task

import (
	"encoding/json"
	"time"

	"studio-tcc/model"
	"studio-tcc/service"
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
	list := t.svc.TaskGetTransactionList()

	return list
}

func (t *Task) exec(trans *model.Transaction) {
	bus := t.svc.GetBus(trans.Busid)
	req := model.DoingReq{}
	if err := json.Unmarshal([]byte(trans.Param), &req); err != nil {
		return
	}

	switch trans.Status {
	case 1:
		t.svc.Confirm(trans.Id, &req, bus)
		//t.svc.Confirm(trans.Id, trans.Param)

	case 4:
	case 6:

	}
}
