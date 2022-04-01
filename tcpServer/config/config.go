package config

import (
	"encoding/json"
	"io/ioutil"
)

const (
	path string = "/home/scottk.zhang/work/et/et/tcpServer/config/server.json"
)

type ServerConfig struct {
	RedisAddr         string `json:"redisAddr"`
	DBIP              string `json:"DBIP"`
	DBPort            int32  `json:"DBPort"`
	DBUserName        string `json:"DBUserName"`
	DBPassword        string `json:"BDPassword"`
	DBName            string `json:"DBName"`
	DBDriverName      string `json:"DBDriverName"`
	ServerAddr        string `json:"serverAddr"`
	SessionExpireTime int64  `json:"sessionExpireTime"`
}

var Config ServerConfig

// loadConfig 加载配置
func loadConfig(path string) error {
	buf, err := ioutil.ReadFile(path)
	if err != nil {
		return err
	}

	err = json.Unmarshal(buf, &Config)
	if err != nil {
		return err
	}
	return nil
}

func Init() error {
	return loadConfig(path)
}
