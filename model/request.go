package model

type DoingReq struct {
	Index string `json:"index"`
	Param string `json:"param"`
}

type GenOrderReq struct {
	Uid int64 `json:"uid"`
	Gid int64 `json:"gid"`
	Num int   `json:"num"`
}

type Response struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
}
