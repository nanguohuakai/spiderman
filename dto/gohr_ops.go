package dto

type JobChannelData struct {
	Total       int              `json:"total"`
	CurrentPage int              `json:"current_page"`
	PerPage     int              `json:"per_page"`
	List        []jobChannelList `json:"list"`
}

type jobChannelList struct {
	JobChannel string `json:"T_JOB_CHANNEL"`
	Effdt      string `json:"EFFDT"`
	EffStatus  string `json:"EFF_STATUS"`
	Descr      string `json:"DESCR"`
	DescrShort string `json:"DESCRSHORT"`
}

type JobcData struct {
	Total       int        `json:"total"`
	CurrentPage int        `json:"current_page"`
	PerPage     int        `json:"per_page"`
	List        []jobcList `json:"list"`
}

type jobcList struct {
	JobCode        string       `json:"JOBCODE"`
	Effdt          string       `json:"EFFDT"`
	EffStatus      string       `json:"EFF_STATUS"`
	Descr          string       `json:"DESCR"`
	JobFamily      string       `json:"JOB_FAMILY"`
	JobFunction    string       `json:"JOB_FUNCTION"`
	JobSubFunc     string       `json:"JOB_SUB_FUNC"`
	CJobFamily     string       `json:"C_JOB_FAMILY"`
	CJobSeq        string       `json:"C_JOB_SEQ"`
	SubSequence    string       `json:"T_SUB_SEQUENCE"`
	JobChannel     string       `json:"T_JOB_CHANNEL"`
	ProChannel     string       `json:"T_PRO_CHANNEL"`
	ProChannelChit string       `json:"T_PRO_CHANNEL_CHIL"`
	Lvl            []jobcLvlMap `json:"LVL"`
	NewLvl         []jobcLvlMap `json:"NEW_LVL"`
}

type jobcLvlMap struct {
	Id   string `json:"id"`
	Name string `json:"name"`
}

type JobSeqData struct {
	Total       int          `json:"total"`
	CurrentPage int          `json:"current_page"`
	PerPage     int          `json:"per_page"`
	List        []jobSeqList `json:"list"`
}

type jobSeqList struct {
	CJobSeq    string `json:"C_JOB_SEQ"`
	Effdt      string `json:"EFFDT"`
	EffStatus  string `json:"EFF_STATUS"`
	Descr      string `json:"DESCR"`
	DescrShort string `json:"DESCRSHORT"`
}
