package api

import (
	"fmt"
	"mmm3w/go-proxy/proxy"
	"mmm3w/go-proxy/subscribe"
	"mmm3w/go-proxy/support"
	"net/http"
	"net/url"
)

func HandleConfig(w http.ResponseWriter, r *http.Request) {
	var code int
	var message string
	if r.Method == "POST" {
		fmt.Println("Post coming")
		r.ParseMultipartForm(32 << 20)
		key := support.GetValue(r.PostForm, "key", "")
		value := support.GetValue(r.PostForm, "value", "")
		if key == "" {
			code, message = 500, "缺少Key"
		} else {
			support.PutC(key, value)
			code, message = 200, ""
		}
	} else if r.Method == "GET" {
		fmt.Println("Get coming")
		r.ParseForm()
		queryForm, err := url.ParseQuery(r.URL.RawQuery)
		if err != nil {
			code, message = 500, err.Error()
		} else {
			key := support.GetValue(queryForm, "key", "")
			if key == "" {
				code, message = 500, "缺少Key"
			} else {
				code, message = 200, support.GetC(key)
			}
		}
	} else {
		code, message = 403, "Error"
	}
	w.WriteHeader(code)
	fmt.Fprint(w, message)
}

func UpdateSub(w http.ResponseWriter, r *http.Request) {
	support.Post(w, r, func(data url.Values) (int, string) {
		t := support.GetValue(data, "type", "")
		if len(t) <= 0 {
			return 500, "订阅更新类型错误"
		} else {
			err := subscribe.UpdateConfig(t)
			if err != nil {
				return 500, err.Error()
			} else {
				return 200, "更新成功"
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
		return 200, data
	})
}

func CurrentProxyConfig(w http.ResponseWriter, r *http.Request) {
	support.GetJson(w, r, func(form url.Values) (int, string) {
		key := support.GetValue(form, "key", "")
		if key == "" {
			return 500, "缺少Key"
		} else {
			return 200, support.GetC(key)
		}
	})
}

func ProxyRunInfo(w http.ResponseWriter, r *http.Request) {
	support.Get(w, r, func(form url.Values) (int, string) {
		t := support.GetValue(form, "type", "")
		data, err := proxy.GetPid(t)
		if err != nil {
			return 500, err.Error()
		}
		return 200, data
	})
}

func StartProxy(w http.ResponseWriter, r *http.Request) {
	support.Post(w, r, func(form url.Values) (int, string) {
		t := support.GetValue(form, "type", "")
		data, err := proxy.StartUp(t)
		if err != nil {
			return 500, err.Error()
		}
		return 200, data
	})
}

func StopProxy(w http.ResponseWriter, r *http.Request) {
	support.Post(w, r, func(form url.Values) (int, string) {
		t := support.GetValue(form, "type", "")
		pid := support.GetValue(form, "pid", "")
		err := proxy.StopProxy(t, pid)
		if err != nil {
			return 500, err.Error()
		}
		return 200, ""
	})
}

func ApplyConfig(w http.ResponseWriter, r *http.Request) {
	support.Post(w, r, func(data url.Values) (int, string) {
		d := support.GetValue(data, "data", "")
		t := support.GetValue(data, "type", "")
		err := proxy.ApplyProxyConfig(t, d)
		if err != nil {
			return 500, err.Error()
		} else {
			return 200, "sucess"
		}
	})
}

func JustForward(w http.ResponseWriter, r *http.Request) {
	support.Post(w, r, func(data url.Values) (int, string) {
		t := support.GetValue(data, "tag", "")
		err := proxy.ForwardSwitch(t)
		if err != nil {
			return 500, err.Error()
		} else {
			return 200, "sucess"
		}
	})
}
