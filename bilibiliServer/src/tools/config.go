// 配置文件解析
package tools

import (
	"bufio"
	"encoding/json"
	"errors"
	"os"
)

type ConfigType struct {
	DbUsername string `json:"dbUsername"`
	DbPassword string `json:"dbPassword"`
	DbName     string `json:"dbName"`
	Socket     string `json:"socket"`
}

var Config *ConfigType = nil

func ParseConfig(path string) error {
	file, err := os.Open(path)
	if err != nil {
		return errors.New("配置文件打开错误")
	}
	defer file.Close()

	reader := bufio.NewReader(file)
	decoder := json.NewDecoder(reader)

	if err := decoder.Decode(&Config); err != nil {
		return errors.New("配置文件解析错误")
	}
	return nil
}
