/*
 * @Date: 2021-03-09 09:57:02
 * @LastEditors: viletyy
 * @LastEditTime: 2021-04-07 10:35:47
 * @FilePath: /egg/global/global.go
 */
package global

import (
	"github.com/go-redis/redis"
	"github.com/jinzhu/gorm"
	"go.uber.org/zap"
)

var (
	DB    *gorm.DB
	REDIS *redis.Client
	LOG   *zap.Logger
)
