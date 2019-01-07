package main

import (
	_ "DayDayshengxian/tiantianshengxian/daydayfresh/models"
	_ "DayDayshengxian/tiantianshengxian/daydayfresh/routers"

	"github.com/astaxie/beego"
)

func main() {
	beego.Run()
}
