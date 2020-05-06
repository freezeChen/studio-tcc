/*
   @Time : 2019-07-12 15:19
   @Author : frozenChen
   @File : userController
   @Software: studio-tcc
*/
package controller

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"studio-tcc/model"
	"xorm.io/xorm"
)

type UserController struct {
	model.JsonResult
}

func NewUserController() *UserController {
	c := &UserController{}
	c.Code = -1
	return c
}

func (self UserController) Router(router *gin.Engine) {
	userGroup := router.Group("/user")
	{
		userGroup.POST("/try", self.try)

		userGroup.POST("/confirm", self.confirm)

		userGroup.POST("/cancel", self.cancel)
	}
}

func (self UserController) try(ctx *gin.Context) {
	defer func() {
		self.Response(ctx)
	}()

	var param model.CallReq

	if err := ctx.ShouldBind(&param); err != nil {
		self.Msg = err.Error()
		return
	}



	var oParam model.GenOrderReq

	fmt.Println(param.Param,param)
	err := json.Unmarshal([]byte(param.Param), &oParam)
	if err != nil {
		fmt.Println("oparam")
		self.Msg = err.Error()
		return
	}

	if ExistTrans(param.TransId, 1, 1) {
		self.Code = 0
		return
	}

	session := Db.NewSession()
	session.Begin()
	defer session.Close()

	if err := AddTrans(session, param.TransId, 1, 1); err != nil {
		self.Msg = err.Error()
		return
	}

	result, err := session.Exec("update user set money_lock = money_lock+? where money>money_lock+? and id=?", oParam.Num*10, oParam.Num*10, oParam.Uid)
	if err != nil {
		self.Msg = err.Error()
		return
	}
	affected, _ := result.RowsAffected()
	if affected == 0 {
		self.Msg = "not affect "
		return
	}

	if err := session.Commit(); err != nil {
		self.Msg = err.Error()
		return
	}

	self.Code = 0

}

func (self UserController) confirm(ctx *gin.Context) {
	defer func() {
		self.Response(ctx)
	}()

	var param model.CallReq

	if err := ctx.ShouldBind(&param); err != nil {
		self.Msg = err.Error()
		return
	}

	var oParam model.GenOrderReq

	err := json.Unmarshal([]byte(param.Param), &oParam)
	if err != nil {
		self.Msg = err.Error()
		return
	}

	if ExistTrans(param.TransId, 1, 2) {
		self.Code = 0
		return
	}

	session := Db.NewSession()
	session.Begin()
	defer session.Close()

	if err := AddTrans(session, param.TransId, 1, 2); err != nil {
		self.Msg = err.Error()
		return
	}
	result, err := session.Exec("update user set money=money-?,money_lock=money_lock-? where money_lock>? and id=?", oParam.Num*10, oParam.Num*10, oParam.Num*10, oParam.Uid)
	if err != nil {
		self.Msg = err.Error()
		return
	}
	affected, _ := result.RowsAffected()
	if affected == 0 {
		self.Msg = "not affect"
		return
	}

	if err := session.Commit(); err != nil {
		self.Msg = err.Error()
		return
	}

	self.Code = 0

}

func (self UserController) cancel(ctx *gin.Context) {
	defer func() {
		self.Response(ctx)
	}()

	var param model.CallReq

	if err := ctx.ShouldBind(&param); err != nil {
		return
	}

	var oParam model.GenOrderReq

	err := json.Unmarshal([]byte(param.Param), &oParam)
	if err != nil {
		return
	}

	if ExistTrans(param.TransId, 1, 3) {
		self.Code = 0
		return
	}

	session := Db.NewSession()
	session.Begin()
	defer session.Close()

	if err := AddTrans(session, param.TransId, 1, 3); err != nil {
		return
	}
	result, err := session.Exec("update user set money_lock=money_lock-? where money_lock>? and id=?", oParam.Num*10, oParam.Num*10, oParam.Uid)
	if err != nil {
		return
	}
	affected, _ := result.RowsAffected()
	if affected == 0 {
		return
	}

	if err := session.Commit(); err != nil {
		return
	}

	self.Code = 0

}

func AddTrans(session *xorm.Session, transId, tye, reqTye int64) (err error) {

	one, err := session.InsertOne(&model.Trans{
		TransId: transId,
		Type:    tye,
		ReqType: reqTye,
	})
	if err != nil {
		return
	}
	if one == 0 {
		return errors.New("not affect")
	}

	return
}

func ExistTrans(transId, tye, reqTye int64) bool {
	exist, err := Db.Exist(&model.Trans{
		TransId: transId,
		Type:    tye,
		ReqType: reqTye,
	})

	if err != nil {
		return false
	}

	return exist
}
