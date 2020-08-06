package controllers

import (
	"github.com/astaxie/beego"
	"youblog/service"
)

type SortWareController struct {
	beego.Controller
}

func (c *SortWareController) Get() {
	c.TplName = "page_soft.html"

}
func (c *SortWareController) RedirectDevtool() {
	c.TplName = "page_devtool.html"

}
func (c *SortWareController) ListSoftWares() {
	c.Ctx.ResponseWriter.Header().Set("Access-Control-Allow-Origin", "*") //允许跨域
	c.Data["json"] = service.ListSoftWares()
	c.ServeJSON()
}
func (c *SortWareController) ListDevTools() {
	c.Ctx.ResponseWriter.Header().Set("Access-Control-Allow-Origin", "*") //允许跨域
	c.Data["json"] = service.ListDevTools()
	c.ServeJSON()
}
