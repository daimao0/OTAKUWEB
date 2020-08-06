package controllers

import (
	"github.com/astaxie/beego"
)

type SysMainController struct {
	beego.Controller
}

func (c *SysMainController) Get() {

	c.TplName = "sys/sys_main.html"
}
