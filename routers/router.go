/*
 * @Date: 2021-03-11 16:33:58
 * @LastEditors: viletyy
 * @LastEditTime: 2021-03-12 18:19:55
 * @FilePath: /hello/routers/router.go
 */
package routers

import (
	"github.com/viletyy/egg/controllers"

	beego "github.com/beego/beego/v2/server/web"
)

func init() {
	beego.Router("/", &controllers.IndexController{}, "get:Index")
	beego.Router("/login", &controllers.UserController{}, "get:Login;post:LoginHandle")
	beego.Router("/register", &controllers.UserController{}, "get:Register;post:RegisterHandle")
	beego.Router("/loginout", &controllers.UserController{}, "get:Loginout")
}
