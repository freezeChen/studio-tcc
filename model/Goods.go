package model


type Goods struct {

    Id int64 `json:"id" xorm:"'id'"`   //null
    Name string `json:"name" xorm:"'name'"`   //null
    Num int64 `json:"num" xorm:"'num'"`   //null
    NumLock int64 `json:"num_lock" xorm:"'num_lock'"`   //null
    Price int64 `json:"price" xorm:"'price'"`   //null
}
func (Goods) TableName() string {
 return "goods" 
}
