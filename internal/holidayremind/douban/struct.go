package douban

type CollectionResponse struct {
	Count int              `json:"count"`
	Items []CollectionItem `json:"subject_collection_items"`
	Total int              `json:"total"`
	Start int              `json:"start"`
}

type Collection struct {
	SubjectType           string      `json:"subject_type"`
	SubTitle              string      `json:"subtitle"`
	BackgroundColorScheme ColorScheme `json:"background_color_scheme"`
	IsFollow              bool        `json:"is_follow"`
	UpdatedAt             string      `json:"updated_at"`
	ScreenshotTitle       string      `json:"screenshot_title"`
	ScreenshotUrl         string      `json:"screenshot_url"`
	Total                 string      `json:"total"`
	ScreenshotType        string      `json:"screenshot_type"`
	TypeIconBgText        string      `json:"type_icon_bg_text"`
	Category              string      `json:"category"`
	IsOfficial            bool        `json:"is_official"`
	IsMergedCover         bool        `json:"is_merged_cover"`
	HeaderFgImage         string      `json:"header_fg_image"`
	Title                 string      `json:"title"`
	WxQrCode              string      `json:"wx_qr_code"`
	IconText              string      `json:"icon_text"`
	Id                    string      `json:"id"`
	FollowersCount        int         `json:"followers_count"`
	ShowHeaderMask        bool        `json:"show_header_mask"`
	MediumName            string      `json:"medium_name"`
	RankType              string      `json:"rank_type"`
	Description           string      `json:"description"`
	ShortName             string      `json:"short_name"`
	NFollowers            int         `json:"n_followers"`
	CoverUrl              string      `json:"cover_url"`
	HeaderBgImage         string      `json:"header_bg_image"`
	CanFollow             bool        `json:"can_follow"`
	ShowRank              bool        `json:"show_rank"`
	ChartId               int         `json:"chart_id"`
	Name                  string      `json:"name"`
	DoneCount             int         `json:"done_count"`
	SharingUrl            string      `json:"sharing_url"`
	WxAppCode             string      `json:"wx_app_code"`
	SubjectCount          int         `json:"subject_count"`
	ItemsCount            int         `json:"items_count"`
	WechatTimelineShare   string      `json:"wechat_timeline_share"`
	CollectCount          int         `json:"collect_count"`
	Url                   string      `json:"url"`
	Type                  string      `json:"type"`
	IsBadgeChart          bool        `json:"is_badge_chart"`
	Uri                   string      `json:"uri"`
	MiniProgramPage       string      `json:"mini_program_page"`
	IconFgImage           string      `json:"icon_fg_image"`
	FinishSoon            bool        `json:"finish_soon"`
	MoreDescription       string      `json:"more_description"`
	ListType              string      `json:"list_type"`
	MiniProgramName       string      `json:"mini_program_name"`
	Display               Display     `json:"display"`
}

type Display struct {
	Layout string `json:"layout"`
}

type CollectionItem struct {
	Comment           string      `json:"comment"`
	Rating            Rating      `json:"rating"`
	ControversyReason string      `json:"controversy_reason"`
	Picture           Picture     `json:"pic"`
	Rank              int         `json:"rank"`
	Uri               string      `json:"uri"`
	IsShow            bool        `json:"is_show"`
	VendorIcons       []string    `json:"vendor_icons"`
	CardSubtitle      string      `json:"card_subtitle"`
	Id                string      `json:"id"`
	Title             string      `json:"title"`
	HasLineWatch      bool        `json:"has_linewatch"`
	IsReleased        bool        `json:"is_released"`
	ColorScheme       ColorScheme `json:"color_scheme"`
	Type              string      `json:"type"`
	Description       string      `json:"description"`
	Tags              []Tag       `json:"tags"`
	CoverUrl          string      `json:"cover_url"`
	Photos            []string    `json:"photos"`
	Actions           []string    `json:"actions"`
	SharingUrl        string      `json:"sharing_url"`
	Url               string      `json:"url"`
	HonorInfos        []HonorInfo `json:"honor_infos"`
	GoodRatingStats   int         `json:"good_rating_stats"`
	SubType           string      `json:"subtype"`
	NullRatingReason  string      `json:"null_rating_reason"`
}

type Rating struct {
	Count     int     `json:"count"`
	Max       int     `json:"max"`
	StarCount float32 `json:"star_count"`
	Value     float32 `json:"value"`
}

type Picture struct {
	Large  string `json:"large"`
	Normal string `json:"normal"`
}

type ColorScheme struct {
	IsDark            bool      `json:"is_dark"`
	PrimaryColorLight string    `json:"primary_color_light"`
	BaseColor         []float32 `json:"_base_color"`
	SecondaryColor    string    `json:"secondary_color"`
	AvgColor          []float32 `json:"_avg_color"`
	PrimaryColorDark  string    `json:"primary_color_dark"`
}

type Tag struct {
	Type string `json:"type"`
	Name string `json:"name"`
}

type HonorInfo struct {
	Kind  string `json:"kind"`
	Uri   string `json:"uri"`
	Rank  int    `json:"rank"`
	Title string `json:"title"`
}

type CollectionParams struct {
	Start     int `json:"start"`
	Count     int `json:"count"`
	ItemsOnly int `json:"items_only"`
	ForMobile int `json:"for_mobile"`
}
