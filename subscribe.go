package main

import (
	"encoding/base64"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"path"
	"strconv"
	"strings"
)

var configFile = "sub.temp"

type Sub struct {
	Url  string `json:"url"`
	Path string `json:"path"`
}

type V2ray struct {
	Path string `json:"path"`
	Tls  string `json:"tls"`
	Add  string `json:"add"`
	Port int    `json:"port"`
	Aid  int    `json:"aid"`
	Net  string `json:"net"`
	Id   string `json:"id"`
	Host string `json:"host"`
	Ps   string `json:"ps"`
	V    string `json:"v"`
	Type string `json:"type"`
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
	err = ioutil.WriteFile("./"+configFile, jsonBytes, 0644)

	if err != nil {
		return 500, err.Error()
	}

	return 200, ""
}

func obtainSubConfig() (code int, message string) {
	content, err := ioutil.ReadFile("./" + configFile)
	if err != nil {
		return 500, err.Error()
	}
	return 200, string(content)
}

func updateConfig() (code int, message string) {
	content, err := ioutil.ReadFile("./" + configFile)
	if err != nil {
		return 500, err.Error()
	}
	var s Sub
	json.Unmarshal(content, &s)
	if len(s.Url) == 0 {
		return 500, "缺少订阅地址"
	}
	if len(s.Path) == 0 {
		return 500, "缺少配置存储路径"
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
	firstLevelSource, err := base64.StdEncoding.DecodeString(string(body))
	if err != nil {
		return 500, err.Error()
	}

	if strings.HasPrefix(string(firstLevelSource), "vmess") {
		v2rCodeList := strings.Split(strings.ReplaceAll(string(firstLevelSource), "\n", ""), "vmess://")
		for _, item := range v2rCodeList {
			if len(item) <= 0 {
				continue
			}
			finalData, err := base64.StdEncoding.DecodeString(item)
			if err != nil {
				continue
			}

			var v V2ray
			json.Unmarshal(finalData, &v)
			fileName := v.Add + "~" + strconv.Itoa(v.Port) + ".json"
			ioutil.WriteFile(path.Join(s.Path, fileName), finalData, 0644)
		}
	} else {
		return 500, "不支持的解析协议"
	}

	return 200, ""
}
