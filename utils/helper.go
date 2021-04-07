/*
 * @Date: 2021-03-12 17:37:03
 * @LastEditors: viletyy
 * @LastEditTime: 2021-03-12 22:12:59
 * @FilePath: /egg/utils/helper.go
 */
package utils

import "strconv"

func ToString(i interface{}) string {
	switch i.(type) {
	case string:
		return i.(string)
	case int:
		return strconv.Itoa(i.(int))
	case int64:
		return strconv.FormatInt(i.(int64), 10)
	case uint:
		return strconv.Itoa(int(i.(uint)))
	}
	return ""
}
