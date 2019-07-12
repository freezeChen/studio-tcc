package util

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"studio-tcc/model"
)

func HttpPost(url string, param *model.CallReq) (*model.Response, error) {

	marshal, err := json.Marshal(param)
	if err != nil {
		return nil, err
	}

	fmt.Println(string(marshal))


	rsp, err := http.DefaultClient.Post(url,"application/json",bytes.NewReader(marshal))

	if err != nil {
		return nil, err
	}

	if rsp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf(rsp.Status)
	}

	body, err := ioutil.ReadAll(rsp.Body)
	if err != nil {
		return nil, err
	}

	var res model.Response

	err = json.Unmarshal(body, &res)
	if err != nil {
		return nil, err
	}

	return &res, nil
}
