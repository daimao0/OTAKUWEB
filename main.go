package main

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context"
	"net/http"
	_ "youblog/routers"
	"youblog/service"
)

func main() {
	//后台管理oss
	beego.InsertFilter("/sys/manager/*", beego.BeforeRouter, func(ctx *context.Context) {
		cookie, err := ctx.Request.Cookie("Authorization")
		//utils.CheckToken(cookie.Value)
		if err != nil || !service.CheckTokenFromRedis(cookie.Value) {
			http.Redirect(ctx.ResponseWriter, ctx.Request, "/sys/login", http.StatusMovedPermanently)
		}
	})
	beego.Run()
}
