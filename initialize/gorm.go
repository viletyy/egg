/*
 * @Date: 2021-03-09 09:57:02
 * @LastEditors: viletyy
 * @LastEditTime: 2021-04-07 10:14:54
 * @FilePath: /egg/initialize/gorm.go
 */
package initialize

import (
	"fmt"
	"github.com/viletyy/egg/conf/config"
	"github.com/viletyy/egg/model"

	"github.com/beego/beego/v2/core/logs"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

var dbConfig = config.InitDatabaseConfig()

func Gorm() *gorm.DB {
	switch dbConfig.Type {
	case "mysql":
		return GormMysql()
	case "postgresql":
		return GormPostgresql()
	default:
		return GormMysql()
	}
}

func GormMysql() *gorm.DB {
	db, err := gorm.Open("mysql", fmt.Sprintf("%s:%s@(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", dbConfig.User, dbConfig.Password, dbConfig.Host, dbConfig.Port, dbConfig.Name))
	if err != nil {
		logs.Error(fmt.Sprintf("Mysql Gorm Open Error: %v", err))
	}
	GormSet(db)
	return db
}

func GormPostgresql() *gorm.DB {
	db, err := gorm.Open("postgres", fmt.Sprintf("host=%s user=%s dbname=%s port=%s sslmode=disable password=%s", dbConfig.Host, dbConfig.User, dbConfig.Name, dbConfig.Port, dbConfig.Password))
	if err != nil {
		logs.Error(fmt.Sprintf("Postgresql Gorm Open Error: %v", err))
	}
	GormSet(db)
	return db
}

func GormSet(db *gorm.DB) {
	// 设置表前缀
	gorm.DefaultTableNameHandler = func(db *gorm.DB, defaultTableName string) string {
		return dbConfig.TablePrefix + defaultTableName
	}

	// 设置日志
	db.LogMode(true)

	// 设置迁移
	db.AutoMigrate(&model.User{})

	// 设置空闲连接池中的最大连接数
	db.DB().SetMaxIdleConns(10)

	// 设置打开数据库连接的最大数量
	db.DB().SetMaxOpenConns(100)
}
