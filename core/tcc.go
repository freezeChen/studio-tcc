package core

import "studio-tcc/model"

type Tcc interface {
	Registry(transId int64)

	Try(transId int64, req *model.DoingReq, bus *model.TCCBus) (successSteps []*model.TryStep, err error)
	Confirm(transId int64, req *model.DoingReq, tcc *model.TCC) (*model.Response, error)
	Cancel(transId int64, req *model.DoingReq, bus *model.TCCBus, steps []*model.TryStep) (ids []int64, err error)
}
