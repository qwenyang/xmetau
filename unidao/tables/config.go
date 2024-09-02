package tables

import (
	"encoding/json"
	"io/ioutil"
	"log"
)

var configPath string = "/www/wwwroot/xmetau/conf/config.json"

type MainServiceConfig struct {
	DataBaseName     string `json:"database_name"`
	DataBaseUserName string `json:"database_user_name"`
	DataBasePassword string `json:"database_password"`
	DataBaseIP       string `json:"database_ip"`
	DataBasePort     int    `json:"database_port"`

	WxChessAppId               string `json:"wx_chess_appid"`
	WxChessAppSecret           string `json:"wx_chess_appsecret"`
	WxBilliardAppId            string `json:"wx_billiard_appid"`
	WxBilliardAppSecret        string `json:"wx_billiard_appsecret"`
	WxTankAppId                string `json:"wx_tank_appid"`
	WxTankAppSecret            string `json:"wx_tank_appsecret"`
	ByteDanceBilliardAppId     string `json:"bytedance_billiard_appid"`
	ByteDanceBilliardAppSecret string `json:"bytedance_billiard_appsecret"`
}

var MainConfig *MainServiceConfig

// 从配置文件中载入json字符串
func LoadConfig(path string) *MainServiceConfig {
	buf, err := ioutil.ReadFile(path)
	if err != nil {
		log.Panicln("load config conf failed: ", err)
	}
	mainConfig := &MainServiceConfig{}
	err = json.Unmarshal(buf, mainConfig)
	if err != nil {
		log.Panicln("decode config file failed:", string(buf), err)
	}

	return mainConfig
}

// 初始化，只能运行一次
func init() {
	MainConfig = LoadConfig(configPath)
}
