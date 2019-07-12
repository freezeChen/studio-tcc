/*
   @Time : 2019-07-12 15:25:04
   @Author :
   @File : main
   @Software: example
*/
package main

import (
	"fmt"
	"github.com/freezeChen/studio-library/database/mysql"
	"github.com/freezeChen/studio-library/zlog"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"studio-tcc/example/conf"
	"studio-tcc/example/controller"
)

func main() {
	cfg, err := conf.Init()
	if err != nil {
		panic(err)
	}

	fmt.Println(fmt.Sprintf("%+v", cfg))

	zlog.InitLogger(cfg.Log)

	engine := gin.Default()

	controller.Db = mysql.New(cfg.Mysql)
	controller.NewUserController().Router(engine)
	controller.NewGoodsController().Router(engine)
	controller.NewOrderController().Router(engine)

	if err := engine.Run(":8081"); err != nil {
		return
	}

}
