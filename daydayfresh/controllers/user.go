package controllers

import (
	"DayDayshengxian/tiantianshengxian/daydayfresh/models"
	_ "DayDayshengxian/tiantianshengxian/daydayfresh/models"
	"regexp"
	"strconv"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"github.com/astaxie/beego/utils"
)

//创建userController
type UserController struct {
	//继承自beego
	beego.Controller
}

//显示注册页面
func (this *UserController) ShowRegister() {
	//显示注册页面
	this.TplName = "register.html"
}

//处理注册业务
func (this *UserController) HandleRegister() {
	//获取数据
	userName := this.GetString("user_name")
	pwd := this.GetString("pwd")
	cpwd := this.GetString("cpwd")
	email := this.GetString("email")

	//校验数据
	if userName == "" || pwd == "" || cpwd == "" || email == "" {
		beego.Error("注册数据不能为空")
		this.TplName = "register.html"
		return
	}

	//匹配邮箱格式
	reg, err := regexp.Compile(`^[A-Za-z\d]+([-_.][A-Za-z\d]+)*@([A-Za-z\d]+[-.])+[A-Za-z\d]{2,4}$`)
	if err != nil {
		beego.Error("邮箱格式不正确")
		this.TplName = "register.html"
		return
	}

	res := reg.MatchString(email)
	if res == false {
		beego.Error("邮箱格式不正确")
		this.TplName = "register.html"
		return
	}

	//判断两次输入的密码是否正确
	if cpwd != pwd {
		beego.Error("两次输入的密码不正确")
		this.TplName = "register.html"
		return
	}

	//处理数据
	//获取orm对象
	o := orm.NewOrm()
	//创建可操作对象
	var user models.User
	//给对象赋值
	user.UserName = userName
	user.Pwd = pwd
	user.Email = email

	//执行插入操作
	_, err = o.Insert(&user)
	if err != nil {
		beego.Error("插入数据出错", err)
		this.TplName = "register.html"
		return
	}

	//注册成功的时候发送激活邮件
	config := `{"username":"1352727102@qq.com","password":"boolzhpxrkckiacj","host":"smtp.qq.com", "port":587}`
	//创建email对象
	emailSend := utils.NewEMail(config)
	emailSend.From = "1352727102@qq.com"
	emailSend.To = []string{email}
	emailSend.Subject = "天天生鲜用户激活"
	emailSend.HTML = `<a href="http://192.168.79.21:8080/active?userId=` + strconv.Itoa(user.Id) + `">点击激活</a>`

	emailSend.Send()

	this.Ctx.WriteString("注册成功，请前往邮箱激活!")
}

//显示登录页面
func (this *UserController) ShowLogin() {
	//获取cookie
	userName := this.Ctx.GetCookie("userName")
	if userName != "" {
		this.Data["userName"] = userName
		this.Data["checked"] = "checked"
	} else {
		this.Data["userName"] = ""
		this.Data["checked"] = ""
	}

	this.TplName = "login.html"
}

//处理登录业务
func (this *UserController) HandleLogin() {
	//获取数据
	userName := this.GetString("username")
	pwd := this.GetString("pwd")

	//校验数据
	if userName == "" || pwd == "" {
		beego.Error("登录数据不能为空")
		this.TplName = "login.html"
		return
	}

	//处理数据
	//查询数据
	//创建orm对象
	o := orm.NewOrm()
	//创建可操作对象
	var user models.User
	//给查询条件赋值
	user.UserName = userName
	//读一下
	err := o.Read(&user, "UserName")
	if err != nil {
		beego.Error("查询出错")
		this.TplName = "login.html"
		return
	}

	if user.Pwd != pwd {
		beego.Error("密码不匹配")
		this.TplName = "login.html"
		return
	}

	if user.Active == 0 {
		beego.Error("用户未激活")
		this.TplName = "login.html"
		return
	}

	//利用cookie记住用户名
	remember := this.GetString("remember")
	if remember == "on" {
		this.Ctx.SetCookie("userName", userName, 3600)
	} else {
		this.Ctx.SetCookie("userName", userName, -1)
	}

	this.SetSession("userName", userName)

	//返回数据
	this.Redirect("/index", 302)
}

//激活用户
func (this *UserController) ActiveUser() {
	beego.Info(123123123123123123)
	//获取数据
	userId, err := this.GetInt("userId")
	if err != nil {
		beego.Error("获取用户数据出错")
		this.TplName = "register.html"
		return
	}

	//获取orm对象
	o := orm.NewOrm()
	//获取可操作对象
	var user models.User
	//给对象赋值
	user.Id = userId

	//去取数据库
	err = o.Read(&user)
	if err != nil {
		beego.Error("用户不存在，激活出错")
		this.TplName = "register.html"
		return
	}

	user.Active = 1
	_, err = o.Update(&user)
	if err != nil {
		beego.Error("激活用户出错")
		this.TplName = "register.html"
		return
	}

	//跳转到登录页面
	this.Redirect("/login", 302)
}

//退出登录
func (this *UserController) Logout() {
	//删除session
	this.DelSession("userName")

	//跳转页面
	this.Redirect("/index", 302)
}

//显示用户中心
func (this *UserController) ShowUserCenterInfo() {

	//获取用户名
	userName := this.GetSession("userName")
	this.Data["userName"] = userName.(string)

	this.Layout = "layout.html"
	this.TplName = "user_center_info.html"

}

//显示用户中心订单页面
func (this *UserController) UserCenterOrder() {
	this.Layout = "layout.html"
	this.TplName = "user_center_order.html"
}

//显示收件人地址页面
func (this *UserController) ShowUserCenterSite() {
	// beego.Info("啦啦啦啦啦啦啦")
	this.Layout = "layout.html"
	this.TplName = "user_center_site.html"
}

//添加用户信息
func (this *UserController) HandleSite() {
	//获取数据
	receiverName := this.GetString("receiver")
	zipCode := this.GetString("zipCode")
	phone := this.GetString("phone")
	addr := this.GetString("addr")

	//校验数据
	if receiverName == "" || zipCode == "" || phone == "" || addr == "" {
		beego.Error("获取信息失败")
		this.Redirect("/goods/userCenterInfo", 302)
		return
	}

	//处理数据
	//创建orm对象
	o := orm.NewOrm()
	//获取操作对象
	var receiver models.Receiver
	//给插入对象赋值
	receiver.Name = receiverName
	receiver.ZipCode = zipCode
	receiver.Phone = phone
	receiver.Addr = addr

	//获取user对象
	userName := this.GetSession("userName")
	//查询数据库，给user赋值
	var user models.User
	user.UserName = userName.(string)
	//读一下
	o.Read(&user, "UserName")

	receiver.User = &user

	//每次新插入的地址为默认地址,需要把以前的默认地址更新为非默认地址
	//获取receiver对象
	var oldReceiver models.Receiver
	//查询当前用户是否有默认地址，查询当前用户的所有收件人地址
	qs := o.QueryTable("Receiver").RelatedSel("User").Filter("User__Id", user.Id)
	//查询是否有默认地址
	err := qs.Filter("IsDefault", true).One(&oldReceiver)
	if err == nil {
		//把默认地址改为非默认地址
		oldReceiver.IsDefault = false
		//更新
		o.Update(&oldReceiver)
	}

	//把新地址作为默认地址插入
	receiver.IsDefault = true
	o.Insert(&receiver)

	//重定向页面
	this.Redirect("/goods/userCenterInfo", 302)
}
