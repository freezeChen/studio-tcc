/*
   @Time : 2019-07-09 21:57:17
   @Author :
   @File : service
   @Software: server
*/
package service

import (
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

func (svc Service) HandlerRequest(req *model.DoingReq) *model.Bus {

	bus := svc.dao.GetOrderBus()

}

func (svc Service) Try(req *model.DoingReq, bus *model.Bus) error {


	for _, v := range bus.TCCS {
		util.HttpPost(v.Try.Url, []byte(req.Param))
	}

}
