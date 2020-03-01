package controller

import "tesou.io/platform/brush-parent/brush-core/common/base/controller"

type MatchLastController struct {
	controller.BaseController
}

func (th *MatchLastController) Get() {
	th.Data["json"] = "match"
	th.ServeJSON()
}
