package holiday

// 假期配置
type holidayConfig struct {
	Holidays           map[string]string //假期
	SpecialWorkingDays map[string]string //补班日
}

// DayProperty 日期属性
type DayProperty struct {
	IsHoliday   bool   // 是否假期
	Description string // 日期描述
}

func getDayProperty(isHoliday bool, description string) *DayProperty {
	return &DayProperty{
		IsHoliday:   isHoliday,
		Description: description,
	}
}
