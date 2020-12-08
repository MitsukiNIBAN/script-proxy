package main

import (
	"bytes"
	"io/ioutil"
	"os/exec"
	"strings"
)

func startV2ray() error {
	// path, err := ioutil.ReadFile("./" + V2rayTempFile)
	// if err != nil {
	// 	return err
	// }
	// if len(string(path)) <= 0 {
	// 	panic("缺少v2ray配置路径")
	// }
	// cmd := exec.Command("nohup", "v2ray", "--config="+string(path), ">/dev/null", "2>&1", "&")
	cmd := exec.Command("sudo", "systemctl", "start", "v2ray")
	return cmd.Run()
}

func stopV2ray(pid string) error {
	// cmd := exec.Command("sudo", "kill", "-9", pid)
	cmd := exec.Command("sudo", "systemctl", "stop", "v2ray")
	return cmd.Run()
}

func isV2rayRunning() string {
	var outInfo bytes.Buffer
	cmd := exec.Command("sudo", "pidof", "v2ray")
	cmd.Stdout = &outInfo
	err := cmd.Run()
	if err != nil {
		return ""
	}
	return strings.TrimSpace(outInfo.String())
}

func obtainV2rayConfigPath() (int, string) {
	content, err := ioutil.ReadFile("./" + V2rayTempFile)
	if err != nil {
		return 500, "获取配置路径失败:" + err.Error()
	}
	return 200, string(content)
}

func saveV2rayConfigPath(path string) (int, string) {
	if len(path) == 0 {
		return 403, "配置保存失败:未填写配置路径"
	} else {
		err := ioutil.WriteFile("./"+V2rayTempFile, []byte(path), 0644)
		if err != nil {
			return 500, "配置保存失败:" + err.Error()
		}
	}
	return 200, ""
}

func obtainV2rayStatus() (int, string) {
	return 200, isV2rayRunning()
}

func controlV2ray(isStartUp bool) (int, string) {
	var msg string
	if isStartUp {
		msg = "进程已启动"
	} else {
		msg = "进程已停止"
	}

	pid := isV2rayRunning()
	if (len(pid) > 0) == isStartUp {
		return 200, msg
	}

	if isStartUp {
		err := startV2ray()
		if err != nil {
			return 500, "启动进程失败:" + err.Error()
		}
	} else {
		err := stopV2ray(pid)
		if err != nil {
			return 500, "停止进程失败:" + err.Error()
		}
	}

	return 200, msg
}
