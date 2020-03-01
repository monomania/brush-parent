package vo

import "tesou.io/platform/brush-parent/brush-api/module/analy/pojo"

type AnalyResultVO struct {
	pojo.AnalyResult `xorm:"extends"`
	/**
	联赛名称
	 */
	LeagueName string
	/**
	 * 主队id
	 */
	MainTeamId string
	/**
	 * 客队id
	 */
	GuestTeamId string

}
