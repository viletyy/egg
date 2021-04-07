/*
 * @Date: 2021-03-12 17:25:08
 * @LastEditors: viletyy
 * @LastEditTime: 2021-03-12 17:25:30
 * @FilePath: /hello/validations/UserLoginValidation.go
 */
package validations

type UserLoginValidation struct {
	Username string `form:"username" valid:"Required; MaxSize(64)"`
	Password string `form:"password" valid:"Required; MinSize(6); MaxSize(16)"`
}
