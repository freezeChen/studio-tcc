package model


type UserRecord struct {

    Id int64 `json:"id" xorm:"'id'"`   //null
    Status int64 `json:"status" xorm:"'status'"`   //(1:new;2:run;3:finish

}
func (UserRecord) TableName() string {
 return "user_record" 
}
