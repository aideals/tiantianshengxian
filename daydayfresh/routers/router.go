package routers

import (
	"DayDayshengxian/tiantianshengxian/daydayfresh/controllers"

	"github.com/astaxie/beego/context"

	"github.com/astaxie/beego"
)

func init() {
	//过滤器
	beego.InsertFilter("/goods/*", beego.BeforeExec, filterFunc)

	beego.Router("/index", &controllers.GoodsController{}, "get:ShowIndex")
	//注册页面
	beego.Router("/register", &controllers.UserController{}, "get:ShowRegister;post:HandleRegister")
	//登录页面
	beego.Router("/login", &controllers.UserController{}, "get:ShowLogin;post:HandleLogin")
	//激活用户
	beego.Router("/active", &controllers.UserController{}, "get:ActiveUser")
	//退出登录
	beego.Router("/logout", &controllers.UserController{}, "get:Logout")
	//用户中心
	beego.Router("/goods/userCenterInfo", &controllers.UserController{}, "get:ShowUserCenterInfo")
	//用户中心订单页
	beego.Router("/goods/userCenterOrder", &controllers.UserController{}, "get:UserCenterOrder")
	//用户中心地址页面
	beego.Router("/goods/userCenterSite", &controllers.UserController{}, "get:ShowUserCenterSite")
	//添加收件人地址页面
	beego.Router("/goods/addSite", &controllers.UserController{}, "post:HandleSite")
}

func filterFunc(ctx *context.Context) {
	//获取seesion
	userName := ctx.Input.Session("userName")
	if userName == nil {
		ctx.Redirect(302, "/login")
		return
	}
}
