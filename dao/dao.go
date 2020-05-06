/*
   @Time : 2019-07-09 21:57:17
   @Author :
   @File : dao
   @Software: server
*/
package dao

import (
	"errors"
	"fmt"

	"github.com/freezeChen/studio-library/database/mysql"
	"github.com/freezeChen/studio-library/redis"
	_ "github.com/go-sql-driver/mysql"
	"studio-tcc/conf"
	"studio-tcc/model"
	"studio-tcc/pkg/snowflake"
	"studio-tcc/pkg/util"
	"xorm.io/xorm"
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
		fmt.Println(err)
		return nil
	}

	return trans
}

//(1:try成功;2:try失败;3:cancel成功;4:cancel失败;5:confirm成功;6:confirm失败;7:人工干预)
func (d Dao) GetExTransactionList() []*model.Transaction {
	var list = make([]*model.Transaction, 0)
	if err := d.Db.SQL(`select * from transaction.transaction 
			where status=1 or (status=2 and try_times<5) or (status=4 and cancel_times<5) or (status=6 or confirm_times<5)
			and mtime>date_add(now(),interval -1 minute )`).Find(&list); err != nil {
		return nil
	}
	return list
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
	var sql string

	switch status {
	case model.Trans_cancel_fail:
		sql = "update transaction.transaction set status = ?,cancel_times=cancel_times+1 where id=?;"
	case model.Trans_confirm_fail:
		sql = "update transaction.transaction set status = ?,confirm_times=confirm_times+1 where id=?;"
	default:
		sql = "update transaction.transaction set status = ? where id=?;"
	}

	result, err := d.Db.Exec(sql, status, id)
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
