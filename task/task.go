/*
   @Time : 2019-07-12 14:51
   @Author : frozenChen
   @File : task
   @Software: studio-tcc
*/
package task

import (
	"encoding/json"
	"github.com/freezeChen/studio-library/zlog"
	"studio-tcc/model"
	"studio-tcc/service"
	"sync"
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

func (t *Task) Run() {
	once := sync.Once{}
	once.Do(func() {
		go func() {
			for {
				timer := time.NewTimer(5 * time.Second)
				<-timer.C
				list := t.getExTransactionList()
				if list == nil {
					continue
				}

				for _, transaction := range list {
					t.Do(transaction)
				}

			}

		}()

	})

}

func (t *Task) getExTransactionList() []*model.Transaction {
	list := t.svc.TaskGetTransactionList()
	return list
}

func (t *Task) exec(trans []*model.Transaction) {

}

func (t *Task) Do(trans *model.Transaction) {
	var (
		err error
		bus = t.svc.GetBus(trans.Busid)
		req = model.DoingReq{}
	)

	if err = json.Unmarshal([]byte(trans.Param), &req); err != nil {
		return
	}

	switch trans.Status {
	case 1, 6:
		err = t.svc.Confirm(trans.Id, &req, bus)

	case 4:
		var step = make([]*model.TryStep, 0)

		for _, tcc := range bus.TCCS {

			step = append(step, &model.TryStep{
				TransId: trans.Id,
				NodeId:  tcc.Id,
				Url:     tcc.Try.Url,
				Param:   req.Param,
				Msg:     "",
				Status:  1,
				Tcc: &model.TCC{
					Cancel: &model.Node{
						Url: tcc.Cancel.Url,
					},
				},
			})
		}

		err = t.svc.Cancel(trans.Id, &req, bus, step)

	}

	if err != nil {
		zlog.Errorf("Failed dong auto:error(%v)", err)
	}
}
