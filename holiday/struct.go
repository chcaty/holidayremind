package holiday

type holidayConfig struct {
	Holidays           map[string]string
	SpecialWorkingDays map[string]string
}

type DayProperty struct {
	IsHoliday   bool
	Description string
}
