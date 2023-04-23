package tools

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/url"
	"toolbox/common"
)

var Client = &http.Client{}

func Post(url string, data interface{}, headerInfo map[string]string) (out []byte, err error) {
	client := Client
	jsonStr, _ := json.Marshal(data)
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonStr))
	if err != nil {
		return out, err
	}
	for k, v := range headerInfo {
		if len(v) > common.Zero {
			req.Header.Add(k, v)
		}
	}
	res, err := client.Do(req)
	if err != nil {
		return out, err
	}
	defer res.Body.Close()
	result, _ := ioutil.ReadAll(res.Body)
	return result, nil
}

func PostForm(urlStr string, data url.Values) (out []byte, err error) {
	client := Client
	resp, err := client.PostForm(urlStr, data)

	if err != nil {
		return out, err
	}

	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)
	return body, nil
}
