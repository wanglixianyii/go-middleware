package main

import (
	"fmt"
	"github.com/wanglixianyii/go-middleware/go-es/wire"
	"os"

	"gopkg.in/yaml.v2"

	"github.com/wanglixianyii/go-middleware/go-es/config"
)

type UserSvr struct {
	conf config.ServerConfig
}

func (u *UserSvr) init() {
	file, err := os.ReadFile("./config.yaml")
	if err != nil {
		fmt.Println("read yaml file failed")
	}

	err = yaml.UnmarshalStrict(file, &u.conf)
	if err != nil {
		fmt.Println("yaml unmarshal failed")
	}
}

func (u *UserSvr) Run() {
	handler := wire.InitializeHandler(&u.conf)
	handler.Run()
}
