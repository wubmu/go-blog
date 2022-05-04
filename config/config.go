package config

import (
	"github.com/BurntSushi/toml"
	"log"
	"os"
)

// 映射配置文件
type tomlConfig struct {
	//外部访问开头大写
	Viewer Viewer
	System SystemConfig
}

type Viewer struct {
	Title       string
	Description string
	Logo        string
	Navigation  []string
	Bilibili    string
	Avatar      string
	UserName    string
	UserDesc    string
}
type SystemConfig struct {
	AppName         string
	Version         float32
	CurrentDir      string
	CdnURL          string
	QiniuAccessKey  string
	QiniuSecretKey  string
	Valine          bool
	ValineAppid     string
	ValineAppkey    string
	ValineServerURL string
}

var Cfg *tomlConfig

func init() {
	// 程序启动的时候，就会执行init
	Cfg = new(tomlConfig)

	Cfg.System.AppName = "go-blog"
	Cfg.System.Version = 1.0
	currentDir, _ := os.Getwd()
	Cfg.System.CurrentDir = currentDir
	println("---------------------", currentDir)
	_, err := toml.DecodeFile("config/config.toml", &Cfg)
	if err != nil {
		//panic(err)
		log.Println("读取配置文件出错", err)
	}
}
