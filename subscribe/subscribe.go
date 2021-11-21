package subscribe

import (
	"errors"
	"fmt"
	"mmm3w/go-proxy/support"
	"os"
	"path"
)

func UpdateConfig(tag int) error {
	cacheFolder := support.JsonCacheFolder()
	if len(cacheFolder) <= 0 {
		return errors.New("no cache folder")
	}

	if tag&1 == 1 {
		v2rayUrl := support.V2raySubUrl()

		if len(v2rayUrl) > 0 {
			if !support.Exists(cacheFolder) {
				os.MkdirAll(cacheFolder, os.ModePerm)
			}
			return updateV2raySub(v2rayUrl, cacheFolder)
		} else {
			return fmt.Errorf("no v2ray url.(tag:%d)", tag)
		}
	}

	if tag&2 == 2 {
		ssrUrl := support.SsrSubUrl()
		if len(ssrUrl) > 0 {
			if !support.Exists(cacheFolder) {
				os.MkdirAll(cacheFolder, os.ModePerm)
			}
			return updateSsrSub(ssrUrl, cacheFolder)
		} else {
			return fmt.Errorf("no ssr url.(tag:%d)", tag)
		}
	}

	return fmt.Errorf("error tag:%d", tag)
}

func GetConfigSet(t string) ([]map[string]string, error) {
	cacheFolder := support.JsonCacheFolder()
	if len(cacheFolder) <= 0 {
		return nil, errors.New("no cache folder")
	}

	if t == "v2ray" {
		return getConfigSet(path.Join(cacheFolder, support.V2rayConfigListCache))
	}

	if t == "ssr" {
		return getConfigSet(path.Join(cacheFolder, support.SsrConfigListCache))
	}

	return nil, errors.New("error type")
}
