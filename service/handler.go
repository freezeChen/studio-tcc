/*
   @Time : 2019-07-09 21:57:17
   @Author : 
   @File : service
   @Software: server
*/
package service

import (
	"context"
	"fmt"
	"steam/proto"
)

func (Service) Hello(ctx context.Context, req *proto.Req, reply *proto.Reply) error {
	reply.Message = fmt.Sprintf("hello %s, Congratulations you success call rpc service!", req.S)
	return nil
}

