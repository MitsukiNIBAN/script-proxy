package main

import (
	"bufio"
	"encoding/json"
	"io"
	"io/ioutil"
	"os"
	"strings"
)

type Config struct {
	Mode               string `json:"mode"`
	Tcponly            bool   `json:"tcponly"`
	Selfonly           bool   `json:"selfonly"`
	Proxy_svraddr4     string `json:"proxy_svraddr4"`
	Proxy_svrport      string `json:"proxy_svrport"`
	Proxy_startcmd     string `json:"proxy_startcmd"`
	Proxy_stopcmd      string `json:"proxy_stopcmd"`
	Dnsmasq_log_enable bool   `json:"dnsmasq_log_enable"`
	Chinadns_verbose   bool   `json:"chinadns_verbose"`
	Dns2tcp_verbose    bool   `json:"dns2tcp_verbose"`
	File_ignlist_ext   string `json:"file_ignlist_ext"`
}

//脚本整体状态
func scriptStatus() {

}

func startScript() {

}

func stopScript() {

}

//是否在正在运行
func isRunning() bool {
	return false
}

func obtainConfigPath() (code int, message string) {
	content, err := ioutil.ReadFile("./" + StpcTempFile)
	if err != nil {
		return 500, err.Error()
	}
	return 200, string(content)
}

func obtainConfig(path string) (code int, message string) {
	err := ioutil.WriteFile("./"+StpcTempFile, []byte(path), 0644)
	if err != nil {
		return 500, err.Error()
	}

	file, err := os.OpenFile(path, os.O_RDWR, 0666)
	if err != nil {
		return 500, "Read config error:" + err.Error()
	}
	defer file.Close()

	buf := bufio.NewReader(file)
	var config Config
	for {
		line, err := buf.ReadString('\n')
		line = strings.TrimSpace(line)

		if strings.HasPrefix(line, "mode") {
			config.Mode = line[strings.Index(line, "'")+1 : strings.LastIndex(line, "'")]
		}

		if strings.Contains(line, "tcponly=") {
			config.Tcponly = line[strings.Index(line, "'")+1:strings.LastIndex(line, "'")] == "true"
		}

		if strings.Contains(line, "selfonly=") {
			config.Selfonly = line[strings.Index(line, "'")+1:strings.LastIndex(line, "'")] == "true"
		}

		if strings.Contains(line, "proxy_svraddr4=") {
			config.Proxy_svraddr4 = line[strings.Index(line, "(")+1 : strings.LastIndex(line, ")")]
		}

		if strings.Contains(line, "proxy_svrport=") {
			config.Proxy_svrport = line[strings.Index(line, "'")+1 : strings.LastIndex(line, "'")]
		}

		if strings.Contains(line, "proxy_startcmd=") {
			config.Proxy_startcmd = line[strings.Index(line, "'")+1 : strings.LastIndex(line, "'")]
		}

		if strings.Contains(line, "proxy_stopcmd=") {
			config.Proxy_stopcmd = line[strings.Index(line, "'")+1 : strings.LastIndex(line, "'")]
		}

		if strings.Contains(line, "dnsmasq_log_enable=") {
			config.Dnsmasq_log_enable = line[strings.Index(line, "'")+1:strings.LastIndex(line, "'")] == "true"
		}

		if strings.Contains(line, "chinadns_verbose=") {
			config.Chinadns_verbose = line[strings.Index(line, "'")+1:strings.LastIndex(line, "'")] == "true"
		}

		if strings.Contains(line, "dns2tcp_verbose=") {
			config.Dns2tcp_verbose = line[strings.Index(line, "'")+1:strings.LastIndex(line, "'")] == "true"
		}

		if strings.Contains(line, "file_ignlist_ext=") {
			config.File_ignlist_ext = line[strings.Index(line, "'")+1 : strings.LastIndex(line, "'")]
		}

		if err != nil {
			if err == io.EOF {
				break
			} else {
				return 500, "Read config error:" + err.Error()
			}
		}
	}

	jsonBytes, err := json.Marshal(config)
	if err != nil {
		return 500, err.Error()
	}

	return 200, string(jsonBytes)
}
