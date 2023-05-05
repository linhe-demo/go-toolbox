package config

import (
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/rest"
)

type Config struct {
	rest.RestConf
	BaiduOauth BaiduOauth      `json:"BaiduOauth"`
	MysqlConf  MysqlConf       `json:"mysqlConf"`
	RedisConf  RedisConf       `json:"RedisConf"`
	CacheRedis cache.CacheConf // redis缓存
}

type BaiduOauth struct {
	AppID     int    `json:"AppID"`
	AppKey    string `json:"AppKey"`
	AppSecret string `json:"AppSecret"`
	OauthUrl  string `json:"OauthUrl"`
}

type RedisConf struct {
	Host     string `json:"Host"`
	Port     int    `json:"Port"`
	User     string `json:"User"`
	Password string `json:"Password"`
}

type MysqlConf struct {
	Host     string `json:"Host"`
	Port     int    `json:"Port"`
	User     string `json:"User"`
	Password string `json:"Password"`
	DbName   string `json:"DbName"`
}
