package main

import (
	"os"
	"path"
	"path/filepath"
)

const StpcTempFile string = "tpc.temp"
const SubTempFile string = "sub.temp"
const ConfigSaveFolder string = "csf.temp"
const V2rayTempFile string = "v2ray.temp"

var tempCurrent string = ""

func currentFolder() string {
	if len(tempCurrent) == 0 {
		tempCurrent, _ = filepath.Abs(filepath.Dir(os.Args[0]))
	}
	return tempCurrent
}

func subTempFilePath() string {
	return path.Join(currentFolder(), SubTempFile)
}

func configSaveFolderTempFilePath() string {
	return path.Join(currentFolder(), ConfigSaveFolder)
}

func v2rayTempFilePath() string {
	return path.Join(currentFolder(), V2rayTempFile)
}

func stpcTempFilePath() string {
	return path.Join(currentFolder(), StpcTempFile)
}
