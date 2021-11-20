package support

import (
	"os"
	"path"
)

// import (
// 	"os"
// 	"path"
// 	"path/filepath"
// )

const ConfigFileName string = "server.conf"
const ConfigTempFile string = "config.temp"

const KeySsrSubUrl string = "ssrSubUrl"
const KeyV2raySubUrl string = "v2raySubUrl"
const KeySsrJsonCacheFolder string = "ssrJsonCacheFolder"
const KeyV2rayJsonCacheFolder string = "v2rayJsonCacheFolder"

var ssrSubUrl string = ""
var v2raySubUrl string = ""
var ssrJsonCacheFolder string = ""
var v2rayJsonCacheFolder string = ""

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

func SsrJsonCacheFolder() string {
	if len(ssrJsonCacheFolder) <= 0 {
		source := loadConfigTempFile()
		ssrJsonCacheFolder = source[KeySsrJsonCacheFolder]
	}
	return ssrJsonCacheFolder
}

func V2rayJsonCacheFolder() string {
	if len(v2rayJsonCacheFolder) <= 0 {
		source := loadConfigTempFile()
		v2rayJsonCacheFolder = source[KeyV2rayJsonCacheFolder]
	}
	return v2rayJsonCacheFolder
}

// var tempCurrent string = ""

// func currentFolder() string {
// 	if len(tempCurrent) == 0 {
// 		tempCurrent, _ = filepath.Abs(filepath.Dir(os.Args[0]))
// 	}
// 	return tempCurrent
// }

// func subTempFilePath() string {
// 	return path.Join(currentFolder(), SubTempFile)
// }

// func configSaveFolderTempFilePath() string {
// 	return path.Join(currentFolder(), ConfigSaveFolder)
// }

// func v2rayTempFilePath() string {
// 	return path.Join(currentFolder(), V2rayTempFile)
// }

// func stpcTempFilePath() string {
// 	return path.Join(currentFolder(), StpcTempFile)
// }
