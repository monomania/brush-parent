module tesou.io/platform/brush-parent/brush-core

require (
	github.com/PuerkitoBio/goquery v1.5.0
	github.com/astaxie/beego v1.12.0
	github.com/chanxuehong/wechat v0.0.0-20190521093015-fafb751f9916
	github.com/go-sql-driver/mysql v1.4.1
	github.com/go-xorm/core v0.6.3
	github.com/go-xorm/xorm v0.7.9
	github.com/shiena/ansicolor v0.0.0-20151119151921-a422bbe96644 // indirect
	github.com/smartystreets/goconvey v1.6.4 // indirect
	google.golang.org/appengine v1.6.5 // indirect
	gopkg.in/ini.v1 v1.51.0
	gopkg.in/mgo.v2 v2.0.0-20190816093944-a6b53ec6cb22
	tesou.io/platform/brush-parent/brush-api v1.0.0
)

replace (
	github.com/go-xorm/core v0.6.3 => github.com/go-xorm/core v0.6.2
	tesou.io/platform/brush-parent/brush-api => ../brush-api
)

go 1.13
