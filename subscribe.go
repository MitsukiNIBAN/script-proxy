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
	} else {
		err := ioutil.WriteFile("./"+SubTempFile, []byte(url), 0644)
		if err != nil {
			return 500, err.Error()
		}
	}

	if len(path) == 0 {
		return 403, "未填写配置保存路径"
	} else {
		err := ioutil.WriteFile("./"+ConfigSaveFolder, []byte(path), 0644)
		if err != nil {
			return 500, err.Error()
		}
	}

	return 200, ""
}

func obtainSubConfig() (code int, message string) {
	url, _ := ioutil.ReadFile("./" + SubTempFile)
	path, _ := ioutil.ReadFile("./" + ConfigSaveFolder)
	jsonBytes, err := json.Marshal(Sub{string(url), string(path)})
	if err != nil {
		return 500, err.Error()
	}
	return 200, string(jsonBytes)
}

func updateConfig() (code int, message string) {
	url, err := ioutil.ReadFile("./" + SubTempFile)
	if err != nil || len(string(url)) == 0 {
		return 500, "缺少订阅地址"
	}

	folder, err := ioutil.ReadFile("./" + ConfigSaveFolder)
	if err != nil || len(string(folder)) == 0 {
		return 500, "缺少配置存储路径"
	}

	resp, err := http.Get(string(url))
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
			ioutil.WriteFile(path.Join(string(folder), fileName), finalData, 0644)
		}
	} else {
		return 500, "不支持解析的协议"
	}

	return 200, ""
}
