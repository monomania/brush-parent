module tesou.io/platform/brush-parent

require (
	github.com/astaxie/beego v1.12.0
	github.com/go-vgo/robotgo v0.0.0-20191226160149-28f256a4c5a0
	github.com/lxn/walk v0.0.0-20191128110447-55ccb3a9f5c1
	gopkg.in/Knetic/govaluate.v3 v3.0.0 // indirect
	tesou.io/platform/brush-parent/brush-api v1.0.0
	tesou.io/platform/brush-parent/brush-core v1.0.0
	tesou.io/platform/brush-parent/brush-web v1.0.0
)

replace (
	github.com/go-xorm/core v0.6.3 => github.com/go-xorm/core v0.6.2
	opensource.io/go_spider => ../../../../opensource.io/go_spider
	tesou.io/platform/brush-parent/brush-api => ./brush-api
	tesou.io/platform/brush-parent/brush-core => ./brush-core
	tesou.io/platform/brush-parent/brush-web => ./brush-web
)

go 1.13
