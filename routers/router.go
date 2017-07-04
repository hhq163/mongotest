package routers

import (
	"mongotest/controllers"

	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/", &controllers.MainController{})

	beego.Router("/datamake", &controllers.WelcomeController{})
	beego.Router("/datamake/make_gri", &controllers.DataMakeController{})

}
