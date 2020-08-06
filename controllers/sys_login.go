package controllers

import (
	"github.com/astaxie/beego"
	"log"
	"youblog/common/api"
	"youblog/common/utils"
	"youblog/mapper"
	"youblog/modules"
)

type SysLoginController struct {
	beego.Controller
}

func (c *SysLoginController) Get() {

	c.TplName = "sys/sys_login.html"
}

func (c *SysLoginController) Login() {
	email := c.GetString("username")
	password := c.GetString("password")
	remember := c.GetString("remember")
	var su modules.SysUser
	sysUser := mapper.GetUserByEmailAndPassword(email, password)
	//查询不到用户,返回204
	if sysUser == su {
		c.Data["json"] = api.SendResponse(204, "验证失败", sysUser)
		c.ServeJSON()
		return
	}
	//查询到用户后，选择记住我之后
	if remember == "remember" {
		token := utils.GenerateToken(&sysUser)
		/*		cookie := http.Cookie{Name: "Authorization", Value: token, Path: "/sys/", MaxAge: 2592000}	//最大时间30天
				http.SetCookie(c.Ctx.ResponseWriter, &cookie)*/
		c.Ctx.SetCookie("Authorization", token, 604800, "/sys/") //最大时间30天
		RedisConn := utils.InitRedis().Get()
		defer RedisConn.Close()
		n, err := RedisConn.Do("set", "token:"+string(sysUser.UserId), token)
		if err != nil {
			log.Println("redis存入string类型token失败", err)
		}
		if n == int64(1) {
			log.Println("redis存储string类型成功")
		}

		n, err = RedisConn.Do("expire", "token:"+string(sysUser.UserId), 604800) //设置过期时间30天
		if err != nil {
			log.Println("redis设置token过期时间失败", err)
		}
		if n == int64(1) {
			log.Println("redis设置token过期时间成功")
		}
	}

	c.Data["json"] = api.SendResponse(201, "登陆成功", sysUser)
	c.ServeJSON()

}
