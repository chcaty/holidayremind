package moyuduck

// Response 摸鱼鸭返回结果
type Response[T string | Holiday | HotTopSite] struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
	Data T      `json:"data"`
}

// Holiday 假期信息
type Holiday struct {
	Time    string       `json:"time"`
	Week    string       `json:"week"`
	WeekDay string       `json:"weekDay"`
	List    []HolidayDay `json:"list"`
}

// HolidayDay 详细假期信息
type HolidayDay struct {
	Date string `json:"date"`
	Name string `json:"name"`
	Day  int    `json:"day"`
}

// HotTopSite 热点平台列表
type HotTopSite struct {
	ZhiHu    HotTopInfo `json:"zhihu"`
	DouYin   HotTopInfo `json:"douyin"`
	Weibo    HotTopInfo `json:"weibo"`
	Baidu    HotTopInfo `json:"baidu"`
	BiliBili HotTopInfo `json:"bilibili"`
	History  HotTopInfo `json:"history"`
	TieBa    HotTopInfo `json:"tieba"`
	TouTiao  HotTopInfo `json:"toutiao"`
	PoJie52  HotTopInfo `json:"pojie52"`
}

// HotTopInfo 热点信息列表
type HotTopInfo struct {
	HotTops []HotTopDetail
	Time    string
}

// HotTopDetail 热点详细信息
type HotTopDetail struct {
	Title    string
	Url      string
	HotValue string
}

type ResponseData interface {
	Response[string] | Response[Holiday] | Response[HotTopSite]
}
