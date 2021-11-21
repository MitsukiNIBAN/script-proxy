package api

import (
	"encoding/json"
	"mmm3w/go-proxy/proxy"
	"mmm3w/go-proxy/subscribe"
	"mmm3w/go-proxy/support"
	"net/http"
	"net/url"
	"os"
	"path"
	"strconv"
)

func SaveConfig(w http.ResponseWriter, r *http.Request) {
	support.Post(w, r, func(data url.Values) (int, string) {
		currentFolder, _ := os.Getwd()
		targetFile := path.Join(currentFolder, support.ConfigTempFile)
		source := make(map[string]string)
		for k := range data {
			v := support.GetValue(data, k, "")
			if len(v) > 0 {
				source[k] = v
			}
		}
		err := support.AppendConfig(targetFile, source)
		if err != nil {
			return 500, err.Error()
		}
		return 200, "success"
	})
}

func GetConfig(w http.ResponseWriter, r *http.Request) {
	support.GetJson(w, r, func(data url.Values) (int, string) {
		currentFolder, _ := os.Getwd()
		targetFile := path.Join(currentFolder, support.ConfigTempFile)
		source, err := support.LoadConfig(targetFile)
		if err != nil {
			return 500, err.Error()
		}
		jsonStr, err := json.Marshal(source)
		if err != nil {
			return 500, err.Error()
		}
		return 200, string(jsonStr)
	})
}

func ProxyStatus(w http.ResponseWriter, r *http.Request) {
	support.GetJson(w, r, func(form url.Values) (int, string) {
		data, err := proxy.ComponentStatus()
		if err != nil {
			return 500, err.Error()
		}
		jsonStr, err := json.Marshal(data)
		if err != nil {
			return 500, err.Error()
		}
		return 200, string(jsonStr)
	})
}

func EnableRouteConfig(w http.ResponseWriter, r *http.Request) {
	support.PostJson(w, r, func(form url.Values) (int, string) {
		isEnable, _ := strconv.ParseBool(support.GetValue(form, "isEnable", "false"))

		err := proxy.ApplyRule(isEnable)
		if err != nil {
			return 500, err.Error()
		}

		data, err := proxy.ComponentStatus()
		if err != nil {
			return 500, err.Error()
		}
		jsonStr, err := json.Marshal(data)
		if err != nil {
			return 500, err.Error()
		}
		return 200, string(jsonStr)
	})
}

func KillProxyProcess(w http.ResponseWriter, r *http.Request) {
	support.PostJson(w, r, func(form url.Values) (int, string) {
		pid := support.GetValue(form, "pid", "")
		if len(pid) <= 0 {
			return 500, "No pid"
		}
		err := proxy.KillProcess(pid)
		if err != nil {
			return 500, err.Error()
		}

		data, err := proxy.ComponentStatus()
		if err != nil {
			return 500, err.Error()
		}
		jsonStr, err := json.Marshal(data)
		if err != nil {
			return 500, err.Error()
		}
		return 200, string(jsonStr)
	})
}

func StartUpProxyProcess(w http.ResponseWriter, r *http.Request) {
	support.PostJson(w, r, func(form url.Values) (int, string) {
		t := support.GetValue(form, "type", "")

		err := proxy.StartUp(t)
		if err != nil {
			return 500, err.Error()
		}

		data, err := proxy.ComponentStatus()
		if err != nil {
			return 500, err.Error()
		}
		jsonStr, err := json.Marshal(data)
		if err != nil {
			return 500, err.Error()
		}
		return 200, string(jsonStr)
	})
}

func GetLog(w http.ResponseWriter, r *http.Request) {
	support.Get(w, r, func(form url.Values) (int, string) {
		t := support.GetValue(form, "type", "")
		result, err := proxy.GetLog(t)
		if err != nil {
			return 500, err.Error()
		}
		return 200, result
	})
}

func ClearLog(w http.ResponseWriter, r *http.Request) {
	support.Post(w, r, func(form url.Values) (int, string) {
		t := support.GetValue(form, "type", "")
		err := proxy.ClearLog(t)
		if err != nil {
			return 500, err.Error()
		}
		return 200, "sucess"
	})
}

func UpdateSub(w http.ResponseWriter, r *http.Request) {
	support.Post(w, r, func(data url.Values) (int, string) {
		tag := support.GetValue(data, "tag", "")
		if len(tag) <= 0 {
			return 500, "no tag"
		} else {
			i, err := strconv.Atoi(tag)
			if err != nil {
				return 500, err.Error()
			}
			err = subscribe.UpdateConfig(i)
			if err != nil {
				return 500, err.Error()
			} else {
				return 200, "sucess"
			}
		}
	})
}

func GetProxyConfig(w http.ResponseWriter, r *http.Request) {
	support.GetJson(w, r, func(form url.Values) (int, string) {
		t := support.GetValue(form, "type", "")
		data, err := subscribe.GetConfigSet(t)
		if err != nil {
			return 500, err.Error()
		}
		jsonStr, err := json.Marshal(data)
		if err != nil {
			return 500, err.Error()
		}
		return 200, string(jsonStr)
	})
}

func ApplyConfig(w http.ResponseWriter, r *http.Request) {
	support.Post(w, r, func(data url.Values) (int, string) {
		key := support.GetValue(data, "key", "")
		t := support.GetValue(data, "type", "")
		err := proxy.ApplyConfig(t, key)
		if err != nil {
			return 500, err.Error()
		} else {
			return 200, "sucess"
		}
	})
}
