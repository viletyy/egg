/*
 * @Date: 2021-03-12 17:25:57
 * @LastEditors: viletyy
 * @LastEditTime: 2021-04-07 10:25:24
 * @FilePath: /egg/validations/UserRegisterValidation.go
 */
package validations

import (
	"github.com/viletyy/egg/model"

	"github.com/beego/beego/v2/core/validation"
)

type UserRegisterValidation struct {
	Username             string `form:"username" valid:"Required; MaxSize(64)"`
	Password             string `form:"password" valid:"Required; MinSize(6); MaxSize(16)"`
	PasswordConfirmation string `form:"password_confirmation"`
}

func (urv *UserRegisterValidation) Valid(valid *validation.Validation) {
	if urv.Password != urv.PasswordConfirmation {
		valid.SetError("password", "两次输入密码不一致")
		return
	}

	if model.ExistByUsername(urv.Username) {
		valid.SetError("username", "用户名已经存在")
		return
	}
}
