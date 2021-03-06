package tcc

import (
	context "context"
	"errors"

	"studio-tcc/model"
	"studio-tcc/pkg/util"

	"github.com/freezeChen/studio-library/lib/errgroup"
)

type Tcc interface {
	Try(transId int64, req *model.DoingReq, bus *model.TCCBus) (successSteps []*model.TryStep, err error)
	Confirm(transId int64, req *model.DoingReq, tcc *model.TCC) (*model.Response, error)
	Cancel(transId int64, req *model.DoingReq, bus *model.TCCBus, steps []*model.TryStep) (ids []int64, err error)
}

type tcc struct {
}

func New() Tcc {
	return &tcc{}
}

func (tcc) Try(transId int64, req *model.DoingReq, bus *model.TCCBus) (successSteps []*model.TryStep, err error) {
	group := errgroup.Group{}

	for _, vo := range bus.TCCS {
		v := vo
		group.Go(func(ctx context.Context) error {
			var try model.TryStep
			var response = new(model.Response)
			response, err = util.HttpPost(v.Try.Url, &model.CallReq{
				TransId: transId,
				Param:   req.Param,
			})
			try.Url = v.Try.Url
			try.NodeId = v.Id
			try.Param = req.Param
			if err != nil {
				return err
			}
			if response.Code != 0 {
				err = errors.New(response.Msg)
				return err
			}

			try.Status = 1
			try.Tcc = v
			successSteps = append(successSteps, &try)
			return nil
		})

	}

	err = group.Wait()

	return
}

func (tcc) Confirm(transId int64, req *model.DoingReq, tcc *model.TCC) (*model.Response, error) {
	return util.HttpPost(tcc.Confirm.Url, &model.CallReq{TransId: transId, Param: req.Param})
}

func (tcc) Cancel(transId int64, req *model.DoingReq, bus *model.TCCBus, steps []*model.TryStep) (ids []int64, err error) {
	var response *model.Response
	for _, v := range steps {

		response, err = util.HttpPost(v.Tcc.Cancel.Url, &model.CallReq{TransId: transId, Param: req.Param})
		if err != nil {
			return
		}

		if response.Code == 0 {
			err = errors.New(response.Msg)
			return
		}
		ids = append(ids, v.Tcc.Id)
	}
	return
}
