/*
 * @Date: 2021-03-09 09:57:02
 * @LastEditors: viletyy
 * @LastEditTime: 2021-04-07 09:50:19
 * @FilePath: /egg/global/model.go
 */
package global

import (
	"time"
)

type Model struct {
	ID        uint `gorm:"primarykey"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time `gorm:"index" json:"-"`
}

type PageInfo struct {
	Page     int `json:"page_num"`
	PageSize int `json:"page_size"`
}

type Search struct {
	Maps interface{} `json:"maps"`
	PageInfo
}

type SearchResult struct {
	List     interface{} `json:"list"`
	Total    int64       `json:"total"`
	Page     int         `json:"page"`
	PageSize int         `json:"page_size"`
}
