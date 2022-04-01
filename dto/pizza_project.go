package dto

import "time"

type OwnedProject struct {
	ID       string `json:"id"`
	Level    string `json:"level"`
	Name     string `json:"name"`
	Owner    string `json:"owner"`
	URL      string `json:"url"`
	Readable uint8  `json:"readable"`
}

type ProjectGeneral struct {
	Workcode       string                    `json:"workcode"`        // 工号
	Username       string                    `json:"username"`        // 姓名
	Deptname       string                    `json:"deptname"`        // 一级部门
	Leaderid       string                    `json:"leaderid"`        // 汇报人
	LeaderAvgscore float64                   `json:"leader_avgscore"` // 上级合作分平均值
	LevelAvgscore  float64                   `json:"level_avgscore"`  // 平级别合作分平均值
	DeptAvgscore   float64                   `json:"dept_avgscore"`   // 一级部门合作分平均值
	ProjectNum     map[string]int            `json:"project_num"`
	ProjectStatus  map[string]int            `json:"project_status"`
	RoleNum        map[string]int            `json:"role_num"`
	RoleLevelNum   map[string]map[string]int `json:"role_level_num"`
	UpdateTime     string                    `json:"update_time"`
}

type ProjectListWithPage struct {
	CurrentPage int           `json:"current_page"`
	LastPage    int           `json:"last_page"`
	Total       int64         `json:"total"`
	Data        []ProjectInfo `json:"data"`
}

type ProjectInfo struct {
	Workcode          string                     `json:"workcode"`            // 工号
	ProjectStory      string                     `json:"project_story"`       // 点赞故事
	ProjectAdvice     string                     `json:"project_advice"`      // 有待提升
	ProjectRole       string                     `json:"project_role"`        // 人员项目角色
	ProjectCode       string                     `json:"project_code"`        // 项目code
	ProjectName       string                     `json:"project_name"`        // 项目名称
	ProjectLevel      string                     `json:"project_level"`       // 项目级别
	ProjectLevelShort string                     `json:"project_level_short"` // 项目级别符号
	ProjectLevelNum   string                     `json:"project_level_num"`   // 项目级别数字
	ProjectStatus     string                     `json:"project_status"`      // 项目状态
	ProjectResult     string                     `json:"project_result"`      // 目标完成情况
	ProjectUsetime    string                     `json:"project_usetime"`     // 项目所用时间情况
	ProjectEvaluate   string                     `json:"project_evaluate"`    // 项目评优奖项
	ProjectOkr        string                     `json:"project_okr"`         // 项目关联OKR
	ProjectStarttime  string                     `json:"project_starttime"`   // 项目开始时间
	ProjectEnddate    time.Time                  `json:"project_enddate"`     // 项目结束时间
	ProjectDetail     map[string][]UserBasicInfo `json:"project_detail"`
}

type UserBasicInfo struct {
	Workcode string `json:"workcode"` // 工号
	Username string `json:"username"` // 姓名
}