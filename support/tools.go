package support

import (
	"bufio"
	"bytes"
	"errors"
	"fmt"
	"io"
	"net/url"
	"os"
	"os/exec"
	"strings"
)

func RepairConfig(path string) (map[string]string, error) {
	data, err := LoadConfig(path)
	if err != nil {
		return data, err
	}
	err = SaveConfig(path, data)
	return data, err
}

//后面的项会覆盖前面的项目
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

		kv := strings.SplitAfterN(line, "=", 2)
		if len(kv) < 2 {
			continue
		}
		data[kv[0][:len(kv[0])-1]] = kv[1]
	}
	return data, nil
}

//会覆盖所有配置
func SaveConfig(path string, data map[string]string) error {
	file, err := os.OpenFile(path, os.O_WRONLY|os.O_TRUNC|os.O_CREATE, 0777)
	if err != nil {
		return err
	}
	defer file.Close()

	w := bufio.NewWriter(file)
	for k, v := range data {
		updateConfigMemoryCache(k, v)
		fmt.Fprintf(w, "%s=%s\n", k, v)
	}
	return w.Flush()
}

//在后面追加数据
func AppendConfig(path string, data map[string]string) error {
	file, err := os.OpenFile(path, os.O_WRONLY|os.O_APPEND|os.O_CREATE, 0777)
	if err != nil {
		return err
	}
	defer file.Close()

	w := bufio.NewWriter(file)
	for k, v := range data {
		updateConfigMemoryCache(k, v)
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
	var stderr bytes.Buffer
	cmd := exec.Command(name, arg...)
	cmd.Stdout = &outInfo
	cmd.Stderr = &stderr
	err := cmd.Run()
	if err != nil {
		errInfo := stderr.String()
		if len(errInfo) > 0 {
			return "", errors.New(err.Error() + ":" + errInfo)
		}
	}
	return outInfo.String(), nil
}
