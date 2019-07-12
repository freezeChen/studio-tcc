/*
   @Time : 2019-07-12 15:20
   @Author : frozenChen
   @File : goods
   @Software: studio-tcc
*/
package controller

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"studio-tcc/model"
)

type GoodsController struct {
	model.JsonResult
}

func NewGoodsController() *GoodsController {
	c := &GoodsController{}
	c.Code = -1
	return c
}

func (self GoodsController) Router(router *gin.Engine) {
	userGroup := router.Group("/goods")
	{
		userGroup.POST("/try", self.try)

		userGroup.POST("/confirm", self.confirm)

		userGroup.POST("/cancel", self.cancel)
	}
}

func (self GoodsController) try(ctx *gin.Context) {
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

	if ExistTrans(param.TransId, 2, 1) {
		self.Code = 0
		return
	}

	session := Db.NewSession()
	session.Begin()
	defer session.Close()

	if err := AddTrans(session, param.TransId, 2, 1); err != nil {
		return
	}

	result, err := session.Exec("update goods set num_lock = num_lock+? where num>num_lock+? and id=?;", oParam.Num, oParam.Num, oParam.Gid)
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

func (self GoodsController) confirm(ctx *gin.Context) {
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

	if ExistTrans(param.TransId, 2, 2) {
		self.Code = 0
		return
	}

	session := Db.NewSession()
	session.Begin()
	defer session.Close()

	if err := AddTrans(session, param.TransId, 2, 2); err != nil {
		return
	}
	result, err := session.Exec("update goods set num = num-? ,num_lock=num_lock-? where num_lock>? and id=?;", oParam.Num, oParam.Num, oParam.Num, oParam.Gid)
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

func (self GoodsController) cancel(ctx *gin.Context) {
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

	if ExistTrans(param.TransId, 2, 3) {
		self.Code = 0
		return
	}

	session := Db.NewSession()
	session.Begin()
	defer session.Close()

	if err := AddTrans(session, param.TransId, 2, 3); err != nil {
		return
	}
	result, err := session.Exec("update goods set num_lock=num_lock-? where num_lock>? and id=?", oParam.Num, oParam.Num, oParam.Gid)
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
