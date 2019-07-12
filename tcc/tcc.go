package tcc

import (
	"errors"
	"studio-tcc/model"
	"studio-tcc/pkg/util"
)

type Tcc interface {
	Try(req *model.DoingReq, bus *model.TCCBus) (successSteps []*model.TryStep, err error)
	Confirm()
	Cancel(transId int64, req *model.DoingReq, bus *model.TCCBus, steps []*model.TryStep) (err error)
}

type tcc struct {
}

func New() *tcc {
	return &tcc{}
}

func (tcc) Try(req *model.DoingReq, bus *model.TCCBus) (successSteps []*model.TryStep, err error) {
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
		successSteps = append(successSteps, &try)
	}
	return
}

func (tcc) Confirm() {
	panic("implement me")
}

func (tcc) Cancel(transId int64, req *model.DoingReq, bus *model.TCCBus, steps []*model.TryStep) (err error){
	panic("implement me")
}
