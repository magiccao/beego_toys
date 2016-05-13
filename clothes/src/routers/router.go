package routers

import "github.com/astaxie/beego"

import "controllers"

func init() {
	beego.Router("/", &controllers.MainController{}, "get:Get")
	beego.Router("/css/*", &controllers.MainController{}, "get:Get")
	beego.Router("/index", &controllers.MainController{}, "get:Get")
	beego.Router("/statistics", &controllers.MainController{}, "post:Statistics")
	beego.Router("/show", &controllers.MainController{}, "get:Show")
}
