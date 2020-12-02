package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type Sub struct {
	Url  string `json:"url"`
	Path string `json:"path"`
}

func saveSubConfig(url string, path string) (code int, message string) {
	if len(url) == 0 {
		return 403, "未填写订阅地址"
	}

	if len(path) == 0 {
		return 403, "未填写配置保存路径"
	}

	jsonBytes, err := json.Marshal(Sub{url, path})
	if err != nil {
		return 500, err.Error()
	}
	err = ioutil.WriteFile("./sub", jsonBytes, 0644)

	if err != nil {
		return 500, err.Error()
	}

	return 200, ""
}

func obtainSubConfig() (code int, message string) {
	content, err := ioutil.ReadFile("./sub")
	if err != nil {
		return 500, err.Error()
	}
	return 200, string(content)
}

func updateConfig() (code int, message string) {
	content, err := ioutil.ReadFile("./sub")
	if err != nil {
		return 500, err.Error()
	}
	var s Sub
	json.Unmarshal(content, &s)
	if len(s.Url) == 0 {
		return 500, "缺少订阅地址"
	}
	resp, err := http.Get(s.Url)
	if err != nil {
		return 500, err.Error()
	}
	defer resp.Body.Close()
	if resp.StatusCode != 200 {
		return resp.StatusCode, ""
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return 500, err.Error()
	}

	//这里拿到body后解析数据并写入文件
	fmt.Println(string(body))

	return 200, ""
}
