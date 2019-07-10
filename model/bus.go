package model

type Bus struct {
	Id     int64  `json:"id" xorm:"pk 'id'"`      //null
	Status int64  `json:"status" xorm:"'status'"` //null
	Remark string `json:"remark" xorm:"'remark'"` //null
}

func (Bus) TableName() string {
	return "bus"
}
