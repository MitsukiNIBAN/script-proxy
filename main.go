package main

import (
	"fmt"
	"net/http"
)

func subConfig(w http.ResponseWriter, r *http.Request) {
	var code int
	var message string
	if r.Method == "POST" { //POST用于提交配置信息
		r.ParseForm()
		code, message = saveSubConfig(r.Form["url"][0], r.Form["path"][0])
	} else if r.Method == "GET" { //GET用于获取配置信息
		code, message = obtainSubConfig()
	} else {
		code, message = 403, "Error"
	}
	w.WriteHeader(code)
	fmt.Fprintf(w, string(message))
}

func updateSub(w http.ResponseWriter, r *http.Request) {
	var code int
	var message string
	if r.Method == "GET" {
		code, message = updateConfig()
	} else {
		code, message = 403, "Error"
	}
	w.WriteHeader(code)
	fmt.Fprintf(w, string(message))
}

func main() {
	http.Handle("/", http.FileServer(http.Dir("./"))) //首页路由
	http.HandleFunc("/subConfig", subConfig)          //保存订阅配置
	http.HandleFunc("/updateSub", updateSub)          //更新订阅
	http.ListenAndServe(":80", nil)                   // 设置监听的端口
}
