package main

import (
	_ "demo/routers"
	beego "github.com/beego/beego/v2/server/web"
)

func main() {
	//StaticDir["/static"] = "static"
	beego.Run()
}

