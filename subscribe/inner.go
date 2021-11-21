package subscribe

import (
	"crypto/md5"
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"mmm3w/go-proxy/support"
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
	resp, err := http.Get(url)
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

	configMapping := make(map[string]string)

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

		md5str := fmt.Sprintf("%x", md5.Sum(finalData))
		fileName := md5str + ".json"

		err = ioutil.WriteFile(path.Join(folder, fileName), finalData, 0644)
		if err == nil {
			configMapping[md5str] = v.Ps + "(" + v.Add + ":" + strconv.Itoa(v.Port) + ")"
		}
	}

	//保存相关内容方便查询
	return support.SaveConfig(path.Join(folder, support.V2rayConfigListCache), configMapping)
}

func updateSsrSub(url string, folder string) error {
	return nil
}

func getConfigSet(loadFile string) ([]map[string]string, error) {
	configMapping, err := support.LoadConfig(loadFile)
	if err != nil {
		return nil, err
	}
	outData := make([]map[string]string, 0)
	for k, v := range configMapping {
		item := make(map[string]string)
		item["name"] = v
		item["key"] = k
		outData = append(outData, item)
	}
	return outData, nil
}
