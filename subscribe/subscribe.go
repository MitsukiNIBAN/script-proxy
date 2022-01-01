package subscribe

import (
	"errors"
	"fmt"
	"mmm3w/go-proxy/support"
	"os"
	"path"
)

func UpdateConfig(t string) error {
	cacheFolder := support.SubCacheFolder()
	if len(cacheFolder) <= 0 {
		return errors.New("未配置缓存目录")
	}
	if !support.Exists(cacheFolder) {
		os.MkdirAll(cacheFolder, os.ModePerm)
	}

	if t == "v2ray" {
		v2rayUrl := support.V2raySubUrl()
		if len(v2rayUrl) > 0 {
			return updateV2raySub(v2rayUrl, cacheFolder)
		} else {
			return fmt.Errorf("缺少v2ray订阅地址(type:%s)", t)
		}
	} else if t == "ssr" {
		ssrUrl := support.SsrSubUrl()
		if len(ssrUrl) > 0 {
			return updateSsrSub(ssrUrl, cacheFolder)
		} else {
			return fmt.Errorf("缺少ssr订阅地址.(type:%s)", t)
		}
	}

	return fmt.Errorf("错误类型:%s", t)
}

func GetConfigSet(t string) (string, error) {
	cacheFolder := support.SubCacheFolder()
	if len(cacheFolder) <= 0 {
		return "[]", errors.New("未配置缓存目录")
	}

	if t == "v2ray" {
		return getConfigSet(path.Join(cacheFolder, support.V2rayConfigListCache))
	}

	if t == "ssr" {
		return getConfigSet(path.Join(cacheFolder, support.SsrConfigListCache))
	}

	return "[]", fmt.Errorf("错误类型:%s", t)
}
