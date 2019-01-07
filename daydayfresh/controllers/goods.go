package controllers

import (
	"github.com/astaxie/beego"
)

type GoodsController struct {
	beego.Controller
}

func (this *GoodsController) ShowIndex() {

	userName := this.GetSession("userName")
	if userName == nil {
		this.Data["userName"] = ""
		beego.Error("userName=空字符串")
	} else {
		this.Data["userName"] = userName.(string)
	}

	this.TplName = "index.html"
}
