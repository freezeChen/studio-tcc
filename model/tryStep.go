/*
   @Time : 2019-07-10 11:03
   @Author : frozenChen
   @File : tryStep
   @Software: studio-tcc
*/
package model

const (
	Step_try_success = iota + 1
	Step_cancel_fail
	Step_cancel_success
	Step_confirm_fail
	step_comfirm_success
)

//(1:try成功;2:cancel失败;3:cancel成功;4:confirm失败;5:confirm成功)
type TryStep struct {
	Id      int64  `json:"id" xorm:"pk 'id'"`
	TransId int64  `json:"trans_id" xorm:"trans_id"`
	NodeId  int64  `json:"node_id" xorm:"node_id"`
	Url     string `json:"url" xorm:"url"`
	Param   string `json:"param" xorm:"param"`
	Msg     string `json:"msg" xorm:"msg"`
	Status  uint8  `json:"status" xorm:"status"`
	Tcc     *TCC   `xorm:"-"`
}

func (TryStep) TableName() string {
	return "try_step"
}
