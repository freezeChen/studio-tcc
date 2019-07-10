package util

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"steam/model"
)

func HttpPost(url string, param []byte) (*model.Response, error) {

	r, err := http.NewRequest(http.MethodPost, url, bytes.NewReader(param))
	if err != nil {
		return nil, err
	}

	rsp, err := (&http.Client{}).Do(r)

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
