package routers

import (
	"DayDayshengxian/tiantianshengxian/controllers"
	"github.com/astaxie/beego"
)

func init() {
    beego.Router("/", &controllers.MainController{})
}
