package main

import (
	"bufio"
	"bytes"
	"encoding/json"
	"io"
	"io/ioutil"
	"os"
	"os/exec"
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

//启动
func startScript() error {
	cmd := exec.Command("sudo", "ss-tproxy", "start")
	return cmd.Run()
}

//关闭
func stopScript() error {
	cmd := exec.Command("sudo", "ss-tproxy", "stop")
	return cmd.Run()
}

//运行状态
func scriptStatus() (map[string]bool, error) {
	var outInfo bytes.Buffer
	cmd := exec.Command("sudo", "ss-tproxy", "status")
	cmd.Stdout = &outInfo
	err := cmd.Run()
	if err != nil {
		return nil, err
	}
	var temp = make(map[string]bool)
	for _, s := range strings.Split(outInfo.String(), "\n") {
		if strings.Contains(s, "mode") {
			continue
		}
		var item = strings.Split(s, ":")
		if len(item) != 2 {
			continue
		}
		temp[item[0]] = strings.Contains(item[1], "running")
	}
	return temp, nil
}

//是否在正在运行
func isRunning() (bool, error) {
	status, err := scriptStatus()
	if len(status) != 4 {
		return false, err
	}
	var count = 0
	for k, v := range status {
		if strings.Contains(k, "pxy") {
			continue
		}
		if v {
			count++
		}
	}
	return count == 3, nil
}

//获取配置路径
func obtainConfigPath() (int, string) {
	content, err := ioutil.ReadFile("./" + StpcTempFile)
	if err != nil {
		return 500, "获取配置路径失败:" + err.Error()
	}
	return 200, string(content)
}

//读取配置
func obtainConfig(path string) (int, string) {
	err := ioutil.WriteFile("./"+StpcTempFile, []byte(path), 0644)
	if err != nil {
		return 500, "配置读取失败:" + err.Error()
	}

	file, err := os.OpenFile(path, os.O_RDWR, 0666)
	if err != nil {
		return 500, "配置读取失败:" + err.Error()
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
				return 500, "配置读取失败:" + err.Error()
			}
		}
	}

	jsonBytes, err := json.Marshal(config)
	if err != nil {
		return 500, "配置读取失败:" + err.Error()
	}

	return 200, string(jsonBytes)
}

//控制启停
func controlScript(isStartUp bool) (int, string) {
	var msg string
	if isStartUp {
		msg = "脚本已启动"
	} else {
		msg = "脚本已停止"
	}

	runStatus, err := isRunning()
	if err != nil {
		return 500, "脚本运行状态获取失败:" + err.Error()
	}
	if runStatus == isStartUp {
		return 200, msg
	}

	if isStartUp {
		err := startScript()
		if err != nil {
			return 500, "启动脚本失败:" + err.Error()
		}
	} else {
		err := stopScript()
		if err != nil {
			return 500, "停止脚本失败:" + err.Error()
		}
	}

	return 200, msg
}

//获取状态
func obtainStatus() (int, string) {
	status, err := scriptStatus()
	if err != nil {
		return 500, "获取状态失败:" + err.Error()
	}
	jsonBytes, err := json.Marshal(status)
	if err != nil {
		return 500, "获取状态失败:" + err.Error()
	}
	return 200, string(jsonBytes)
}

//保存配置
func saveConf(data string) (int, string) {
	runningStatus, err := isRunning()
	if err != nil {
		return 500, "保存配置失败:" + err.Error()
	}

	if runningStatus {
		return 500, "保存配置失败:脚本正在运行中"
	}

	filePath, err := ioutil.ReadFile("./" + StpcTempFile)
	if err != nil {
		return 500, "保存配置失败:" + err.Error()
	}

	var config Config
	err = json.Unmarshal([]byte(data), &config)
	if err != nil {
		return 500, "保存配置失败:" + err.Error()
	}

	f, err := os.OpenFile(string(filePath), os.O_RDONLY, 0644)
	if err != nil {
		f.Close()
		return 500, "保存配置失败:" + err.Error()
	}
	reader := bufio.NewReader(f)

	var outStr = ""
	for {
		line, err := reader.ReadString('\n')
		line = strings.TrimSpace(line)

		if strings.Contains(line, "tcponly=") {
			var temp = ""
			if config.Tcponly {
				temp = "true"
			} else {
				temp = "false"
			}
			line = line[0:strings.Index(line, "'")+1] + temp + line[strings.LastIndex(line, "'"):len(line)]
		}

		if strings.Contains(line, "selfonly=") {
			var temp = ""
			if config.Selfonly {
				temp = "true"
			} else {
				temp = "false"
			}
			line = line[0:strings.Index(line, "'")+1] + temp + line[strings.LastIndex(line, "'"):len(line)]
		}

		if strings.Contains(line, "proxy_svraddr4=") {
			line = line[0:strings.Index(line, "(")+1] + config.Proxy_svraddr4 + line[strings.LastIndex(line, ")"):len(line)]
		}

		if strings.Contains(line, "proxy_svrport=") {
			line = line[0:strings.Index(line, "'")+1] + config.Proxy_svrport + line[strings.LastIndex(line, "'"):len(line)]
		}

		if strings.Contains(line, "proxy_startcmd=") {
			line = line[0:strings.Index(line, "'")+1] + config.Proxy_startcmd + line[strings.LastIndex(line, "'"):len(line)]
		}

		if strings.Contains(line, "proxy_stopcmd=") {
			line = line[0:strings.Index(line, "'")+1] + config.Proxy_stopcmd + line[strings.LastIndex(line, "'"):len(line)]
		}

		if strings.Contains(line, "dnsmasq_log_enable=") {
			var temp = ""
			if config.Dnsmasq_log_enable {
				temp = "true"
			} else {
				temp = "false"
			}
			line = line[0:strings.Index(line, "'")+1] + temp + line[strings.LastIndex(line, "'"):len(line)]
		}

		if strings.Contains(line, "chinadns_verbose=") {
			var temp = ""
			if config.Chinadns_verbose {
				temp = "true"
			} else {
				temp = "false"
			}
			line = line[0:strings.Index(line, "'")+1] + temp + line[strings.LastIndex(line, "'"):len(line)]
		}

		if strings.Contains(line, "dns2tcp_verbose=") {
			var temp = ""
			if config.Dns2tcp_verbose {
				temp = "true"
			} else {
				temp = "false"
			}
			line = line[0:strings.Index(line, "'")+1] + temp + line[strings.LastIndex(line, "'"):len(line)]
		}

		outStr = outStr + line + "\n"

		if err != nil {
			if err == io.EOF {
				break
			} else {
				return 500, "保存配置失败:" + err.Error()
			}
		}
	}
	f.Close()

	f, err = os.OpenFile(string(filePath), os.O_WRONLY|os.O_TRUNC, 0600)
	if err != nil {
		return 500, "保存配置失败:" + err.Error()
	}
	defer f.Close()

	writer := bufio.NewWriter(f)
	_, err = writer.Write([]byte(outStr))
	if err != nil {
		return 500, "保存配置失败:" + err.Error()
	}
	writer.Flush()
	return 200, ""
}
