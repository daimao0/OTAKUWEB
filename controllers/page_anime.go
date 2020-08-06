package controllers

import (
	"github.com/astaxie/beego"
)

type AnimeController struct {
	beego.Controller
}

func (c *AnimeController) Get() {
	c.TplName = "page_soft.html"

}
