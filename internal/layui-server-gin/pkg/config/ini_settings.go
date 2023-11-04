package config

import (
	"fmt"

	"github.com/go-ini/ini"
)

// Config is config entry
var Config *Settings

// APISettings is config mapping
type Settings struct {
	EnvSetting    Environment
	APISetting    API
	LogSetting    Log
	MysqlSettings Mysql
}

// InitConfig returns error type
func InitConfig(config string) (err error) {
	var cfg *ini.File
	cfg, err = ini.Load(config)
	if err != nil {
		fmt.Println("Read config file error: " + config)
		return err
	}
	cfg.NameMapper = ini.TitleUnderscore

	Config = new(Settings)

	err = cfg.Section("environment").MapTo(&Config.EnvSetting)
	if err != nil {
		return err
	}
	err = cfg.Section("api").MapTo(&Config.APISetting)
	if err != nil {
		return err
	}
	err = cfg.Section("log").MapTo(&Config.LogSetting)
	if err != nil {
		return err
	}
	err = cfg.Section("mysql").MapTo(&Config.MysqlSettings)
	if err != nil {
		return err
	}
	return nil
}
