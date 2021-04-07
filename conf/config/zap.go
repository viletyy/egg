/*
 * @Date: 2021-04-07 10:37:55
 * @LastEditors: viletyy
 * @LastEditTime: 2021-04-07 10:42:53
 * @FilePath: /egg/conf/config/zap.go
 */
package config

import (
	"github.com/beego/beego/v2/core/logs"
	beego "github.com/beego/beego/v2/server/web"
)

type Zap struct {
	Level     string
	Format     string
	Prefix       string
	Director string
	LinkName      string 
	ShowLine      bool   
	EncodeLevel   string 
	StacktraceKey string 
	LogInConsole  bool   
}

func InitZapConfig() *Zap {
	zLevel, zLevelErr := beego.AppConfig.String("zap::level")
	if zLevelErr != nil {
		logs.Error("Get Config Error: [zap] level")
	}
	zFormat, zFormatErr := beego.AppConfig.String("zap::format")
	if zFormatErr != nil {
		logs.Error("Get Config Error: [zap] format")
	}
	zPrefix, zPrefixErr := beego.AppConfig.String("zap::format")
	if zPrefixErr != nil {
		logs.Error("Get Config Error: [zap] format")
	}
	zDirector, zDirectorErr := beego.AppConfig.String("zap::director")
	if zDirectorErr != nil {
		logs.Error("Get Config Error: [zap] director")
	}
	zLinkName, zLinkNameErr := beego.AppConfig.String("zap::link_name")
	if zLinkNameErr != nil {
		logs.Error("Get Config Error: [zap] linkname")
	}
	zShowLine, zShowLineErr := beego.AppConfig.Bool("zap::show_line")
	if zShowLineErr != nil {
		logs.Error("Get Config Error: [zap] showline")
	}
	zEncodeLevel, zEncodeLevelErr := beego.AppConfig.String("zap::encode_level")
	if zEncodeLevelErr != nil {
		logs.Error("Get Config Error: [zap] encode_level")
	}
	zStacktraceKey, zStacktraceKeyErr := beego.AppConfig.String("zap::stacktrace_key")
	if zStacktraceKeyErr != nil {
		logs.Error("Get Config Error: [zap] stacktrace_key")
	}
	zLogInConsole, zLogInConsoleErr := beego.AppConfig.Bool("zap::log_in_console")
	if zLogInConsoleErr != nil {
		logs.Error("Get Config Error: [zap] log_in_console")
	}
	return &Zap{
		Level: zLevel,
		Format: zFormat,
		Prefix: zPrefix,
		Director: zDirector,
		LinkName: zLinkName,
		ShowLine: zShowLine,
		EncodeLevel: zEncodeLevel,
		StacktraceKey: zStacktraceKey,
		LogInConsole: zLogInConsole,
	}
}