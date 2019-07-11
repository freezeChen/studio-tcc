package model

const (
	Trans_try_success = iota + 1
	Trans_try_fail
	Trans_cancel_success
	Trans_cancel_fail
	Trans_confirm_success
	Trans_confirm_fail
	Trans_manual
)

//(1:try成功;2:try失败;3:cancel成功;4:cancel失败;5:confirm成功;6:confirm失败;7:人工干预)

type Transaction struct {
	Id     int64  `json:"id" xorm:"pk 'id'"`      //null
	Busid  int64  `json:"busid" xorm:"'busid'"`   //null
	Status int64  `json:"status" xorm:"'status'"` //null
	Param  string `json:"param"`
}

func (Transaction) TableName() string {
	return "transaction"
}
