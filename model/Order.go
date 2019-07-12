package model

type Order struct {
	Id     int64  `json:"id" xorm:"'id'"`         //null
	Uid    int64  `json:"uid" xorm:"'uid'"`       //null
	Gid    int64  `json:"gid" xorm:"'gid'"`       //null
	Num    int64  `json:"num" xorm:"'num'"`       //null
	Price  int64  `json:"price" xorm:"'price'"`   //null
	Time   string `json:"time" xorm:"'time'"`     //null
	Status int64  `json:"status" xorm:"'status'"` //null
}

func (Order) TableName() string {
	return "order"
}
