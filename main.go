package main

import (
	_ "lucky-draw/routers"
	beego "github.com/beego/beego/v2/server/web"
)

func main() {
	beego.Run()
}

