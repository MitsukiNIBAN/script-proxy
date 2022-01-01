package support

import (
	"bufio"
	"io/ioutil"
	"os"
	"path"
)

var kvCache = make(map[string]string)

func GetC(key string) string {
	data := kvCache[key]
	if data == "" {
		localData := getC(key)
		kvCache[key] = localData
		return localData
	}
	return data
}

func PutC(key string, value string) {
	kvCache[key] = value
	go putC(key, value)
}

func getC(key string) string {
	currentFolder, _ := os.Getwd()
	targetFolder := path.Join(currentFolder, ConfigFolder)

	file, err := os.Open(path.Join(targetFolder, key))
	if err != nil {
		return ""
	}
	defer file.Close()
	content, err := ioutil.ReadAll(file)
	if err != nil {
		return ""
	}
	return string(content)
}

func putC(key string, value string) {
	currentFolder, _ := os.Getwd()

	targetFolder := path.Join(currentFolder, ConfigFolder)
	if !Exists(targetFolder) {
		os.MkdirAll(targetFolder, os.ModePerm)
	}

	file, err := os.OpenFile(path.Join(targetFolder, key), os.O_WRONLY|os.O_TRUNC|os.O_CREATE, 0777)
	if err != nil {
		return
	}
	defer file.Close()

	w := bufio.NewWriter(file)
	w.WriteString(value)
	w.Flush()
}

//协定好的key
//sub_cache_folder:所有配置缓存目录
func SubCacheFolder() string {
	return GetC("sub_cache_folder")
}

//v2ray_sub_url:v2ray订阅地址
func V2raySubUrl() string {
	return GetC("v2ray_sub_url")
}

//v2ray:当前v2ray应用的配置文件
func V2rayCurrentProxy() string {
	return GetC("v2ray")
}

//v2ray_config_file:v2ray的config文件
func V2rayConfigFile() string {
	return GetC("v2ray_config_file")
}

//ssr_sub_url:ssr订阅地址
func SsrSubUrl() string {
	return GetC("ssr_sub_url")
}

//ssr:当前ssr应用的配置文件
func SsrCurrentProxy() string {
	return GetC("ssr")
}

//ssr_config_file:ssr的config文件
func SsrConfigFile() string {
	return GetC("ssr_config_file")
}

//start_script_file:启动脚本路径
func StartScriptFile() string {
	return GetC("start_script_file")
}

//stop_script_file:停用脚本路径
func StopScriptFile() string {
	return GetC("stop_script_file")
}

//v2ray_access_log_file
func AccessLogFile() string {
	return GetC("v2ray_access_log_files")
}

//v2ray_error_log_file
func ErrorLogFile() string {
	return GetC("v2ray_error_log_file")
}

//v2ray_log_level
func V2rayLogLevel() string {
	return GetC("v2ray_log_level")
}
