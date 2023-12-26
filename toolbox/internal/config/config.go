package config

import (
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/rest"
)

type Config struct {
	rest.RestConf
	BaiduOauth   BaiduOauth      `json:"BaiduOauth"`
	MysqlConf    MysqlConf       `json:"mysqlConf"`
	RedisConf    RedisConf       `json:"RedisConf"`
	RocketMqConf RocketMqConf    `json:"RocketMqConf"`
	CacheRedis   cache.CacheConf // redis缓存
	QiNiuConf    QiNiuConf       `json:"QiNiuConf"`
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
	Host       string `json:"Host"`
	Port       int    `json:"Port"`
	User       string `json:"User"`
	Password   string `json:"Password"`
	LogDbName  string `json:"LogDbName"`
	LifeDbName string `json:"LifeDbName"`
}

type RocketMqConf struct {
	Host string `json:"host"`
	Port int    `json:"port"`
}

type QiNiuConf struct {
	AccessKey string `json:"AccessKey"`
	SecretKey string `json:"SecretKey"`
	Bucket    string `json:"Bucket"`
}
