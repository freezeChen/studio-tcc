package http

import (
	"github.com/freezeChen/studio-library/zlog"
	"github.com/gin-gonic/gin"
	"studio-tcc/model"
	"studio-tcc/service"
)

var (
	svc *service.Service
)

type Controller struct {
	model.JsonResult
}

func InitRouter(engine *gin.Engine, s *service.Service) {
	svc = s
	c := &Controller{}
	engine.POST("/doing", c.doing)

}

func (self Controller) doing(ctx *gin.Context) {
	defer func() {
		self.Response(ctx)
	}()
	self.Code = -1
	var param model.DoingReq
	if err := ctx.ShouldBind(&param); err != nil {
		zlog.Errorf("Binding error(%v)", err)
		self.Msg = "参数错误"
		return
	}

	err := svc.HandlerRequest(&param)
	if err != nil {
		self.Msg = err.Error()
		return
	}

	self.Code = 0

}
