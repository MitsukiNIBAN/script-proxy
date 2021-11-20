package main

import (
	"fmt"
	"mmm3w/go-proxy/api"
	"mmm3w/go-proxy/support"
	"net/http"
	"os"
	"path"
)

// func subConfig(w http.ResponseWriter, r *http.Request) {
// 	var code int
// 	var message string
// 	if r.Method == "POST" { //POST用于提交配置信息
// 		r.ParseForm()
// 		code, message = saveSubConfig(r.Form["url"][0], r.Form["path"][0])
// 	} else if r.Method == "GET" { //GET用于获取配置信息
// 		code, message = obtainSubConfig()
// 	} else {
// 		code, message = 403, "Error"
// 	}
// 	w.Header().Set("Access-Control-Allow-Origin", "*")
// 	w.WriteHeader(code)
// 	fmt.Fprintf(w, string(message))
// }

// func updateSub(w http.ResponseWriter, r *http.Request) {
// 	var code int
// 	var message string
// 	if r.Method == "GET" {
// 		code, message = updateConfig()
// 	} else {
// 		code, message = 403, "Error"
// 	}
// 	w.Header().Set("Access-Control-Allow-Origin", "*")
// 	w.WriteHeader(code)
// 	fmt.Fprintf(w, string(message))
// }

// func tproxyConfigPath(w http.ResponseWriter, r *http.Request) {
// 	var code int
// 	var message string
// 	if r.Method == "GET" {
// 		code, message = obtainConfigPath()
// 	} else {
// 		code, message = 403, "Error"
// 	}
// 	w.Header().Set("Access-Control-Allow-Origin", "*")
// 	w.WriteHeader(code)
// 	fmt.Fprintf(w, string(message))
// }

// func tproxyConfig(w http.ResponseWriter, r *http.Request) {
// 	var code int
// 	var message string
// 	if r.Method == "GET" {
// 		r.ParseForm()
// 		code, message = obtainConfig(r.Form["path"][0])
// 	} else if r.Method == "POST" {
// 		r.ParseForm()
// 		code, message = saveConf(r.Form["data"][0])
// 	} else {
// 		code, message = 403, "Error"
// 	}
// 	w.Header().Set("Access-Control-Allow-Origin", "*")
// 	w.WriteHeader(code)
// 	fmt.Fprintf(w, string(message))
// }

// func tproxyStatus(w http.ResponseWriter, r *http.Request) {
// 	var code int
// 	var message string
// 	if r.Method == "GET" {
// 		code, message = obtainStatus()
// 	} else {
// 		code, message = 403, "Error"
// 	}
// 	w.Header().Set("Access-Control-Allow-Origin", "*")
// 	w.WriteHeader(code)
// 	fmt.Fprintf(w, string(message))
// }

// func tproxyControl(w http.ResponseWriter, r *http.Request) {
// 	var code int
// 	var message string
// 	if r.Method == "POST" {
// 		code, message = controlScript(strings.Contains(r.URL.String(), "Start"))
// 	} else {
// 		code, message = 403, "Error"
// 	}
// 	w.Header().Set("Access-Control-Allow-Origin", "*")
// 	w.WriteHeader(code)
// 	fmt.Fprintf(w, string(message))
// }

// func v2rayConfigPath(w http.ResponseWriter, r *http.Request) {
// 	var code int
// 	var message string
// 	if r.Method == "GET" {
// 		code, message = obtainV2rayConfigPath()
// 	} else if r.Method == "POST" {
// 		r.ParseForm()
// 		code, message = saveV2rayConfigPath(r.Form["path"][0])
// 	} else {
// 		code, message = 403, "Error"
// 	}
// 	w.Header().Set("Access-Control-Allow-Origin", "*")
// 	w.WriteHeader(code)
// 	fmt.Fprintf(w, string(message))
// }

// func v2rayConfigInfo(w http.ResponseWriter, r *http.Request) {
// 	var code int
// 	var message string
// 	if r.Method == "GET" {
// 		code, message = obtainCurrentProxyName()
// 	} else {
// 		code, message = 403, "Error"
// 	}
// 	w.Header().Set("Access-Control-Allow-Origin", "*")
// 	w.WriteHeader(code)
// 	fmt.Fprintf(w, string(message))
// }

// func v2rayStatus(w http.ResponseWriter, r *http.Request) {
// 	var code int
// 	var message string
// 	if r.Method == "GET" {
// 		code, message = obtainV2rayStatus()
// 	} else {
// 		code, message = 403, "Error"
// 	}
// 	w.Header().Set("Access-Control-Allow-Origin", "*")
// 	w.WriteHeader(code)
// 	fmt.Fprintf(w, string(message))
// }

// func v2rayControl(w http.ResponseWriter, r *http.Request) {
// 	var code int
// 	var message string
// 	if r.Method == "POST" {
// 		code, message = controlV2ray(strings.Contains(r.URL.String(), "Start"))
// 	} else {
// 		code, message = 403, "Error"
// 	}
// 	w.Header().Set("Access-Control-Allow-Origin", "*")
// 	w.WriteHeader(code)
// 	fmt.Fprintf(w, string(message))
// }

// func serverSet(w http.ResponseWriter, r *http.Request) {
// 	var code int
// 	var message string
// 	if r.Method == "GET" {
// 		code, message = obtainServerSet()
// 	} else {
// 		code, message = 403, "Error"
// 	}
// 	w.Header().Set("Access-Control-Allow-Origin", "*")
// 	w.WriteHeader(code)
// 	fmt.Fprintf(w, string(message))
// }

// func portSet(w http.ResponseWriter, r *http.Request) {
// 	var code int
// 	var message string
// 	if r.Method == "GET" {
// 		code, message = obtainPortSet()
// 	} else {
// 		code, message = 403, "Error"
// 	}
// 	w.Header().Set("Access-Control-Allow-Origin", "*")
// 	w.WriteHeader(code)
// 	fmt.Fprintf(w, string(message))
// }

// func configSet(w http.ResponseWriter, r *http.Request) {
// 	var code int
// 	var message string
// 	if r.Method == "GET" {
// 		code, message = obtainConfigSet()
// 	} else {
// 		code, message = 403, "Error"
// 	}
// 	w.Header().Set("Access-Control-Allow-Origin", "*")
// 	w.WriteHeader(code)
// 	fmt.Fprintf(w, string(message))
// }

// func v2rayConfig(w http.ResponseWriter, r *http.Request) {
// 	var code int
// 	var message string
// 	if r.Method == "POST" {
// 		r.ParseForm()
// 		code, message = modifyV2rayConfig(r.Form["data"][0])
// 	} else {
// 		code, message = 403, "Error"
// 	}
// 	w.Header().Set("Access-Control-Allow-Origin", "*")
// 	w.WriteHeader(code)
// 	fmt.Fprintf(w, string(message))
// }

func main() {
	fmt.Println("start...")
	currentFolder, _ := os.Getwd()
	fmt.Println("load config file...")

	configData, err := support.LoadConfig(path.Join(currentFolder, support.ConfigFileName))

	if err != nil {
		fmt.Println("load config fail:" + err.Error())
		return
	}

	//相关配置信息
	//需要包含 v2ray订阅地址，v2ray配置存放地址，ssr订阅地址，ssr配置存放地址
	http.HandleFunc("/status", api.ProxyStatus)            //各项代理状态
	http.HandleFunc("/saveConfig", api.SaveConfig)         //保存各项配置
	http.HandleFunc("/getConfig", api.GetConfig)           //获取各项配置
	http.HandleFunc("/updateSub", api.UpdateSub)           //更新订阅
	http.HandleFunc("/enableRoute", api.EnableRouteConfig) //应用或清除iptables相关配置

	http.Handle("/", http.FileServer(http.Dir(configData["index"]))) //首页路由
	fmt.Println("start listen port...")
	err = http.ListenAndServe(configData["port"], nil) // 设置监听的端口
	if err != nil {
		fmt.Println(err.Error())
	}

	//需要重新提供的接口
	//查询代理状态，包括路由表配置情况，v2ray进程情况，ssr进程情况
	//都做成单独的开关

	// fmt.Println("Work folder : " + currentFolder())

	// http.HandleFunc("/updateSub", updateSub)                     //更新订阅
	// http.HandleFunc("/tproxyConfigPath", tproxyConfigPath)       //stp配置文件
	// http.HandleFunc("/tproxyConfig", tproxyConfig)               //获取保存配置
	// http.HandleFunc("/tproxyStatus", tproxyStatus)               //获取脚本状态
	// http.HandleFunc("/tproxyStart", tproxyControl)               //启动脚本
	// http.HandleFunc("/tproxyStop", tproxyControl)                //关闭脚本
	// http.HandleFunc("/v2rayConfigPath", v2rayConfigPath)         //v2ray配置文件
	// http.HandleFunc("/v2rayConfigInfo", v2rayConfigInfo)         //v2ray配置简息
	// http.HandleFunc("/v2rayConfig", v2rayConfig)                 //切换v2ray配置
	// http.HandleFunc("/v2rayStatus", v2rayStatus)                 //v2ray进程状态
	// http.HandleFunc("/v2rayStart", v2rayControl)                 //启动v2ray进程
	// http.HandleFunc("/v2rayStop", v2rayControl)                  //关闭v2ray进程
	// http.HandleFunc("/serverSet", serverSet)                     //尝试获取服务器集合
	// http.HandleFunc("/portSet", portSet)                         //尝试获取端口集合
	// http.HandleFunc("/configSet", configSet)                     //获取配置集合

}
