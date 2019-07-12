/*
   @Time : 2019-07-09 21:57:17
   @Author :
   @File : dao
   @Software: server
*/
package dao

import (
	"errors"
	"github.com/freezeChen/studio-library/database/mysql"
	"github.com/freezeChen/studio-library/redis"
	_ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/xorm"
	"studio-tcc/conf"
	"studio-tcc/model"
	"studio-tcc/pkg/snowflake"
	"studio-tcc/pkg/util"
)

const _URL = "http://localhost:8081/"

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

//获取事务节点
func (d Dao) GetOrderBus() *model.TCCBus {

	var b = new(model.TCCBus)
	b.Id = 1
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

	return b

}

//新增事务
func (d Dao) GentTransaction(busId int64, param string) *model.Transaction {
	var trans = &model.Transaction{
		Id:     snowflake.GenID(),
		Busid:  busId,
		Status: 0,
		Param:  param,
	}

	affect, err := d.Db.InsertOne(trans)
	if err != nil || affect == 0 {
		return nil
	}

	return trans
}

//保存try步骤
func (d Dao) SaveTryStep(ts []*model.TryStep) error {

	session := d.Db.NewSession()
	session.Begin()
	defer session.Close()

	_, err := session.Insert(&ts)
	if err != nil {
		return err
	}

	if err := session.Commit(); err != nil {
		return err
	}

	return nil
}

func (d Dao) DoCancel(transId int64, req *model.DoingReq, steps []*model.TryStep) (ids []int64, err error) {
	for _, v := range steps {
		response, err1 := util.HttpPost(v.Tcc.Cancel.Url, &model.CallReq{TransId: transId, Param: req.Param})
		if err1 != nil {
			return
		}

		if response.Code == 0 {
			err = errors.New(response.Msg)
			return
		}
		ids = append(ids, v.Tcc.Id)
	}
	return
}

//修改事务状态
func (d Dao) SetTransactionStatus(id int64, status int64) error {
	result, err := d.Db.Exec("update transaction.transaction set status = ? where id=?;", status, id)
	if err != nil {
		return err
	}

	affected, _ := result.RowsAffected()
	if affected == 0 {
		return errors.New("not affected")
	}

	return nil
}

func (d Dao) SetStepStatus(id int64, status int) error {

	result, err := d.Db.Exec("update try_step set status = ? where id =?;", status, id)
	if err != nil {
		return err
	}

	affected, _ := result.RowsAffected()
	if affected == 0 {
		return errors.New("not affected")
	}
	return nil
}
