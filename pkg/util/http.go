package util

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"
)

func HttpPost(url string,param []byte) ([]byte, error) {

	r, err := http.NewRequest(http.MethodPost, url, bytes.NewReader(param))
	if err != nil {
		return nil, err
	}


	rsp, err := (&http.Client{

	}).Do(r)
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

	return body, nil
}
