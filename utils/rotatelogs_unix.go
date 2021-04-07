/*
 * @Date: 2021-03-22 16:44:06
 * @LastEditors: viletyy
 * @LastEditTime: 2021-04-06 18:30:01
 * @FilePath: /egg/utils/rotatelogs_unix.go
 */
// +build !windows

package utils

import (
	"github.com/viletyy/egg/conf/config"
	"os"
	"path"
	"time"

	zaprotatelogs "github.com/lestrrat-go/file-rotatelogs"
	"go.uber.org/zap/zapcore"
)

var zapConfig = config.InitZapConfig()

//@author: [SliverHorn](https://github.com/SliverHorn)
//@function: GetWriteSyncer
//@description: zap logger中加入file-rotatelogs
//@return: zapcore.WriteSyncer, error

func GetWriteSyncer() (zapcore.WriteSyncer, error) {
	fileWriter, err := zaprotatelogs.New(
		path.Join(zapConfig.Director, "%Y-%m-%d.log"),
		zaprotatelogs.WithLinkName(zapConfig.LinkName),
		zaprotatelogs.WithMaxAge(7*24*time.Hour),
		zaprotatelogs.WithRotationTime(24*time.Hour),
	)
	if zapConfig.LogInConsole {
		return zapcore.NewMultiWriteSyncer(zapcore.AddSync(os.Stdout), zapcore.AddSync(fileWriter)), err
	}
	return zapcore.AddSync(fileWriter), err
}
