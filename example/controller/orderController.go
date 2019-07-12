/*
   @Time : 2019-07-12 15:20
   @Author : frozenChen
   @File : orderController
   @Software: studio-tcc
*/
package controller

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"studio-tcc/model"
)

type OrderController struct {
	model.JsonResult
}

func NewOrderController() *OrderController {
	c := &OrderController{}
	c.Code = -1
	return c
}

func (self OrderController) Router(router *gin.Engine) {
	userGroup := router.Group("/order")
	{
		userGroup.POST("/try", self.try)

		userGroup.POST("/confirm", self.confirm)

		userGroup.POST("/cancel", self.cancel)
	}
}

func (self OrderController) try(ctx *gin.Context) {
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

	if ExistTrans(param.TransId, 3, 1) {
		self.Code = 0
		return
	}

	session := Db.NewSession()
	session.Begin()
	defer session.Close()

	if err := AddTrans(session, param.TransId, 3, 1); err != nil {
		return
	}

	result, err := session.Exec("insert into `order` (id, uid, gid, num, price, time, status) values (?,?,?,?,10,now(),1);", param.TransId, oParam.Uid, oParam.Gid, oParam.Num)
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

func (self OrderController) confirm(ctx *gin.Context) {
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

	if ExistTrans(param.TransId, 3, 2) {
		self.Code = 0
		return
	}

	session := Db.NewSession()
	session.Begin()
	defer session.Close()

	if err := AddTrans(session, param.TransId, 3, 2); err != nil {
		return
	}
	result, err := session.Exec("update `order` set status = 2 where id=?;", param.TransId)
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

func (self OrderController) cancel(ctx *gin.Context) {
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

	if ExistTrans(param.TransId, 3, 3) {
		self.Code = 0
		return
	}

	session := Db.NewSession()
	session.Begin()
	defer session.Close()

	if err := AddTrans(session, param.TransId, 3, 3); err != nil {
		return
	}
	result, err := session.Exec("delete from `order` where id=?", param.TransId)
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
