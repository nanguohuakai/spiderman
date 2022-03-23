package dto

type EmployeeInfo struct {
	ID               uint64   `json:"id"`
	EmplID           string   `json:"workcode"`
	Name             string   `json:"name,omitempty"`               // 姓名
	HrStatus         string   `json:"hr_status,omitempty"`          // 员工状态 （A 在职 I 离职）
	Yachid           string   `json:"yachid,omitempty"`             // yachid
	Email            string   `json:"email,omitempty"`              // 邮箱
	Phone            string   `json:"phone,omitempty"`              // 手机号
	DeptDescr        string   `json:"dept_descr,omitempty"`         // 部门名称
	JobSeqDescr      string   `json:"job_seq_descr,omitempty"`      // 职务序列描述
	Sl               *float64 `json:"sl,omitempty"`                 // 司龄（年）
	EthnicGrpCd      string   `json:"ethnic_grp_cd,omitempty"`      // 民族描述
	SexDescr         string   `json:"sex_descr,omitempty"`          // 性别描述
	MarDescr         string   `json:"mar_descr,omitempty"`          // 婚姻描述
	Avatar           string   `json:"avatar,omitempty"`             // 头像
	Lvl              *int64   `json:"lvl,omitempty"`                // 职级 code
	ChannelLvl       string   `json:"channel_lvl,omitempty"`        // 新职级
	Level            string   `json:"level,omitempty"`              // 旧职级
	NativePlaceChn   string   `json:"native_place_chn,omitempty"`   // 籍贯
	Age              *int64   `json:"age,omitempty"`                // 年龄
	JobcodeDescrLast string   `json:"jobcode_descr_last,omitempty"` // 职位描述
	ReportLine       string   `json:"report_line"`                  // 汇报线
	SuperiorCode     string   `json:"superior_code"`                // 上级
	SuperiorName     string   `json:"superior_name"`                // 上级姓名
	DetpID           string   `json:"detp_id"`                      // 部门id
	AccountID        string   `json:"account_id"`                   // 账号id，每个用户唯一，36个长度字符串
	IsZZB            string   `json:"is_zzb"`                       // 是否高管

	Jobcode            string `json:"job_code"`
	JobFamily          string `json:"job_family"`           // 职务族
	JobFamilyDesc      string `json:"job_family_desc"`      // 职务族描述
	JobSeq             string `json:"job_seq"`              // 职务序列
	SubSequence        string `json:"sub_sequence"`         // 职务子序列
	SubSequenceDesc    string `json:"sub_sequence_desc"`    // 职务子序列描述
	JobChannel         string `json:"job_channel"`          // 职务通道
	JobChannelDesc     string `json:"job_channel_desc"`     // 职务通道描述
	Conversion         string `json:"conversion"`           // 是否转正
	LevelExperienceDay *int   `json:"level_experience_day"` // 当前职级晋升时间(日)
	EmplClass          string `json:"empl_class"`           // 员工类别id
	EmplClassDesc      string `json:"empl_class_desc"`      // 员工类别描述
	JobcodeDescr       string `json:"jobcode_descr,omitempty"`
	TJobcodeDescr      string `json:"t_jobcode_descr,omitempty"`

	ChannelChilDesc   string `json:"channel_chil_desc"` // 专业子通道
	ChannelChil       string `json:"channel_chil"`      // 专业子通道编码
	ChannelDesc       string `json:"channel_desc"`      // 专业通道
	Channel           string `json:"channel"`           // 专业通道编码
	HrExpctPstvDt     string `json:"hr_expct_pstv_dt"`  // 预计转正日期
	DeptPath          string `json:"dept_path"`         //部门路径
	LastHireDt        string `json:"last_hire_dt"`
	Location          string `json:"location"`
	JobExperienceTime *int64 `json:"job_experience_time"` //当前岗位时间
	Septum            string `json:"septum"`              // 隔级工号
	SeptumName        string `json:"septum_name"`         // 隔级姓名
	Day               int64  `json:"day"`

	ThisLevelOn         string `json:"this_level_on"`
	LevelExperienceTime *int   `json:"level_experience_time"`

	HrWhetherEffdt string `json:"hr_whether_effdt"` // 成为高管日期
	RusiDt         string `json:"rusi_dt"`          // 入司日期
	Birthdate      string `json:"birthdate"`        // 生日
	Nickname       string `json:"nickname"`         // 昵称
	WYears         string `json:"w_years"`          // 工龄年数
	PermAddr       string `json:"perm_addr"`        // 永久(户口所在地)
	PoliticalDescr string `json:"political_descr"`  // 政治面貌描述
	NationalID     string `json:"national_id"`      // 身份证号
	NationalIDType string `json:"national_id_type"` // 身份证类型
	TerminationDt string `json:"termination_dt"`
}