package subscribe

import (
	"encoding/base64"
	"encoding/json"
	"errors"
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

func updateV2raySub(url string, folder string) error {
	resp, err := http.Get(string(url))
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	if resp.StatusCode != 200 {
		return errors.New("更新订阅失败:未请求到配置信息")
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	firstLevelSource, err := base64.StdEncoding.DecodeString(string(body))
	if err != nil {
		return err
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
		return errors.New("更新订阅失败:不支持解析的协议")
	}
	return nil
}

func updateSsrSub(url string, folder string) error {
	return nil
}

// func obtainServerSet() (int, string) {
// 	dirPth, err := ioutil.ReadFile(configSaveFolderTempFilePath())
// 	if err != nil {
// 		return 500, err.Error()
// 	}
// 	dir, err := ioutil.ReadDir(string(dirPth))
// 	if err != nil {
// 		return 500, err.Error()
// 	}
// 	serverStr := ""
// 	for _, fi := range dir {
// 		content, _ := ioutil.ReadFile(path.Join(string(dirPth), fi.Name()))
// 		var item V2ray
// 		json.Unmarshal(content, &item)
// 		if !strings.Contains(serverStr, item.Add) {
// 			if len(serverStr) <= 0 {
// 				serverStr = serverStr + item.Add
// 			} else {
// 				serverStr = serverStr + " " + item.Add
// 			}
// 		}
// 	}
// 	return 200, serverStr
// }

// func obtainPortSet() (int, string) {
// 	dirPth, err := ioutil.ReadFile(configSaveFolderTempFilePath())
// 	if err != nil {
// 		return 500, err.Error()
// 	}
// 	dir, err := ioutil.ReadDir(string(dirPth))
// 	if err != nil {
// 		return 500, err.Error()
// 	}
// 	portStr := ""
// 	for _, fi := range dir {
// 		content, _ := ioutil.ReadFile(path.Join(string(dirPth), fi.Name()))
// 		var item V2ray
// 		json.Unmarshal(content, &item)

// 		p := strconv.Itoa(item.Port)

// 		if !strings.Contains(portStr, p) {
// 			if len(portStr) <= 0 {
// 				portStr = portStr + p
// 			} else {
// 				portStr = portStr + "," + p
// 			}
// 		}
// 	}
// 	return 200, portStr
// }

// func obtainConfigSet() (int, string) {
// 	dirPth, err := ioutil.ReadFile(configSaveFolderTempFilePath())
// 	if err != nil {
// 		return 500, err.Error()
// 	}
// 	dir, err := ioutil.ReadDir(string(dirPth))
// 	if err != nil {
// 		return 500, err.Error()
// 	}

// 	setStr := ""
// 	for _, fi := range dir {
// 		content, _ := ioutil.ReadFile(path.Join(string(dirPth), fi.Name()))
// 		if len(setStr) <= 0 {
// 			setStr = setStr + string(content)
// 		} else {
// 			setStr = setStr + "," + string(content)
// 		}
// 	}

// 	return 200, "[" + setStr + "]"
// }
