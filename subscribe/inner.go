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

type Ssr struct {
	LocalAddress  string `json:"local_address"`
	LocalPort     int    `json:"local_port"`
	Server        string `json:"server"`
	ServerPort    int    `json:"server_port"`
	Timeout       int    `json:"timeout"`
	Workers       int    `json:"workers"`
	Password      string `json:"password"`
	Method        string `json:"method"`
	Obfs          string `json:"obfs"`
	ObfsParam     string `json:"obfs_param"`
	Protocol      string `json:"protocol"`
	ProtocolParam string `json:"protocol_param"`
	Remarks       string `json:"remarks"`
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

	var configData []map[string]string

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
			summaryItem := make(map[string]string)
			summaryItem["id"] = md5str
			summaryItem["name"] = v.Ps
			configData = append(configData, summaryItem)
		}
	}
	jsonBytes, err := json.Marshal(configData)
	if err != nil {
		return err
	}
	if jsonBytes == nil {
		return errors.New("无配置信息")
	}
	return support.Write(path.Join(folder, support.V2rayConfigListCache), string(jsonBytes))
}

func updateSsrSub(url string, folder string) error {
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

	var configData []map[string]string

	ssrCodeList := strings.Split(strings.ReplaceAll(string(firstLevelSource), "\n", ""), "ssr://")
	for _, item := range ssrCodeList {
		if len(item) <= 0 {
			continue
		}
		finalData, err := base64.RawURLEncoding.DecodeString(item)
		if err != nil {
			continue
		}

		result := strings.Split(string(finalData), ":")
		if len(result) < 6 {
			continue
		}

		var s Ssr

		s.LocalAddress = "0.0.0.0"
		s.LocalPort = 33334
		s.Workers = 1
		s.Timeout = 300

		s.Server = result[0]
		s.ServerPort, _ = strconv.Atoi(result[1])
		s.Protocol = result[2]
		s.Method = result[3]
		s.Obfs = result[4]

		ns := strings.Split(result[5], "/?")

		if len(result) < 2 {
			continue
		}

		s.Password = support.NoErrorBase64(ns[0])

		params := strings.Split(ns[1], "&")

		for _, p := range params {
			pkv := strings.Split(p, "=")
			if len(pkv) < 2 {
				continue
			}
			if pkv[0] == "obfs_param" {
				s.ObfsParam = support.NoErrorBase64(pkv[1])
				continue
			}
			if pkv[0] == "protoparam" {
				s.ProtocolParam = support.NoErrorBase64(pkv[1])
				continue
			}
			if pkv[0] == "remarks" {
				s.Remarks = support.NoErrorBase64(pkv[1])
				continue
			}
		}

		md5str := fmt.Sprintf("%x", md5.Sum(finalData))
		fileName := md5str + ".json"
		jsonData, _ := json.Marshal(s)
		err = ioutil.WriteFile(path.Join(folder, fileName), jsonData, 0644)
		if err == nil {
			summaryItem := make(map[string]string)
			summaryItem["id"] = md5str
			summaryItem["name"] = s.Remarks
			configData = append(configData, summaryItem)
		}
	}

	jsonBytes, err := json.Marshal(configData)
	if err != nil {
		return err
	}
	if jsonBytes == nil {
		return errors.New("无配置信息")
	}
	return support.Write(path.Join(folder, support.SsrConfigListCache), string(jsonBytes))
}

func getConfigSet(loadFile string) (string, error) {
	data, err := support.Read(loadFile)
	if err != nil {
		return "[]", err
	}
	if data == "" {
		return "[]", nil
	}
	return data, nil
}
