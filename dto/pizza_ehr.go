package dto

type DimissionInfo struct {
	ID            uint64 `json:"id"`
	LeaderCode    string `json:"leader_code"`     // 部门领导工号
	LeaderID      int64  `json:"leader_id"`       // 部门领导ID
	EffStatus     string `json:"eff_status"`      // 生效状态（A:有效 I:无效）
	EmplID        string `json:"workcode"`        // 员工工号
	Name          string `json:"name"`            // 姓名
	HrStatus      string `json:"hr_status"`       // 在离职状态 （A 在职， I 离职）
	DeptID        string `json:"dept_id"`         // 部门id
	DeptDescr     string `json:"dept_descr"`      // 部门描述（部门名称）
	DimissionDate string `json:"dimission_date"`  // 离职日期
	Level         string `json:"level"`           // 老职级
	ChannelLvl    string `json:"channel_lvl"`     // 新职级
	Lvl           string `json:"lvl"`             // 职级id
	EmpClass      string `json:"emp_class"`       // 员工类别id
	EmpClassDescr string `json:"emp_class_descr"` // 员工类别描述
	Status        int8   `json:"status"`          // 删除状态，0有效，1已删除
}

type DeptPromotionsInfo struct {
	ID                  uint64 `json:"id"`
	EmplID              string `json:"workcode"`              // 员工工号
	Name                string `json:"name"`                  // 姓名
	HrStatus            string `json:"hr_status"`             // 在离职状态 （A 在职， I 离职）
	DeptID              string `json:"dept_id"`               // 部门id
	LeaderCode          string `json:"leader_code"`           // 部门领导工号
	LeaderID            int64  `json:"leader_id"`             // 部门领导ID
	EffStatus           string `json:"eff_status"`            // 生效状态（A:有效 I:无效）
	DeptDescr           string `json:"dept_descr"`            // 部门描述（部门名称）
	StartDate           string `json:"start_date"`            // 调入日期
	EndDate             string `json:"end_date"`              // 调离日期
	Level               string `json:"level"`                 // 当前老职级
	ChannelLvl          string `json:"channel_lvl"`           // 当前新职级
	Lvl                 string `json:"lvl"`                   // 当前职级id
	PromotionLevel      string `json:"promotion_level"`       // 晋升时老职级
	PromotionChannelLvl string `json:"promotion_channel_lvl"` // 晋升时新职级
	PromotionLvl        string `json:"promotion_lvl"`         // 晋升时职级id
	PromotionDate       string `json:"promotion_date"`        // 晋升日期
	Post                string `json:"post"`                  // 当前岗位
	Status              int8   `json:"status"`                // 删除状态，0有效，1已删除
}

type LeadershipsInfo struct {
	ID        uint64 `json:"id"`
	EmplID    string `json:"empl_id"`    // 部门领导ID
	DeptID    string `json:"dept_id"`    // 部门id
	Descr     string `json:"descr"`      // 部门描述（部门名称）
	EffDt     string `json:"eff_dt"`     // 生效日期
	EffStatus string `json:"eff_status"` // 生效状态（A:有效 I:无效）
	EndTime   string `json:"end_time"`   // 区间结束时间
	Status    int8   `json:"status"`     // 删除状态，0有效，1已删除
}

type EhrChangeItem struct {
	ID                 uint64 `json:"id"`
	EmplID             string `json:"empl_id"`
	StartDate          string `json:"start_date"`
	EndDate            string `json:"end_date"`
	DepartmentId       string `json:"department_id"`
	DepartmentName     string `json:"department_name"`
	DepartmentPosition string `json:"department_position"`
	Position           string `json:"position"`
	Level              string `json:"level"` // 职级
	Reason             string `json:"reason"`
	Type               string `json:"type"`
	FullReason         string `json:"full_reason"`
	JobSeq             string `json:"job_seq"`
	JobSeqDesc         string `json:"job_seq_desc"`
	JobFamily          string `json:"job_family"`
	JobFamilyDescr     string `json:"job_family_descr"`
	SuperiorName       string `json:"superior_name"`
	SuperiorCode       string `json:"superior_code"`
	EmpRcd             string `json:"emp_rcd"` // 岗位序号
	EffSeq             string `json:"eff_seq"`
}