/*
 * @Date: 2021-03-09 09:57:02
 * @LastEditors: viletyy
 * @LastEditTime: 2021-04-07 10:33:44
 * @FilePath: /egg/main.go
 */
package main

import (
	"github.com/viletyy/egg/global"
	"github.com/viletyy/egg/initialize"
	_ "github.com/viletyy/egg/routers"

	beego "github.com/beego/beego/v2/server/web"
)

func main() {
	global.LOG = initialize.Zap()

	// 初始化数据库
	global.DB = initialize.Gorm()
	defer global.DB.Close()

	// 初始化redis
	global.REDIS = initialize.Redis()
	defer global.REDIS.Close()

	// 启动应用
	beego.Run()
}
