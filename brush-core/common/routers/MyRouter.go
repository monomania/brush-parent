package routers

import (
	"github.com/astaxie/beego"
	"tesou.io/platform/brush-parent/brush-core/module/index/controller"
	controller2 "tesou.io/platform/brush-parent/brush-core/module/match/controller"
	controller3 "tesou.io/platform/brush-parent/brush-core/module/wechat/controller"
)

type MyRouter struct {

}

func init() {
	beego.Router("/", &controller.IndexController{})

	//match
	beego.AutoRouter(&controller2.MatchController{})
	beego.AutoRouter(&controller2.MatchLastController{})

	//wechat
	beego.AutoRouter(&controller3.WechatController{})
	beego.AutoRouter(&controller3.MaterialController{})
}

func (this *MyRouter) Hello(){

}
