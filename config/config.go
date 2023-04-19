package config

/**
* @Author: Xenolies
* @Date: 2023/4/16 12:10
* @Version: 1.0
 */

import (
	"fmt"
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
	"log"
)

/*
读取配置文件
*/

type Config struct {
	Admin Admin
}

type Admin struct {
	IP   string
	Port string

	Name string
	Pwd  string
	Mail string

	DbName     string
	DbLoginID  string
	DbLoginPwd string
	DbAddr     string
	DbConfig   string

	JwtSecret      string
	ExpirationTime string
}

var GlobalConfig Config

func init() {
	fmt.Println("Load Config...... ")
	v := viper.New()
	v.SetConfigName("bolg_config") // 配置文件名
	v.AddConfigPath("./conf/")     // 配置文件搜索路径
	v.SetConfigType("json")        // 配置文件读取类型
	err := v.ReadInConfig()        // 读取配置文件
	if err != nil {
		panic(fmt.Errorf("Fatal error config file: %s \n", err))
	}

	// 热更新
	v.WatchConfig()
	v.OnConfigChange(func(e fsnotify.Event) {
		log.Print("Config file updated: ", e.Name)
	})

	if err := v.Unmarshal(&GlobalConfig); err != nil {
		fmt.Println(err)
	}

	//adminName := GlobalConfig.Admin.Name
	//adminPwd := GlobalConfig.Admin.Pwd
	//adminMail := GlobalConfig.Admin.Mail
}
