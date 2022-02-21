package controllers

import (
	beego "github.com/beego/beego/v2/server/web"
)

type UserController struct {
	beego.Controller
}

func (c *UserController) Get() {
	web.AppConfig.String("mysqluser")
	c.Ctx.WriteString("hello")
}

func (c *UserController) User() {
	c.Ctx.WriteString("demo1111")
}
