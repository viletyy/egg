/*
 * @Date: 2021-03-09 09:57:02
 * @LastEditors: viletyy
 * @LastEditTime: 2021-03-10 17:49:52
 * @FilePath: /hello/conf/config/database.go
 */
package config

import (
	"github.com/beego/beego/v2/core/logs"
	beego "github.com/beego/beego/v2/server/web"
)

type Database struct {
	Type        string
	User        string
	Password    string
	Host        string
	Port        string
	Name        string
	TablePrefix string
}

func InitDatabaseConfig() *Database {
	cType, cTypeErr := beego.AppConfig.String("database::type")
	if cTypeErr != nil {
		logs.Error("Get Config Error: [database] type")
	}
	cUser, cUserErr := beego.AppConfig.String("database::user")
	if cUserErr != nil {
		logs.Error("Get Config Error: [database] user")
	}
	cPassword, cPasswordErr := beego.AppConfig.String("database::password")
	if cPasswordErr != nil {
		logs.Error("Get Config Error: [database] password")
	}
	cHost, cHostErr := beego.AppConfig.String("database::host")
	if cHostErr != nil {
		logs.Error("Get Config Error: [database] host")
	}
	cPort, cPortErr := beego.AppConfig.String("database::port")
	if cPortErr != nil {
		logs.Error("Get Config Error: [database] port")
	}
	cName, cNameErr := beego.AppConfig.String("database::name")
	if cNameErr != nil {
		logs.Error("Get Config Error: [database] name")
	}
	cTablePrefix, cTablePrefixErr := beego.AppConfig.String("database::table_prefix")
	if cTablePrefixErr != nil {
		logs.Error("Get Config Error: [database] table_prefix")
	}
	return &Database{
		Type:        cType,
		User:        cUser,
		Password:    cPassword,
		Host:        cHost,
		Port:        cPort,
		Name:        cName,
		TablePrefix: cTablePrefix,
	}
}
