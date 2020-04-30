package model

type User struct {
	Id        int64  `json:"id" xorm:"'id'"`                 //null
	Name      string `json:"name" xorm:"'name'"`             //null
	Point     int64  `json:"point" xorm:"'point'"`           //null
	PointLock int64  `json:"point_lock" xorm:"'point_lock'"` //null
	Money     int64  `json:"money" xorm:"'money'"`           //null
	MoneyLock int64  `json:"money_lock" xorm:"'money_lock'"` //null
}

func (User) TableName() string {
	return "user"
}
