package support

import (
	"os"
	"path"
)

const ConfigFileName string = "server.conf"
const ConfigTempFile string = "config.temp"
const V2rayConfigListCache string = "v2ray.temp"
const SsrConfigListCache string = "ssr.temp"

const KeySsrSubUrl string = "ssrSubUrl"
const KeyV2raySubUrl string = "v2raySubUrl"
const KeyProxyJsonCacheFolder = "jsonCacheFolder"
const KeyErrorLogFile = "errorLogFile"
const KeyAccessLogFile = "accessLogFile"
const KeyV2rayLogLevel = "v2rayLogLevel"
const KeyV2rayConfigJsonPath = "v2rayConfigJsonPath"
const KeyCurrentV2rayConfig = "currentV2rayConfig"

var currentV2rayConfig string = ""
var v2rayConfigJsonPath string = ""
var v2rayLogLevel string = ""
var errorLogFile string = ""
var accessLogFile string = ""
var ssrSubUrl string = ""
var v2raySubUrl string = ""
var proxyJsonCacheFolder string = ""

func updateConfigMemoryCache(k string, v string) {
	switch k {
	case KeySsrSubUrl:
		ssrSubUrl = v
	case KeyV2raySubUrl:
		v2raySubUrl = v
	case KeyProxyJsonCacheFolder:
		proxyJsonCacheFolder = v
	case KeyErrorLogFile:
		errorLogFile = v
	case KeyAccessLogFile:
		accessLogFile = v
	case KeyV2rayLogLevel:
		v2rayLogLevel = v
	case KeyV2rayConfigJsonPath:
		v2rayConfigJsonPath = v
	}
}

func loadConfigTempFile() map[string]string {
	currentFolder, _ := os.Getwd()
	targetFile := path.Join(currentFolder, ConfigTempFile)
	data, _ := LoadConfig(targetFile)
	return data
}

func SsrSubUrl() string {
	if len(ssrSubUrl) <= 0 {
		source := loadConfigTempFile()
		ssrSubUrl = source[KeySsrSubUrl]
	}
	return ssrSubUrl
}

func V2raySubUrl() string {
	if len(v2raySubUrl) <= 0 {
		source := loadConfigTempFile()
		v2raySubUrl = source[KeyV2raySubUrl]
	}
	return v2raySubUrl
}

func JsonCacheFolder() string {
	if len(proxyJsonCacheFolder) <= 0 {
		source := loadConfigTempFile()
		proxyJsonCacheFolder = source[KeyProxyJsonCacheFolder]
	}
	return proxyJsonCacheFolder
}

func ErrorLogFile() string {
	if len(errorLogFile) <= 0 {
		source := loadConfigTempFile()
		errorLogFile = source[KeyErrorLogFile]
	}
	return path.Join(errorLogFile, "error.log")
}

func AccessLogFile() string {
	if len(accessLogFile) <= 0 {
		source := loadConfigTempFile()
		accessLogFile = source[KeyAccessLogFile]
	}
	return path.Join(accessLogFile, "access.log")
}

func V2rayLogLevel() string {
	if len(v2rayLogLevel) <= 0 {
		source := loadConfigTempFile()
		v2rayLogLevel = source[KeyV2rayLogLevel]
	}
	return v2rayLogLevel
}

func V2rayConfigJsonPath() string {
	if len(v2rayConfigJsonPath) <= 0 {
		source := loadConfigTempFile()
		v2rayConfigJsonPath = source[KeyV2rayConfigJsonPath]
	}
	return v2rayConfigJsonPath
}

func CurrentV2rayConfig() string {
	if len(currentV2rayConfig) <= 0 {
		source := loadConfigTempFile()
		currentV2rayConfig = source[KeyCurrentV2rayConfig]
	}
	return currentV2rayConfig
}
