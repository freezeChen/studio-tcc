/*
   @Time : 2019-07-09 21:57:17
   @Author :
   @File : dao
   @Software: server
*/
package dao

import (
	"github.com/freezeChen/studio-library/database/mysql"
	"github.com/freezeChen/studio-library/redis"
	_ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/xorm"
	"steam/conf"
	"steam/model"
)

const _URL = "http://localhost:8080/"

type Dao struct {
	Db    xorm.EngineInterface
	Redis *redis.Redis
}

func New(c *conf.Config) (dao *Dao) {
	dao = &Dao{
		Db:    mysql.New(c.Mysql),
		Redis: redis.New(c.Redis),
	}
	return
}

func (d Dao) GetOrderBus() *model.Bus {

	var b = new(model.Bus)
	b.TCCS = make([]*model.TCC, 0)

	//user
	b.TCCS = append(b.TCCS, &model.TCC{
		Name: "用户信息检查",
		Try: &model.Node{
			Url: _URL + "user/try",
		},
		Confirm: &model.Node{
			Url: _URL + "user/confirm",
		},
		Cancel: &model.Node{
			Url: _URL + "user/cancel",
		},
	})

	//goods
	b.TCCS = append(b.TCCS, &model.TCC{
		Name: "库存扣减",
		Try: &model.Node{
			Url: _URL + "goods/try",
		},
		Confirm: &model.Node{
			Url: _URL + "goods/confirm",
		},
		Cancel: &model.Node{
			Url: _URL + "goods/cancel",
		},
	})

	b.TCCS = append(b.TCCS, &model.TCC{
		Name: "生成订单",
		Try: &model.Node{
			Url: _URL + "order/try",
		},
		Confirm: &model.Node{
			Url: _URL + "order/confirm",
		},
		Cancel: &model.Node{
			Url: _URL + "order/cancel",
		},
	})

}
