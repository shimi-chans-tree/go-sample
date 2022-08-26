package config

import (
	"log"
	"os"

	"gopkg.in/ini.v1"
)

type ConfigList struct {
	Port    int
	LogFile string
}

var Config ConfigList

func init() {
	cfg, err := ini.Load("config.ini")
	if err != nil {
		log.Printf("Failed to read file: %v", err)
		os.Exit(1)
	}

	Config = ConfigList{
		Port:    cfg.Section("web").Key("port").MustInt(),
		LogFile: cfg.Section("log").Key("log_file").String(),
	}
}
