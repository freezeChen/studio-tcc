package http

import (
	"github.com/gin-gonic/gin"
	"steam/model"
	"steam/service"
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
	engine.GET("/doing", c.doing)

}

func (self Controller) doing(ctx *gin.Context) {
	defer func() {
		self.Response(ctx)
	}()
	var param model.DoingReq
	if err := ctx.ShouldBind(&param); err != nil {
		self.Msg = "参数错误"
		return
	}





}
