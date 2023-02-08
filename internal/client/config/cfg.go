package config

import "time"

var Cfg = &cfg{}

type ClientCfg struct {
	Addr    string        `json:"addr"`
	Timeout time.Duration `json:"timeout"`
	AppId   string        `json:"appId"`
	AppKey  string        `json:"appKey"`
}

type cfg struct {
	UserClient  ClientCfg `json:"userClient"`
	OrderClient ClientCfg `json:"orderClient"`
	PayClient   ClientCfg `json:"payClient"`
}
