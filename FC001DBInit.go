package main

import "tesou.io/platform/brush-parent/brush-core/launch"

func main() {

	//生成数据库表
	launch.GenTable()
	////清空数据表
	//launch.TruncateTable()
}
