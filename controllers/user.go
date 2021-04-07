/*
 * @Date: 2021-03-11 10:26:08
 * @LastEditors: viletyy
 * @LastEditTime: 2021-04-07 10:28:02
 * @FilePath: /egg/controllers/user.go
 */
package controllers

import (
	"github.com/viletyy/egg/model"
	"github.com/viletyy/egg/utils"
	"github.com/viletyy/egg/utils/crypt"
	"github.com/viletyy/egg/validations"

	beego "github.com/beego/beego/v2/server/web"
)

type UserController struct {
	BaseController
}

func (c *UserController) Login() {
	if c.CurrentLoginUser != nil {
		c.RedirectTo("/")
	}
}

func (c *UserController) LoginHandle() {
	c.RedirectUrl = beego.URLFor("UserController.Login")
	loginUser := validations.UserLoginValidation{}
	c.ValidatorAuto(&loginUser)

	isExist := model.ExistByUsername(loginUser.Username)
	if !isExist {
		c.FlashError("用户不存在")
		c.RedirectTo(c.RedirectUrl)
	}

	dbUser, _ := model.GetUserByUsername(loginUser.Username)

	if !dbUser.CheckPassword(loginUser.Password) {
		c.FlashError("密码不正确")
		c.RedirectTo(c.RedirectUrl)
	}

	if c.GetString("remember_me") != "" {
		c.Ctx.SetCookie("login_user_id", utils.ToString(dbUser.ID), 3600*24*7)
		c.Ctx.SetCookie("login_user_sign", utils.AuthSign(int(dbUser.ID), dbUser.Username, dbUser.Password), 3600*24*7)
	} else {
		c.Ctx.SetCookie("login_user_id", utils.ToString(dbUser.ID), 3600)
		c.Ctx.SetCookie("login_user_sign", utils.AuthSign(int(dbUser.ID), dbUser.Username, dbUser.Password), 3600)
	}

	c.FlashSuccess("登陆成功")
	c.RedirectTo("/")
}

func (c *UserController) Register() {

}

func (c *UserController) RegisterHandle() {
	c.RedirectUrl = beego.URLFor("UserController.Register")
	registerUser := validations.UserRegisterValidation{}
	c.ValidatorAuto(&registerUser)

	dbUser := model.User{
		Username: registerUser.Username,
		Password: crypt.Md5Encode(registerUser.Password),
	}

	if err := model.CreateUser(dbUser); err != nil {
		c.FlashError("注册失败")
		c.RedirectTo(c.RedirectUrl)
	} else {
		c.FlashSuccess("注册成功")
		c.RedirectTo(beego.URLFor("UserController.Login"))
	}
}

func (c *UserController) Loginout() {
	c.CurrentLoginUser = nil
	c.Ctx.SetCookie("login_user_id", "", -3600)
	c.Ctx.SetCookie("login_user_sign", "", -3600)
	c.FlashSuccess("已安全登出")
	c.RedirectTo(c.RedirectUrl)
}
