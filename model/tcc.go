package model

//每组业务请求
type TCC struct {
	Id      int64  `json:"id"`
	Name    string `json:"Name"`
	Try     *Node  `json:"try"`
	Confirm *Node  `json:"confirm"`
	Cancel  *Node  `json:"cancel"`
}

//节点
type Node struct {
	Url string `json:"url"`
}

type TCCBus struct {
	Id   int64  `json:"id"`
	TCCS []*TCC `json:"tccs"`
}
