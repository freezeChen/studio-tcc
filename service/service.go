/*
   @Time : 2019-07-09 21:57:17
   @Author :
   @File : service
   @Software: server
*/
package service

import (
	"errors"
	"steam/conf"
	"steam/dao"
	"steam/model"
	"steam/pkg/util"
)

type Service struct {
	dao *dao.Dao

}

func New(c *conf.Config) (s *Service) {
	s = &Service{
		dao: dao.New(c),
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

	trySteps, err := svc.Try(req, bus)

	err2 := svc.dao.SaveTryStep(trySteps)
	if err2 != nil {
		//TODO 数据库保存失败 写入日志保存
	}

	if err != nil || err2 != nil {
		svc.Cancel(transaction.Id, req, bus, trySteps)
	}

	return
}

//执行try操作 返回操作成功的请求
func (svc Service) Try(req *model.DoingReq, bus *model.TCCBus) (successStep []*model.TryStep, err error) {
	for _, v := range bus.TCCS {
		var try model.TryStep
		var response = new(model.Response)
		response, err = util.HttpPost(v.Try.Url, []byte(req.Param))
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
	ids, err := svc.dao.DoCancel(req, steps)
	if err != nil {

		//cancel 操作失败
		if err := svc.dao.SetTransactionStatus(transId, model.Trans_cancel_fail); err != nil {
			//TODO 事务状态修改失败 操作
			return err
		}
		return
	}

	for _, v := range ids {
		if err := svc.dao.SetStepStatus(v, model.Step_cancel_success); err != nil {
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




}
