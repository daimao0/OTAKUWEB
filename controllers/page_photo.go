package controllers

import (
	"github.com/astaxie/beego"
)

type PhotoController struct {
	beego.Controller
}

func (c *PhotoController) Get() {
	c.TplName = "page_soft.html"

}
