package controllers

import (
	"time"

	"github.com/astaxie/beego"
)

type MainController struct {
	beego.Controller
}

func (c *MainController) Get() {
	c.Data["title"] = "技术验证列表"
	c.Data["Ntime"] = time.Now().String()
	c.Data["Email"] = "hhq163@gmail.com"
	c.Data["Website"] = "blog.csdn.net/hhq163"
	c.TplName = "index.tpl"

}
