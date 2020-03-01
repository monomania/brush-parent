package main

import (
	"tesou.io/platform/brush-parent/brush-api/common/base"
	_ "tesou.io/platform/brush-parent/brush-core/common/fliters"
	_ "tesou.io/platform/brush-parent/brush-core/common/routers"
	"tesou.io/platform/brush-parent/brush-core/module/wechat/controller"
	"time"
)

func main() {
	materialController := new(controller.MaterialController)
	materialController.ModifyNewsOnly()
	base.Log.Info("--------发布公众号周期结束--------")
	time.Sleep(10 * time.Minute)

}
