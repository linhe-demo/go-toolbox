package config

import (
	"github.com/zeromicro/go-zero/rest"
)

type Config struct {
	rest.RestConf
	BaiduOauth BaiduOauth `json:"BaiduOauth"`
	RedisConf  RedisConf  `json:"RedisConf"`
}

type RedisConf struct {
	Host     string `json:"Host"`
	Port     int    `json:"Port"`
	User     string `json:"User"`
	Password string `json:"Password"`
}

type BaiduOauth struct {
	AppID     int    `json:"AppID"`
	AppKey    string `json:"AppKey"`
	AppSecret string `json:"AppSecret"`
	OauthUrl  string `json:"OauthUrl"`
}
