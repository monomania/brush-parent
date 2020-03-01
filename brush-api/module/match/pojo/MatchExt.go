package pojo

import (
	entity2 "tesou.io/platform/brush-parent/brush-api/module/core/pojo"
)

/**
比赛信息配置
*/
type MatchExt struct {
	entity2.SourceConfig `xorm:"extends"`
}
