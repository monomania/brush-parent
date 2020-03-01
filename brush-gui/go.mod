module tesou.io/platform/brush-parent/brush-gui

replace (
	github.com/go-xorm/core v0.6.3 => github.com/go-xorm/core v0.6.2
	opensource.io/go_spider => ../../../../opensource.io/go_spider
	tesou.io/platform/brush-parent/brush-api => ../brush-api
	tesou.io/platform/brush-parent/brush-core => ../brush-core
	tesou.io/platform/brush-parent/brush-spider => ../brush-spider
)

go 1.13
