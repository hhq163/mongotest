package controllers

import (
	"math/rand"
	"time"

	"github.com/astaxie/beego"
)

type BaseController struct {
	beego.Controller
}

/**
* 获取一个随机字符
**/
func (this *BaseController) GenerateMixed(n int) string {
	var chars = [36]string{"0", "1", "2", "3", "4", "5", "6", "7", "8", "9", "A", "B", "C", "D", "E", "F", "G", "H", "I", "J", "K", "L", "M", "N", "O", "P", "Q", "R", "S", "T", "U", "V", "W", "X", "Y", "Z"}
	var res string = ""

	rand.Seed(int64(time.Now().Nanosecond()))

	for i := 0; i < n; i++ {
		id := rand.Intn(35)
		res += chars[id]
	}
	return res
}

func (this *BaseController) Go404() {
	this.TplName = "404.tpl"
	return
}
