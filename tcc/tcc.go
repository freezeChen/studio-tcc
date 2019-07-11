package tcc

import "steam/model"

type tcc interface {
	Try() (successSteps []*model.TryStep, err error)
	Confirm()
	Cancel(teq model.DoingReq, steps []*model.TryStep) (err error)
}

type Tcc struct {
}

func (t *Tcc) Try() {

}
