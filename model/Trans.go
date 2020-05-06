package model

type Trans struct {
	Id      int64 `json:"id" xorm:"pk 'id'"`          //null
	TransId int64 `json:"trans_id" xorm:"'trans_id'"` //事务Id
	Type    int64 `json:"type" xorm:"'type'"`         //(1:用户;2:库存;3:订单)
	ReqType int64 `json:"req_type" xorm:"'req_type'"` //(1:try;2:confirm:3cancel)
}

func (Trans) TableName() string {
	return "trans"
}
