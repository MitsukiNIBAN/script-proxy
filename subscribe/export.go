package subscribe

import (
	"mmm3w/go-proxy/support"
	"os"
)

func UpdateConfig(tag int) (int, string) {
	if tag&1 == 1 {
		v2rayUrl := support.V2raySubUrl()
		v2rayFolder := support.V2rayJsonCacheFolder()
		if len(v2rayUrl) > 0 && len(v2rayFolder) > 0 {
			if !support.Exists(v2rayFolder) {
				os.MkdirAll(v2rayFolder, os.ModePerm)
			}
			err := updateV2raySub(v2rayUrl, v2rayFolder)
			if err != nil {
				return 500, err.Error()
			}
		}
	}
	if tag&3 == 3 {
		ssrUrl := support.SsrSubUrl()
		ssrFolder := support.SsrJsonCacheFolder()
		if len(ssrUrl) > 0 && len(ssrFolder) > 0 {
			if !support.Exists(ssrFolder) {
				os.MkdirAll(ssrFolder, os.ModePerm)
			}
			err := updateSsrSub(ssrUrl, ssrFolder)
			if err != nil {
				return 500, err.Error()
			}
		}
	}
	return 200, ""
}
