/*
   @Time : 2019-07-09 21:57:17
   @Author :
   @File : main
   @Software: server
*/
package main

import (
	"github.com/freezeChen/studio-library/zlog"
	"github.com/gin-gonic/gin"
	"studio-tcc/conf"
	"studio-tcc/server/http"
	"studio-tcc/service"
)

func main() {
	cfg, err := conf.Init()
	if err != nil {
		panic(err)
	}
	zlog.InitLogger(cfg.Log)

	s := service.New(cfg)
	engine := gin.Default()
	http.InitRouter(engine,s)

	if err := engine.Run(":8080"); err != nil {
		panic("gin run:" + err.Error())
		return
	}

}
