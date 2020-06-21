/*
   @Time : 2019-07-09 21:57:17
   @Author :
   @File : model
   @Software: server
*/
package model

import (
	"fmt"
	"github.com/freezeChen/studio-library/zlog"
	"github.com/gin-gonic/gin"
	jsoniter "github.com/json-iterator/go"
	"net/http"
	"net/url"
)

type JsonResult struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data,omitempty"`
}

func (self *JsonResult) Response(ctx *gin.Context) {
	var (
		params url.Values
	)

	//获取传入参数
	if ctx.Request.Method == "GET" {
		params = ctx.Request.URL.Query()
	} else {
		if err := ctx.Request.ParseMultipartForm(32 << 20); err == nil {
			params = ctx.Request.PostForm
		}
	}

	var p string
	for k, v := range params {
		p += fmt.Sprintf("|%s=%s", k, v[0])
	}

	//获取错误码对应文字
	if self.Code == 0 {
		self.Msg = "success"
		zlog.ApiInfof(p, "success")
	} else {
		if len(self.Msg) == 0 {
			self.Msg = "操作失败"
		}
		zlog.ApiErrorf(p, "错误:%s", self.Msg)
	}

	str, _ := jsoniter.Marshal(self)
	ctx.Data(http.StatusOK, "application/json; charset=utf-8", str)
}
