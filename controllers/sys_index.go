package controllers

import (
	"github.com/astaxie/beego"
)

type SysController struct {
	beego.Controller
}

func (c *SysController) Get() {

	c.TplName = "sys/sys_index.html"
}
