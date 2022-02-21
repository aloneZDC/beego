package routers

import (
	"demo/controllers"
	beego "github.com/beego/beego/v2/server/web"
)

func init() {
    beego.Router("/", &controllers.MainController{})
	beego.Router("/user", &controllers.UserController{}, "get:User")
}