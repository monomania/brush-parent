package controller

import "tesou.io/platform/brush-parent/brush-core/common/base/controller"

type MatchController struct {
	controller.BaseController
}

func (th *MatchController) Get() {
	th.Data["json"] = "match"
	th.ServeJSON()
}
