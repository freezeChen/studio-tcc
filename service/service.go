/*
   @Time : 2019-07-09 21:57:17
   @Author :
   @File : service
   @Software: server
*/
package service

import (
	"errors"
	"github.com/freezeChen/studio-library/zlog"
	"studio-tcc/conf"
	"studio-tcc/dao"
	"studio-tcc/model"
	"studio-tcc/pkg/util"
	"studio-tcc/tcc"
)

type Service struct {
	dao *dao.Dao
	tcc tcc.Tcc
}

func New(c *conf.Config) (s *Service) {
	s = &Service{
		dao: dao.New(c),
		tcc: tcc.New(),
	}
	return s
}

//处理请求
func (svc Service) HandlerRequest(req *model.DoingReq) (err error) {
	bus := svc.dao.GetOrderBus()

	transaction := svc.dao.GentTransaction(bus.Id, req.Param)
	if transaction == nil {
		err = errors.New("事务启动失败")
		return
	}

	trySteps, err := svc.tcc.Try(transaction.Id, req, bus)

	err2 := svc.dao.SaveTryStep(trySteps)
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
	ids, err := svc.tcc.Cancel(transId, req, bus, steps)
	if err != nil {
		//cancel 操作失败
		if err = svc.dao.SetTransactionStatus(transId, model.Trans_cancel_fail); err != nil {
			//TODO 事务状态修改失败 操作
			return err
		}
		return
	}

	for _, v := range ids {
		if err = svc.dao.SetStepStatus(v, model.Step_cancel_success); err != nil {
			//TODO 状态修改失败
			return
		}
	}

	if err = svc.dao.SetTransactionStatus(transId, model.Trans_cancel_success); err != nil {
		//TODO
		return
	}

	return
}

func (svc Service) Confirm(transId int64, req *model.DoingReq, bus *model.TCCBus) (err error) {

	var response *model.Response
	for _, v := range bus.TCCS {

		response, err = svc.tcc.Confirm(transId, req, v)
		if err != nil {
			if err = svc.dao.SetTransactionStatus(transId, model.Trans_confirm_fail); err != nil {
				zlog.Infof("set transaction confirm_fail error(%v)", err)
				return
			}

			return
		}

		if response.Code != 0 {
			zlog.Infof("do confirm error(%s)", response.Msg)
			if err = svc.dao.SetTransactionStatus(transId, model.Trans_confirm_fail); err != nil {
				zlog.Infof("set transaction confirm_fail error(%v)", err)
				return
			}
			return
		}

	}

	if err := svc.dao.SetTransactionStatus(transId, model.Trans_confirm_success); err != nil {
	}

	return
}

func (svc Service) TaskGetTransactionList() []*model.Transaction {
	svc.dao.GentTransaction()
}
