package api

import (
	"encoding/json"
	"fmt"
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
		source, _ := support.LoadConfig(targetFile)
		for k := range data {
			v := support.GetValue(data, k, "")
			if len(v) > 0 {
				source[k] = v
			}
		}
		fmt.Println(source)
		err := support.SaveConfig(targetFile, source)
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

func UpdateSub(w http.ResponseWriter, r *http.Request) {
	support.Post(w, r, func(data url.Values) (int, string) {
		tag := support.GetValue(data, "tag", "")
		if len(tag) <= 0 {
			return subscribe.UpdateConfig(0)
		} else {
			i, err := strconv.Atoi(tag)
			if err != nil {
				return 500, err.Error()
			}
			return subscribe.UpdateConfig(i)
		}
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
