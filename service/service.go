/*
   @Time : 2019-07-09 21:57:17
   @Author :
   @File : service
   @Software: server
*/
package service

import (
	context "context"
	"errors"

	"studio-tcc/conf"
	"studio-tcc/model"
	"studio-tcc/pkg/util"
	"studio-tcc/repository"
	"studio-tcc/tcc"

	"github.com/freezeChen/studio-library/lib/errgroup"
	"github.com/freezeChen/studio-library/zlog"
)

type Service struct {
	repo repository.Repository
	tcc  tcc.Tcc
}

func New(c *conf.Config) (s *Service) {
	s = &Service{
		repo: repository.New(c),
		tcc:  tcc.New(),
	}
	return s
}

//处理请求
func (svc Service) HandlerRequest(req *model.DoingReq) (err error) {
	bus := svc.repo.GetBus()

	transaction := svc.repo.GentTransaction(bus.Id, req.Param)
	if transaction == nil {
		err = errors.New("事务启动失败")
		return
	}

	trySteps, err := svc.tcc.Try(transaction.Id, req, bus)

	err2 := svc.repo.SaveTryStep(trySteps)
	if err2 != nil {
		//TODO 数据库保存失败 写入日志保存
	}

	if err != nil || err2 != nil {
		err = svc.Cancel(transaction.Id, req, bus, trySteps)
		if err != nil {
			return
		}
		err = errors.New("订单生成失败")
		return
	} else {
		err = svc.Confirm(transaction.Id, req, bus)
		if err != nil {
			return
		}

	}

	return
}

func (svc Service) GetBus(busId int64) *model.TCCBus {
	return svc.repo.GetBus()
}

//执行try操作 返回操作成功的请求
func (svc Service) Try(transId int64, req *model.DoingReq, bus *model.TCCBus) (successStep []*model.TryStep, err error) {
	for _, v := range bus.TCCS {
		var try model.TryStep
		var response = new(model.Response)

		response, err = util.HttpPost(v.Try.Url, &model.CallReq{TransId: transId, Param: req.Param})
		try.Url = v.Try.Url
		try.NodeId = v.Id
		try.Param = req.Param
		if err != nil {
			return
		}
		if response.Code != 0 {
			err = errors.New(response.Msg)
			return
		}

		try.Status = 1
		try.Tcc = v
		successStep = append(successStep, &try)
	}
	return
}

func (svc Service) Cancel(transId int64, req *model.DoingReq, bus *model.TCCBus, steps []*model.TryStep) (err error) {
	_, err = svc.tcc.Cancel(transId, req, bus, steps)
	if err != nil {
		//cancel 操作失败
		if err = svc.repo.SetTransactionStatus(transId, model.Trans_cancel_fail); err != nil {
			//TODO 事务状态修改失败 操作
			return err
		}
		return
	}

	//for _, v := range ids {
	//	if err = svc.repository.SetStepStatus(v, model.Step_cancel_success); err != nil {
	//		//TODO 状态修改失败
	//		return
	//	}
	//}

	if err = svc.repo.SetTransactionStatus(transId, model.Trans_cancel_success); err != nil {
		//TODO
		return
	}

	return
}

func (svc Service) Confirm(transId int64, req *model.DoingReq, bus *model.TCCBus) (err error) {

	var response *model.Response
	group := errgroup.Group{}

	for _, vo := range bus.TCCS {
		v := vo
		group.Go(func(ctx context.Context) error {
			response, err = svc.tcc.Confirm(transId, req, v)
			if err != nil {
				if err = svc.repo.SetTransactionStatus(transId, model.Trans_confirm_fail); err != nil {
					zlog.Infof("set transaction confirm_fail error(%v)", err)
					return err
				}

				return err
			}

			if response.Code != 0 {
				zlog.Infof("do confirm error(%s)", response.Msg)
				if err = svc.repo.SetTransactionStatus(transId, model.Trans_confirm_fail); err != nil {
					zlog.Infof("set transaction confirm_fail error(%v)", err)
					return err
				}
				return err
			}
			return nil
		})

	}

	group.Wait()
	if err := svc.repo.SetTransactionStatus(transId, model.Trans_confirm_success); err != nil {
	}

	return
}

func (svc Service) TaskGetTransactionList() []*model.Transaction {
	return svc.repo.GetExTransactionList()
}
