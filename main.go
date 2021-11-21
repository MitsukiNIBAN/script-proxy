package main

import (
	"fmt"
	"mmm3w/go-proxy/api"
	"mmm3w/go-proxy/support"
	"net/http"
	"os"
	"path"
)

func main() {
	fmt.Println("start...")
	currentFolder, _ := os.Getwd()
	fmt.Println("load config file...")

	configData, err := support.RepairConfig(path.Join(currentFolder, support.ConfigFileName))

	if err != nil {
		fmt.Println("load config fail:" + err.Error())
		return
	}

	//这里可能还需要一些启动代理额内容

	//相关配置信息
	//需要包含 v2ray订阅地址，v2ray配置存放地址，ssr订阅地址，ssr配置存放地址
	http.HandleFunc("/status", api.ProxyStatus)                      //各项代理状态（路由是否已配置，v2ray代理进程是否已启动，ssr代理进程是否已启动）
	http.HandleFunc("/enableRoute", api.EnableRouteConfig)           //应用或清除iptables相关配置（清楚或应用路由配置，返回各项代理状态数据，后期区分两个模式，游戏模式能够支持游戏udp代理但不支持本机代理，本机模式支持本机代理，但是不支持游戏udp流量代理）
	http.HandleFunc("/killProxyProcess", api.KillProxyProcess)       //关闭相关代理进程（传入pid关闭相关进程，返回各项代理状态数据）
	http.HandleFunc("/startUpProxyProcess", api.StartUpProxyProcess) //开启相关代理进程（传入key启动相关进程，返回各项代理状态数据）
	http.HandleFunc("/saveConfig", api.SaveConfig)                   //保存各项配置（键值对）
	http.HandleFunc("/getConfig", api.GetConfig)                     //获取各项配置（键值对）
	http.HandleFunc("/getLog", api.GetLog)                           //获取日志
	http.HandleFunc("/clearLog", api.ClearLog)                       //删除日志
	http.HandleFunc("/getProxyConfig", api.GetProxyConfig)           //获取代理的所有配置信息
	http.HandleFunc("/updateSub", api.UpdateSub)                     //更新订阅（请求订阅地址获取相关配置，还缺少ssr的订阅获取）
	http.HandleFunc("/applyConfig", api.ApplyConfig)                 //应用相关配置

	//能切换 v2ray or ssr 配置
	//需要能更新订阅，获取配置列表，修改config配置，至于重启手动来就行

	http.Handle("/", http.FileServer(http.Dir(configData["index"]))) //首页路由
	fmt.Println("start listen port...")
	err = http.ListenAndServe(configData["port"], nil) // 设置监听的端口
	if err != nil {
		fmt.Println(err.Error())
	}
}
