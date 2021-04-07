/*
 * @Date: 2021-03-09 14:45:26
 * @LastEditors: viletyy
 * @LastEditTime: 2021-04-07 10:24:53
 * @FilePath: /egg/initialize/redis.go
 */
package initialize

import (
	"fmt"
	"github.com/viletyy/egg/conf/config"

	"github.com/beego/beego/v2/core/logs"
	"github.com/go-redis/redis"
)

var redisConfig = config.InitRedisConfig()

func Redis() *redis.Client {
	rdb := redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%s", redisConfig.Host, redisConfig.Port),
		Password: redisConfig.Password,
		DB:       redisConfig.Db,
	})

	RedisSet(rdb)
	return rdb
}

func RedisSet(rdb *redis.Client) {
	_, pingErr := rdb.Ping().Result()
	if pingErr != nil {
		logs.Error(fmt.Sprintf("Redis Connection Error: %v", pingErr))
	}
}
