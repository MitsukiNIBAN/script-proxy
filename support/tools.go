package support

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"net/url"
	"os"
	"os/exec"
	"strings"
)

func LoadConfig(path string) (map[string]string, error) {
	data := make(map[string]string)

	file, err := os.Open(path)
	if err != nil {
		return data, err
	}
	defer file.Close()

	readBuf := bufio.NewReader(file)
	for {
		content, _, err := readBuf.ReadLine()
		line := strings.TrimSpace(string(content))
		if err == io.EOF {
			break
		}
		if strings.HasPrefix(line, "#") {
			continue
		}
		kv := strings.Split(line, "=")
		if len(kv) < 2 {
			continue
		}
		data[kv[0]] = kv[1]
	}
	return data, nil
}

func SaveConfig(path string, data map[string]string) error {
	file, err := os.OpenFile(path, os.O_WRONLY|os.O_TRUNC|os.O_CREATE, 0777)
	if err != nil {
		return err
	}
	defer file.Close()

	w := bufio.NewWriter(file)
	for k, v := range data {
		//注意随时更新一下
		if k == KeySsrSubUrl {
			ssrSubUrl = v
		}
		if k == KeyV2raySubUrl {
			v2raySubUrl = v
		}
		if k == KeySsrJsonCacheFolder {
			ssrJsonCacheFolder = v
		}
		if k == KeyV2rayJsonCacheFolder {
			v2rayJsonCacheFolder = v
		}

		fmt.Fprintf(w, "%s=%s\n", k, v)
	}
	return w.Flush()
}

func Exists(path string) bool {
	_, err := os.Stat(path)
	if err != nil {
		return os.IsExist(err)
	}
	return true
}

func GetValue(values url.Values, key string, def string) string {
	if vs := values[key]; len(vs) > 0 {
		return vs[0]
	} else {
		return def
	}
}

func ExecCommand(name string, arg ...string) (string, error) {
	var outInfo bytes.Buffer
	cmd := exec.Command(name, arg...)
	cmd.Stdout = &outInfo
	err := cmd.Run()
	if err != nil {
		return "", err
	}
	return outInfo.String(), nil
}
