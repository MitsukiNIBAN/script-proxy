package proxy

import (
	"errors"
	"mmm3w/go-proxy/support"
	"strings"
)

// func InitV2rayConfig(){

// }

func ApplyRule(isEnable bool) error {
	var err error
	if isEnable {
		_, err = support.ExecCommand("bash", "./enable.sh")
	} else {
		_, err = support.ExecCommand("bash", "./disable.sh")
	}
	return err
}

func StartUp(t string) error {
	if t == "ssr" {
		err := startUpSsr()
		if err != nil {
			return err
		}
		return nil
	}
	if t == "v2ray" {
		err := startUpV2ray()
		if err != nil {
			return err
		}
		return nil
	}

	return errors.New("error type")
}

func KillProcess(pid string) error {
	_, err := support.ExecCommand("kill", "-9", pid)
	return err
}

func ComponentStatus() (map[string]string, error) {
	var temp = make(map[string]string)

	// iptables相关的配置以ip_forward为准
	str, err := support.ExecCommand("sysctl", "-n", "net.ipv4.ip_forward")
	if err != nil {
		return nil, err
	}
	temp["routeConfig"] = strings.TrimSpace(str)

	str, err = support.ExecCommand("pidof", "v2ray")
	if err != nil {
		return nil, err
	}
	temp["v2rayPid"] = strings.TrimSpace(str)
	return temp, nil
}

func ClearLog(t string) error {
	if t == "access" {
		_, err := support.ExecCommand("rm", "-rf", support.AccessLogFile())
		if err != nil {
			return err
		}
	}
	if t == "error" {
		_, err := support.ExecCommand("rm", "-rf", support.ErrorLogFile())
		if err != nil {
			return err
		}
	}
	return nil
}

func GetLog(t string) (string, error) {
	if t == "access" {
		return support.ExecCommand("cat", support.AccessLogFile())
	}
	if t == "error" {
		return support.ExecCommand("cat", support.ErrorLogFile())
	}
	return "", nil
}

func ApplyConfig(t string, key string) error {
	if len(key) <= 0 {
		return errors.New("error key")
	}
	if t == "access" {
		return applyV2rayConfig(key)
	}
	if t == "error" {
		return applySsrConfig(key)
	}
	return errors.New("error type")
}
