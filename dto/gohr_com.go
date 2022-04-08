package dto

type UserListResponse struct {
	CurrentPage  int           `json:"current_page"`
	Data         []OkrListData `json:"data"`
	FirstPageUrl string        `json:"first_page_url"`
	From         int           `json:"from"`
	LastPage     int           `json:"last_page"`
	LastPageUrl  string        `json:"last_page_url"`
	NextPageUrl  string        `json:"next_page_url"`
	Path         string        `json:"path"`
	PerPage      int           `json:"per_page"`
	PrevPageUrl  string        `json:"prev_page_url"`
	To           int           `json:"to"`
	Total        int           `json:"total"`
}

type OkrListData struct {
	Id             int    `json:"id"`
	Title          string `json:"title"`
	CommentsCount  int    `json:"comments_count"`
	AlignsCount    int    `json:"aligns_count"`
	FollowersCount int    `json:"followers_count"`
	UpdateAt       string `json:"updated_at"`
}

type RecordJournalYachResponse struct {
	List []RecordJournalData `json:"list"`
	Url  string              `json:"url"`
}

type RecordJournalData struct {
	Title           string          `json:"title"`
	TemplateId      string          `json:"template_id"`
	CreateTime      string          `json:"create_time"`
	ContentInfoList []ContentInfo   `json:"content_info"`
	UserInfo        JournalUserInfo `json:"user_info"`
}
type ContentInfo struct {
	Title   string `json:"title"`
	Content string `json:"content"`
}
type JournalUserInfo struct {
	Id       int      `json:"id"`
	UserId   string   `json:"user_id"`
	Workcode string   `json:"work_code"`
	Avatar   string   `json:"avatar"`
	Name     string   `json:"name"`
	Mobile   string   `json:"mobile"`
	Nickname string   `json:"nickname"`
	OgIDs    []string `json:"og_ids"`
}

type UserInfoResponse struct {
	GroupName         string           `json:"groupName"`
	Capability        []CapabilityItem `json:"capability"`
	WorkCode          string           `json:"workCode"`
	GroupId           string           `json:"groupId"`
	UserName          string           `json:"userName"`
	BusinessLabelList []string         `json:"businessLabelList"`
	JishuLabelList    []string         `json:"jishuLabelList"`
}

type CapabilityItem struct {
	Key   string  `json:"key"`
	Score float64 `json:"score"`
}

type WeekReportYachResponse struct {
	Url  *string           `json:"url"`
	Data *[]WeekReportData `json:"data"`
}

type WeekReportData struct {
	Content       string  `json:"content"`
	Content_full  string  `json:"content_full"`
	Create_at     float64 `json:"create_at"`
	Update_at     float64 `json:"update_at"`
	ZanNumber     int     `json:"zan_number"`
	CommentNumber int     `json:"comment_number"`
}
