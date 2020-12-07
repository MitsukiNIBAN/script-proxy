package main

import (
	"fmt"
	"net/http"
	"strings"
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
	w.Header().Set("Access-Control-Allow-Origin", "*")
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
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.WriteHeader(code)
	fmt.Fprintf(w, string(message))
}

func tproxyConfigPath(w http.ResponseWriter, r *http.Request) {
	var code int
	var message string
	if r.Method == "GET" {
		r.ParseForm()
		code, message = obtainConfigPath()
	} else {
		code, message = 403, "Error"
	}
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.WriteHeader(code)
	fmt.Fprintf(w, string(message))
}

func tproxyConfig(w http.ResponseWriter, r *http.Request) {
	var code int
	var message string
	if r.Method == "GET" {
		r.ParseForm()
		code, message = obtainConfig(r.Form["path"][0])
	} else if r.Method == "POST" {
		r.ParseForm()
		code, message = saveConf(r.Form["data"][0])
	} else {
		code, message = 403, "Error"
	}
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.WriteHeader(code)
	fmt.Fprintf(w, string(message))
}

func tproxyStatus(w http.ResponseWriter, r *http.Request) {
	var code int
	var message string
	if r.Method == "GET" {
		code, message = obtainStatus()
	} else {
		code, message = 403, "Error"
	}
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.WriteHeader(code)
	fmt.Fprintf(w, string(message))
}

func tproxyControl(w http.ResponseWriter, r *http.Request) {
	var code int
	var message string
	if r.Method == "POST" {
		code, message = controlScript(strings.Contains(r.URL.String(), "Start"))
	} else {
		code, message = 403, "Error"
	}
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.WriteHeader(code)
	fmt.Fprintf(w, string(message))
}

func main() {
	http.Handle("/", http.FileServer(http.Dir("./")))      //首页路由
	http.HandleFunc("/subConfig", subConfig)               //保存订阅配置
	http.HandleFunc("/updateSub", updateSub)               //更新订阅
	http.HandleFunc("/tproxyConfigPath", tproxyConfigPath) //stp配置文件
	http.HandleFunc("/tproxyConfig", tproxyConfig)         //获取保存配置
	http.HandleFunc("/tproxyStatus", tproxyStatus)         //获取脚本状态
	http.HandleFunc("/tproxyStart", tproxyControl)         //启动脚本
	http.HandleFunc("/tproxyStop", tproxyControl)          //关闭脚本

	//脚本是否正在运行中
	//脚本运行状态信息
	//主要执行脚本后针对返回的信息进行解析相关的数据

	http.ListenAndServe(":80", nil) // 设置监听的端口
}
