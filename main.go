package main

import (
	"fmt"
	"mmm3w/go-proxy/api"
	"mmm3w/go-proxy/proxy"
	"mmm3w/go-proxy/support"
	"net/http"
	"os"
	"path"
)

func main() {
	fmt.Println("start...")
	currentFolder, _ := os.Getwd()
	fmt.Println("load config file...")

	configData, err := support.LoadConfig(path.Join(currentFolder, support.ServerConf))

	if err != nil {
		fmt.Println("load config fail:" + err.Error())
		return
	}

	fmt.Println("start proxy")
	temp, _ := proxy.GetPid("script")
	if temp == "" {
		proxy.StartUp("script")
	}
	temp, _ = proxy.GetPid("v2ray")
	if temp == "" {
		proxy.StartUp("v2ray")
	}
	temp, _ = proxy.GetPid("ssr")
	if temp == "" {
		proxy.StartUp("ssr")
	}

	//这里可能还需要一些启动代理额内容

	http.HandleFunc("/config", api.HandleConfig)                   //处理配置
	http.HandleFunc("/updateSub", api.UpdateSub)                   //更新订阅（请求订阅地址获取相关配置, ssr的未实现）
	http.HandleFunc("/proxyConfig", api.GetProxyConfig)            //获取代理的所有配置信息
	http.HandleFunc("/currentProxyConfig", api.CurrentProxyConfig) //获取当前正在使用的代理信息
	http.HandleFunc("/proxyRunInfo", api.ProxyRunInfo)             //获取代理信息
	http.HandleFunc("/startProxy", api.StartProxy)                 //启用相关代理内容
	http.HandleFunc("/stopProxy", api.StopProxy)                   //停用相关代理内容
	http.HandleFunc("/applyConfig", api.ApplyConfig)               //应用配置
	http.HandleFunc("/justForward", api.JustForward)               //仅转发

	http.Handle("/", http.FileServer(http.Dir(configData["index"]))) //首页路由
	fmt.Println("start listen port...")
	err = http.ListenAndServe(configData["port"], nil) // 设置监听的端口
	if err != nil {
		fmt.Println(err.Error())
	}
}
