/*
 * @Date: 2021-03-09 14:47:21
 * @LastEditors: viletyy
 * @LastEditTime: 2021-03-10 17:48:39
 * @FilePath: /hello/conf/config/redis.go
 */
package config

import (
	"github.com/beego/beego/v2/core/logs"
	beego "github.com/beego/beego/v2/server/web"
)

type Redis struct {
	Host     string
	Port     string
	Db       int
	Password string
}

func InitRedisConfig() *Redis {
	rHost, rHostErr := beego.AppConfig.String("redis::host")
	if rHostErr != nil {
		logs.Error("Get Config Error: [redis] host")
	}
	rPort, rPortErr := beego.AppConfig.String("redis::port")
	if rPortErr != nil {
		logs.Error("Get Config Error: [redis] port")
	}
	rDb, rDbErr := beego.AppConfig.Int("redis::db")
	if rDbErr != nil {
		logs.Error("Get Config Error: [redis] db")
	}
	rPassword, rPasswordErr := beego.AppConfig.String("redis::password")
	if rPasswordErr != nil {
		logs.Error("Get Config Error: [redis] password")
	}
	return &Redis{
		Host:     rHost,
		Port:     rPort,
		Db:       rDb,
		Password: rPassword,
	}
}
