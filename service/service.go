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

	tryRes, err := svc.Try(req, bus)

	err2 := svc.dao.SaveTryStep(tryRes)
	if err2 != nil {
		//TODO 数据库保存失败 写入日志保存

	}

	if err != nil || err2 != nil {
		svc.Cancel()
	}

}

//执行try操作 返回操作成功的请求
func (svc Service) Try(req *model.DoingReq, bus *model.TCCBus) (successStep []*model.TryStep, err error) {
	for _, v := range bus.TCCS {
		var try model.TryStep
		var response = new(model.Response)
		response, err = util.HttpPost(v.Try.Url, []byte(req.Param))
		try.Url = v.Try.Url
		try.NodeId = try.Id
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

func (svc Service) Cancel(req *model.DoingReq, bus *model.TCCBus) (err error) {
	
}
