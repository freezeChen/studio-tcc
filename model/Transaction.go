package model

type Transaction struct {
	Id     int64  `json:"id" xorm:"pk 'id'"`      //null
	Busid  int64  `json:"busid" xorm:"'busid'"`   //null
	Status int64  `json:"status" xorm:"'status'"` //null
	Param  string `json:"param"`
}

func (Transaction) TableName() string {
	return "transaction"
}
