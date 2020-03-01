package pojo

import (
	"tesou.io/platform/brush-parent/brush-api/common/base/pojo"
	"tesou.io/platform/brush-parent/brush-api/module/suggest/enums"
)

/**
发布记录
 */
type Pub struct {
	//发布源类型
	Source enums.PubSourceLevel
	//发布的帐号
	Account string
	//发布的赛事ID
	MatchId string
	//发布的analy表的ID
	AnalyId string

	pojo.BasePojo `xorm:"extends"`
}
