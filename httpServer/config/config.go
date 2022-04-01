package config

import (
	"encoding/json"
	"io/ioutil"
)

const (
	path string = "/home/scottk.zhang/work/et/et/httpServer/config/server.json"
)

type ServerConfig struct {
	RpcServerIP   string `json:"rpcServerIP"`
	RpcServerPort int32  `json:"rpcServerPort"`
	PicturePath   string `json:"picturePath"`
	PictureUrl    string `json:"pictureUrl"`
	IndexPage     string `json:"indexPage"`
	LoginPage     string `json:"loginPage"`
	UserInfoPage  string `json:"userInfoPage"`
	ConnCap       int32  `json:"connCap"`
	ConnTimeOut   int32  `json:"connTimeOut"`
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
