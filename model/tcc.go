package model

//每组业务请求
type TCC struct {
	Name   string   `json:"Name"`
	Try     *Node `json:"try"`
	Confirm *Node `json:"confirm"`
	Cancel  *Node `json:"cancel"`
}

//节点
type Node struct {
	Url string `json:"url"`
}

type Bus struct {
	TCCS []*TCC `json:"tccs"`
}
