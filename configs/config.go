//Package configs 配置文件读取
/**
* @Author: Kanesong
* @Date: 2019/4/8 10:35 AM
 */
package configs

import (
	"log"
	"os"
	"time"

	"github.com/go-ini/ini"
)

//Cfg 配置文件
var Cfg *ini.File

//ServerConfig 服务端配置数据结构
type ServerConfig struct {
	RunMode      string
	HTTPPort     int
	ReadTimeout  time.Duration
	WriteTimeout time.Duration
}

// 数据库配置
func LoadDBConfig(section string) *ini.Section {
	var err error
	var configPath = "configs/app.ini"
	_, err = os.Stat(configPath)
	if err != nil {
		configPath = "../configs/app.ini"
	}
	Cfg, err = ini.Load(configPath)
	if err != nil {
		log.Fatal(2, "Fail to parse 'app.ini': %v", err)
	}
	//server配置节点读取
	dbSection, err := Cfg.GetSection(section)
	if err != nil {
		log.Fatal(2, "Fail to get section 'server': %v", err)
	}
	return dbSection
}

//LoadServerConfig 加载服务端配置
func LoadServerConfig() ServerConfig {

	var err error
	var configPath = "configs/app.ini"
	_, err = os.Stat(configPath)
	if err != nil {
		configPath = "../configs/app.ini"
	}
	Cfg, err = ini.Load(configPath)
	if err != nil {
		log.Fatal(2, "Fail to parse 'app.ini': %v", err)
	}
	//server配置节点读取
	server, err := Cfg.GetSection("server")
	if err != nil {
		log.Fatal(2, "Fail to get section 'server': %v", err)
	}

	//app配置节点读取
	// app, err := Cfg.GetSection("app")
	// if err != nil {
	// 	log.Fatal(2, "Fail to get section 'app': %v", err)
	// }

	Config := ServerConfig{
		RunMode:      Cfg.Section("").Key("RUN_MODE").MustString("debug"),
		HTTPPort:     server.Key("HTTP_PORT").MustInt(),
		ReadTimeout:  time.Duration(server.Key("READ_TIMEOUT").MustInt(60)) * time.Second,
		WriteTimeout: time.Duration(server.Key("WRITE_TIMEOUT").MustInt(60)) * time.Second,
	}

	return Config
}
