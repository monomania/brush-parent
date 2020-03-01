package controller

import "tesou.io/platform/brush-parent/brush-core/common/base/controller"

type IndexController struct {
	controller.BaseController
}

// @router / [get]
func (this *IndexController) Get() {
	this.Data["json"] = "hello"
	this.ServeJSON()
}
