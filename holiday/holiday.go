package holiday

import (
	"encoding/json"
	"fmt"
	"github.com/golang-module/carbon"
	"os"
	"strconv"
	"time"
)

type holidayConfig struct {
	Holidays           map[string]string
	SpecialWorkingDays map[string]string
}

type DayProperty struct {
	IsHoliday   bool
	Description string
}

var nowCarbon = carbon.Time2Carbon(time.Now())

func GetHolidayConfig(holidayConfig *holidayConfig) {
	file, err := os.ReadFile("./json/holiday" + strconv.Itoa(nowCarbon.Year()) + ".json")
	if err != nil {
		fmt.Printf("get holiday%s.json file fail. error: %s", strconv.Itoa(nowCarbon.Year()), err.Error())
	}
	err = json.Unmarshal(file, &holidayConfig)
	if err != nil {
		fmt.Printf("get holidayConfig struct fail. error: %s", err.Error())
	}
}

func CreateCalendar(calendar map[string]DayProperty) {
	// 本年开始时间
	yearStartDate := nowCarbon.StartOfYear()
	// 本年结束时间
	yearEndDate := nowCarbon.EndOfYear()
	currentDate := yearStartDate
	for currentDate.DiffInDays(yearEndDate) > 0 {
		currentDateStr := currentDate.ToDateString()
		property := DayProperty{}
		SetHoliday(&property, currentDateStr)
		calendar[currentDateStr] = property
		currentDate = currentDate.AddDay()
	}
}

func SetHoliday(property *DayProperty, currentDay string) {
	current := carbon.Parse(currentDay)
	// 根据周几初始化日历
	property.IsHoliday = false
	property.Description = "工作"
	if current.IsSaturday() {
		flag := IsBigWeek(current.Carbon2Time())
		if flag {
			property.IsHoliday = true
			property.Description = "休息"
		} else {
			property.IsHoliday = false
			property.Description = "工作"
		}
	}
	if current.IsSunday() {
		property.IsHoliday = true
		property.Description = "休息"
	}

	// 根据法定节假日设置日历
	config := &holidayConfig{}
	GetHolidayConfig(config)
	if value, ok := config.Holidays[currentDay]; ok {
		property.Description = value
		property.IsHoliday = true
	}

	if value, ok := config.SpecialWorkingDays[currentDay]; ok {
		property.Description = value
		property.IsHoliday = false
	}
}

func IsBigWeek(currentDate time.Time) bool {
	// 挑选一个大周作为大周的判断日期
	flagDate := time.Date(2022, 8, 8, 0, 0, 0, 0, time.Local)
	weekCount := carbon.Time2Carbon(flagDate).DiffAbsInWeeks(carbon.Time2Carbon(currentDate))
	if weekCount%2 == 0 {
		return true
	}
	return false
}
