package model


type BusDetail struct {

    Id int64 `json:"id" xorm:"'id'"`   //null
    BusId int64 `json:"bus_id" xorm:"'bus_id'"`   //null
    TryUrl string `json:"try_url" xorm:"'try_url'"`   //null
    ConfirmUrl string `json:"confirm_url" xorm:"'confirm_url'"`   //null
    CancelUrl string `json:"cancel_url" xorm:"'cancel_url'"`   //null
}
func (BusDetail) TableName() string {
 return "bus_detail" 
}
