package proxy

import (
	"mmm3w/go-proxy/support"
)

func ApplyRule(isEnable bool) error {
	var err error
	if isEnable {
		_, err = support.ExecCommand("bash", "./enable.sh")
	} else {
		_, err = support.ExecCommand("bash", "./disable.sh")
	}
	return err
}

func ComponentStatus() (map[string]string, error) {
	var temp = make(map[string]string)

	//iptables相关的配置以ip_forward为准
	str, err := support.ExecCommand("sysctl", "-n", "net.ipv4.ip_forward")
	if err != nil {
		return nil, err
	}
	temp["routeConfig"] = str

	str, err = support.ExecCommand("pidof", "v2ray")
	if err != nil {
		return nil, err
	}
	temp["v2rayPid"] = str
	return temp, nil
}
