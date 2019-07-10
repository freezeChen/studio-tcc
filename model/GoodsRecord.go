package model


type GoodsRecord struct {

    Id int64 `json:"id" xorm:"'id'"`   //null
    Status int64 `json:"status" xorm:"'status'"`   //null
}
func (GoodsRecord) TableName() string {
 return "goods_record" 
}
