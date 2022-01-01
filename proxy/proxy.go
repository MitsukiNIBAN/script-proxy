package proxy

import (
	"encoding/json"
	"errors"
	"mmm3w/go-proxy/support"
	"path"
)

func GetPid(t string) (string, error) {
	if t == "ssr" {
		return getSSRPid()
	} else if t == "v2ray" {
		return getV2rayPid()
	} else if t == "script" {
		return getScriptInfo()
	}
	return "", errors.New("错误的类型:" + t)
}

func StartUp(t string) (string, error) {
	if t == "ssr" {
		err := startUpSsr()
		if err != nil {
			return "", err
		}
		return getSSRPid()
	} else if t == "v2ray" {
		err := startUpV2ray()
		if err != nil {
			return "", err
		}
		return getV2rayPid()
	} else if t == "script" {
		err := startScript()
		if err != nil {
			return "", err
		}
		return getScriptInfo()
	}

	return "", errors.New("错误的类型:" + t)
}

func StopProxy(t string, pid string) error {
	if t == "ssr" {
		return killProcess(pid)
	} else if t == "v2ray" {
		return killProcess(pid)
	} else if t == "script" {
		return stopScript()
	}
	return errors.New("错误的类型:" + t)
}

func ApplyProxyConfig(t string, data string) error {
	if len(data) <= 0 {
		return errors.New("缺少数据")
	}

	var kv map[string]string
	json.Unmarshal([]byte(data), &kv)
	key := kv["id"]

	cpath := path.Join(support.SubCacheFolder(), key+".json")
	if !support.Exists(cpath) {
		return errors.New("配置文件不存在，请更新订阅")
	}

	if t == "v2ray" {
		err := applyV2rayConfig(cpath)
		if err != nil {
			return err
		}
		support.PutC("v2ray", data)
		return nil
	} else if t == "ssr" {
		applySsrConfig(cpath)
		support.PutC("ssr", data)
		return nil
	}

	return errors.New("错误的类型")
}

func ForwardSwitch(tag string) error {
	return forwardSwitch(tag == "1")
}
