package launch

import (
	"tesou.io/platform/brush-parent/brush-core/common/base/service/mysql"
)


func TruncateTable() {
	//开启SQL输出

	opsService := new(mysql.DBOpsService)
	//指定需要清空的数据表
	opsService.TruncateTable([]string{"t_match_last", "t_asia_last", "t_euro_last"})
}

func GenTable() {
	//开启SQL输出

	generateService := new(mysql.DBOpsService)
	generateService.SyncTableStruct()
}
