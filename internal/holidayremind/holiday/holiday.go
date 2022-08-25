package holiday

import (
	"github.com/golang-module/carbon"
	"holidayRemind/internal/pkg/uxconfig"
	"strconv"
	"time"
)

var nowCarbon = carbon.Time2Carbon(time.Now())

func CreateCalendar(calendar *map[string]DayProperty) error {
	var err error
	// 根据法定节假日设置日历
	config := holidayConfig{}
	err = getHolidayConfig(&config)
	if err != nil {
		return err
	}
	// 本年开始时间
	yearStartDate := nowCarbon.StartOfYear()
	// 本年结束时间
	yearEndDate := nowCarbon.EndOfYear()
	currentDate := yearStartDate
	for currentDate.DiffInDays(yearEndDate) > 0 {
		currentDateStr := currentDate.ToDateString()
		property := DayProperty{}
		err = setHoliday(currentDateStr, config, &property)
		if err != nil {
			return err
		}
		(*calendar)[currentDateStr] = property
		currentDate = currentDate.AddDay()
	}
	return nil
}

func setHoliday(currentDay string, config holidayConfig, property *DayProperty) error {
	current := carbon.Parse(currentDay)
	property = getDayProperty(false, "工作")
	// 根据周几初始化日历
	if current.IsSaturday() {
		flag := isBigWeek(current.Carbon2Time())
		if flag {
			property = getDayProperty(true, "休息")
		}
	} else if current.IsSunday() {
		property = getDayProperty(true, "休息")
	}

	if value, ok := config.Holidays[currentDay]; ok {
		property = getDayProperty(true, value)
	}

	if value, ok := config.SpecialWorkingDays[currentDay]; ok {
		property = getDayProperty(false, value)
	}
	return nil
}

func getHolidayConfig(holidayConfig *holidayConfig) error {
	fileName := "holiday" + strconv.Itoa(nowCarbon.Year())
	err := uxconfig.GetValue(fileName, uxconfig.Json, uxconfig.Path, holidayConfig)
	if err != nil {
		return err
	}
	return nil
}

func isBigWeek(currentDate time.Time) bool {
	// 挑选一个大周作为大周的判断日期
	flagDate := time.Date(2022, 8, 8, 0, 0, 0, 0, time.Local)
	weekCount := carbon.Time2Carbon(flagDate).DiffAbsInWeeks(carbon.Time2Carbon(currentDate))
	if weekCount%2 == 0 {
		return true
	}
	return false
}
