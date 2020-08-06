package routers

import (
	"github.com/astaxie/beego"
	"youblog/controllers"
)

func init() {
	beego.Router("/", &controllers.IndexController{})
	beego.Router("/main", &controllers.MainController{})
	beego.Router("/sys/login", &controllers.SysLoginController{})
	beego.Router("/sys/loginCheck", &controllers.SysLoginController{}, "post:Login")
	beego.Router("/sys/manager/index", &controllers.SysController{})
	beego.Router("/sys/manager/index/main", &controllers.SysMainController{})
	beego.Router("/page_soft", &controllers.SortWareController{})
	beego.Router("/page_devtool", &controllers.SortWareController{}, "get:RedirectDevtool")
	beego.Router("/page_soft/ListSoftWares", &controllers.SortWareController{}, "get:ListSoftWares")
	beego.Router("/page_soft/ListDevTools", &controllers.SortWareController{}, "get:ListDevTools")
}
