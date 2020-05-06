package model

import "time"

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
	Id           int64     `json:"id" xorm:"pk 'id'"`                    //null
	Busid        int64     `json:"busid" xorm:"'busid'"`                 //null
	Status       int64     `json:"status" xorm:"'status'"`               //(1:try成功;2:try失败;3:cancel成功;4:cancel失败;5:confirm成功;6:confirm失败;7:人工干预)
	Param        string    `json:"param" xorm:"'param'"`                 //null
	TryTimes     int64     `json:"try_times" xorm:"'try_times'"`         //尝试次数
	ConfirmTimes int64     `json:"confirm_times" xorm:"'confirm_times'"` //null
	CancelTimes  int64     `json:"cancel_times" xorm:"'cancel_times'"`   //null
	Mtime        time.Time `json:"mtime" xorm:"'mtime' created"`
	CTime        time.Time `json:"ctime" xorm:"'ctime' updated"` //null
}

func (Transaction) TableName() string {
	return "transaction"
}
