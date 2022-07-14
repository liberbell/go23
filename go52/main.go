package main

import (
	"fmt"

	"gopkg.in/go-ini/ini.v1"
)

type Configlist struct {
	Port      int
	DBname    string
	SQLDriver string
}

var Config Configlist

func init() {
	cfg, _ := ini.Load("config.init")
	Config = Configlist{
		Port:      cfg.Section("web").Key("port").MustInt(),
		DBname:    cfg.Section("db").Key("name").MustString("example.sql"),
		SQLDriver: cfg.Section("db").Key("driver").String(),
	}
}

func main() {
	fmt.Printf("%T %v\n", Config.Port, Config.Port)
	fmt.Printf("%T %v\n", Config.DBname, Config.DBname)
	fmt.Printf("%T %v\n", Config.SQLDriver, Config.SQLDriver)
}
