package dto

type DeptInfo struct {
	ID             uint64 `json:"id"`
	DeptID         string `json:"dept_id"`      // 部门ID
	Desc           string `json:"desc"`         // 部门描述
	EffStatus      string `json:"eff_status"`   // 生效状态
	DeptDeepLv     string `json:"dept_deep_lv"` // 部门层级
	Path           string `json:"path"`
	ShortDesc      string `json:"short_desc"`       // 部门短描述
	ParentNodeName string `json:"parent_node_name"` // 父节点名称
	ManagerName    string `json:"manager_name"`     // 姓名
	ManagerID      string `json:"manager_id"`       // 工号
	DeptLvl        string `json:"dept_lvl"`         // 部门级别
	DeptXes        string `json:"dept_xes"`         // 级别='集团'
	DeptSyb        string `json:"dept_syb"`         // 级别='事业部'
	DeptJ00        string `json:"dept_j00"`         // 级别='集团总部'
	DeptS00        string `json:"dept_s00"`         // 级别='事业部总部'
	DeptF00        string `json:"dept_f00"`         // 级别='分校'
	DeptJ01        string `json:"dept_j01"`         // 级别='集团一级部门'
	DeptJ02        string `json:"dept_j02"`         // 级别='集团二级部门'
	DeptJ03        string `json:"dept_j03"`         // 级别='集团三级部门'
	DeptJ04        string `json:"dept_j04"`         // 级别='集团四级部门'
	DeptS01        string `json:"dept_s01"`         // 级别='事业部一级部门'
	DeptS02        string `json:"dept_s02"`         // 级别='事业部二级部门'
	DeptS03        string `json:"dept_s03"`         // 级别='事业部三级部门'
	DeptS04        string `json:"dept_s04"`         // 级别='事业部四级部门'
	DeptS05        string `json:"dept_s05"`         // 级别='事业部五级部门'
	DeptF01        string `json:"dept_f01"`         // 级别='分校一级部门'
	DeptF02        string `json:"dept_f02"`         // 级别='分校二级部门'
	DeptF03        string `json:"dept_f03"`         // 级别='分校三级部门'
	DeptF04        string `json:"dept_f04"`         // 级别='分校四级部门'
	DeptF05        string `json:"dept_f05"`         // 级别='分校五级部门'
	DeptVig        string `json:"dept_vig"`         // 级别='大虚拟事业部'
	DeptVis        string `json:"dept_vis"`         // 级别='虚拟事业部'
	DeptVit        string `json:"dept_vit"`         // 级别='虚拟部门'
	DeptLvl1       string `json:"dept_lvl1"`        // 前台/中台/后台
	DeptLvl2       string `json:"dept_lvl2"`        // 虚拟事业部
	DeptLvl3       string `json:"dept_lvl3"`        // 事业部/项目组
	DeptLvl4       string `json:"dept_lvl4"`        // 分校/前中后台一级部门
	DeptLvl5       string `json:"dept_lvl5"`        // 分校一级部门/前中后台二级部门
	HrbpCode       string `json:"hrbp_code"`        // 部门属性id
	HrbpName       string `json:"hrbp_name"`        // 部门属性id值
	HrbpMgrCode    string `json:"hrbp_mgr_code"`    // 部门属性id
	HrbpMgrName    string `json:"hrbp_mgr_name"`    // 部门属性id值
	HrdCode        string `json:"hrd_code"`         // 部门属性id
	HrdName        string `json:"hrd_name"`         // 部门属性id值
}
