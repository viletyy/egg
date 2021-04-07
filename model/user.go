/*
 * @Date: 2021-03-09 09:57:02
 * @LastEditors: viletyy
 * @LastEditTime: 2021-04-07 10:25:08
 * @FilePath: /egg/model/user.go
 */
package model

import (
	"github.com/viletyy/egg/global"
	"github.com/viletyy/egg/utils/crypt"
)

type User struct {
	global.Model
	Username string `json:"username"`
	Password string `json:"-"`
	IsAdmin  bool   `json:"is_admin" gorm:"default: false"`
}

func GetUsers(search *global.Search) (searchResult global.SearchResult, err error) {
	var users []User
	offset := search.PageInfo.PageSize * (search.PageInfo.Page - 1)
	limit := search.PageInfo.Page
	db := global.DB.Where(search.Maps)
	err = db.Count(&searchResult.Total).Error
	if err != nil {
		return
	}
	err = db.Offset(offset).Limit(limit).Find(&users).Error
	if err != nil {
		return
	}
	searchResult.Page = search.PageInfo.Page
	searchResult.PageSize = search.PageInfo.PageSize
	searchResult.List = users
	return
}

func GetUserById(id int) (user User, err error) {
	err = global.DB.Where("id = ?", id).First(&user).Error
	return
}

func GetUserByUsername(username string) (user User, err error) {
	err = global.DB.Where("username = ?", username).First(&user).Error
	return
}

func ExistByUsername(username string) bool {
	var user User
	global.DB.Where("username = ?", username).First(&user)

	return user.ID > 0
}

func CreateUser(user User) (err error) {
	err = global.DB.Create(&user).Error

	return
}

func UpdateUser(user *User) (err error) {
	err = global.DB.Save(&user).Error
	return
}

func DeleteUser(user *User) (err error) {
	err = global.DB.Delete(&user).Error
	return
}

func (user *User) CheckPassword(password string) bool {
	return crypt.Md5Check(password, user.Password)
}
